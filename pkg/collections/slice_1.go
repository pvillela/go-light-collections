package collections

func (s SliceT0) FlatMapT1(f func(T0) SliceT1) SliceT1 {
	return s.FlatMap(f)
}

// FoldT1 returns the accumulated value obtained by applying the operation op to z,
// and the first element of the given slice, then applying op to the result of the first
// operation and the second element of the given slice, and so on.
// Returns z if the slice is empty.
func (s SliceT0) FoldT1(z T1, op func(T1, T0) T1) T1 {
	return s.Fold(z, op)
}

func (s SliceT0) GroupByT1(keySelector func(T0) T1) MapT0SliceT1 {
	return s.GroupBy(keySelector)
}

// MapT1 returns a new slice resulting from the application of a given function to
// each element of a given slice.
func (s SliceT0) MapT1(f func(T0) T1) SliceT1 {
	return s.Map(f)
}
