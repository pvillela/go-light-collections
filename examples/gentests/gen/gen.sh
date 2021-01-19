#!/bin/bash

export COLL_MODULE_DIR=../../..
export DEST_DIR=../generated
export PACKAGE=collections

./copy_aliases.sh

# Slice tests
./copy_notest.sh prelim_0 Dat
./copy_sub.sh types_0 Dat
./copy_sub.sh slice_0 Dat
./copy_sub.sh slice_0_test Dat
./copy_sub.sh types_01 Dat int
./copy_sub.sh slice_01 Dat int
./copy_sub.sh slice_01_test Dat int

# Map tests
./copy_sub.sh types_0 int
./copy_sub.sh types_01 int string
# ./copy_sub.sh map_01_test int string
