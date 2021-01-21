// Code generated -- DO NOT EDIT.

package coll

// PairPersonint defines a pseudo-generic pair.
type PairPersonint struct {
	X1 Person
	X2 int
}

// SliceOfPairPersonint is a type wrapper.
type SliceOfPairPersonint []PairPersonint

// MapPersonint is a type wrapper, implements Map interface.
type MapPersonint map[Person]int

// MapPersonSliceint is a type wrapper.
type MapPersonSliceint map[Person]Sliceint

// MapPersonSetint is a type wrapper.
type MapPersonSetint map[Person]Setint

// SetOfPairPersonint is a type wrapper, implements Set interface.
type SetOfPairPersonint map[PairPersonint]bool
