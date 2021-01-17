#!/bin/bash

echo "PWD: $PWD"

mkdir -p "${DEST_DIR}/${PACKAGE}"

cat ./aliases.go | \
    sed -e 's/package .*/package '${PACKAGE}'/g' \
    > "${DEST_DIR}/${PACKAGE}/aliases.go"
