package collections

import (
	"errors"

	"github.com/pvillela/go-light-collections/pkg/util/math"
)

/////////////////////
// Helper functions

/////////////////////
// Slice method that returns a Set

// ToSet returns a set containing the values in the receiver.
func (s SliceT0) ToSet() SetT0 {
	set := make(SetT0, len(s)) // optimize for speed vs space
	for _, x := range s {
		set.Put(x)
	}
	return set
}

/////////////////////
// Set methods

// Put mutates the receiver by adding the argument if it is not already in the receiver.
func (s SetT0) Put(e T0) {
	s[e] = true
}

// Copy returns a copy fo the receiver.
func (s SetT0) Copy() SetT0 {
	if s == nil {
		return nil
	}
	s1 := make(SetT0, len(s))
	for e := range s {
		s1[e] = true
	}
	return s1
}

// Length returns the number of items in the receiver.
func (s SetT0) Length() int {
	return len(s)
}

// Size returns the number of items in the receiver. Same as Length.
func (s SetT0) Size() int {
	return len(s)
}

// Contains returns true if the element argment is in the receiver, false otherwise.
func (s SetT0) Contains(elem T0) bool {
	_, ok := s[elem]
	return ok
}

// ContainsSet returns true if all the elements in the argument set are in the receiver,
// false otherwise.
func (s SetT0) ContainsSet(elems SetT0) bool {
	for e := range elems {
		if !s.Contains(e) {
			return false
		}
	}
	return true
}

// ContainsSlice returns true if all the elements in the argument slice are in the receiver,
// false otherwise.
func (s SetT0) ContainsSlice(elems SliceT0) bool {
	for _, e := range elems {
		if !s.Contains(e) {
			return false
		}
	}
	return true
}

// IsEmpty returns true if the receiver is empty, false otherwise.
func (s SetT0) IsEmpty() bool {
	return s == nil || len(s) == 0
}

// All returns true if all elements in the receiver satisfy the predicate, false otherwise.
func (s SetT0) All(pred func(T0) bool) bool {
	for e := range s {
		if !pred(e) {
			return false
		}
	}
	return true
}

// Any returns true if at least one element in the receiver satisfy the predicate, false otherwise.
func (s SetT0) Any(pred func(T0) bool) bool {
	for e := range s {
		if pred(e) {
			return true
		}
	}
	return false
}

// Count returns the number of items in the receiver that satisfy the predicate.
func (s SetT0) Count(pred func(T0) bool) int {
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
func (s SetT0) Filter(pred func(T0) bool) SetT0 {
	output := make(SetT0, len(s)/2) // optimizing for speed vs space
	for e := range s {
		if pred(e) {
			output[e] = true
		}
	}
	return output
}

// FilterNot returns a new set containing only the elements in the receiver that
// do not satisfy the predicate.
func (s SetT0) FilterNot(pred func(T0) bool) SetT0 {
	return s.Filter(func(a T0) bool { return !pred(a) })
}

// ForEach applies the argument function to each item in the receiver.
func (s SetT0) ForEach(f func(T0)) {
	for e := range s {
		f(e)
	}
}

// Intersect returns a new set that contains the elements that are in both the receiver
// and the other set.
func (s SetT0) Intersect(other SetT0) SetT0 {
	s1 := make(SetT0, math.MinInt(len(s), len(other)))
	for e := range other {
		_, ok := s[e]
		if ok {
			s1[e] = true
		}
	}
	return s1
}

// IsNotEmpty returns true if the receiver is not empty, false otherwise.
func (s SetT0) IsNotEmpty() bool {
	return !s.IsEmpty()
}

// MaxWith uses a comparator function to determine the maximum value. If the set is
// non-empty, returns the element in the set with maximum value.
// Otherwise, returns an error.
func (s SetT0) MaxWith(comparator func(T0, T0) int) (T0, error) {
	var max T0

	if len(s) == 0 {
		return max, errors.New("empty set")
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
func (s SetT0) MinusElement(elem T0) SetT0 {
	s1 := s.Copy()
	delete(s1, elem)
	return s1
}

// MinusSet returns a new set which contains the elements of the receiver except for the
// elements of the other set.
func (s SetT0) MinusSet(other SetT0) SetT0 {
	s1 := s.Copy()
	for e := range other {
		delete(s1, e)
	}
	return s1
}

// MinusSlice returns a new set which contains the elements of the receiver except for the
// elements of the slice.
func (s SetT0) MinusSlice(slice SliceT0) SetT0 {
	s1 := s.Copy()
	for _, e := range slice {
		delete(s1, e)
	}
	return s1
}

// MinWith uses a comparator function to determine the maximum value. If the set is
// non-empty, returns theelement in the set with minimum value.
// Returns an error if the set is empty.
func (s SetT0) MinWith(comparator func(T0, T0) int) (T0, error) {
	reverseComp := func(a1 T0, a2 T0) int { return -comparator(a1, a2) }
	return s.MaxWith(reverseComp)
}

// Partition returns two sets, the first containing all items in the receiver that
// satisfy the argument predicate and the second containing all other items in the receiver.
func (s SetT0) Partition(pred func(T0) bool) (SetT0, SetT0) {
	output1 := make(SetT0, len(s)/2) // optimizing for speed vs space
	output2 := make(SetT0, len(s)/2) // optimizing for speed vs space
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
func (s SetT0) PlusElement(elem T0) SetT0 {
	s1 := s.Copy()
	s1[elem] = true
	return s1
}

// PlusSet returns a copy of the receiver with the elements of the other set added to it.
func (s SetT0) PlusSet(other SetT0) SetT0 {
	s1 := s.Copy()
	for e := range other {
		s1[e] = true
	}
	return s1
}

// PlusSlice returns a copy of the receiver with the elements of the slice added to it.
func (s SetT0) PlusSlice(slice SliceT0) SetT0 {
	s1 := s.Copy()
	for _, e := range slice {
		s1[e] = true
	}
	return s1
}

// ToSlice returns a slice containing the elements of the receiver.
func (s SetT0) ToSlice() SliceT0 {
	slice := make(SliceT0, len(s))
	i := 0
	for e := range s {
		slice[i] = e
		i++
	}
	return slice
}
