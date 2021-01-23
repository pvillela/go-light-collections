// Code generated -- DO NOT EDIT.

package coll

// PairPersonint defines a pseudo-generic pair.
type PairPersonint struct {
	X1 Person
	X2 int
}

// MapPersonint is a type wrapper, implements IMap interfaces.
type MapPersonint map[Person]int

// SliceOfPairPersonint is a type wrapper.
type SliceOfPairPersonint []PairPersonint

// SetOfPairPersonint is a type wrapper, implements ISetOfPairPersonint interface.
type SetOfPairPersonint map[PairPersonint]bool
