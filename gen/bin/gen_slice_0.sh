#!/bin/bash

T0=$1
PACKAGE=$2
DEST_DIR=$3

echo $PWD
#pushd gen/bin

mkdir -p "${DEST_DIR}"

cat ${COLL_MODULE_DIR}/pkg/collections/slice_0.go | \
    sed -e 's/T0/'${T0}'/g; s/package collections/package '${PACKAGE}'/g' \
        -e '1 i // Code generated -- DO NOT EDIT.\n' \
    > "${DEST_DIR}/slice_0_${T0}.go"

#popd
