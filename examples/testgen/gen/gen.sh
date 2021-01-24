#!/bin/bash

export COLL_MODULE_DIR=../../..
export DEST_DIR=../generated
export PACKAGE=collections

# Types
./copy_aliases.sh
./copy_notest.sh prelim_0 Dat

# Slice tests
./copy_sub.sh slice_0_test Dat
./copy_sub.sh slice_0x_test Dat
./copy_sub.sh slice_01_test Dat int
./copy_sub.sh slice_01x_test Dat int
./copy_sub.sh slice_of_pair_01_test Dat int

# Slice test dependencies
./copy_sub.sh slice_0 Dat
./copy_sub.sh slice_0x Dat
./copy_sub.sh slice_01 Dat int
./copy_sub.sh slice_01x Dat int
./copy_sub.sh slice_of_pair_01 Dat int

# Map tests
./copy_sub.sh map_01_test int string
./copy_sub.sh map_012_test int string int
./copy_sub.sh map_012x_test int string int

# Dependencies for Map tests
./copy_sub.sh map_01 int string
./copy_sub.sh map_012 int string int
./copy_sub.sh map_012x int string int

# Set tests
./copy_sub.sh set_0_test int
./copy_sub.sh set_01_test int string
./copy_sub.sh set_of_pair_01_test int string

# Additional dependencies for Set tests
./copy_sub.sh set_0 int
./copy_sub.sh set_01 int string
./copy_sub.sh set_of_pair_01 int string
