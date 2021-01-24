package glc

// PairSlT0T1 is a type alias used only in Slice methods to avoid code generation issues.
type PairSlT0T1 = struct {
	X1 T0
	X2 T1
}

// FlatMapT1 returns the slice obtained by applying the argument f to each item in the
// receiver and concatenating the results.
func (s SliceT0) FlatMapT1(f func(T0) []T1) []T1 {
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
func (s SliceT0) FoldT1(z T1, op func(T1, T0) T1) T1 {
	result := z
	for _, a := range s {
		result = op(result, a)
	}
	return result
}

// MapT1 returns a new slice resulting from the application of a given function to
// each element of a given slice.
func (s SliceT0) MapT1(f func(T0) T1) []T1 {
	if s == nil {
		return nil
	}
	r := make([]T1, len(s))
	for i, a := range s {
		r[i] = f(a)
	}
	return r
}

func (s SliceT0) ZipT1(other []T1) []PairSlT0T1 {
	if s == nil {
		return nil
	}
	size := len(s)
	if size > len(other) {
		size = len(other)
	}
	r := make([]PairSlT0T1, size)
	for i := 0; i < size; i++ {
		r[i] = PairSlT0T1{s[i], other[i]}
	}
	return r
}
