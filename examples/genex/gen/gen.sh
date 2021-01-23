#!/bin/bash

export COLL_MODULE_DIR=../../..
export DEST_DIR=../generated
export PACKAGE=coll

./copy_aliases.sh

# Types

./copy_sub.sh types_0 Person
./copy_sub.sh types_0 int
./copy_sub.sh types_0 String

./copy_sub.sh types_01 Person String
./copy_sub.sh types_01 Person int
./copy_sub.sh types_01 String int

# Slices

./copy_sub.sh slice_0 Person  # requires types_0 Person

./copy_sub.sh slice_01 Person String  # requires types_0 Person and String; types_01 Person String
./copy_sub.sh slice_01 Person int  # requires types_0 Person and int; types_01 Person int
./copy_sub.sh slice_01 String int  # requires types_0 String and int; types_01 String int

# Maps

./copy_sub.sh map_01 String int  # requires types_0 String and int; types_01 String int
./copy_sub.sh map_012 String int int  # requires types_0 String, int, and int; types_01 String int

# Nested collection
./copy_sub.sh types_01 int SlicePerson
./copy_sub.sh map_01 int SlicePerson
