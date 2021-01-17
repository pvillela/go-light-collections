// Code generated -- DO NOT EDIT.

package coll

// PairPersonString defines a pseudo-generic pair.
type PairPersonString struct {
	X1 Person
	X2 String
}

// SliceOfPairPersonString is a type wrapper.
type SliceOfPairPersonString []PairPersonString

// MapPersonString is a type wrapper, implements Map interface.
type MapPersonString map[Person]String

// MapStringSlicePerson is a type wrapper.
type MapStringSlicePerson map[String]SlicePerson

// SetOfPairPersonString is a type wrapper, implements Set interface.
type SetOfPairPersonString map[PairPersonString]bool
