// Code generated -- DO NOT EDIT.

package collections

// PairintDat defines a pseudo-generic pair.
type PairintDat struct {
	X1 int
	X2 Dat
}

// MapintDat is a type wrapper, implements IMap interfaces.
type MapintDat map[int]Dat

// SliceOfPairintDat is a type wrapper.
type SliceOfPairintDat []PairintDat

// SetOfPairintDat is a type wrapper, implements ISetOfPairintDat interface.
type SetOfPairintDat map[PairintDat]bool
