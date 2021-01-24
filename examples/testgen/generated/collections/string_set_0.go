// Code generated -- DO NOT EDIT.

package collections

import (
	"errors"

	"github.com/pvillela/go-light-collections/pkg/util/util"
)

////
// Type

// Setstring is a type wrapper, implements Set interface.
type Setstring map[string]bool

/////////////////////
// Slice method that returns a Set

// ToSet returns a set containing the values in the receiver.
func (s Slicestring) ToSet() Setstring {
	if s == nil {
		return nil
	}
	set := make(Setstring, len(s)) // optimize for speed vs space
	for _, x := range s {
		set.Put(x)
	}
	return set
}

/////////////////////
// Set methods

// Put mutates the receiver by adding the argument if it is not already in the receiver.
// Panics if the receiver is nil.
func (s Setstring) Put(e string) {
	s[e] = true
}

// Copy returns a copy fo the receiver.
func (s Setstring) Copy() Setstring {
	if s == nil {
		return nil
	}
	s1 := make(Setstring, len(s))
	for e := range s {
		s1[e] = true
	}
	return s1
}

// Length returns the number of items in the receiver.
func (s Setstring) Length() int {
	return len(s)
}

// Size returns the number of items in the receiver. Same as Length.
func (s Setstring) Size() int {
	return len(s)
}

// Contains returns true if the element argment is in the receiver, false otherwise.
func (s Setstring) Contains(elem string) bool {
	_, ok := s[elem]
	return ok
}

// ContainsSet returns true if all the elements in the argument set are in the receiver,
// false otherwise.
func (s Setstring) ContainsSet(elems Setstring) bool {
	for e := range elems {
		if !s.Contains(e) {
			return false
		}
	}
	return true
}

// ContainsSlice returns true if all the elements in the argument slice are in the receiver,
// false otherwise.
func (s Setstring) ContainsSlice(elems Slicestring) bool {
	for _, e := range elems {
		if !s.Contains(e) {
			return false
		}
	}
	return true
}

// IsEmpty returns true if the receiver is empty, false otherwise.
func (s Setstring) IsEmpty() bool {
	return s == nil || len(s) == 0
}

// All returns true if all elements in the receiver satisfy the predicate, false otherwise.
func (s Setstring) All(pred func(string) bool) bool {
	for e := range s {
		if !pred(e) {
			return false
		}
	}
	return true
}

// Any returns true if at least one element in the receiver satisfy the predicate, false otherwise.
func (s Setstring) Any(pred func(string) bool) bool {
	for e := range s {
		if pred(e) {
			return true
		}
	}
	return false
}

// Count returns the number of items in the receiver that satisfy the predicate.
func (s Setstring) Count(pred func(string) bool) int {
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
func (s Setstring) Filter(pred func(string) bool) Setstring {
	if s == nil {
		return nil
	}
	output := make(Setstring, len(s)/2) // optimizing for speed vs space
	for e := range s {
		if pred(e) {
			output[e] = true
		}
	}
	return output
}

// FilterNot returns a new set containing only the elements in the receiver that
// do not satisfy the predicate.
func (s Setstring) FilterNot(pred func(string) bool) Setstring {
	return s.Filter(func(a string) bool { return !pred(a) })
}

// ForEach applies the argument function to each item in the receiver.
func (s Setstring) ForEach(f func(string)) {
	for e := range s {
		f(e)
	}
}

// Intersect returns a new set that contains the elements that are in both the receiver
// and the other set.
func (s Setstring) Intersect(other Setstring) Setstring {
	if s == nil {
		return nil
	}
	s1 := make(Setstring, util.MinInt(len(s), len(other)))
	for e := range other {
		_, ok := s[e]
		if ok {
			s1[e] = true
		}
	}
	return s1
}

// IsNotEmpty returns true if the receiver is not empty, false otherwise.
func (s Setstring) IsNotEmpty() bool {
	return !s.IsEmpty()
}

// MaxWith uses a comparator function to determine the maximum value. If the set is
// nonempty, returns the element in the set with maximum value.
// Otherwise, returns an error.
func (s Setstring) MaxWith(comparator func(string, string) int) (string, error) {
	var max string

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
func (s Setstring) MinusElement(elem string) Setstring {
	s1 := s.Copy()
	delete(s1, elem)
	return s1
}

// MinusSet returns a new set which contains the elements of the receiver except for the
// elements of the other set.
func (s Setstring) MinusSet(other Setstring) Setstring {
	s1 := s.Copy()
	for e := range other {
		delete(s1, e)
	}
	return s1
}

// MinusSlice returns a new set which contains the elements of the receiver except for the
// elements of the slice.
func (s Setstring) MinusSlice(slice Slicestring) Setstring {
	s1 := s.Copy()
	for _, e := range slice {
		delete(s1, e)
	}
	return s1
}

// MinWith uses a comparator function to determine the maximum value. If the set is
// nonempty, returns theelement in the set with minimum value.
// Returns an error if the set is empty.
func (s Setstring) MinWith(comparator func(string, string) int) (string, error) {
	reverseComp := func(a1 string, a2 string) int { return -comparator(a1, a2) }
	return s.MaxWith(reverseComp)
}

// Partition returns two sets, the first containing all items in the receiver that
// satisfy the argument predicate and the second containing all other items in the receiver.
func (s Setstring) Partition(pred func(string) bool) (Setstring, Setstring) {
	output1 := make(Setstring, len(s)/2) // optimizing for speed vs space
	output2 := make(Setstring, len(s)/2) // optimizing for speed vs space
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
func (s Setstring) PlusElement(elem string) Setstring {
	s1 := s.Copy()
	if s1 == nil {
		s1 = Setstring{}
	}
	s1[elem] = true
	return s1
}

// PlusSet returns a copy of the receiver with the elements of the other set added to it.
func (s Setstring) PlusSet(other Setstring) Setstring {
	s1 := s.Copy()
	if s1 == nil {
		if other == nil {
			return nil
		}
		s1 = Setstring{}
	}
	for e := range other {
		s1[e] = true
	}
	return s1
}

// PlusSlice returns a copy of the receiver with the elements of the slice added to it.
func (s Setstring) PlusSlice(slice Slicestring) Setstring {
	s1 := s.Copy()
	if s1 == nil {
		if slice == nil {
			return nil
		}
		s1 = Setstring{}
	}
	for _, e := range slice {
		s1[e] = true
	}
	return s1
}

// ToSlice returns a slice containing the elements of the receiver.
func (s Setstring) ToSlice() Slicestring {
	if s == nil {
		return nil
	}
	slice := make(Slicestring, len(s))
	i := 0
	for e := range s {
		slice[i] = e
		i++
	}
	return slice
}
