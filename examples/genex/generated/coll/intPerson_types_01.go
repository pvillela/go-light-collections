// Code generated -- DO NOT EDIT.

package coll

// PairintPerson defines a pseudo-generic pair.
type PairintPerson struct {
	X1 int
	X2 Person
}

// SliceOfPairintPerson is a type wrapper.
type SliceOfPairintPerson []PairintPerson

// MapintPerson is a type wrapper, implements Map interface.
type MapintPerson map[int]Person

// MapintSlicePerson is a type wrapper.
type MapintSlicePerson map[int]SlicePerson

// MapintSetPerson is a type wrapper.
type MapintSetPerson map[int]SetPerson

// SetOfPairintPerson is a type wrapper, implements Set interface.
type SetOfPairintPerson map[PairintPerson]bool
