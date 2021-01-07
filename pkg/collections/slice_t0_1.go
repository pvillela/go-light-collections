package collections

func (s SliceT0) FlatMap(f func(AnyT0) []AnyT1) []AnyT1 {
	r := make([]AnyT1, 0)
	for _, x := range s {
		r = append(r, f(x))
	}
	return r
}

// Fold returns the accumulated value obtained by applying the operation op to z,
// and the first element of the given slice, then applying op to the result of the first
// operation and the second element of the given slice, and so on.
func (s SliceT0) Fold(z AnyT1, op func(AnyT1, AnyT0) AnyT1) AnyT1 {
	result := z
	for _, a := range s {
		result = op(result, a)
	}
	return result
}

func (s SliceT0) GroupBy(keySelector func(AnyT0) AnyT1) map[AnyT0][]AnyT1 {
	m := make(map[AnyT0][]AnyT1)
	for _, x := range s {
		k := keySelector(x)
		lst, ok := m[k]
		if !ok {
			lst = make(SliceT1, 1)
			m[k] = lst
		}
		lst = append(lst, x)
	}
	return m
}

// Map returns a new slice resulting from the application of a given function to
// each element of a given slice.
func (s SliceT0) Map(f func(AnyT0) AnyT1) []AnyT1 {
	output := make([]AnyT1, len(s))
	for i, a := range s {
		output[i] = f(a)
	}
	return output
}
