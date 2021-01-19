// Code generated -- DO NOT EDIT.

package collections

// Pairintstring defines a pseudo-generic pair.
type Pairintstring struct {
	X1 int
	X2 string
}

// SliceOfPairintstring is a type wrapper.
type SliceOfPairintstring []Pairintstring

// Mapintstring is a type wrapper, implements Map interface.
type Mapintstring map[int]string

// MapstringSliceint is a type wrapper.
type MapstringSliceint map[string]Sliceint

// MapstringSetint is a type wrapper.
type MapstringSetint map[string]Setint

// SetOfPairintstring is a type wrapper, implements Set interface.
type SetOfPairintstring map[Pairintstring]bool
