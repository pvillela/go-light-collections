// Code generated -- DO NOT EDIT.

package coll

// PairintString defines a pseudo-generic pair.
type PairintString struct {
	X1 int
	X2 String
}

// SliceOfPairintString is a type wrapper.
type SliceOfPairintString []PairintString

// MapintString is a type wrapper, implements Map interface.
type MapintString map[int]String

// MapintSliceString is a type wrapper.
type MapintSliceString map[int]SliceString

// MapintSetString is a type wrapper.
type MapintSetString map[int]SetString

// SetOfPairintString is a type wrapper, implements Set interface.
type SetOfPairintString map[PairintString]bool
