// Code generated -- DO NOT EDIT.

/*
 * Copyright © 2021 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package collections

import (
	"errors"

	"github.com/pvillela/go-light-collections/pkg/util/util"
)

////
// Type

// Setint is a type wrapper, implements Set interface.
type Setint map[int]bool

////
// Set methods

// Put mutates the receiver by adding the argument if it is not already in the receiver.
// Panics if the receiver is nil.
func (s Setint) Put(e int) {
	s[e] = true
}

// Copy returns a copy fo the receiver.
func (s Setint) Copy() Setint {
	if s == nil {
		return nil
	}
	s1 := make(Setint, len(s))
	for e := range s {
		s1[e] = true
	}
	return s1
}

// Length returns the number of items in the receiver.
func (s Setint) Length() int {
	return len(s)
}

// Size returns the number of items in the receiver. Same as Length.
func (s Setint) Size() int {
	return len(s)
}

// Contains returns true if the element argment is in the receiver, false otherwise.
func (s Setint) Contains(elem int) bool {
	_, ok := s[elem]
	return ok
}

// ContainsSet returns true if all the elements in the argument set are in the receiver,
// false otherwise.
func (s Setint) ContainsSet(elems Setint) bool {
	for e := range elems {
		if !s.Contains(e) {
			return false
		}
	}
	return true
}

// ContainsSlice returns true if all the elements in the argument slice are in the receiver,
// false otherwise.
func (s Setint) ContainsSlice(elems []int) bool {
	for _, e := range elems {
		if !s.Contains(e) {
			return false
		}
	}
	return true
}

// IsEmpty returns true if the receiver is empty, false otherwise.
func (s Setint) IsEmpty() bool {
	return s == nil || len(s) == 0
}

// All returns true if all elements in the receiver satisfy the predicate, false otherwise.
func (s Setint) All(pred func(int) bool) bool {
	for e := range s {
		if !pred(e) {
			return false
		}
	}
	return true
}

// Any returns true if at least one element in the receiver satisfy the predicate, false otherwise.
func (s Setint) Any(pred func(int) bool) bool {
	for e := range s {
		if pred(e) {
			return true
		}
	}
	return false
}

// Count returns the number of items in the receiver that satisfy the predicate.
func (s Setint) Count(pred func(int) bool) int {
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
func (s Setint) Filter(pred func(int) bool) Setint {
	if s == nil {
		return nil
	}
	output := make(Setint, len(s)/2) // optimizing for speed vs space
	for e := range s {
		if pred(e) {
			output[e] = true
		}
	}
	return output
}

// FilterNot returns a new set containing only the elements in the receiver that
// do not satisfy the predicate.
func (s Setint) FilterNot(pred func(int) bool) Setint {
	return s.Filter(func(a int) bool { return !pred(a) })
}

// ForEach applies the argument function to each item in the receiver.
func (s Setint) ForEach(f func(int)) {
	for e := range s {
		f(e)
	}
}

// Intersect returns a new set that contains the elements that are in both the receiver
// and the other set.
func (s Setint) Intersect(other Setint) Setint {
	if s == nil {
		return nil
	}
	s1 := make(Setint, util.MinInt(len(s), len(other)))
	for e := range other {
		_, ok := s[e]
		if ok {
			s1[e] = true
		}
	}
	return s1
}

// IsNotEmpty returns true if the receiver is not empty, false otherwise.
func (s Setint) IsNotEmpty() bool {
	return !s.IsEmpty()
}

// MaxWith uses a comparator function to determine the maximum value. If the set is
// nonempty, returns the element in the set with maximum value.
// Otherwise, returns an error.
func (s Setint) MaxWith(comparator func(int, int) int) (int, error) {
	var max int

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
func (s Setint) MinusElement(elem int) Setint {
	s1 := s.Copy()
	delete(s1, elem)
	return s1
}

// MinusSet returns a new set which contains the elements of the receiver except for the
// elements of the other set.
func (s Setint) MinusSet(other Setint) Setint {
	s1 := s.Copy()
	for e := range other {
		delete(s1, e)
	}
	return s1
}

// MinusSlice returns a new set which contains the elements of the receiver except for the
// elements of the slice.
func (s Setint) MinusSlice(slice []int) Setint {
	s1 := s.Copy()
	for _, e := range slice {
		delete(s1, e)
	}
	return s1
}

// MinWith uses a comparator function to determine the maximum value. If the set is
// nonempty, returns theelement in the set with minimum value.
// Returns an error if the set is empty.
func (s Setint) MinWith(comparator func(int, int) int) (int, error) {
	reverseComp := func(a1 int, a2 int) int { return -comparator(a1, a2) }
	return s.MaxWith(reverseComp)
}

// Partition returns two sets, the first containing all items in the receiver that
// satisfy the argument predicate and the second containing all other items in the receiver.
func (s Setint) Partition(pred func(int) bool) (Setint, Setint) {
	output1 := make(Setint, len(s)/2) // optimizing for speed vs space
	output2 := make(Setint, len(s)/2) // optimizing for speed vs space
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
func (s Setint) PlusElement(elem int) Setint {
	s1 := s.Copy()
	if s1 == nil {
		s1 = Setint{}
	}
	s1[elem] = true
	return s1
}

// PlusSet returns a copy of the receiver with the elements of the other set added to it.
func (s Setint) PlusSet(other Setint) Setint {
	s1 := s.Copy()
	if s1 == nil {
		if other == nil {
			return nil
		}
		s1 = Setint{}
	}
	for e := range other {
		s1[e] = true
	}
	return s1
}

// PlusSlice returns a copy of the receiver with the elements of the slice added to it.
func (s Setint) PlusSlice(slice []int) Setint {
	s1 := s.Copy()
	if s1 == nil {
		if slice == nil {
			return nil
		}
		s1 = Setint{}
	}
	for _, e := range slice {
		s1[e] = true
	}
	return s1
}

// ToSlice returns a slice containing the elements of the receiver.
func (s Setint) ToSlice() []int {
	if s == nil {
		return nil
	}
	slice := make([]int, len(s))
	i := 0
	for e := range s {
		slice[i] = e
		i++
	}
	return slice
}
