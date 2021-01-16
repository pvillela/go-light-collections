// Code generated -- DO NOT EDIT.

package coll

// FlatMapint returns the slice obtained by applying the argument f to each item in the
// receiver and concatenating the results.
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

// GroupByint returns a map whose keys are outputs of the keySelector function applied to
// the items in the receiver and whose values are slices containing the items in the
// receiver that correspond to each key obtained with the keySelector function.
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

func (s SliceFoo) Zipint(other Sliceint) SliceOfPairFooint {
	size := len(s)
	if size > len(other) {
		size = len(other)
	}
	r := make([]PairFooint, size)
	for i := 0; i < size; i++ {
		r[i] = PairFooint{s[i], other[i]}
	}
	return r
}

// ToMap returns a map whose keys are the first components in the items of the receiver and
// whose values are the corresonding second components in the items of the receiver.
// If multiple items in the receiver have the same first component, the corresponding
// value in the resulting map will be taken from the last such item in the receiver.
func (s SliceOfPairFooint) ToMap() MapFooint {
	m := make(map[Foo]int, len(s))
	for _, p := range s {
		m[p.X1] = p.X2
	}
	return m
}
