// Code generated -- DO NOT EDIT.

package collections

// PairintDat defines a pseudo-generic pair.
type PairintDat struct {
	X1 int
	X2 Dat
}

// SliceOfPairintDat is a type wrapper.
type SliceOfPairintDat []PairintDat

// MapintDat is a type wrapper, implements Map interface.
type MapintDat map[int]Dat

// MapintSliceDat is a type wrapper.
type MapintSliceDat map[int]SliceDat

// MapintSetDat is a type wrapper.
type MapintSetDat map[int]SetDat

// SetOfPairintDat is a type wrapper, implements Set interface.
type SetOfPairintDat map[PairintDat]bool
