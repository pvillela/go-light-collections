// Code generated -- DO NOT EDIT.

package coll

// PairSlPersonint is a type alias used only in Slice methods to avoid code generation issues.
type PairSlPersonint = struct {
	X1 Person
	X2 int
}

// FlatMapint returns the slice obtained by applying the argument f to each item in the
// receiver and concatenating the results.
func (s SlicePerson) FlatMapint(f func(Person) []int) []int {
	if s == nil {
		return nil
	}
	r := make([]int, 0, len(s)) // optimizing for speed vs space
	for _, x := range s {
		r = append(r, f(x)...)
	}
	return r
}

// Foldint returns the accumulated value obtained by applying the operation op to z,
// and the first element of the given slice, then applying op to the result of the first
// operation and the second element of the given slice, and so on.
// Returns z if the slice is empty.
func (s SlicePerson) Foldint(z int, op func(int, Person) int) int {
	result := z
	for _, a := range s {
		result = op(result, a)
	}
	return result
}

// Mapint returns a new slice resulting from the application of a given function to
// each element of a given slice.
func (s SlicePerson) Mapint(f func(Person) int) []int {
	if s == nil {
		return nil
	}
	r := make([]int, len(s))
	for i, a := range s {
		r[i] = f(a)
	}
	return r
}

func (s SlicePerson) Zipint(other []int) []PairSlPersonint {
	if s == nil {
		return nil
	}
	size := len(s)
	if size > len(other) {
		size = len(other)
	}
	r := make([]PairSlPersonint, size)
	for i := 0; i < size; i++ {
		r[i] = PairSlPersonint{s[i], other[i]}
	}
	return r
}
