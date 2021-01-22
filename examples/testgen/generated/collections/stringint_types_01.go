// Code generated -- DO NOT EDIT.

package collections

// Pairstringint defines a pseudo-generic pair.
type Pairstringint struct {
	X1 string
	X2 int
}

// SliceOfPairstringint is a type wrapper.
type SliceOfPairstringint []Pairstringint

// Mapstringint is a type wrapper, implements Map interface.
type Mapstringint map[string]int

// MapstringSliceint is a type wrapper.
type MapstringSliceint map[string]Sliceint

// MapstringSetint is a type wrapper.
type MapstringSetint map[string]Setint

// SetOfPairstringint is a type wrapper, implements Set interface.
type SetOfPairstringint map[Pairstringint]bool
