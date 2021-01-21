#!/bin/bash

export COLL_MODULE_DIR=../../..
export DEST_DIR=../generated
export PACKAGE=collections

./copy_aliases.sh

# Types
./copy_notest.sh prelim_0 Dat
./copy_sub.sh types_0 Dat
./copy_sub.sh types_0 int
./copy_sub.sh types_0 string
./copy_sub.sh types_01 Dat int
./copy_sub.sh types_01 int Dat
./copy_sub.sh types_01 int string
./copy_sub.sh types_01 int int
./copy_sub.sh types_01 string int

# Slice tests
./copy_sub.sh slice_0 Dat
./copy_sub.sh slice_0_test Dat
./copy_sub.sh slice_01 Dat int
./copy_sub.sh slice_01_test Dat int

# Map tests
./copy_sub.sh map_01 int string
./copy_sub.sh map_01_test int string
./copy_sub.sh map_012 int string int
./copy_sub.sh map_012_test int string int

# Set tests
./copy_sub.sh set_0 int
./copy_sub.sh set_0_test int
./copy_sub.sh set_01 int string
./copy_sub.sh set_01_test int string
