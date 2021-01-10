package collections

/////////////////////
// Simple names of methods without mangling.

func (s SliceT0) FlatMap(f func(T0) SliceT1) SliceT1 {
	r := make([]T1, 0, len(s)) // optimizing for speed vs space
	for _, x := range s {
		r = append(r, f(x)...)
	}
	return r
}

// Fold returns the accumulated value obtained by applying the operation op to z,
// and the first element of the given slice, then applying op to the result of the first
// operation and the second element of the given slice, and so on.
// Returns z if the slice is empty.
func (s SliceT0) Fold(z T1, op func(T1, T0) T1) T1 {
	result := z
	for _, a := range s {
		result = op(result, a)
	}
	return result
}

func (s SliceT0) GroupBy(keySelector func(T0) T1) MapT0SliceT1 {
	var m MapT0SliceT1 = make(map[T0]SliceT1)
	for _, x := range s {
		k := keySelector(x)
		lst, ok := m[k]
		if !ok {
			lst = make(SliceT1, 0, 1)
		}
		lst = append(lst, x)
		m[k] = lst
	}
	return m
}

// Map returns a new slice resulting from the application of a given function to
// each element of a given slice.
func (s SliceT0) Map(f func(T0) T1) SliceT1 {
	output := make([]T1, len(s))
	for i, a := range s {
		output[i] = f(a)
	}
	return output
}
