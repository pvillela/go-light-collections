// Code generated -- DO NOT EDIT.

package collections

// Pairintstring defines a pseudo-generic pair.
type Pairintstring struct {
	X1 int
	X2 string
}

// Mapintstring is a type wrapper, implements IMap interfaces.
type Mapintstring map[int]string

// SliceOfPairintstring is a type wrapper.
type SliceOfPairintstring []Pairintstring

// SetOfPairintstring is a type wrapper, implements ISetOfPairintstring interface.
type SetOfPairintstring map[Pairintstring]bool
