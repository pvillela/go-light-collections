#!/bin/bash

export COLL_MODULE_DIR=../../..
export DEST_DIR=../generated
export PACKAGE=coll

./copy_aliases.sh

# Types

./copy_sub.sh types_0 Person
./copy_sub.sh types_0 int
./copy_sub.sh types_0 String

./copy_sub.sh types_01 Person Person
./copy_sub.sh types_01 Person int
./copy_sub.sh types_01 Person String
./copy_sub.sh types_01 int String
./copy_sub.sh types_01 String int

# ./copy_sub.sh types_012 Person int String

# Slices

./copy_sub.sh slice_0 Person
./copy_sub.sh slice_0 int
./copy_sub.sh slice_0 String

./copy_sub.sh slice_01 Person int
./copy_sub.sh slice_01 Person String
./copy_sub.sh slice_01 Person Person
./copy_sub.sh slice_01 int String
./copy_sub.sh slice_01 String int

# Maps

# ./copy_sub.sh map_01 Person int
# ./copy_sub.sh map_012 Person int String

# Sets