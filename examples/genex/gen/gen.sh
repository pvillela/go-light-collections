#!/bin/bash

export COLL_MODULE_DIR=../../..
export DEST_DIR=../generated
export PACKAGE=coll

./copy_aliases.sh

# Types

./copy_sub.sh types_0 Person  # required for slice_0 Person
./copy_sub.sh types_0 int  # required for slice_01 Person int and slice_01 String int
./copy_sub.sh types_0 String  # required for slice_01 Person String and slice_01 String int

./copy_sub.sh types_01 Person String  # required for slice_01 Person String
./copy_sub.sh types_01 String Person  # required for slice_01 Person String
./copy_sub.sh types_01 Person int  # required for slice_01 Person int
./copy_sub.sh types_01 int Person  # required for slice_01 Person int
./copy_sub.sh types_01 String int  # required for slice_01 String int and map_012 String int int
./copy_sub.sh types_01 int String  # required for slice_01 String int, map_01 String int, and map_012 String int int
./copy_sub.sh types_01 int int  # required for map_012 String int int

# Slices

./copy_sub.sh slice_0 Person

./copy_sub.sh slice_01 Person String
./copy_sub.sh slice_01 Person int
./copy_sub.sh slice_01 String int

# Maps

./copy_sub.sh map_01 String int
./copy_sub.sh map_012 String int int
