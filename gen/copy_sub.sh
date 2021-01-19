#!/bin/bash

FILE=$1
T0=$2
T1=$3
T2=$4

echo "PWD: $PWD"

mkdir -p "${DEST_DIR}/${PACKAGE}"

cat ${COLL_MODULE_DIR}/pkg/glc/${FILE}.go | \
    sed -e 's/package glc/package '${PACKAGE}'/g' \
        -e 's/T0/'${T0}'/g' \
        -e 's/T1/'${T1}'/g' \
        -e 's/T2/'${T2}'/g' \
        -e '1 i // Code generated -- DO NOT EDIT.\n' \
    > "${DEST_DIR}/${PACKAGE}/${T0}${T1}${T2}_${FILE}.go"
