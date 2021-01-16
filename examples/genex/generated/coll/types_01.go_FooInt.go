// Code generated -- DO NOT EDIT.

package coll

// PairFooInt defines a pseudo-generic pair.
type PairFooInt struct {
	X1 Foo
	X2 Int
}

// SliceOfPairFooInt is a type wrapper.
type SliceOfPairFooInt []PairFooInt

// MapFooInt is a type wrapper, implements Map interface.
type MapFooInt map[Foo]Int

// MapIntSliceFoo is a type wrapper.
type MapIntSliceFoo map[Int]SliceFoo

// SetOfPairFooInt is a type wrapper, implements Set interface.
type SetOfPairFooInt map[PairFooInt]bool
