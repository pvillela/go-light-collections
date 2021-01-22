// Code generated -- DO NOT EDIT.

package coll

// PairStringPerson defines a pseudo-generic pair.
type PairStringPerson struct {
	X1 String
	X2 Person
}

// SliceOfPairStringPerson is a type wrapper.
type SliceOfPairStringPerson []PairStringPerson

// MapStringPerson is a type wrapper, implements Map interface.
type MapStringPerson map[String]Person

// MapStringSlicePerson is a type wrapper.
type MapStringSlicePerson map[String]SlicePerson

// MapStringSetPerson is a type wrapper.
type MapStringSetPerson map[String]SetPerson

// SetOfPairStringPerson is a type wrapper, implements Set interface.
type SetOfPairStringPerson map[PairStringPerson]bool
