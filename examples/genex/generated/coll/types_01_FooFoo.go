// Code generated -- DO NOT EDIT.

package coll

// PairFooFoo defines a pseudo-generic pair.
type PairFooFoo struct {
	X1 Foo
	X2 Foo
}

// SliceOfPairFooFoo is a type wrapper.
type SliceOfPairFooFoo []PairFooFoo

// MapFooFoo is a type wrapper, implements Map interface.
type MapFooFoo map[Foo]Foo

// MapFooSliceFoo is a type wrapper.
type MapFooSliceFoo map[Foo]SliceFoo

// SetOfPairFooFoo is a type wrapper, implements Set interface.
type SetOfPairFooFoo map[PairFooFoo]bool
