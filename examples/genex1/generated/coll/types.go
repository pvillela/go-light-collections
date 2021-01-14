package coll

import "github.com/pvillela/go-light-collections/examples/genex1/somepkg"

// Foo is a type alias.
type Foo = somepkg.Foo

// SliceFoo is a type wrapper.
type SliceFoo []Foo

// Sliceint is a type wrapper.
type Sliceint []int

// Slicestring is a type alias.
type Slicestring = somepkg.Slicestring

// MapintSliceFoo is a type wrapper.
type MapintSliceFoo map[int]SliceFoo

// MapstringSliceint is a type wrapper.
type MapstringSliceint map[string]Sliceint
