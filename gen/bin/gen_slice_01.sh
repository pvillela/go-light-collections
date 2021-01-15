#!/bin/bash

T0=$1
T1=$2
PACKAGE=$3
DEST_DIR=$4

echo "PWD: $PWD"

mkdir -p "${DEST_DIR}"

cat ${COLL_MODULE_DIR}/pkg/collections/slice_01.go | \
    sed -e 's/T0/'${T0}'/g; s/package collections/package '${PACKAGE}'/g' \
        -e 's/T1/'${T1}'/g; s/package collections/package '${PACKAGE}'/g' \
        -e '1 i // Code generated -- DO NOT EDIT.\n' \
    > "${DEST_DIR}/slice_01_${T0}.go"
