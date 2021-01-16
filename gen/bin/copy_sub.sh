#!/bin/bash

FILE=$1
T0=$2
T1=$3

echo "PWD: $PWD"

mkdir -p "${DEST_DIR}"

cat ${COLL_MODULE_DIR}/pkg/collections/${FILE} | \
    sed -e 's/T0/'${T0}'/g; s/package collections/package '${PACKAGE}'/g' \
        -e 's/T1/'${T1}'/g; s/package collections/package '${PACKAGE}'/g' \
        -e '1 i // Code generated -- DO NOT EDIT.\n' \
    > "${DEST_DIR}/${FILE}_${T0}${T1}.go"
