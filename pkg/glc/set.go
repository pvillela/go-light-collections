/*
 *  Copyright © 2021 Paulo Villela. All rights reserved.
 *  Use of this source code is governed by the Apache 2.0 license
 *  that can be found in the LICENSE file.
 */

package glc

import (
	"errors"
	"github.com/pvillela/go-light-collections/legacy/pkg/util/util"
)

////
// Type

// Set[T0] is a type wrapper, implements Set interface.
type Set[T0 comparable] map[T0]bool

////
// Set methods

// Put mutates the receiver by adding the argument if it is not already in the receiver.
// Panics if the receiver is nil.
func (s Set[T0]) Put(e T0) {
	s[e] = true
}

// Copy returns a copy fo the receiver.
func (s Set[T0]) Copy() Set[T0] {
	if s == nil {
		return nil
	}
	s1 := make(Set[T0], len(s))
	for e := range s {
		s1[e] = true
	}
	return s1
}

// Length returns the number of items in the receiver.
func (s Set[T0]) Length() int {
	return len(s)
}

// Size returns the number of items in the receiver. Same as Length.
func (s Set[T0]) Size() int {
	return len(s)
}

// Contains returns true if the element argment is in the receiver, false otherwise.
func (s Set[T0]) Contains(elem T0) bool {
	_, ok := s[elem]
	return ok
}

// ContainsSet returns true if all the elements in the argument set are in the receiver,
// false otherwise.
func (s Set[T0]) ContainsSet(elems Set[T0]) bool {
	for e := range elems {
		if !s.Contains(e) {
			return false
		}
	}
	return true
}

// ContainsSlice returns true if all the elements in the argument slice are in the receiver,
// false otherwise.
func (s Set[T0]) ContainsSlice(elems []T0) bool {
	for _, e := range elems {
		if !s.Contains(e) {
			return false
		}
	}
	return true
}

// IsEmpty returns true if the receiver is empty, false otherwise.
func (s Set[T0]) IsEmpty() bool {
	return s == nil || len(s) == 0
}

// All returns true if all elements in the receiver satisfy the predicate, false otherwise.
func (s Set[T0]) All(pred func(T0) bool) bool {
	for e := range s {
		if !pred(e) {
			return false
		}
	}
	return true
}

// Any returns true if at least one element in the receiver satisfy the predicate, false otherwise.
func (s Set[T0]) Any(pred func(T0) bool) bool {
	for e := range s {
		if pred(e) {
			return true
		}
	}
	return false
}

// Count returns the number of items in the receiver that satisfy the predicate.
func (s Set[T0]) Count(pred func(T0) bool) int {
	count := 0
	for e := range s {
		if pred(e) {
			count++
		}
	}
	return count
}

// Filter returns a new set containing only the elements in the receiver that
// satisfy the predicate.
func (s Set[T0]) Filter(pred func(T0) bool) Set[T0] {
	if s == nil {
		return nil
	}
	output := make(Set[T0], len(s)/2) // optimizing for speed vs space
	for e := range s {
		if pred(e) {
			output[e] = true
		}
	}
	return output
}

// FilterNot returns a new set containing only the elements in the receiver that
// do not satisfy the predicate.
func (s Set[T0]) FilterNot(pred func(T0) bool) Set[T0] {
	return s.Filter(func(a T0) bool { return !pred(a) })
}

// ForEach applies the argument function to each item in the receiver.
func (s Set[T0]) ForEach(f func(T0)) {
	for e := range s {
		f(e)
	}
}

// Intersect returns a new set that contains the elements that are in both the receiver
// and the other set.
func (s Set[T0]) Intersect(other Set[T0]) Set[T0] {
	if s == nil {
		return nil
	}
	s1 := make(Set[T0], util.MinInt(len(s), len(other)))
	for e := range other {
		_, ok := s[e]
		if ok {
			s1[e] = true
		}
	}
	return s1
}

// IsNotEmpty returns true if the receiver is not empty, false otherwise.
func (s Set[T0]) IsNotEmpty() bool {
	return !s.IsEmpty()
}

// MaxWith uses a comparator function to determine the maximum value. If the set is
// nonempty, returns the element in the set with maximum value.
// Otherwise, returns an error.
func (s Set[T0]) MaxWith(comparator func(T0, T0) int) (T0, error) {
	var max T0

	if len(s) == 0 {
		return max, errors.New("empty or nil set")
	}

	first := true
	for e := range s {
		if first {
			max = e
			first = false
			continue
		}
		if comparator(max, e) < 0 {
			max = e
		}
	}
	return max, nil
}

// MinusElement -- if the element passed as an argument is present in the receiver, this
// function returns a new set with the contents of the receiver minus that element.
// Otherwise, it returns a copy of the original set.
func (s Set[T0]) MinusElement(elem T0) Set[T0] {
	s1 := s.Copy()
	delete(s1, elem)
	return s1
}

// MinusSet returns a new set which contains the elements of the receiver except for the
// elements of the other set.
func (s Set[T0]) MinusSet(other Set[T0]) Set[T0] {
	s1 := s.Copy()
	for e := range other {
		delete(s1, e)
	}
	return s1
}

// MinusSlice returns a new set which contains the elements of the receiver except for the
// elements of the slice.
func (s Set[T0]) MinusSlice(slice []T0) Set[T0] {
	s1 := s.Copy()
	for _, e := range slice {
		delete(s1, e)
	}
	return s1
}

// MinWith uses a comparator function to determine the maximum value. If the set is
// nonempty, returns theelement in the set with minimum value.
// Returns an error if the set is empty.
func (s Set[T0]) MinWith(comparator func(T0, T0) int) (T0, error) {
	reverseComp := func(a1 T0, a2 T0) int { return -comparator(a1, a2) }
	return s.MaxWith(reverseComp)
}

// Partition returns two sets, the first containing all items in the receiver that
// satisfy the argument predicate and the second containing all other items in the receiver.
func (s Set[T0]) Partition(pred func(T0) bool) (Set[T0], Set[T0]) {
	output1 := make(Set[T0], len(s)/2) // optimizing for speed vs space
	output2 := make(Set[T0], len(s)/2) // optimizing for speed vs space
	for e := range s {
		if pred(e) {
			output1[e] = true
		} else {
			output2[e] = true
		}
	}
	return output1, output2
}

// PlusElement returns a copy of the receiver with the element added to it if the element
// is not already in the receiver. If the element is already in the receiver, returns a
// copy of the receiver.
func (s Set[T0]) PlusElement(elem T0) Set[T0] {
	s1 := s.Copy()
	if s1 == nil {
		s1 = Set[T0]{}
	}
	s1[elem] = true
	return s1
}

// PlusSet returns a copy of the receiver with the elements of the other set added to it.
func (s Set[T0]) PlusSet(other Set[T0]) Set[T0] {
	s1 := s.Copy()
	if s1 == nil {
		if other == nil {
			return nil
		}
		s1 = Set[T0]{}
	}
	for e := range other {
		s1[e] = true
	}
	return s1
}

// PlusSlice returns a copy of the receiver with the elements of the slice added to it.
func (s Set[T0]) PlusSlice(slice []T0) Set[T0] {
	s1 := s.Copy()
	if s1 == nil {
		if slice == nil {
			return nil
		}
		s1 = Set[T0]{}
	}
	for _, e := range slice {
		s1[e] = true
	}
	return s1
}

// ToSlice returns a slice containing the elements of the receiver.
func (s Set[T0]) ToSlice() []T0 {
	if s == nil {
		return nil
	}
	slice := make([]T0, len(s))
	i := 0
	for e := range s {
		slice[i] = e
		i++
	}
	return slice
}

// FlatMapT1 returns the set obtained by applying the argument function to each item in the
// receiver and taking the union of the results.
func SetFlatMap[T0 comparable, T1 comparable](s Set[T0], f func(T0) map[T1]bool) map[T1]bool {
	if s == nil {
		return nil
	}
	r := make(map[T1]bool, len(s)) // optimizing for speed vs space
	for x := range s {
		for e := range f(x) {
			r[e] = true
		}
	}
	return r
}

// GroupByT1 returns a map whose keys are outputs of the keySelector function applied to
// the elements in the receiver and whose values are sets containing the elements in the
// receiver that correspond to each key obtained with the keySelector function.
func SetGroupBy[T0 comparable, T1 comparable](s Set[T0], keySelector func(T0) T1) map[T1]Set[T0] {
	if s == nil {
		return nil
	}
	m := make(map[T1]Set[T0], len(s)/2) // optimizing for speed vs space
	for x := range s {
		k := keySelector(x)
		set, ok := m[k]
		if !ok {
			set = make(Set[T0], 1)
		}
		set[x] = true
		m[k] = set
	}
	return m
}

// MapT1 returns a new set resulting from the application of a given function to
// each element of a given set.
func SetMap[T0 comparable, T1 comparable](s Set[T0], f func(T0) T1) map[T1]bool {
	if s == nil {
		return nil
	}
	r := make(map[T1]bool, len(s))
	for a := range s {
		r[f(a)] = true
	}
	return r
}

// ToMap returns a map whose keys are the first components in the elements of the receiver and
// whose values are the corresonding second components in the elements of the receiver.
// If multiple elements in the receiver have the same first component, the corresponding
// value in the resulting map will be picked from one of them.
func SetToMap[T0 comparable, T1 comparable](s Set[Pair[T0, T1]]) Map[T0, T1] {
	if s == nil {
		return nil
	}
	m := make(map[T0]T1, len(s))
	for p := range s {
		m[p.X1] = p.X2
	}
	return m
}
