// Code generated -- DO NOT EDIT.

package coll

func (s Sliceint) FlatMapString(f func(int) SliceString) SliceString {
	r := make([]String, 0, len(s)) // optimizing for speed vs space
	for _, x := range s {
		r = append(r, f(x)...)
	}
	return r
}

// FoldString returns the accumulated value obtained by applying the operation op to z,
// and the first element of the given slice, then applying op to the result of the first
// operation and the second element of the given slice, and so on.
// Returns z if the slice is empty.
func (s Sliceint) FoldString(z String, op func(String, int) String) String {
	result := z
	for _, a := range s {
		result = op(result, a)
	}
	return result
}

func (s Sliceint) GroupByString(keySelector func(int) String) MapStringSliceint {
	m := make(MapStringSliceint)
	for _, x := range s {
		k := keySelector(x)
		lst, ok := m[k]
		if !ok {
			lst = make(Sliceint, 0, 1)
		}
		lst = append(lst, x)
		m[k] = lst
	}
	return m
}

// MapString returns a new slice resulting from the application of a given function to
// each element of a given slice.
func (s Sliceint) MapString(f func(int) String) SliceString {
	output := make([]String, len(s))
	for i, a := range s {
		output[i] = f(a)
	}
	return output
}
