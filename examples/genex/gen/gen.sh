#!/bin/bash

export COLL_MODULE_DIR=../../..
export DEST_DIR=../generated
export PACKAGE=coll

./copy_aliases.sh

# Slices

./copy_sub.sh slice_0a Person
./copy_sub.sh slice_01a Person String
./copy_sub.sh slice_01a Person int
./copy_sub.sh slice_0a String
./copy_sub.sh slice_01a String int

# Maps

./copy_sub.sh map_01a String int
./copy_sub.sh map_012a String int int
./copy_sub.sh map_012b String int int

# Nested collection
./copy_sub.sh slice_0a MapStringint
./copy_sub.sh slice_01a MapStringint Person
