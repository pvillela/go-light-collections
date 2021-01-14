// Code generated -- DO NOT EDIT.

package coll

func (s Sliceint) FlatMapstring(f func(int) Slicestring) Slicestring {
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
func (s Sliceint) Foldstring(z string, op func(string, int) string) string {
	result := z
	for _, a := range s {
		result = op(result, a)
	}
	return result
}

func (s Sliceint) GroupBystring(keySelector func(int) string) MapstringSliceint {
	m := make(MapstringSliceint)
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

// Mapstring returns a new slice resulting from the application of a given function to
// each element of a given slice.
func (s Sliceint) Mapstring(f func(int) string) Slicestring {
	output := make([]string, len(s))
	for i, a := range s {
		output[i] = f(a)
	}
	return output
}
