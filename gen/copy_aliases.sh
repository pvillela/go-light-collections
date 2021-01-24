#!/bin/bash

FULL_DEST_PATH="$PWD/${DEST_DIR}/${PACKAGE}/aliases.go"

echo "${FULL_DEST_PATH}"
# echo "PWD: $PWD"

mkdir -p "${DEST_DIR}/${PACKAGE}"

cat ./aliases.go | \
    sed -e 's/package .*/package '${PACKAGE}'/g' \
    > "${FULL_DEST_PATH}"
