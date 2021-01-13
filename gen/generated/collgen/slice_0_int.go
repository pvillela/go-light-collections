package collgen

import (
	"errors"
	"reflect"
	"sort"
)

// Sliceint is a type wrapper, implements List interface.
type Sliceint []int

/////////////////////
// Methods

// Copy returns a copy fo the receiver.
func (s Sliceint) Copy() Sliceint {
	if s == nil {
		var zero Sliceint
		return zero
	}
	s1 := make(Sliceint, len(s))
	copy(s1, s)
	return s1
}

// Length returns the number of items in the receiver.
func (s Sliceint) Length() int {
	return len(s)
}

// Contains returns true if the element argment is in the receiver, false otherwise.
func (s Sliceint) Contains(elem int) bool {
	return s.IndexOf(elem) >= 0
}

// ContainsAll returns true if all the elements in the argument slice are in the receiver,
// false otherwise.
func (s Sliceint) ContainsAll(elems Sliceint) bool {
	for i := range elems {
		e := elems[i]
		if !s.Contains(e) {
			return false
		}
	}
	return true
}

// Get returns the element at the specified index and true if index is within the bounds of
// the slice.
// Returns the int zero value and false if the index is outside the bounds of the slice.
func (s Sliceint) Get(index int) (int, bool) {
	if 0 <= index && index < len(s) {
		return s[index], true
	}
	var zero int
	return zero, false
}

// IndexOf -- if the argument element is in the receiver, this function returns the
// first index with which the element appears in the slice; otherwise, returns -1.
func (s Sliceint) IndexOf(elem int) int {
	pred := func(a int) bool { return reflect.DeepEqual(elem, a) }
	return s.IndexOfFirst(pred)
}

// IsEmpty returns true if the receiver is empty, false otherwise.
func (s Sliceint) IsEmpty() bool {
	return s == nil || len(s) == 0
}

// LastIndexOf -- if the argument element is in the receiver, this function returns the
// first index with which the element appears in the slice; otherwise, returns -1.
func (s Sliceint) LastIndexOf(elem int) int {
	pred := func(a int) bool { return reflect.DeepEqual(elem, a) }
	return s.IndexOfLast(pred)
}

// SubSlice returns a copy of the sub-slice of the receiver for the index arguments.
// Panics if the index arguments do not form a valid sub-slice.
func (s Sliceint) SubSlice(fromIndex int, toIndex int) Sliceint {
	return s[fromIndex:toIndex].Copy()
}

// All returns true if all elements in the receiver satisfy the predicate, false otherwise.
func (s Sliceint) All(pred func(int) bool) bool {
	for _, x := range s {
		if !pred(x) {
			return false
		}
	}
	return true
}

// Any returns true if at least one element in the receiver satisfy the predicate, false otherwise.
func (s Sliceint) Any(pred func(int) bool) bool {
	for _, x := range s {
		if pred(x) {
			return true
		}
	}
	return false
}

// Count returns the number of items in the receiver that satisfy the predicate.
func (s Sliceint) Count(pred func(int) bool) int {
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
func (s Sliceint) Drop(n int) Sliceint {
	if n > len(s) {
		return s[:0].Copy()
	}
	return s[n:].Copy()
}

// DropLast returns a copy of the receiver with the last n elements removed.  If n is greater
// than the length of the receiver then the empty slice is returned.
func (s Sliceint) DropLast(n int) Sliceint {
	if n > len(s) {
		return s[:0].Copy()
	}
	return s[:len(s)-n].Copy()
}

// DropLastWhile returns a copy of the receiver minus all the contiguous elements at the
// end of the slice that satisfy the predicate.
func (s Sliceint) DropLastWhile(pred func(int) bool) Sliceint {
	last := len(s)
	for i := len(s) - 1; i >= 0 && pred(s[i]); i-- {
		last = i
	}
	return s[:last].Copy()
}

// DropWhile returns a copy of the receiver minus all the contiguous elements at the
// beginning of the slice that satisfy the predicate.
func (s Sliceint) DropWhile(pred func(int) bool) Sliceint {
	first := 0
	for i := 0; i < len(s) && pred(s[i]); i++ {
		first = i + 1
	}
	return s[first:].Copy()
}

// Filter returns a new slice containing only the elements in the receiver that
// satisfy the predicate.
func (s Sliceint) Filter(pred func(int) bool) Sliceint {
	output := make(Sliceint, 0, len(s)/2) // optimizing for speed vs space
	for i, a := range s {
		if pred(s[i]) {
			output = append(output, a)
		}
	}
	return output
}

// FilterNot returns a new slice containing only the elements in the receiver that
// do not satisfy the predicate.
func (s Sliceint) FilterNot(pred func(int) bool) Sliceint {
	return s.Filter(func(a int) bool { return !pred(a) })
}

// First returns the first element in the slice, error if the slice is empty/nil.
func (s Sliceint) First() (int, error) {
	if len(s) == 0 {
		var zero int
		return zero, errors.New("empty slice")
	}
	return s[0], nil
}

// ForEach applies the argument function to each item in the receiver.
func (s Sliceint) ForEach(f func(int)) {
	for _, x := range s {
		f(x)
	}
}

// IndexOfFirst -- if some item in the receiver satisfies the argument predicate then this
// function returns the index of the first such item; otherwise, it returns -1.
func (s Sliceint) IndexOfFirst(pred func(int) bool) int {
	for i, x := range s {
		if pred(x) {
			return i
		}
	}
	return -1
}

// IndexOfLast -- if some item in the receiver satisfies the argument predicate then this
// function returns the index of the last such item; otherwise, it returns -1.
func (s Sliceint) IndexOfLast(pred func(int) bool) int {
	for i := len(s) - 1; i >= 0; i-- {
		if pred(s[i]) {
			return i
		}
	}
	return -1
}

// IsNotEmpty returns true if the receiver is not empty, false otherwise.
func (s Sliceint) IsNotEmpty() bool {
	return !s.IsEmpty()
}

// Last returns the last element in the slice, error if the slice is empty/nil.
func (s Sliceint) Last() (int, error) {
	if len(s) == 0 {
		var zero int
		return zero, errors.New("empty slice")
	}
	return s[len(s)-1], nil
}

// MaxWith uses a comparator function to determine the maximum value. If the slice is
// non-empty, returns the first element in the slice with maximum value.
// Otherwise, returns an error.
func (s Sliceint) MaxWith(comparator func(int, int) int) (int, error) {
	if len(s) == 0 {
		var zero int
		return zero, errors.New("empty slice")
	}
	max := s[0]
	for i := 1; i < len(s); i++ {
		if comparator(max, s[i]) < 0 {
			max = s[i]
		}
	}
	return max, nil
}

func (s Sliceint) minusAllElement(elem int) Sliceint {
	return s.FilterNot(func(a int) bool { return reflect.DeepEqual(a, elem) })
}

// Minus returns a new slice which contains the elements of the receiver except for all
// instances of the elements of the other slice.
func (s Sliceint) Minus(other Sliceint) Sliceint {
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
func (s Sliceint) MinusElement(elem int) Sliceint {
	index := s.IndexOfFirst(func(a int) bool { return reflect.DeepEqual(a, elem) })
	if index == -1 {
		return s.Copy()
	}
	return append(s[:index].Copy(), s[index+1:]...)
}

// MinWith uses a comparator function to determine the maximum value. If the slice is
// non-empty, returns the first element in the slice with minimum value.
// Returns an error is the slice is empty.
func (s Sliceint) MinWith(comparator func(int, int) int) (int, error) {
	reverseComp := func(a1 int, a2 int) int { return -comparator(a1, a2) }
	return s.MaxWith(reverseComp)
}

// Partition returns two slices, the first containing all items in the receiver that
// satisfy the argument predicate and the second containing all other items in the receiver.
func (s Sliceint) Partition(pred func(int) bool) (Sliceint, Sliceint) {
	output1 := make(Sliceint, 0, len(s)/2) // optimizing for speed vs space
	output2 := make(Sliceint, 0, len(s)/2) // optimizing for speed vs space
	for i, a := range s {
		if pred(s[i]) {
			output1 = append(output1, a)
		} else {
			output2 = append(output2, a)
		}
	}
	return output1, output2
}

// Plus returns a copy of the receiver with the elements of the other slice appended to it.
func (s Sliceint) Plus(other Sliceint) Sliceint {
	return append(s.Copy(), other...)
}

// PlusElement returns a copy of the receiver with the element appended to it.
func (s Sliceint) PlusElement(elem int) Sliceint {
	return append(s.Copy(), elem)
}

// Reduce returns the accumulated value obtained by applying the operation op to the first
// two elements of the given slice, then applying op to the result of the first
// operation and the third element of the given slice, and so on.
// If the slice has length 1, returns the only element in the slice.
// It is a special case of Fold where the z value is the first element of the receiver and
// the fold is executed on the original slice minus the first element.
// If the slice is empty, returns an error.
func (s Sliceint) Reduce(op func(int, int) int) (int, error) {
	if len(s) == 0 {
		var zero int
		return zero, errors.New("empty slice")
	}
	z := s[0]
	for i := 1; i < len(s); i++ {
		z = op(z, s[i])
	}
	return z, nil
}

// Reversed returns a copy of the receiver with the elements in reverse sequence.
func (s Sliceint) Reversed() Sliceint {
	r := make(Sliceint, len(s))
	size := len(s)
	for i, x := range s {
		r[size-i-1] = x
	}
	return r
}

type sortable struct {
	comparator func(int, int) int
	slice      Sliceint
}

func (x sortable) Len() int               { return len(x.slice) }
func (x sortable) Less(i int, j int) bool { return x.comparator(x.slice[i], x.slice[j]) < 0 }
func (x sortable) Swap(i int, j int)      { x.slice[i], x.slice[j] = x.slice[j], x.slice[i] }

// SortedWith returns a copy of the receiver with its elements sorted in increasing order
// based on the comparator argument.
func (s Sliceint) SortedWith(comparator func(int, int) int) Sliceint {
	r := s.Copy()
	srt := sortable{comparator: comparator, slice: r}
	sort.Sort(srt)
	return r
}

// Take returns a copy of the receiver containing the first n elements. If n is greater
// than the length of the receiver then a copy of the receiver is returned.
func (s Sliceint) Take(n int) Sliceint {
	if n > len(s) {
		return s.Copy()
	}
	return s[:n].Copy()
}

// TakeLast returns a copy of the receiver containing the last n elements. If n is greater
// than the length of the receiver then a copy of the receiver is returned.
func (s Sliceint) TakeLast(n int) Sliceint {
	if n > len(s) {
		return s.Copy()
	}
	return s[len(s)-n:].Copy()
}

// TakeLastWhile returns a copy of the receiver containing all the contiguous elements at the
// end of the slice that satisfy the predicate.
func (s Sliceint) TakeLastWhile(pred func(int) bool) Sliceint {
	first := len(s)
	for i := len(s) - 1; i >= 0 && pred(s[i]); i-- {
		first = i
	}
	return s[first:].Copy()
}

// TakeWhile returns a copy of the receiver containing all the contiguous elements at the
// beginning of the slice that satisfy the predicate.
func (s Sliceint) TakeWhile(pred func(int) bool) Sliceint {
	last := 0
	for i := 0; i < len(s) && pred(s[i]); i++ {
		last = i + 1
	}
	return s[:last].Copy()
}

// ToSlice returns the underlying native Go slice.
func (s Sliceint) ToSlice() Sliceint {
	return s
}
