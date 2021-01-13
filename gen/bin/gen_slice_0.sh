#!/bin/bash

T0=$1
PACKAGE=$2
DESTDIR="../generated/${PACKAGE}"

echo $PWD
#pushd gen/bin

mkdir -p "${DESTDIR}"

cat ../../pkg/collections/slice_0.go | \
    sed -e 's/T0/'${T0}'/g; s/package collections/package '${PACKAGE}'/g' > \
    "${DESTDIR}/slice_0_${T0}.go"

#popd
