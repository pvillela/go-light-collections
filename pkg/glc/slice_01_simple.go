package glc

/////////////////////
// Simple names of methods without mangling.

func (s SliceT0) FlatMap(f func(T0) []T1) []T1 {
	return s.FlatMapT1(f)
}

// Fold returns the accumulated value obtained by applying the operation op to z,
// and the first element of the given slice, then applying op to the result of the first
// operation and the second element of the given slice, and so on.
// Returns z if the slice is empty.
func (s SliceT0) Fold(z T1, op func(T1, T0) T1) T1 {
	return s.FoldT1(z, op)
}

func (s SliceT0) GroupBy(keySelector func(T0) T1) map[T1]SliceT0 {
	return s.GroupByT1(keySelector)
}

// Map returns a new slice resulting from the application of a given function to
// each element of a given slice.
func (s SliceT0) Map(f func(T0) T1) []T1 {
	return s.MapT1(f)
}

func (s SliceT0) Zip(other []T1) SliceOfPairT0T1 {
	return s.ZipT1(other)
}
