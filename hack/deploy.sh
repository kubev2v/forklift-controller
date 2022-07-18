#!/bin/bash
set -e

function help() {
   echo "Deploy a custom image of the forklift-controller to a custom registry."
   echo
   echo "Syntax: deploy.sh [-h|n|i|r]"
   echo "options:"
   echo "h     Print this help"
   echo "n     Namespace to use (Default: konveyor-forklift)"
   echo "i     Image name       (Default: forklift-controller:latest)"
   echo "r     Registry         (Default: uses default openshift route)"
   echo
}

while getopts n:i:r:h flag
do
    case "${flag}" in
	n) namespace=${OPTARG};;
        i) image=${OPTARG};;
	r) registry=${OPTARG};;
	h) help
	   exit;;
    esac
done

NAMESPACE=${namespace:-konveyor-forklift}
IMG=${image:-forklift-controller:latest}
REGISTRY=${registry:-"$(oc get route/default-route -n openshift-image-registry -o=jsonpath='{.spec.host}')/openshift"}

podman login --tls-verify=false -u unused -p $(oc whoami -t) ${REGISTRY}
podman build -t ${REGISTRY}/${IMG} .
podman push --tls-verify=false ${REGISTRY}/${IMG}

CSV=$(oc get csv -n konveyor-forklift -o custom-columns=":metadata.name" --no-headers=true)
oc get csv "${CSV}" -n "${NAMESPACE}" -o json | \
        jq ".spec.install.spec.deployments[0].spec.template.spec.containers[0].env[1].value = \"${REGISTRY}/${IMG}\"" | \
        oc replace -f -
