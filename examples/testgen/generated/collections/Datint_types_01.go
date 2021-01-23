// Code generated -- DO NOT EDIT.

package collections

// PairDatint defines a pseudo-generic pair.
type PairDatint struct {
	X1 Dat
	X2 int
}

// MapDatint is a type wrapper, implements IMap interfaces.
type MapDatint map[Dat]int

// SliceOfPairDatint is a type wrapper.
type SliceOfPairDatint []PairDatint

// SetOfPairDatint is a type wrapper, implements ISetOfPairDatint interface.
type SetOfPairDatint map[PairDatint]bool
