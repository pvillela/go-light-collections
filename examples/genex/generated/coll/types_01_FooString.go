// Code generated -- DO NOT EDIT.

package coll

// PairFooString defines a pseudo-generic pair.
type PairFooString struct {
	X1 Foo
	X2 String
}

// SliceOfPairFooString is a type wrapper.
type SliceOfPairFooString []PairFooString

// MapFooString is a type wrapper, implements Map interface.
type MapFooString map[Foo]String

// MapStringSliceFoo is a type wrapper.
type MapStringSliceFoo map[String]SliceFoo

// SetOfPairFooString is a type wrapper, implements Set interface.
type SetOfPairFooString map[PairFooString]bool
