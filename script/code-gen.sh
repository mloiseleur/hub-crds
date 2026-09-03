#!/usr/bin/env bash
# shellcheck disable=SC2046

set -euo pipefail

PROJECT_MODULE="github.com/traefik/hub-crds"
KUBE_VERSION=v0.37.0
CONTROLLER_GEN_VERSION=v0.21.0
CURRENT_DIR="$(pwd)"
BOILERPLATE="$(dirname "${BASH_SOURCE[0]}")/boilerplate.go.tmpl"

# Populate the module cache so kube_codegen.sh can be sourced from it.
# The other generators are installed by kube_codegen.sh itself.
go install "k8s.io/code-generator/cmd/deepcopy-gen@${KUBE_VERSION}"
go install "sigs.k8s.io/controller-tools/cmd/controller-gen@${CONTROLLER_GEN_VERSION}"

CODEGEN_PKG="$(go env GOPATH)/pkg/mod/k8s.io/code-generator@${KUBE_VERSION}"
source "${CODEGEN_PKG}/kube_codegen.sh"

echo "Generating deepcopy code ..."
kube::codegen::gen_helpers \
    --boilerplate "${BOILERPLATE}" \
    "${CURRENT_DIR}/pkg/apis"

echo "Generating Hub clientset, listers and informers code ..."
kube::codegen::gen_client \
    --with-watch \
    --output-dir "${CURRENT_DIR}/pkg/client" \
    --output-pkg "${PROJECT_MODULE}/pkg/client" \
    --boilerplate "${BOILERPLATE}" \
    "${CURRENT_DIR}/pkg/apis"

echo "Generating the CRD definitions ..."
"$(go env GOBIN 2>/dev/null || echo "$(go env GOPATH)/bin")/controller-gen" \
    crd:crdVersions=v1,allowDangerousTypes=true \
    paths=./pkg/apis/hub/v1alpha1/... \
    output:dir=./pkg/apis/hub/v1alpha1/crd
