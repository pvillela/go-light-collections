/*
 * Copyright © 2021 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package slice

import (
	"errors"
	"reflect"
	"sort"

	"github.com/pvillela/go-light-collections/pkg/g2lc/pair"
	//"sort"
)

////
// Types

// Slice[T0] is a type wrapper, implements List interface.
type Slice[T0 any] []T0

////
// Methods

// Copy returns a copy fo the receiver.
func (s Slice[T0]) Copy() Slice[T0] {
	if s == nil {
		return nil
	}
	s1 := make(Slice[T0], len(s))
	copy(s1, s)
	return s1
}

// Length returns the number of items in the receiver.
func (s Slice[T0]) Length() int {
	return len(s)
}

// Size returns the number of items in the receiver. Same as Length.
func (s Slice[T0]) Size() int {
	return len(s)
}

// Contains returns true if the element argment is in the receiver, false otherwise.
func (s Slice[T0]) Contains(elem T0) bool {
	return s.IndexOf(elem) >= 0
}

// ContainsSlice returns true if all the elements in the argument slice are in the receiver,
// false otherwise.
func (s Slice[T0]) ContainsAll(elems Slice[T0]) bool {
	for _, e := range elems {
		if !s.Contains(e) {
			return false
		}
	}
	return true
}

// Get returns the element at the specified index and true if index is within the bounds of
// the slice.
// Returns the T0 zero value and false if the index is outside the bounds of the slice.
func (s Slice[T0]) Get(index int) (T0, bool) {
	if 0 <= index && index < len(s) {
		return s[index], true
	}
	var zero T0
	return zero, false
}

// IndexOf -- if the argument element is in the receiver, this function returns the
// first index with which the element appears in the slice; otherwise, returns -1.
func (s Slice[T0]) IndexOf(elem T0) int {
	pred := func(a T0) bool { return reflect.DeepEqual(elem, a) }
	return s.IndexOfFirst(pred)
}

// IsEmpty returns true if the receiver is empty, false otherwise.
func (s Slice[T0]) IsEmpty() bool {
	return len(s) == 0
}

// LastIndexOf -- if the argument element is in the receiver, this function returns the
// first index with which the element appears in the slice; otherwise, returns -1.
func (s Slice[T0]) LastIndexOf(elem T0) int {
	pred := func(a T0) bool { return reflect.DeepEqual(elem, a) }
	return s.IndexOfLast(pred)
}

// SubSlice returns a copy of the sub-slice of the receiver for the index arguments.
// Panics if the index arguments do not form a valid sub-slice.
func (s Slice[T0]) SubSlice(fromIndex int, toIndex int) Slice[T0] {
	return s[fromIndex:toIndex].Copy()
}

// All returns true if all elements in the receiver satisfy the predicate, false otherwise.
func (s Slice[T0]) All(pred func(T0) bool) bool {
	for _, x := range s {
		if !pred(x) {
			return false
		}
	}
	return true
}

// Any returns true if at least one element in the receiver satisfy the predicate, false otherwise.
func (s Slice[T0]) Any(pred func(T0) bool) bool {
	for _, x := range s {
		if pred(x) {
			return true
		}
	}
	return false
}

// Count returns the number of items in the receiver that satisfy the predicate.
func (s Slice[T0]) Count(pred func(T0) bool) int {
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
func (s Slice[T0]) Drop(n int) Slice[T0] {
	if n > len(s) {
		return s[:0].Copy()
	}
	return s[n:].Copy()
}

// DropLast returns a copy of the receiver with the last n elements removed.  If n is greater
// than the length of the receiver then the empty slice is returned.
func (s Slice[T0]) DropLast(n int) Slice[T0] {
	if n > len(s) {
		return s[:0].Copy()
	}
	return s[:len(s)-n].Copy()
}

// DropLastWhile returns a copy of the receiver minus all the contiguous elements at the
// end of the slice that satisfy the predicate.
func (s Slice[T0]) DropLastWhile(pred func(T0) bool) Slice[T0] {
	last := len(s)
	for i := len(s) - 1; i >= 0 && pred(s[i]); i-- {
		last = i
	}
	return s[:last].Copy()
}

// DropWhile returns a copy of the receiver minus all the contiguous elements at the
// beginning of the slice that satisfy the predicate.
func (s Slice[T0]) DropWhile(pred func(T0) bool) Slice[T0] {
	first := 0
	for i := 0; i < len(s) && pred(s[i]); i++ {
		first = i + 1
	}
	return s[first:].Copy()
}

// Filter returns a new slice containing only the elements in the receiver that
// satisfy the predicate.
func (s Slice[T0]) Filter(pred func(T0) bool) Slice[T0] {
	if s == nil {
		return nil
	}
	output := make(Slice[T0], 0, len(s)/2) // optimizing for speed vs space
	for i, a := range s {
		if pred(s[i]) {
			output = append(output, a)
		}
	}
	return output
}

// FilterNot returns a new slice containing only the elements in the receiver that
// do not satisfy the predicate.
func (s Slice[T0]) FilterNot(pred func(T0) bool) Slice[T0] {
	return s.Filter(func(a T0) bool { return !pred(a) })
}

// First returns the first element in the slice, error if the slice is empty/nil.
func (s Slice[T0]) First() (T0, error) {
	if len(s) == 0 {
		var zero T0
		return zero, errors.New("empty or nil slice")
	}
	return s[0], nil
}

// ForEach applies the argument function to each item in the receiver.
func (s Slice[T0]) ForEach(f func(T0)) {
	for _, x := range s {
		f(x)
	}
}

// IndexOfFirst -- if some item in the receiver satisfies the argument predicate then this
// function returns the index of the first such item; otherwise, it returns -1.
func (s Slice[T0]) IndexOfFirst(pred func(T0) bool) int {
	for i, x := range s {
		if pred(x) {
			return i
		}
	}
	return -1
}

// IndexOfLast -- if some item in the receiver satisfies the argument predicate then this
// function returns the index of the last such item; otherwise, it returns -1.
func (s Slice[T0]) IndexOfLast(pred func(T0) bool) int {
	for i := len(s) - 1; i >= 0; i-- {
		if pred(s[i]) {
			return i
		}
	}
	return -1
}

// IsNotEmpty returns true if the receiver is not empty, false otherwise.
func (s Slice[T0]) IsNotEmpty() bool {
	return !s.IsEmpty()
}

// Last returns the last element in the slice, error if the slice is empty/nil.
func (s Slice[T0]) Last() (T0, error) {
	if len(s) == 0 {
		var zero T0
		return zero, errors.New("empty or nil slice")
	}
	return s[len(s)-1], nil
}

// MaxWith uses a comparator function to determine the maximum value. If the slice is
// nonempty, returns the first element in the slice with maximum value.
// Otherwise, returns an error.
func (s Slice[T0]) MaxWith(comparator func(T0, T0) int) (T0, error) {
	if len(s) == 0 {
		var zero T0
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

func (s Slice[T0]) minusAllElement(elem T0) Slice[T0] {
	return s.FilterNot(func(a T0) bool { return reflect.DeepEqual(a, elem) })
}

// MinusSlice returns a new slice which contains the elements of the receiver except for all
// instances of the elements of the other slice.
func (s Slice[T0]) MinusSlice(other Slice[T0]) Slice[T0] {
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
func (s Slice[T0]) MinusElement(elem T0) Slice[T0] {
	index := s.IndexOfFirst(func(a T0) bool { return reflect.DeepEqual(a, elem) })
	if index == -1 {
		return s.Copy()
	}
	return append(s[:index].Copy(), s[index+1:]...)
}

// MinWith uses a comparator function to determine the maximum value. If the slice is
// nonempty, returns the first element in the slice with minimum value.
// Returns an error if the slice is empty.
func (s Slice[T0]) MinWith(comparator func(T0, T0) int) (T0, error) {
	reverseComp := func(a1 T0, a2 T0) int { return -comparator(a1, a2) }
	return s.MaxWith(reverseComp)
}

// Partition returns two slices, the first containing all items in the receiver that
// satisfy the argument predicate and the second containing all other items in the receiver.
func (s Slice[T0]) Partition(pred func(T0) bool) (Slice[T0], Slice[T0]) {
	output1 := make(Slice[T0], 0, len(s)/2) // optimizing for speed vs space
	output2 := make(Slice[T0], 0, len(s)/2) // optimizing for speed vs space
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
func (s Slice[T0]) PlusElement(elem T0) Slice[T0] {
	return append(s.Copy(), elem)
}

// PlusSlice returns a copy of the receiver with the elements of the other slice appended to it.
func (s Slice[T0]) PlusSlice(other Slice[T0]) Slice[T0] {
	return append(s.Copy(), other...)
}

// Reduce returns the accumulated value obtained by applying the operation op to the first
// two elements of the given slice, then applying op to the result of the first
// operation and the third element of the given slice, and so on.
// If the slice has length 1, returns the only element in the slice.
// It is a special case of Fold where the z value is the first element of the receiver and
// the fold is executed on the original slice minus the first element.
// If the slice is empty, returns an error.
func (s Slice[T0]) Reduce(op func(T0, T0) T0) (T0, error) {
	if len(s) == 0 {
		var zero T0
		return zero, errors.New("empty or nil slice")
	}
	z := s[0]
	for i := 1; i < len(s); i++ {
		z = op(z, s[i])
	}
	return z, nil
}

// Reversed returns a copy of the receiver with the elements in reverse sequence.
func (s Slice[T0]) Reversed() Slice[T0] {
	if s == nil {
		return nil
	}
	r := make(Slice[T0], len(s))
	size := len(s)
	for i, x := range s {
		r[size-i-1] = x
	}
	return r
}

type sortable[T0 any] struct {
	comparator func(T0, T0) int
	slice      Slice[T0]
}

func (x sortable[T0]) Len() int               { return len(x.slice) }
func (x sortable[T0]) Less(i int, j int) bool { return x.comparator(x.slice[i], x.slice[j]) < 0 }
func (x sortable[T0]) Swap(i int, j int)      { x.slice[i], x.slice[j] = x.slice[j], x.slice[i] }

// SortedWith returns a copy of the receiver with its elements sorted in increasing order
// based on the comparator argument.
func (s Slice[T0]) SortedWith(comparator func(T0, T0) int) Slice[T0] {
	r := s.Copy()
	srt := sortable[T0]{comparator: comparator, slice: r}
	sort.Sort(srt)
	return r
}

// Take returns a copy of the receiver containing the first n elements. If n is greater
// than the length of the receiver then a copy of the receiver is returned.
func (s Slice[T0]) Take(n int) Slice[T0] {
	if n > len(s) {
		return s.Copy()
	}
	return s[:n].Copy()
}

// TakeLast returns a copy of the receiver containing the last n elements. If n is greater
// than the length of the receiver then a copy of the receiver is returned.
func (s Slice[T0]) TakeLast(n int) Slice[T0] {
	if n > len(s) {
		return s.Copy()
	}
	return s[len(s)-n:].Copy()
}

// TakeLastWhile returns a copy of the receiver containing all the contiguous elements at the
// end of the slice that satisfy the predicate.
func (s Slice[T0]) TakeLastWhile(pred func(T0) bool) Slice[T0] {
	first := len(s)
	for i := len(s) - 1; i >= 0 && pred(s[i]); i-- {
		first = i
	}
	return s[first:].Copy()
}

// TakeWhile returns a copy of the receiver containing all the contiguous elements at the
// beginning of the slice that satisfy the predicate.
func (s Slice[T0]) TakeWhile(pred func(T0) bool) Slice[T0] {
	last := 0
	for i := 0; i < len(s) && pred(s[i]); i++ {
		last = i + 1
	}
	return s[:last].Copy()
}

func SliceFlatten[T0 any](s Slice[Slice[T0]]) Slice[T0] {
	if s == nil {
		return nil
	}
	r := make(Slice[T0], 0, len(s)) // optimizing for speed vs space
	for _, x := range s {
		r = append(r, x...)
	}
	return r
}

// ToSet returns a set containing the values in the receiver.
func SliceToSet[T0 comparable](s Slice[T0]) map[T0]bool {
	if s == nil {
		return nil
	}
	set := make(map[T0]bool, len(s)) // optimize for speed vs space
	for _, x := range s {
		set[x] = true
	}
	return set
}

// FlatMapT1 returns the slice obtained by applying the argument f to each item in the
// receiver and concatenating the results.
func SliceFlatMap[T0 any, T1 any](s Slice[T0], f func(T0) []T1) []T1 {
	if s == nil {
		return nil
	}
	r := make([]T1, 0, len(s)) // optimizing for speed vs space
	for _, x := range s {
		r = append(r, f(x)...)
	}
	return r
}

// FoldT1 returns the accumulated value obtained by applying the operation op to z,
// and the first element of the given slice, then applying op to the result of the first
// operation and the second element of the given slice, and so on.
// Returns z if the slice is empty.
func SliceFold[T0 any, T1 any](s Slice[T0], z T1, op func(T1, T0) T1) T1 {
	result := z
	for _, a := range s {
		result = op(result, a)
	}
	return result
}

// MapT1 returns a new slice resulting from the application of a given function to
// each element of a given slice.
func SliceMap[T0 any, T1 any](s Slice[T0], f func(T0) T1) []T1 {
	if s == nil {
		return nil
	}
	r := make([]T1, len(s))
	for i, a := range s {
		r[i] = f(a)
	}
	return r
}

func SliceZip[T0 any, T1 any](s Slice[T0], other []T1) []pair.Pair[T0, T1] {
	if s == nil {
		return nil
	}
	size := len(s)
	if size > len(other) {
		size = len(other)
	}
	r := make([]pair.Pair[T0, T1], size)
	for i := 0; i < size; i++ {
		r[i] = pair.Pair[T0, T1]{s[i], other[i]}
	}
	return r
}

// GroupByT1 returns a map whose keys are outputs of the keySelector function applied to
// the items in the receiver and whose values are slices containing the items in the
// receiver that correspond to each key obtained with the keySelector function.
func SliceGroupBy[T0 any, T1 comparable](s Slice[T0], keySelector func(T0) T1) map[T1]Slice[T0] {
	if s == nil {
		return nil
	}
	m := make(map[T1]Slice[T0], len(s)/2) // optimizing for speed vs space
	for _, x := range s {
		k := keySelector(x)
		lst, ok := m[k]
		if !ok {
			lst = make(Slice[T0], 0, 1)
		}
		lst = append(lst, x)
		m[k] = lst
	}
	return m
}

// ToMap returns a map whose keys are the first components in the items of the receiver and
// whose values are the corresonding second components in the items of the receiver.
// If multiple items in the receiver have the same first component, the corresponding
// value in the resulting map will be taken from the last such item in the receiver.
func SliceToMap[T0 comparable, T1 comparable](s Slice[pair.Pair[T0, T1]]) map[T0]T1 {
	if s == nil {
		return nil
	}
	m := make(map[T0]T1, len(s))
	for _, p := range s {
		m[p.X1] = p.X2
	}
	return m
}
