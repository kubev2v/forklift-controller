# Development docs

## Used technologies

Kubernetes operator and controller framework.

Forklift Operator uses ansible and its main task is to install the Forklift into the Kubernetes or Openshift cluster

Forklift Controller is written in go-lang. Its task is to run/orchestate/drive all operations related to the VM migrations. An entrypoint is a reconcile loop which watch events from the cluster and handles received events.

Forklift UI is a React frontend application which communicates with Forklift (an related OpenShif components) via API and provides user interface.

Forklift must-gather provides an image for openshift adm must-gather tool which is used for gatering logs and related debugging information relevant for Forklift

## Related repositories

- forklift-operator
- forklift-controller
- forklift-ui
- forklift-must-gather
and others

## Local development

Controller