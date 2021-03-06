// Code generated -- DO NOT EDIT.

/*
 * Copyright © 2021 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package collections

import (
	"errors"
	"reflect"
	"sort"
)

////
// Types

// SliceDat is a type wrapper, implements List interface.
type SliceDat []Dat

// Slice2Dat is a type wrapper.
type Slice2Dat []SliceDat

////
// Methods

// Copy returns a copy fo the receiver.
func (s SliceDat) Copy() SliceDat {
	if s == nil {
		return nil
	}
	s1 := make(SliceDat, len(s))
	copy(s1, s)
	return s1
}

// Length returns the number of items in the receiver.
func (s SliceDat) Length() int {
	return len(s)
}

// Size returns the number of items in the receiver. Same as Length.
func (s SliceDat) Size() int {
	return len(s)
}

// Contains returns true if the element argment is in the receiver, false otherwise.
func (s SliceDat) Contains(elem Dat) bool {
	return s.IndexOf(elem) >= 0
}

// ContainsSlice returns true if all the elements in the argument slice are in the receiver,
// false otherwise.
func (s SliceDat) ContainsAll(elems SliceDat) bool {
	for _, e := range elems {
		if !s.Contains(e) {
			return false
		}
	}
	return true
}

// Get returns the element at the specified index and true if index is within the bounds of
// the slice.
// Returns the Dat zero value and false if the index is outside the bounds of the slice.
func (s SliceDat) Get(index int) (Dat, bool) {
	if 0 <= index && index < len(s) {
		return s[index], true
	}
	var zero Dat
	return zero, false
}

// IndexOf -- if the argument element is in the receiver, this function returns the
// first index with which the element appears in the slice; otherwise, returns -1.
func (s SliceDat) IndexOf(elem Dat) int {
	pred := func(a Dat) bool { return reflect.DeepEqual(elem, a) }
	return s.IndexOfFirst(pred)
}

// IsEmpty returns true if the receiver is empty, false otherwise.
func (s SliceDat) IsEmpty() bool {
	return len(s) == 0
}

// LastIndexOf -- if the argument element is in the receiver, this function returns the
// first index with which the element appears in the slice; otherwise, returns -1.
func (s SliceDat) LastIndexOf(elem Dat) int {
	pred := func(a Dat) bool { return reflect.DeepEqual(elem, a) }
	return s.IndexOfLast(pred)
}

// SubSlice returns a copy of the sub-slice of the receiver for the index arguments.
// Panics if the index arguments do not form a valid sub-slice.
func (s SliceDat) SubSlice(fromIndex int, toIndex int) SliceDat {
	return s[fromIndex:toIndex].Copy()
}

// All returns true if all elements in the receiver satisfy the predicate, false otherwise.
func (s SliceDat) All(pred func(Dat) bool) bool {
	for _, x := range s {
		if !pred(x) {
			return false
		}
	}
	return true
}

// Any returns true if at least one element in the receiver satisfy the predicate, false otherwise.
func (s SliceDat) Any(pred func(Dat) bool) bool {
	for _, x := range s {
		if pred(x) {
			return true
		}
	}
	return false
}

// Count returns the number of items in the receiver that satisfy the predicate.
func (s SliceDat) Count(pred func(Dat) bool) int {
	count := 0
	for _, x := range s {
		if pred(x) {
			count++
		}
	}
	return count
}

// Drop returns a copy of the receiver with the first n elements removed.  If n is greater
// than the length of the receiver then the empty slice is returned.
func (s SliceDat) Drop(n int) SliceDat {
	if n > len(s) {
		return s[:0].Copy()
	}
	return s[n:].Copy()
}

// DropLast returns a copy of the receiver with the last n elements removed.  If n is greater
// than the length of the receiver then the empty slice is returned.
func (s SliceDat) DropLast(n int) SliceDat {
	if n > len(s) {
		return s[:0].Copy()
	}
	return s[:len(s)-n].Copy()
}

// DropLastWhile returns a copy of the receiver minus all the contiguous elements at the
// end of the slice that satisfy the predicate.
func (s SliceDat) DropLastWhile(pred func(Dat) bool) SliceDat {
	last := len(s)
	for i := len(s) - 1; i >= 0 && pred(s[i]); i-- {
		last = i
	}
	return s[:last].Copy()
}

// DropWhile returns a copy of the receiver minus all the contiguous elements at the
// beginning of the slice that satisfy the predicate.
func (s SliceDat) DropWhile(pred func(Dat) bool) SliceDat {
	first := 0
	for i := 0; i < len(s) && pred(s[i]); i++ {
		first = i + 1
	}
	return s[first:].Copy()
}

// Filter returns a new slice containing only the elements in the receiver that
// satisfy the predicate.
func (s SliceDat) Filter(pred func(Dat) bool) SliceDat {
	if s == nil {
		return nil
	}
	output := make(SliceDat, 0, len(s)/2) // optimizing for speed vs space
	for i, a := range s {
		if pred(s[i]) {
			output = append(output, a)
		}
	}
	return output
}

// FilterNot returns a new slice containing only the elements in the receiver that
// do not satisfy the predicate.
func (s SliceDat) FilterNot(pred func(Dat) bool) SliceDat {
	return s.Filter(func(a Dat) bool { return !pred(a) })
}

// First returns the first element in the slice, error if the slice is empty/nil.
func (s SliceDat) First() (Dat, error) {
	if len(s) == 0 {
		var zero Dat
		return zero, errors.New("empty or nil slice")
	}
	return s[0], nil
}

// ForEach applies the argument function to each item in the receiver.
func (s SliceDat) ForEach(f func(Dat)) {
	for _, x := range s {
		f(x)
	}
}

// IndexOfFirst -- if some item in the receiver satisfies the argument predicate then this
// function returns the index of the first such item; otherwise, it returns -1.
func (s SliceDat) IndexOfFirst(pred func(Dat) bool) int {
	for i, x := range s {
		if pred(x) {
			return i
		}
	}
	return -1
}

// IndexOfLast -- if some item in the receiver satisfies the argument predicate then this
// function returns the index of the last such item; otherwise, it returns -1.
func (s SliceDat) IndexOfLast(pred func(Dat) bool) int {
	for i := len(s) - 1; i >= 0; i-- {
		if pred(s[i]) {
			return i
		}
	}
	return -1
}

// IsNotEmpty returns true if the receiver is not empty, false otherwise.
func (s SliceDat) IsNotEmpty() bool {
	return !s.IsEmpty()
}

// Last returns the last element in the slice, error if the slice is empty/nil.
func (s SliceDat) Last() (Dat, error) {
	if len(s) == 0 {
		var zero Dat
		return zero, errors.New("empty or nil slice")
	}
	return s[len(s)-1], nil
}

// MaxWith uses a comparator function to determine the maximum value. If the slice is
// nonempty, returns the first element in the slice with maximum value.
// Otherwise, returns an error.
func (s SliceDat) MaxWith(comparator func(Dat, Dat) int) (Dat, error) {
	if len(s) == 0 {
		var zero Dat
		return zero, errors.New("empty or nil slice")
	}
	max := s[0]
	for i := 1; i < len(s); i++ {
		if comparator(max, s[i]) < 0 {
			max = s[i]
		}
	}
	return max, nil
}

func (s SliceDat) minusAllElement(elem Dat) SliceDat {
	return s.FilterNot(func(a Dat) bool { return reflect.DeepEqual(a, elem) })
}

// MinusSlice returns a new slice which contains the elements of the receiver except for all
// instances of the elements of the other slice.
func (s SliceDat) MinusSlice(other SliceDat) SliceDat {
	if len(other) == 0 {
		return s.Copy()
	}
	r := s // no need to copy because minusAllElements creates a copy already
	for _, x := range other {
		r = r.minusAllElement(x)
	}
	return r
}

// MinusElement -- if the element passed as an argument is present in the receiver, this
// function returns a new slice with the contents of the receiver minus the first occurrence of
// that element.  Otherwise, it returns a copy of the original slice.
func (s SliceDat) MinusElement(elem Dat) SliceDat {
	index := s.IndexOfFirst(func(a Dat) bool { return reflect.DeepEqual(a, elem) })
	if index == -1 {
		return s.Copy()
	}
	return append(s[:index].Copy(), s[index+1:]...)
}

// MinWith uses a comparator function to determine the maximum value. If the slice is
// nonempty, returns the first element in the slice with minimum value.
// Returns an error if the slice is empty.
func (s SliceDat) MinWith(comparator func(Dat, Dat) int) (Dat, error) {
	reverseComp := func(a1 Dat, a2 Dat) int { return -comparator(a1, a2) }
	return s.MaxWith(reverseComp)
}

// Partition returns two slices, the first containing all items in the receiver that
// satisfy the argument predicate and the second containing all other items in the receiver.
func (s SliceDat) Partition(pred func(Dat) bool) (SliceDat, SliceDat) {
	output1 := make(SliceDat, 0, len(s)/2) // optimizing for speed vs space
	output2 := make(SliceDat, 0, len(s)/2) // optimizing for speed vs space
	for i, a := range s {
		if pred(s[i]) {
			output1 = append(output1, a)
		} else {
			output2 = append(output2, a)
		}
	}
	return output1, output2
}

// PlusElement returns a copy of the receiver with the element appended to it.
func (s SliceDat) PlusElement(elem Dat) SliceDat {
	return append(s.Copy(), elem)
}

// PlusSlice returns a copy of the receiver with the elements of the other slice appended to it.
func (s SliceDat) PlusSlice(other SliceDat) SliceDat {
	return append(s.Copy(), other...)
}

// Reduce returns the accumulated value obtained by applying the operation op to the first
// two elements of the given slice, then applying op to the result of the first
// operation and the third element of the given slice, and so on.
// If the slice has length 1, returns the only element in the slice.
// It is a special case of Fold where the z value is the first element of the receiver and
// the fold is executed on the original slice minus the first element.
// If the slice is empty, returns an error.
func (s SliceDat) Reduce(op func(Dat, Dat) Dat) (Dat, error) {
	if len(s) == 0 {
		var zero Dat
		return zero, errors.New("empty or nil slice")
	}
	z := s[0]
	for i := 1; i < len(s); i++ {
		z = op(z, s[i])
	}
	return z, nil
}

// Reversed returns a copy of the receiver with the elements in reverse sequence.
func (s SliceDat) Reversed() SliceDat {
	if s == nil {
		return nil
	}
	r := make(SliceDat, len(s))
	size := len(s)
	for i, x := range s {
		r[size-i-1] = x
	}
	return r
}

type sortableDat struct {
	comparator func(Dat, Dat) int
	slice      SliceDat
}

func (x sortableDat) Len() int               { return len(x.slice) }
func (x sortableDat) Less(i int, j int) bool { return x.comparator(x.slice[i], x.slice[j]) < 0 }
func (x sortableDat) Swap(i int, j int)      { x.slice[i], x.slice[j] = x.slice[j], x.slice[i] }

// SortedWith returns a copy of the receiver with its elements sorted in increasing order
// based on the comparator argument.
func (s SliceDat) SortedWith(comparator func(Dat, Dat) int) SliceDat {
	r := s.Copy()
	srt := sortableDat{comparator: comparator, slice: r}
	sort.Sort(srt)
	return r
}

// Take returns a copy of the receiver containing the first n elements. If n is greater
// than the length of the receiver then a copy of the receiver is returned.
func (s SliceDat) Take(n int) SliceDat {
	if n > len(s) {
		return s.Copy()
	}
	return s[:n].Copy()
}

// TakeLast returns a copy of the receiver containing the last n elements. If n is greater
// than the length of the receiver then a copy of the receiver is returned.
func (s SliceDat) TakeLast(n int) SliceDat {
	if n > len(s) {
		return s.Copy()
	}
	return s[len(s)-n:].Copy()
}

// TakeLastWhile returns a copy of the receiver containing all the contiguous elements at the
// end of the slice that satisfy the predicate.
func (s SliceDat) TakeLastWhile(pred func(Dat) bool) SliceDat {
	first := len(s)
	for i := len(s) - 1; i >= 0 && pred(s[i]); i-- {
		first = i
	}
	return s[first:].Copy()
}

// TakeWhile returns a copy of the receiver containing all the contiguous elements at the
// beginning of the slice that satisfy the predicate.
func (s SliceDat) TakeWhile(pred func(Dat) bool) SliceDat {
	last := 0
	for i := 0; i < len(s) && pred(s[i]); i++ {
		last = i + 1
	}
	return s[:last].Copy()
}

func (s Slice2Dat) Flatten() SliceDat {
	if s == nil {
		return nil
	}
	r := make([]Dat, 0, len(s)) // optimizing for speed vs space
	for _, x := range s {
		r = append(r, x...)
	}
	return r
}
