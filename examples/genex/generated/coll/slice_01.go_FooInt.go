// Code generated -- DO NOT EDIT.

package coll

func (s SliceFoo) FlatMapInt(f func(Foo) SliceInt) SliceInt {
	r := make([]Int, 0, len(s)) // optimizing for speed vs space
	for _, x := range s {
		r = append(r, f(x)...)
	}
	return r
}

// FoldInt returns the accumulated value obtained by applying the operation op to z,
// and the first element of the given slice, then applying op to the result of the first
// operation and the second element of the given slice, and so on.
// Returns z if the slice is empty.
func (s SliceFoo) FoldInt(z Int, op func(Int, Foo) Int) Int {
	result := z
	for _, a := range s {
		result = op(result, a)
	}
	return result
}

func (s SliceFoo) GroupByInt(keySelector func(Foo) Int) MapIntSliceFoo {
	m := make(MapIntSliceFoo)
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

// MapInt returns a new slice resulting from the application of a given function to
// each element of a given slice.
func (s SliceFoo) MapInt(f func(Foo) Int) SliceInt {
	output := make([]Int, len(s))
	for i, a := range s {
		output[i] = f(a)
	}
	return output
}
