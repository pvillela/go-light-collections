package coll

import "github.com/pvillela/go-light-collections/examples/genex1/otherdomain"

// Foo is an example data structure.
type Foo struct {
	v1 int
	v2 string
}

// SliceFoo is a type wrapper.
type SliceFoo []Foo

// Sliceint is a type wrapper.
type Sliceint []int

// Slicestring is a type alias.
type Slicestring = otherdomain.Slicestring

// MapintSliceFoo is a type wrapper.
type MapintSliceFoo map[int]SliceFoo

// MapstringSliceint is a type wrapper.
type MapstringSliceint map[string]Sliceint
