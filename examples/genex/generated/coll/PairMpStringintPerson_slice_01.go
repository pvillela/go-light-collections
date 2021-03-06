// Code generated -- DO NOT EDIT.

/*
 * Copyright © 2021 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package coll

// PairSlPairMpStringintPerson is a type alias used only in Slice methods to avoid code generation issues.
type PairSlPairMpStringintPerson = struct {
	X1 PairMpStringint
	X2 Person
}

// FlatMapPerson returns the slice obtained by applying the argument f to each item in the
// receiver and concatenating the results.
func (s SlicePairMpStringint) FlatMapPerson(f func(PairMpStringint) []Person) []Person {
	if s == nil {
		return nil
	}
	r := make([]Person, 0, len(s)) // optimizing for speed vs space
	for _, x := range s {
		r = append(r, f(x)...)
	}
	return r
}

// FoldPerson returns the accumulated value obtained by applying the operation op to z,
// and the first element of the given slice, then applying op to the result of the first
// operation and the second element of the given slice, and so on.
// Returns z if the slice is empty.
func (s SlicePairMpStringint) FoldPerson(z Person, op func(Person, PairMpStringint) Person) Person {
	result := z
	for _, a := range s {
		result = op(result, a)
	}
	return result
}

// MapPerson returns a new slice resulting from the application of a given function to
// each element of a given slice.
func (s SlicePairMpStringint) MapPerson(f func(PairMpStringint) Person) []Person {
	if s == nil {
		return nil
	}
	r := make([]Person, len(s))
	for i, a := range s {
		r[i] = f(a)
	}
	return r
}

func (s SlicePairMpStringint) ZipPerson(other []Person) []PairSlPairMpStringintPerson {
	if s == nil {
		return nil
	}
	size := len(s)
	if size > len(other) {
		size = len(other)
	}
	r := make([]PairSlPairMpStringintPerson, size)
	for i := 0; i < size; i++ {
		r[i] = PairSlPairMpStringintPerson{s[i], other[i]}
	}
	return r
}
