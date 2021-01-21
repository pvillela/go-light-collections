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

// MapPersonSliceString is a type wrapper.
type MapPersonSliceString map[Person]SliceString

// MapPersonSetString is a type wrapper.
type MapPersonSetString map[Person]SetString

// SetOfPairPersonString is a type wrapper, implements Set interface.
type SetOfPairPersonString map[PairPersonString]bool
