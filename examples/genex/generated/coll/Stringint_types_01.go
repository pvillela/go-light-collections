// Code generated -- DO NOT EDIT.

package coll

// PairStringint defines a pseudo-generic pair.
type PairStringint struct {
	X1 String
	X2 int
}

// MapStringint is a type wrapper, implements IMap interfaces.
type MapStringint map[String]int

// SliceOfPairStringint is a type wrapper.
type SliceOfPairStringint []PairStringint

// SetOfPairStringint is a type wrapper, implements ISetOfPairStringint interface.
type SetOfPairStringint map[PairStringint]bool
