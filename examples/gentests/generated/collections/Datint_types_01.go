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

// MapintSliceDat is a type wrapper.
type MapintSliceDat map[int]SliceDat

// MapintSetDat is a type wrapper.
type MapintSetDat map[int]SetDat

// SetOfPairDatint is a type wrapper, implements Set interface.
type SetOfPairDatint map[PairDatint]bool
