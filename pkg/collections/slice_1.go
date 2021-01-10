package collections

func (s SliceT0) FlatMapAnyT1(f func(AnyT0) []AnyT1) []AnyT1 {
	return s.FlatMap(f)
}

// FoldAnyT1 returns the accumulated value obtained by applying the operation op to z,
// and the first element of the given slice, then applying op to the result of the first
// operation and the second element of the given slice, and so on.
// Returns z if the slice is empty.
func (s SliceT0) FoldAnyT1(z AnyT1, op func(AnyT1, AnyT0) AnyT1) AnyT1 {
	return s.Fold(z, op)
}

func (s SliceT0) GroupByAnyT1(keySelector func(AnyT0) AnyT1) map[AnyT0][]AnyT1 {
	return s.GroupBy(keySelector)
}

// MapAnyT0AnyT1 returns a new slice resulting from the application of a given function to
// each element of a given slice.
func (s SliceT0) MapAnyT1(f func(AnyT0) AnyT1) []AnyT1 {
	return s.Map(f)
}
