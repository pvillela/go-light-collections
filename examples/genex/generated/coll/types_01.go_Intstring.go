// Code generated -- DO NOT EDIT.

package coll

// PairIntstring defines a pseudo-generic pair.
type PairIntstring struct {
	X1 Int
	X2 string
}

// SliceOfPairIntstring is a type wrapper.
type SliceOfPairIntstring []PairIntstring

// MapIntstring is a type wrapper, implements Map interface.
type MapIntstring map[Int]string

// MapstringSliceInt is a type wrapper.
type MapstringSliceInt map[string]SliceInt

// SetOfPairIntstring is a type wrapper, implements Set interface.
type SetOfPairIntstring map[PairIntstring]bool
