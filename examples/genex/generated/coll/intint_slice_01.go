// Code generated -- DO NOT EDIT.

package coll

// FlatMapint returns the slice obtained by applying the argument f to each item in the
// receiver and concatenating the results.
func (s Sliceint) FlatMapint(f func(int) Sliceint) Sliceint {
	if s == nil {
		return nil
	}
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
func (s Sliceint) Foldint(z int, op func(int, int) int) int {
	result := z
	for _, a := range s {
		result = op(result, a)
	}
	return result
}

// GroupByint returns a map whose keys are outputs of the keySelector function applied to
// the items in the receiver and whose values are slices containing the items in the
// receiver that correspond to each key obtained with the keySelector function.
func (s Sliceint) GroupByint(keySelector func(int) int) MapintSliceint {
	if s == nil {
		return nil
	}
	m := make(MapintSliceint, len(s)/2) // optimizing for speed vs space
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

// Mapint returns a new slice resulting from the application of a given function to
// each element of a given slice.
func (s Sliceint) Mapint(f func(int) int) Sliceint {
	if s == nil {
		return nil
	}
	r := make(Sliceint, len(s))
	for i, a := range s {
		r[i] = f(a)
	}
	return r
}

func (s Sliceint) Zipint(other Sliceint) SliceOfPairintint {
	if s == nil {
		return nil
	}
	size := len(s)
	if size > len(other) {
		size = len(other)
	}
	r := make([]Pairintint, size)
	for i := 0; i < size; i++ {
		r[i] = Pairintint{s[i], other[i]}
	}
	return r
}

// ToMap returns a map whose keys are the first components in the items of the receiver and
// whose values are the corresonding second components in the items of the receiver.
// If multiple items in the receiver have the same first component, the corresponding
// value in the resulting map will be taken from the last such item in the receiver.
func (s SliceOfPairintint) ToMap() Mapintint {
	if s == nil {
		return nil
	}
	m := make(map[int]int, len(s))
	for _, p := range s {
		m[p.X1] = p.X2
	}
	return m
}
