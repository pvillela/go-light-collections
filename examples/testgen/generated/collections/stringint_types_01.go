// Code generated -- DO NOT EDIT.

package collections

// Pairstringint defines a pseudo-generic pair.
type Pairstringint struct {
	X1 string
	X2 int
}

// Mapstringint is a type wrapper, implements IMap interfaces.
type Mapstringint map[string]int

// SliceOfPairstringint is a type wrapper.
type SliceOfPairstringint []Pairstringint

// SetOfPairstringint is a type wrapper, implements ISetOfPairstringint interface.
type SetOfPairstringint map[Pairstringint]bool
