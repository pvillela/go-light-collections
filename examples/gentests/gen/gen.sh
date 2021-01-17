#!/bin/bash

export COLL_MODULE_DIR=../../..
export DEST_DIR=../generated
export PACKAGE=collections

./copy_aliases.sh

./copy_sub.sh data_0_test Dat
./copy_sub.sh types_0 Dat
./copy_sub.sh slice_0 Dat
./copy_test.sh slice_0_test Dat
