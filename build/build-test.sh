#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

GO_FLAGS=
if [[ "${VERBOSE:-0}" = "1" ]]; then
  GO_FLAGS="-v"
  set -o xtrace
fi

if [ -z "${PKG}" ]; then
    echo "PKG must be set"
    exit 1
fi

export CGO_ENABLED=0
export GO111MODULE=on

cd /data/go/src/${PKG}/test/e2e
go test -c ${GO_FLAGS} -o /out/e2e.test .
