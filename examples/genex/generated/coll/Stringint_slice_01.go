// Code generated -- DO NOT EDIT.

package coll

// PairSlStringint is a type alias used only in Slice methods to avoid code generation issues.
type PairSlStringint = struct {
	X1 String
	X2 int
}

// FlatMapint returns the slice obtained by applying the argument f to each item in the
// receiver and concatenating the results.
func (s SliceString) FlatMapint(f func(String) Sliceint) Sliceint {
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
func (s SliceString) Foldint(z int, op func(int, String) int) int {
	result := z
	for _, a := range s {
		result = op(result, a)
	}
	return result
}

// Mapint returns a new slice resulting from the application of a given function to
// each element of a given slice.
func (s SliceString) Mapint(f func(String) int) Sliceint {
	if s == nil {
		return nil
	}
	r := make(Sliceint, len(s))
	for i, a := range s {
		r[i] = f(a)
	}
	return r
}

func (s SliceString) Zipint(other Sliceint) []PairSlStringint {
	if s == nil {
		return nil
	}
	size := len(s)
	if size > len(other) {
		size = len(other)
	}
	r := make([]PairSlStringint, size)
	for i := 0; i < size; i++ {
		r[i] = PairSlStringint{s[i], other[i]}
	}
	return r
}
