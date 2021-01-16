// Code generated -- DO NOT EDIT.

package coll

func (s SliceInt) FlatMapstring(f func(Int) Slicestring) Slicestring {
	r := make([]string, 0, len(s)) // optimizing for speed vs space
	for _, x := range s {
		r = append(r, f(x)...)
	}
	return r
}

// Foldstring returns the accumulated value obtained by applying the operation op to z,
// and the first element of the given slice, then applying op to the result of the first
// operation and the second element of the given slice, and so on.
// Returns z if the slice is empty.
func (s SliceInt) Foldstring(z string, op func(string, Int) string) string {
	result := z
	for _, a := range s {
		result = op(result, a)
	}
	return result
}

func (s SliceInt) GroupBystring(keySelector func(Int) string) MapstringSliceInt {
	m := make(MapstringSliceInt)
	for _, x := range s {
		k := keySelector(x)
		lst, ok := m[k]
		if !ok {
			lst = make(SliceInt, 0, 1)
		}
		lst = append(lst, x)
		m[k] = lst
	}
	return m
}

// Mapstring returns a new slice resulting from the application of a given function to
// each element of a given slice.
func (s SliceInt) Mapstring(f func(Int) string) Slicestring {
	output := make([]string, len(s))
	for i, a := range s {
		output[i] = f(a)
	}
	return output
}
