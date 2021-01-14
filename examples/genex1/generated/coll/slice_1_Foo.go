// Code generated -- DO NOT EDIT.

package coll

func (s SliceFoo) FlatMapint(f func(Foo) Sliceint) Sliceint {
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
func (s SliceFoo) Foldint(z int, op func(int, Foo) int) int {
	result := z
	for _, a := range s {
		result = op(result, a)
	}
	return result
}

func (s SliceFoo) GroupByint(keySelector func(Foo) int) MapintSliceFoo {
	m := make(MapintSliceFoo)
	for _, x := range s {
		k := keySelector(x)
		lst, ok := m[k]
		if !ok {
			lst = make(SliceFoo, 0, 1)
		}
		lst = append(lst, x)
		m[k] = lst
	}
	return m
}

// Mapint returns a new slice resulting from the application of a given function to
// each element of a given slice.
func (s SliceFoo) Mapint(f func(Foo) int) Sliceint {
	output := make([]int, len(s))
	for i, a := range s {
		output[i] = f(a)
	}
	return output
}
