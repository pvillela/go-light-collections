package glc

// FlatMapT1 returns the slice obtained by applying the argument f to each item in the
// receiver and concatenating the results.
func (s SliceT0) FlatMapT1(f func(T0) SliceT1) SliceT1 {
	if s == nil {
		return nil
	}
	r := make([]T1, 0, len(s)) // optimizing for speed vs space
	for _, x := range s {
		r = append(r, f(x)...)
	}
	return r
}

// FoldT1 returns the accumulated value obtained by applying the operation op to z,
// and the first element of the given slice, then applying op to the result of the first
// operation and the second element of the given slice, and so on.
// Returns z if the slice is empty.
func (s SliceT0) FoldT1(z T1, op func(T1, T0) T1) T1 {
	result := z
	for _, a := range s {
		result = op(result, a)
	}
	return result
}

// GroupByT1 returns a map whose keys are outputs of the keySelector function applied to
// the items in the receiver and whose values are slices containing the items in the
// receiver that correspond to each key obtained with the keySelector function.
func (s SliceT0) GroupByT1(keySelector func(T0) T1) map[T1]SliceT0 {
	if s == nil {
		return nil
	}
	m := make(map[T1]SliceT0, len(s)/2) // optimizing for speed vs space
	for _, x := range s {
		k := keySelector(x)
		lst, ok := m[k]
		if !ok {
			lst = make(SliceT0, 0, 1)
		}
		lst = append(lst, x)
		m[k] = lst
	}
	return m
}

// MapT1 returns a new slice resulting from the application of a given function to
// each element of a given slice.
func (s SliceT0) MapT1(f func(T0) T1) SliceT1 {
	if s == nil {
		return nil
	}
	r := make(SliceT1, len(s))
	for i, a := range s {
		r[i] = f(a)
	}
	return r
}

func (s SliceT0) ZipT1(other SliceT1) SliceOfPairT0T1 {
	if s == nil {
		return nil
	}
	size := len(s)
	if size > len(other) {
		size = len(other)
	}
	r := make([]PairT0T1, size)
	for i := 0; i < size; i++ {
		r[i] = PairT0T1{s[i], other[i]}
	}
	return r
}

// ToMap returns a map whose keys are the first components in the items of the receiver and
// whose values are the corresonding second components in the items of the receiver.
// If multiple items in the receiver have the same first component, the corresponding
// value in the resulting map will be taken from the last such item in the receiver.
func (s SliceOfPairT0T1) ToMap() MapT0T1 {
	if s == nil {
		return nil
	}
	m := make(map[T0]T1, len(s))
	for _, p := range s {
		m[p.X1] = p.X2
	}
	return m
}
