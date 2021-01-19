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

// MapintSlicePerson is a type wrapper.
type MapintSlicePerson map[int]SlicePerson

// MapintSetPerson is a type wrapper.
type MapintSetPerson map[int]SetPerson

// SetOfPairPersonint is a type wrapper, implements Set interface.
type SetOfPairPersonint map[PairPersonint]bool
