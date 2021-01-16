// Code generated -- DO NOT EDIT.

package coll

// FlatMapString returns the slice obtained by applying the argument f to each item in the
// receiver and concatenating the results.
func (s SliceFoo) FlatMapString(f func(Foo) SliceString) SliceString {
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
func (s SliceFoo) FoldString(z String, op func(String, Foo) String) String {
	result := z
	for _, a := range s {
		result = op(result, a)
	}
	return result
}

// GroupByString returns a map whose keys are outputs of the keySelector function applied to
// the items in the receiver and whose values are slices containing the items in the
// receiver that correspond to each key obtained with the keySelector function.
func (s SliceFoo) GroupByString(keySelector func(Foo) String) MapStringSliceFoo {
	m := make(MapStringSliceFoo)
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

// MapString returns a new slice resulting from the application of a given function to
// each element of a given slice.
func (s SliceFoo) MapString(f func(Foo) String) SliceString {
	output := make([]String, len(s))
	for i, a := range s {
		output[i] = f(a)
	}
	return output
}

func (s SliceFoo) ZipString(other SliceString) SliceOfPairFooString {
	size := len(s)
	if size > len(other) {
		size = len(other)
	}
	r := make([]PairFooString, size)
	for i := 0; i < size; i++ {
		r[i] = PairFooString{s[i], other[i]}
	}
	return r
}

// ToMap returns a map whose keys are the first components in the items of the receiver and
// whose values are the corresonding second components in the items of the receiver.
// If multiple items in the receiver have the same first component, the corresponding
// value in the resulting map will be taken from the last such item in the receiver.
func (s SliceOfPairFooString) ToMap() MapFooString {
	m := make(map[Foo]String, len(s))
	for _, p := range s {
		m[p.X1] = p.X2
	}
	return m
}
