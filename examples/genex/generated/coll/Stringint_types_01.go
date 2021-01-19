// Code generated -- DO NOT EDIT.

package coll

// PairStringint defines a pseudo-generic pair.
type PairStringint struct {
	X1 String
	X2 int
}

// SliceOfPairStringint is a type wrapper.
type SliceOfPairStringint []PairStringint

// MapStringint is a type wrapper, implements Map interface.
type MapStringint map[String]int

// MapintSliceString is a type wrapper.
type MapintSliceString map[int]SliceString

// MapintSetString is a type wrapper.
type MapintSetString map[int]SetString

// SetOfPairStringint is a type wrapper, implements Set interface.
type SetOfPairStringint map[PairStringint]bool
