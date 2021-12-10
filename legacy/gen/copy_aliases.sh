#!/bin/bash

#
#  Copyright © 2021 Paulo Villela. All rights reserved.
#  Use of this source code is governed by the Apache 2.0 license
#  that can be found in the LICENSE file.
#

FULL_DEST_PATH="$PWD/${DEST_DIR}/${PACKAGE}/aliases.go"

echo "${FULL_DEST_PATH}"
# echo "PWD: $PWD"

mkdir -p "${DEST_DIR}/${PACKAGE}"

cat ./aliases.go | \
    sed -e 's/package .*/package '${PACKAGE}'/g' \
    > "${FULL_DEST_PATH}"
