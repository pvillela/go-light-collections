package glc

// SliceString is a wrapper type to enable extension methods.
type SliceString []string

// ToSliceAny is defined to implement ConvertibleToSliceAny.
func (s SliceString) ToSliceAny() SliceT0 {
	r := make(SliceT0, len(s))
	for i, x := range s {
		r[i] = x
	}
	return r
}

// ToSliceString is a conversion function.
func ToSliceString(s SliceT0) SliceString {
	r := make(SliceString, len(s))
	for i, x := range s {
		r[i] = x.(string)
	}
	return r
}
