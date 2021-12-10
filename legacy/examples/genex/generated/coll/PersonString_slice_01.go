// Code generated -- DO NOT EDIT.

/*
 * Copyright © 2021 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package coll

// PairSlPersonString is a type alias used only in Slice methods to avoid code generation issues.
type PairSlPersonString = struct {
	X1 Person
	X2 String
}

// FlatMapString returns the slice obtained by applying the argument f to each item in the
// receiver and concatenating the results.
func (s SlicePerson) FlatMapString(f func(Person) []String) []String {
	if s == nil {
		return nil
	}
	r := make([]String, 0, len(s)) // optimizing for speed vs space
	for _, x := range s {
		r = append(r, f(x)...)
	}
	return r
}

// FoldString returns the accumulated value obtained by applying the operation op to z,
// and the first element of the given slice, then applying op to the result of the first
// operation and the second element of the given slice, and so on.
// Returns z if the slice is empty.
func (s SlicePerson) FoldString(z String, op func(String, Person) String) String {
	result := z
	for _, a := range s {
		result = op(result, a)
	}
	return result
}

// MapString returns a new slice resulting from the application of a given function to
// each element of a given slice.
func (s SlicePerson) MapString(f func(Person) String) []String {
	if s == nil {
		return nil
	}
	r := make([]String, len(s))
	for i, a := range s {
		r[i] = f(a)
	}
	return r
}

func (s SlicePerson) ZipString(other []String) []PairSlPersonString {
	if s == nil {
		return nil
	}
	size := len(s)
	if size > len(other) {
		size = len(other)
	}
	r := make([]PairSlPersonString, size)
	for i := 0; i < size; i++ {
		r[i] = PairSlPersonString{s[i], other[i]}
	}
	return r
}
