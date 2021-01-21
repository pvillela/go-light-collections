// Code generated -- DO NOT EDIT.

package coll

// Pairintint defines a pseudo-generic pair.
type Pairintint struct {
	X1 int
	X2 int
}

// SliceOfPairintint is a type wrapper.
type SliceOfPairintint []Pairintint

// Mapintint is a type wrapper, implements Map interface.
type Mapintint map[int]int

// MapintSliceint is a type wrapper.
type MapintSliceint map[int]Sliceint

// MapintSetint is a type wrapper.
type MapintSetint map[int]Setint

// SetOfPairintint is a type wrapper, implements Set interface.
type SetOfPairintint map[Pairintint]bool
