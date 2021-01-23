// Code generated -- DO NOT EDIT.

package collections

// Pairintint defines a pseudo-generic pair.
type Pairintint struct {
	X1 int
	X2 int
}

// Mapintint is a type wrapper, implements IMap interfaces.
type Mapintint map[int]int

// SliceOfPairintint is a type wrapper.
type SliceOfPairintint []Pairintint

// SetOfPairintint is a type wrapper, implements ISetOfPairintint interface.
type SetOfPairintint map[Pairintint]bool
