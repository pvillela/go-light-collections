#!/bin/bash

T0=$1
T1=$2
PACKAGE=$3
DESTDIR=$4

echo $PWD
#pushd gen/bin

mkdir -p "${DESTDIR}"

cat ${COLL_MODULE_DIR}/pkg/collections/slice_1.go | \
    sed -e 's/T0/'${T0}'/g; s/package collections/package '${PACKAGE}'/g' \
        -e 's/T1/'${T1}'/g; s/package collections/package '${PACKAGE}'/g' \
        -e '1 i // Code generated -- DO NOT EDIT.\n' \
    > "${DESTDIR}/slice_1_${T0}.go"

#popd
