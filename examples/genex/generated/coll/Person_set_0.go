// Code generated -- DO NOT EDIT.

package coll

import (
	"errors"

	"github.com/pvillela/go-light-collections/pkg/util/math"
)

/////////////////////
// Helper functions

/////////////////////
// Slice method that returns a Set

// ToSet returns a set containing the values in the receiver.
func (s SlicePerson) ToSet() SetPerson {
	if s == nil {
		return nil
	}
	set := make(SetPerson, len(s)) // optimize for speed vs space
	for _, x := range s {
		set.Put(x)
	}
	return set
}

/////////////////////
// Set methods

// Put mutates the receiver by adding the argument if it is not already in the receiver.
// Panics if the receiver is nil.
func (s SetPerson) Put(e Person) {
	s[e] = true
}

// Copy returns a copy fo the receiver.
func (s SetPerson) Copy() SetPerson {
	if s == nil {
		return nil
	}
	s1 := make(SetPerson, len(s))
	for e := range s {
		s1[e] = true
	}
	return s1
}

// Length returns the number of items in the receiver.
func (s SetPerson) Length() int {
	return len(s)
}

// Size returns the number of items in the receiver. Same as Length.
func (s SetPerson) Size() int {
	return len(s)
}

// Contains returns true if the element argment is in the receiver, false otherwise.
func (s SetPerson) Contains(elem Person) bool {
	_, ok := s[elem]
	return ok
}

// ContainsSet returns true if all the elements in the argument set are in the receiver,
// false otherwise.
func (s SetPerson) ContainsSet(elems SetPerson) bool {
	for e := range elems {
		if !s.Contains(e) {
			return false
		}
	}
	return true
}

// ContainsSlice returns true if all the elements in the argument slice are in the receiver,
// false otherwise.
func (s SetPerson) ContainsSlice(elems SlicePerson) bool {
	for _, e := range elems {
		if !s.Contains(e) {
			return false
		}
	}
	return true
}

// IsEmpty returns true if the receiver is empty, false otherwise.
func (s SetPerson) IsEmpty() bool {
	return s == nil || len(s) == 0
}

// All returns true if all elements in the receiver satisfy the predicate, false otherwise.
func (s SetPerson) All(pred func(Person) bool) bool {
	for e := range s {
		if !pred(e) {
			return false
		}
	}
	return true
}

// Any returns true if at least one element in the receiver satisfy the predicate, false otherwise.
func (s SetPerson) Any(pred func(Person) bool) bool {
	for e := range s {
		if pred(e) {
			return true
		}
	}
	return false
}

// Count returns the number of items in the receiver that satisfy the predicate.
func (s SetPerson) Count(pred func(Person) bool) int {
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
func (s SetPerson) Filter(pred func(Person) bool) SetPerson {
	if s == nil {
		return nil
	}
	output := make(SetPerson, len(s)/2) // optimizing for speed vs space
	for e := range s {
		if pred(e) {
			output[e] = true
		}
	}
	return output
}

// FilterNot returns a new set containing only the elements in the receiver that
// do not satisfy the predicate.
func (s SetPerson) FilterNot(pred func(Person) bool) SetPerson {
	return s.Filter(func(a Person) bool { return !pred(a) })
}

// ForEach applies the argument function to each item in the receiver.
func (s SetPerson) ForEach(f func(Person)) {
	for e := range s {
		f(e)
	}
}

// Intersect returns a new set that contains the elements that are in both the receiver
// and the other set.
func (s SetPerson) Intersect(other SetPerson) SetPerson {
	if s == nil {
		return nil
	}
	s1 := make(SetPerson, math.MinInt(len(s), len(other)))
	for e := range other {
		_, ok := s[e]
		if ok {
			s1[e] = true
		}
	}
	return s1
}

// IsNotEmpty returns true if the receiver is not empty, false otherwise.
func (s SetPerson) IsNotEmpty() bool {
	return !s.IsEmpty()
}

// MaxWith uses a comparator function to determine the maximum value. If the set is
// nonempty, returns the element in the set with maximum value.
// Otherwise, returns an error.
func (s SetPerson) MaxWith(comparator func(Person, Person) int) (Person, error) {
	var max Person

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
func (s SetPerson) MinusElement(elem Person) SetPerson {
	s1 := s.Copy()
	delete(s1, elem)
	return s1
}

// MinusSet returns a new set which contains the elements of the receiver except for the
// elements of the other set.
func (s SetPerson) MinusSet(other SetPerson) SetPerson {
	s1 := s.Copy()
	for e := range other {
		delete(s1, e)
	}
	return s1
}

// MinusSlice returns a new set which contains the elements of the receiver except for the
// elements of the slice.
func (s SetPerson) MinusSlice(slice SlicePerson) SetPerson {
	s1 := s.Copy()
	for _, e := range slice {
		delete(s1, e)
	}
	return s1
}

// MinWith uses a comparator function to determine the maximum value. If the set is
// nonempty, returns theelement in the set with minimum value.
// Returns an error if the set is empty.
func (s SetPerson) MinWith(comparator func(Person, Person) int) (Person, error) {
	reverseComp := func(a1 Person, a2 Person) int { return -comparator(a1, a2) }
	return s.MaxWith(reverseComp)
}

// Partition returns two sets, the first containing all items in the receiver that
// satisfy the argument predicate and the second containing all other items in the receiver.
func (s SetPerson) Partition(pred func(Person) bool) (SetPerson, SetPerson) {
	output1 := make(SetPerson, len(s)/2) // optimizing for speed vs space
	output2 := make(SetPerson, len(s)/2) // optimizing for speed vs space
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
func (s SetPerson) PlusElement(elem Person) SetPerson {
	s1 := s.Copy()
	if s1 == nil {
		s1 = SetPerson{}
	}
	s1[elem] = true
	return s1
}

// PlusSet returns a copy of the receiver with the elements of the other set added to it.
func (s SetPerson) PlusSet(other SetPerson) SetPerson {
	s1 := s.Copy()
	if s1 == nil {
		if other == nil {
			return nil
		}
		s1 = SetPerson{}
	}
	for e := range other {
		s1[e] = true
	}
	return s1
}

// PlusSlice returns a copy of the receiver with the elements of the slice added to it.
func (s SetPerson) PlusSlice(slice SlicePerson) SetPerson {
	s1 := s.Copy()
	if s1 == nil {
		if slice == nil {
			return nil
		}
		s1 = SetPerson{}
	}
	for _, e := range slice {
		s1[e] = true
	}
	return s1
}

// ToSlice returns a slice containing the elements of the receiver.
func (s SetPerson) ToSlice() SlicePerson {
	if s == nil {
		return nil
	}
	slice := make(SlicePerson, len(s))
	i := 0
	for e := range s {
		slice[i] = e
		i++
	}
	return slice
}
