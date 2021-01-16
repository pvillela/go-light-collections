#!/bin/bash

export PACKAGE=coll
export DEST_DIR=../generated/coll

./copy_sub.sh types_0.go Foo
./copy_sub.sh slice_0.go Foo
./copy_sub.sh types_01.go Foo Int
./copy_sub.sh slice_01.go Foo Int

./copy_sub.sh types_0.go Int
./copy_sub.sh slice_0.go Int
./copy_sub.sh types_01.go Int string
./copy_sub.sh slice_01.go Int string
