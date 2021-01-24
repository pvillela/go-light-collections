#!/bin/bash

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

# Maps

./copy_sub.sh map_01 String int
./copy_sub.sh map_012 String int int
