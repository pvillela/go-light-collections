// Code generated -- DO NOT EDIT.

package coll

func (s SliceString) FlatMapint(f func(String) Sliceint) Sliceint {
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

func (s SliceString) GroupByint(keySelector func(String) int) MapintSliceString {
	m := make(MapintSliceString)
	for _, x := range s {
		k := keySelector(x)
		lst, ok := m[k]
		if !ok {
			lst = make(SliceString, 0, 1)
		}
		lst = append(lst, x)
		m[k] = lst
	}
	return m
}

// Mapint returns a new slice resulting from the application of a given function to
// each element of a given slice.
func (s SliceString) Mapint(f func(String) int) Sliceint {
	output := make([]int, len(s))
	for i, a := range s {
		output[i] = f(a)
	}
	return output
}
