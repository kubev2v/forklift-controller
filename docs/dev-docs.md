# Development docs

This page should provide basic technical information about the Forklift project helpful for engineering onboarding.

## Used technologies

Konveyor Forklift project is based on Kubernetes/OpenShift operator and controller framework.

### Operator
Forklift Operator uses ansible and its main task is to install the Forklift into the Kubernetes or Openshift cluster.

Repo: https://github.com/konveyor/forklift-operator

### Controller
Forklift Controller is written in go-lang. Its task is to run/orchestate/drive all operations related to the VM migrations. An entrypoint is a reconcile loop which watch events from the cluster and handles received events.

E.g. The VM conversion itself is performed by ```virt-v2v``` executed in a pod (a "conversion pod") created by Forklift.

There is also an Inventory service which explores objects in source and destination cluster and provides such information to the UI, so user can choose from available clusters/VMs/networks/storage.

Repo: https://github.com/konveyor/forklift-controller

### UI
Forklift UI is a React frontend application which communicates with Forklift (an related OpenShif components) via API and provides user interface.

Repo: https://github.com/konveyor/forklift-ui

### Other components
Forklift must-gather provides an image for openshift adm must-gather tool which is used for gatering logs and related debugging information relevant for Forklift

- https://github.com/konveyor/forklift-must-gather
- https://github.com/konveyor/forklift-must-gather-api

- https://github.com/konveyor/forklift-documentation

See all available repos at https://github.com/konveyor?q=forklift#org-repositories

### OpenShift Virtualization/Kubevirt related projects

Forklift relies on several OpenShift/Kubevirt projects. Communication between such components and Forklift is via their CRs (created by Forklift).

## Basic workflow from user point of view

In order to run a migration, a source and destination providers need to exist. Then there is a migration Plan which defines from what VMware/RHV cluster to what Kubevirt/OpenShift Virt will the migration will run, lists VMs which will be migrated and specifies mapping of networks and storage between source and destination clusters.

Optionally, there are also hooks (before and after migration) which can execute an ansible playbook (provided by user).

When the migration Plan is started, a Migration CR is created and that represents the migration and tracks it progress (and updates the Plan status).

All such objects (like plan, migration, virtualmachine CRs) are available also in OpenShift CLI, try e.g. ```$ oc get plans -n openshift-mtv -o yaml```

## Forklift-controller development

Forklift-controller is a go-lang based application. It should be placed under ```$GOPATH``` directory to allow smooth dependencies loading (using ```go mod```).

```
# Get the source code with go get or following:
cd $GOPATH/src/github.com/konveyor
git clone https://github.com/konveyor/forklift-controller
```

For forklift-controller execution, an OpenShift/Kubernetes cluster (or their "lightweigth" versions like https://github.com/code-ready/crc or https://kubevirt.io/quickstart_minikube/) need to be running and have OpenShift Virtualization / Kubevirt Operators installed.

### Run code the locally

Before running the controller locally, scale down the forklift-controller deployment in your cluster.

Then use ```oc login``` to create connection (KUBECONFIG) to your cluster in terminal and finally execute ```make run```.

### Run the code in cluster

As the Openshift/Kubernetes applications consist of bunch of images, the forklift-controller image can be built locally, pushed to the image registry and Forklift installation in the cluster could be dated to use such image.

```
# Create the image:
export IMG=quay.io/my-username/forklift-controller:my-dev-tag
make docker build
make docker push
```

Then update forklift-operator YAML in your cluster to use your image as the ```CONTROLLER_IMAGE```.
