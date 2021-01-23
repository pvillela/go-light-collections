// Code generated -- DO NOT EDIT.

package coll

// PairPersonString defines a pseudo-generic pair.
type PairPersonString struct {
	X1 Person
	X2 String
}

// MapPersonString is a type wrapper, implements IMap interfaces.
type MapPersonString map[Person]String

// SliceOfPairPersonString is a type wrapper.
type SliceOfPairPersonString []PairPersonString

// SetOfPairPersonString is a type wrapper, implements ISetOfPairPersonString interface.
type SetOfPairPersonString map[PairPersonString]bool
