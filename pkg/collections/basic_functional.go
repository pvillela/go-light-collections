package collections

// MapSlice returns a new slice resulting from the application of a given function to
// each element of a given slice.
func MapSlice(sc ConvertibleToSliceAny, f func(AnyT0) AnyT0) SliceT0 {
	s := sc.ToSliceAny()
	output := make(SliceT0, len(s))
	for i, a := range s {
		output[i] = f(a)
	}
	return output
}

// FilterSlice returns a new slice containing only the elements in the given slice for which
// the application of the function pred returns true.
func FilterSlice(sc ConvertibleToSliceAny, pred func(AnyT0) bool) SliceT0 {
	s := sc.ToSliceAny()
	output := make(SliceT0, 0)
	for i, a := range s {
		if pred(s[i]) {
			output = append(output, a)
		}
	}
	return output
}

// FoldLeftSlice returns the accumulated value obtained by applying the operation op to z,
// and the first element of the given slice, then applying op to the result of the first
// operation and the second element of the given slice, and so on.
func FoldLeftSlice(sc ConvertibleToSliceAny, z AnyT0, op func(AnyT0, AnyT1) AnyT1) AnyT0 {
	s := sc.ToSliceAny()
	result := z
	for _, a := range s {
		result = op(result, a)
	}
	return result
}
