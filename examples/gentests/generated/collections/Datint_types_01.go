// Code generated -- DO NOT EDIT.

package collections

// PairDatint defines a pseudo-generic pair.
type PairDatint struct {
	X1 Dat
	X2 int
}

// SliceOfPairDatint is a type wrapper.
type SliceOfPairDatint []PairDatint

// MapDatint is a type wrapper, implements Map interface.
type MapDatint map[Dat]int

// MapDatSliceint is a type wrapper.
type MapDatSliceint map[Dat]Sliceint

// MapDatSetint is a type wrapper.
type MapDatSetint map[Dat]Setint

// SetOfPairDatint is a type wrapper, implements Set interface.
type SetOfPairDatint map[PairDatint]bool
