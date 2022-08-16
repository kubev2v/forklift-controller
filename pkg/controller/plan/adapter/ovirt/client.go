package ovirt

import (
	"fmt"
	"sync"
	"time"

	liberr "github.com/konveyor/controller/pkg/error"
	planapi "github.com/konveyor/forklift-controller/pkg/apis/forklift/v1beta1/plan"
	"github.com/konveyor/forklift-controller/pkg/apis/forklift/v1beta1/ref"
	plancontext "github.com/konveyor/forklift-controller/pkg/controller/plan/context"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/container/ovirt"
	model "github.com/konveyor/forklift-controller/pkg/controller/provider/web/ovirt"
	ovirtsdk "github.com/ovirt/go-ovirt"
	"k8s.io/apimachinery/pkg/util/wait"
	cdi "kubevirt.io/containerized-data-importer-api/pkg/apis/core/v1beta1"
)

const (
	snapshotDesc = "Forklift Operator warm migration precopy"
)

// VM power states
const (
	powerOn      = "On"
	powerOff     = "Off"
	powerUnknown = "Unknown"
)

//
// oVirt VM Client
type Client struct {
	*plancontext.Context
	connection *ovirtsdk.Connection
}

//
// Create a VM snapshot and return its ID.
func (r *Client) CreateSnapshot(vmRef ref.Ref) (snapshot string, err error) {
	_, vmService, err := r.getVM(vmRef)
	if err != nil {
		return
	}
	snapsService := vmService.SnapshotsService()
	snap, err := snapsService.Add().Snapshot(
		ovirtsdk.NewSnapshotBuilder().
			Description(snapshotDesc).
			PersistMemorystate(false).
			MustBuild(),
	).Query("correlation_id", r.Migration.Name).Send()
	if err != nil {
		err = liberr.Wrap(err)
		return
	}
	snapshot = snap.MustSnapshot().MustId()
	return
}

//
// Remove all warm migration snapshots.
func (r *Client) RemoveSnapshots(vmRef ref.Ref, precopies []planapi.Precopy) (err error) {
	// NOOP
	return
}

//
// Set DataVolume checkpoints.
func (r *Client) SetCheckpoints(vmRef ref.Ref, precopies []planapi.Precopy, datavolumes []cdi.DataVolume, final bool) (err error) {
	n := len(precopies)
	previous := ""
	current := precopies[n-1].Snapshot
	if n >= 2 {
		previous = precopies[n-2].Snapshot
	}

	for i := range datavolumes {
		dv := &datavolumes[i]
		var currentDiskSnapshot, previousDiskSnapshot string
		currentDiskSnapshot, err = r.getDiskSnapshot(dv.Spec.Source.Imageio.DiskID, current)
		if err != nil {
			return
		}
		if previous != "" {
			previousDiskSnapshot, err = r.getDiskSnapshot(dv.Spec.Source.Imageio.DiskID, previous)
			if err != nil {
				return
			}
		}

		dv.Spec.Checkpoints = append(dv.Spec.Checkpoints, cdi.DataVolumeCheckpoint{
			Current:  currentDiskSnapshot,
			Previous: previousDiskSnapshot,
		})
		dv.Spec.FinalCheckpoint = final
	}
	return
}

//
// Get the power state of the VM.
func (r *Client) PowerState(vmRef ref.Ref) (state string, err error) {
	vm, _, err := r.getVM(vmRef)
	if err != nil {
		return
	}
	status, _ := vm.Status()
	switch status {
	case ovirtsdk.VMSTATUS_DOWN:
		state = powerOff
	case ovirtsdk.VMSTATUS_UP, ovirtsdk.VMSTATUS_POWERING_UP:
		state = powerOn
	default:
		state = powerUnknown
	}
	return
}

//
// Power on the VM.
func (r *Client) PowerOn(vmRef ref.Ref) (err error) {
	vm, vmService, err := r.getVM(vmRef)
	if err != nil {
		return
	}
	// Request the VM startup if VM is not UP
	if status, _ := vm.Status(); status != ovirtsdk.VMSTATUS_UP {
		_, err = vmService.Start().Send()
		if err != nil {
			err = liberr.Wrap(err)
		}
	}
	return
}

//
// Power off the VM.
func (r *Client) PowerOff(vmRef ref.Ref) (err error) {
	vm, vmService, err := r.getVM(vmRef)
	if err != nil {
		return
	}
	// Request the VM startup if VM is not UP
	if status, _ := vm.Status(); status != ovirtsdk.VMSTATUS_DOWN {
		_, err = vmService.Shutdown().Send()
		if err != nil {
			err = liberr.Wrap(err)
		}
	}
	return
}

//
// Determine whether the VM has been powered off.
func (r *Client) PoweredOff(vmRef ref.Ref) (poweredOff bool, err error) {
	powerState, err := r.PowerState(vmRef)
	if err != nil {
		return
	}
	poweredOff = powerState == powerOff
	return
}

//
// Close the connection to the oVirt API.
func (r *Client) Close() {
	if r.connection != nil {
		_ = r.connection.Close()
		r.connection = nil
	}
}

//
// Get the VM by ref.
func (r *Client) getVM(vmRef ref.Ref) (ovirtVm *ovirtsdk.Vm, vmService *ovirtsdk.VmService, err error) {
	vm := &model.VM{}
	err = r.Source.Inventory.Find(vm, vmRef)
	if err != nil {
		err = liberr.Wrap(
			err,
			"VM lookup failed.",
			"vm",
			vmRef.String())
		return
	}
	vmService = r.connection.SystemService().VmsService().VmService(vm.ID)
	vmResponse, err := vmService.Get().Query("correlation_id", r.Migration.Name).Send()
	if err != nil {
		err = liberr.Wrap(err)
		return
	}
	ovirtVm, ok := vmResponse.Vm()
	if !ok {
		err = liberr.New(
			fmt.Sprintf(
				"VM %s source lookup failed",
				vmRef.String()))
	}
	return
}

//
// Get the disk snapshot for this disk and this snapshot ID.
func (r *Client) getDiskSnapshot(diskID, targetSnapshotID string) (diskSnapshotID string, err error) {
	response, rErr := r.connection.SystemService().DisksService().DiskService(diskID).Get().Query("correlation_id", r.Migration.Name).Send()
	if err != nil {
		err = liberr.Wrap(rErr)
		return
	}
	disk, ok := response.Disk()
	if !ok {
		err = liberr.New("Could not find disk definition in response.", "disk", diskID)
		return
	}

	storageDomains, ok := disk.StorageDomains()
	if !ok {
		err = liberr.New("No storage domains listed for disk.", "disk", diskID)
		return
	}

	for _, sd := range storageDomains.Slice() {
		sdID, ok := sd.Id()
		if !ok {
			continue
		}
		sdService := r.connection.SystemService().StorageDomainsService().StorageDomainService(sdID)
		if sdService == nil {
			err = liberr.New("No service available for storage domain.", "storageDomain", sdID)
			return
		}
		snapshotsResponse, rErr := sdService.DiskSnapshotsService().List().Send()
		if err != nil {
			err = liberr.Wrap(rErr, "Error listing snapshots in storage domain.", "storageDomain", sdID)
			return
		}
		snapshots, ok := snapshotsResponse.Snapshots()
		if !ok || len(snapshots.Slice()) == 0 {
			err = liberr.New("No snapshots listed in storage domain.", "storageDomain", sdID)
			return
		}
		for _, diskSnapshot := range snapshots.Slice() {
			id, ok := diskSnapshot.Id()
			if !ok {
				continue
			}
			snapshotDisk, ok := diskSnapshot.Disk()
			if !ok {
				continue
			}
			snapshotDiskID, ok := snapshotDisk.Id()
			if !ok {
				continue
			}
			snapshot, ok := diskSnapshot.Snapshot()
			if !ok {
				continue
			}
			sid, ok := snapshot.Id()
			if !ok {
				continue
			}
			if snapshotDiskID == diskID && targetSnapshotID == sid {
				diskSnapshotID = id
				return
			}
		}
	}

	err = liberr.New("Could not find disk snapshot.", "disk", diskID, "vmSnapshot", targetSnapshotID)
	return
}

func (r *Client) getJobs(correlationID string) (ovirtJob []*ovirtsdk.Job, err error) {
	jobService := r.connection.SystemService().JobsService().List()
	jobResponse, err := jobService.Search(fmt.Sprintf("correlation_id=%s", correlationID)).Send()
	if err != nil {
		err = liberr.Wrap(err)
		return
	}
	ovirtJobs, ok := jobResponse.Jobs()
	if !ok {
		err = liberr.New(fmt.Sprintf("Job %s source lookup failed", correlationID))
		return
	}
	ovirtJob = ovirtJobs.Slice()
	return
}

func (r *Client) isDiskBeingTransferred(diskID string) (bool, error) {
	transfers, err := r.connection.SystemService().ImageTransfersService().List().Send()
	if err != nil {
		err = liberr.Wrap(err)
		return false, err
	}

	for _, transfer := range transfers.MustImageTransfer().Slice() {
		phase, _ := transfer.Phase()
		if phase != ovirtsdk.IMAGETRANSFERPHASE_FINISHED_FAILURE && phase != ovirtsdk.IMAGETRANSFERPHASE_FINISHED_SUCCESS {
			transferDiskID, _ := transfer.MustImage().Id()
			if transferDiskID == diskID {
				return true, nil
			}
		}
	}

	return false, nil
}

//
// Connect to the oVirt API.
func (r *Client) connect() (err error) {
	URL := r.Source.Provider.Spec.URL
	r.connection, err = ovirtsdk.NewConnectionBuilder().
		URL(URL).
		Username(r.user()).
		Password(r.password()).
		CACert(r.cacert()).
		Insecure(ovirt.GetInsecureSkipVerifyFlag(r.Source.Secret)).
		Build()
	if err != nil {
		err = liberr.Wrap(err)
	}
	return
}

func (r *Client) user() string {
	if user, found := r.Source.Secret.Data["user"]; found {
		return string(user)
	}
	return ""
}

func (r *Client) password() string {
	if password, found := r.Source.Secret.Data["password"]; found {
		return string(password)
	}
	return ""
}

func (r *Client) cacert() []byte {
	if cacert, found := r.Source.Secret.Data["cacert"]; found {
		return cacert
	}
	return nil
}

func (r Client) Finalize(vms []*planapi.VMStatus) {
	r.connect()
	defer r.Close()
	var wg sync.WaitGroup
	wg.Add(len(vms))
	for _, vm := range vms {
		_, vmService, err := r.getVM(vm.Ref)
		if err != nil {
			return
		}

		go r.removePrecopies(vm.Warm.Precopies, vmService, &wg)
	}

	wg.Wait()
}

func (r Client) removePrecopies(precopies []planapi.Precopy, vmService *ovirtsdk.VmService, wg *sync.WaitGroup) {
	if len(precopies) == 0 {
		return
	}

	defer wg.Done()
	snapsService := vmService.SnapshotsService()
	for i := range precopies {
		snapshotID := precopies[i].Snapshot
		snapService := snapsService.SnapshotService(snapshotID)
		correlationID := fmt.Sprintf("%s_finalize", snapshotID[0:8])
		_, err := snapService.Remove().Query("correlation_id", correlationID).Send()
		r.Log.Error(err, "Remove request for snapshot failed", "snapshotID", snapshotID)
		if err != nil {
			// Assume we could not remove because of an ongoing transfer
			backoff := wait.Backoff{
				Duration: time.Second * 3,
				Factor:   0,
				Jitter:   0,
				Steps:    3,
			}

			err = wait.ExponentialBackoff(backoff, func() (bool, error) {
				snapshotDisksResponse, err := snapService.DisksService().List().Send()
				if err != nil {
					r.Log.Error(err, "snapshotDiksService Request failed")
					err = liberr.Wrap(err, "Disks not found")
					return false, err
				}
				snapshotDisks, found := snapshotDisksResponse.Disks()
				if !found {
					r.Log.Error(err, "disks not found")
					err = liberr.Wrap(err, "Disks not found")
					return false, err
				}
				for _, disk := range snapshotDisks.Slice() {
					if transferred, err := r.isDiskBeingTransferred(disk.MustId()); transferred || err != nil {
						r.Log.Info("Disk is being transferred, retrying..")
						return false, nil
					} else {
						// Try to remove the snapshot again
						_, err = snapService.Remove().Query("correlation_id", correlationID).Send()
						if err != nil {
							err = liberr.Wrap(err)
							return true, err
						}
					}
				}
				return false, nil
			})
			if err != nil {
				err = liberr.Wrap(err)
				return
			}

		}

		for {
			jobs, err := r.getJobs(correlationID)
			if err != nil {
				err = liberr.Wrap(err)
			}
			if allJobsFinished(jobs) {
				break
			}

			select {
			case <-time.After(2 * time.Hour):
				r.Log.Info("Timeout out waiting for snapshot removal")
				return
			default:
				time.Sleep(3 * time.Second)
			}
		}
	}
}

func allJobsFinished(jobs []*ovirtsdk.Job) bool {
	for _, job := range jobs {
		status, _ := job.Status()
		if status != ovirtsdk.JOBSTATUS_FINISHED && status != ovirtsdk.JOBSTATUS_FAILED {
			return false
		}
	}

	return true
}
