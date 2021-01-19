// Code generated -- DO NOT EDIT.

package coll

// FlatMapint returns the slice obtained by applying the argument f to each item in the
// receiver and concatenating the results.
func (s SlicePerson) FlatMapint(f func(Person) Sliceint) Sliceint {
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
func (s SlicePerson) Foldint(z int, op func(int, Person) int) int {
	result := z
	for _, a := range s {
		result = op(result, a)
	}
	return result
}

// GroupByint returns a map whose keys are outputs of the keySelector function applied to
// the items in the receiver and whose values are slices containing the items in the
// receiver that correspond to each key obtained with the keySelector function.
func (s SlicePerson) GroupByint(keySelector func(Person) int) MapintSlicePerson {
	m := make(MapintSlicePerson, len(s)/2) // optimizing for speed vs space
	for _, x := range s {
		k := keySelector(x)
		lst, ok := m[k]
		if !ok {
			lst = make(SlicePerson, 0, 1)
		}
		lst = append(lst, x)
		m[k] = lst
	}
	return m
}

// Mapint returns a new slice resulting from the application of a given function to
// each element of a given slice.
func (s SlicePerson) Mapint(f func(Person) int) Sliceint {
	r := make(Sliceint, len(s))
	for i, a := range s {
		r[i] = f(a)
	}
	return r
}

func (s SlicePerson) Zipint(other Sliceint) SliceOfPairPersonint {
	size := len(s)
	if size > len(other) {
		size = len(other)
	}
	r := make([]PairPersonint, size)
	for i := 0; i < size; i++ {
		r[i] = PairPersonint{s[i], other[i]}
	}
	return r
}

// ToMap returns a map whose keys are the first components in the items of the receiver and
// whose values are the corresonding second components in the items of the receiver.
// If multiple items in the receiver have the same first component, the corresponding
// value in the resulting map will be taken from the last such item in the receiver.
func (s SliceOfPairPersonint) ToMap() MapPersonint {
	m := make(map[Person]int, len(s))
	for _, p := range s {
		m[p.X1] = p.X2
	}
	return m
}
