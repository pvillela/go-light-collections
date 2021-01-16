// Code generated -- DO NOT EDIT.

package coll

// PairFooint defines a pseudo-generic pair.
type PairFooint struct {
	X1 Foo
	X2 int
}

// SliceOfPairFooint is a type wrapper.
type SliceOfPairFooint []PairFooint

// MapFooint is a type wrapper, implements Map interface.
type MapFooint map[Foo]int

// MapintSliceFoo is a type wrapper.
type MapintSliceFoo map[int]SliceFoo

// SetOfPairFooint is a type wrapper, implements Set interface.
type SetOfPairFooint map[PairFooint]bool
