#!/bin/bash

export COLL_MODULE_DIR=../../..
export DEST_DIR=../generated
export PACKAGE=collections

./copy_aliases.sh

# Slice tests

./copy_sub.sh slice_0a Dat
./copy_sub.sh slice_0a_test Dat
./copy_sub.sh slice_0b Dat
./copy_sub.sh slice_0b_test Dat
./copy_sub.sh slice_01a Dat int
./copy_sub.sh slice_01a_test Dat int
./copy_sub.sh slice_01b Dat int
./copy_sub.sh slice_01b_test Dat int

# Map tests
./copy_sub.sh map_01a int string
./copy_sub.sh map_01a_test int string
./copy_sub.sh map_01b int string
./copy_sub.sh map_01b_test int string
./copy_sub.sh map_012a int string int
./copy_sub.sh map_012a_test int string int
./copy_sub.sh map_012b int string int
./copy_sub.sh map_012b_test int string int

# Set tests
./copy_sub.sh set_0 int
./copy_sub.sh set_0_test int
./copy_sub.sh set_01 int string
./copy_sub.sh set_01_test int string
