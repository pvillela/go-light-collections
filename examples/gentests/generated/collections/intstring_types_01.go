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

// MapintSlicestring is a type wrapper.
type MapintSlicestring map[int]Slicestring

// MapintSetstring is a type wrapper.
type MapintSetstring map[int]Setstring

// SetOfPairintstring is a type wrapper, implements Set interface.
type SetOfPairintstring map[Pairintstring]bool
