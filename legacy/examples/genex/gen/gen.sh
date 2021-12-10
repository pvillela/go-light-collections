#!/bin/bash

#
#  Copyright © 2021 Paulo Villela. All rights reserved.
#  Use of this source code is governed by the Apache 2.0 license
#  that can be found in the LICENSE file.
#

export COLL_MODULE_DIR=../../..
export DEST_DIR=../generated
export PACKAGE=coll

# Types
./copy_aliases.sh

# Slices

./copy_sub.sh slice_0 Person
./copy_sub.sh slice_0 String
./copy_sub.sh slice_0 int
./copy_sub.sh slice_01 Person String
./copy_sub.sh slice_01 Person int
./copy_sub.sh slice_01 String int
./copy_sub.sh slice_of_pair_01 String int

# Maps

./copy_sub.sh map_01 String int
./copy_sub.sh map_012 String int int
./copy_sub.sh map_012x String int String

# Nested collections
./copy_sub.sh slice_0 MapStringint
./copy_sub.sh slice_01 MapStringint Person
./copy_sub.sh slice_0 PairMpStringint
./copy_sub.sh slice_01 PairMpStringint Person
