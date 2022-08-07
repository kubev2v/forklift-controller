Direct LUNs migration
==============

Summary
-------

Usually a VM contains disks. Some of them might be connected as a direct LUN.
When a user wish to use the LUN disks from another enviroment, he needs them attached to the VM.
The ISCSI storage where the LUNs reside, have to be reachable to the VM environments.
The LUN disk should be used in a single place, or else, it will cause a data corruption.

CDI doesn't support LUNs to that matter, we also don't require copying the data, we just need to re-use the same LUN.

Currently, the feature page intend to oVirt only.

Design
------

The requirements for the LUN migration code are (not in priority order)

1. Gather the LUN information from the source provider.
2. To ease on the user - less manual actions, more automated migration.

Requirements
----------------------------------------

1. Destination (OCP cluster) reachability to the ISCSI storage, where the LUN is.
2. RBAC permission to create `PersistentVolumes` by the controller.
   - Will be added to the operator manifest.

Risks and Mitigations
----------------------------------------

1. A VM that will be used while or after the migration with a direct LUN connected.
   - A warning should be shown before starting the migration.
   - Detaching the LUN from source VM will split into two cases, based on the migration type:
     - Cold migration: The LUN disk will be detached before initiating the copy disk. If the migration will fail, re-attach it to the source VM.
     - Warm migration: Once the migration ends and we shut down the VM on source.
   - (Optional) UI checkbox to the user, before starting the migration.
2. An oVirt version with unsupported API. The [PR](https://github.com/oVirt/ovirt-engine/pull/559) provides LUN details is required.
   - We may check for missing LUN details and block the migration.
   - (Optional) Have a mechanism to check the oVirt version for future work.
3. On non-local destination, the controller needs permission to create the `PersistentVolumes`.
   - It may be enough to be added to the forklift-operator.

Implementation
----------------------------------------

The controller gather data from source provider and saves them to the inventory.
As part of the data gathering we will gather the LUN disks data. We have the VMs disk attachments where we follow to the disks.
Within the disks API, if the _storage\_type_ is `LUN`, we will gather from the API endpoint `api/disks/<disk_id>` the size, address, IQN and the mapping of the LUN.
Those will be found withing the `logical_unit` element of the `lun_storage`.
On the migration plan a new steps will be added, creating a `PersistentVolume`(PV) and a `PersistentVolumeClaim`(PVC) per LUN disk.
Both of them will contain the relevant data to use the LUN, the PVC will be bound the the PV by a unique label, `<vm_id>-<vm_name>-vol-<volume-number>`.
When the VM specification created, the controller will add a LUN disk using the PVC that was created as [described by kubevirt](https://kubevirt.io/user-guide/virtual_machines/disks_and_volumes/#lun).

Relevant validations will be performed to know we have the LUN details:

- Allow `LUN` storage type to pass the validation - until know only `image` type is allowed to migrate.
- LUN details exist on the source provider API, the data is neccassary in order to migrate.
- SCSI persistent reservation (exists), may required further warning when LUN is detected.
- (Optional) Make sure the VM is in `down` on source provider (might be not required and be part of the execution plan).
- In case the storage requires authentication, we may need to add secrets.

Notes:

- CSI is optional, it won't able to select the relevant LUN we wish to use. However, we can add the relevant `StorageClass` of the CSI to the `PVC` specifications.
  - Will require creating a mapping for each LUN to the relevant `StorageClass`.
  - User will need to pre-install the CSI and have the `StorageClass`.
  - `PV` creation still needed to have the specific LUN details.

Execution Flow
----------------------------------------

1. The inventory will be filled based on the provider.
   - On each disk, try to gather `lun_storage` data.
2. Validation rules will be executed.
3. Two steps will be added to the migration:
   - `CreatePersistentVolumes`, which will create the relevant `PV` based on the `LUN` details.
   - `CreatePersistentVolumeClaims`, which will create the relevant `PVC` and **bound** to the above `PV`.
     - Might be possible to skip and set into the VM specification with all the relevant `PVC` details.
On each step, it first will check the existence of the CRs, when they won't be available, it will be created.
4. When building the VM specifications, create the `LUN` disk and volume associated with the `PVC`.
5. Once the migration ends successfully, detach the `LUN` disk from the source VM.
   - Cold migration: The LUN disk will be detached before initiating the copy disk.
   - Warm migration: Once the migration ends and we shut down the VM on source.
6. If the migration failed and it archieved, remove any new CR that was created (PV/PVC).
   - In cold migration, re-attach the LUN on the source VM.

Future
----------------------------------------

- Support SCSI persistent reservation once [kubevirt PR](https://github.com/kubevirt/kubevirt/pull/8210) will be merged.
