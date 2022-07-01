# Development docs

This page should provide basic technical information about a Konveyor Forklift project helpful for engineering onboarding.

## Used technologies

The Konveyor Forklift project is an application based on the Kubernetes/OpenShift operator and controller framework.

### Operator
Forklift Operator uses ansible and its main task is to install the Forklift into the Kubernetes or Openshift cluster.

Repo: https://github.com/konveyor/forklift-operator

### Controller
Forklift Controller is written in golang. Its task is to run/orchestate/drive all operations related to the VM migrations. An entrypoint is a reconcile loop which watches events from the cluster and handles received events.

E.g. The VM conversion itself is performed by ```virt-v2v``` executed in a pod (a "conversion pod") created by Forklift.

There is also an Inventory service which explores objects in source and destination clusters and provides such information to the UI, so users can choose from available clusters/VMs/networks/storage.

Repo: https://github.com/konveyor/forklift-controller
Migration pipeline: https://github.com/konveyor/forklift-controller/blob/main/pkg/controller/plan/migration.go

### UI
Forklift UI is a React frontend application which communicates with Forklift (and related OpenShift components) via API and provides user interface.

Repo: https://github.com/konveyor/forklift-ui

### Other components
There are other components used by the Forklift like a validation service, must-gather, etc.

Forklift must-gather provides an image for OpenShift must-gather (oc adm must-gather command) tool which is used for gathering logs and related debugging information relevant for Forklift.

- https://github.com/konveyor/forklift-must-gather
- https://github.com/konveyor/forklift-must-gather-api

- https://github.com/konveyor/forklift-documentation

See all available repositories at https://github.com/konveyor?q=forklift#org-repositories

### OpenShift Virtualization/Kubevirt related projects

Forklift relies on several OpenShift/Kubevirt projects. Communication between such components and Forklift is via their CRs (created by Forklift).

## Basic workflow from user point of view

In order to run a migration, a source and a destination providers need to exist. Then there is a migration Plan which defines from what VMware/RHV cluster to what Kubevirt/OpenShift Virtualization cluster will the migration run. The plan also lists VMs which will be migrated and specifies mapping of networks and storage between source and destination clusters.

Optionally, there are hooks (before and after migration) which can execute an ansible playbook (provided by user).

When the migration Plan is started, a Migration CRs are created. Each CR represents a VM migration and tracks its progress (and updates the Plan status).

All such objects (like plan, migration, virtualmachine CRs) are available also in OpenShift CLI, try e.g. ```$ oc get plans -n openshift-mtv -o yaml```

## Forklift-controller development

Forklift-controller is a golang application. It should be placed within the ```$GOPATH``` directory to allow smooth dependencies loading (using ```go mod```).

```
# Get the source code with go get or following:
cd $GOPATH/src/github.com/konveyor
git clone https://github.com/konveyor/forklift-controller
```

For forklift-controller execution, an OpenShift/Kubernetes cluster (or their "lightweight" versions like https://github.com/code-ready/crc or https://kubevirt.io/quickstart_minikube/) need to be running and have OpenShift Virtualization / Kubevirt Operators installed.

### Run the controller locally

Before running the controller locally, scale down the forklift-controller deployment in your cluster to zero. Then use ```oc login``` to create a connection (KUBECONFIG) to your cluster in the terminal and finally execute ```make run```.

The controller (in cluster or locally running) listens to events from openshift (its reconcile loop watches events in the cluster and process it, e.g. https://github.com/konveyor/forklift-controller/blob/main/pkg/controller/plan/controller.go#L180). When the controller in the cluster is off, the local is up with ```make run```, it will reconcile events/resources from the cluster.

To interact with the controller, use "normal" oc commands locally, like oc create -f migplan.yaml (or similar), or access Forklift UI running in the cluster directly. The local controller will handle its events.

### Run a custom controller image in cluster

As the Openshift/Kubernetes applications consist of a bunch of images, the forklift-controller image can be built locally, pushed to the image registry and Forklift installation in the cluster could be dated to use the image.

```
# Create the image:
export IMG=quay.io/my-username/forklift-controller:my-dev-tag
make docker build
make docker push
```

Then update forklift-operator YAML in your cluster to use your image as the ```CONTROLLER_IMAGE``` (part of controller deployment env variables, not "related_images" section).
