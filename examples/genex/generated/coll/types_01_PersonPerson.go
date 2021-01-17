// Code generated -- DO NOT EDIT.

package coll

// PairPersonPerson defines a pseudo-generic pair.
type PairPersonPerson struct {
	X1 Person
	X2 Person
}

// SliceOfPairPersonPerson is a type wrapper.
type SliceOfPairPersonPerson []PairPersonPerson

// MapPersonPerson is a type wrapper, implements Map interface.
type MapPersonPerson map[Person]Person

// MapPersonSlicePerson is a type wrapper.
type MapPersonSlicePerson map[Person]SlicePerson

// SetOfPairPersonPerson is a type wrapper, implements Set interface.
type SetOfPairPersonPerson map[PairPersonPerson]bool
