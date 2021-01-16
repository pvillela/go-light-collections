#!/bin/bash

export COLL_MODULE_DIR=../../..
export DEST_DIR=../generated
export PACKAGE=coll

./copy_aliases.sh

./copy_sub.sh types_0 Foo
./copy_sub.sh slice_0 Foo
./copy_sub.sh types_01 Foo int
./copy_sub.sh slice_01 Foo int
./copy_sub.sh types_01 Foo String
./copy_sub.sh slice_01 Foo String
./copy_sub.sh types_01 Foo Foo
./copy_sub.sh slice_01 Foo Foo

./copy_sub.sh types_0 int
./copy_sub.sh slice_0 int
./copy_sub.sh types_01 int String
./copy_sub.sh slice_01 int String

./copy_sub.sh types_0 String
./copy_sub.sh slice_0 String
./copy_sub.sh types_01 String int
./copy_sub.sh slice_01 String int
