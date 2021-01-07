package collections

// SliceInt is a wrapper type to enable extension methods.
type SliceInt []int

// Und converts to the underlying type.
func (s SliceInt) Und() []int { return s }

// ToSliceAny is defined to implement ConvertibleToSliceAny.
func (s SliceInt) ToSliceAny() SliceAny {
	r := make(SliceAny, len(s))
	for i, x := range s {
		r[i] = x
	}
	return r
}

// ToSliceInt is a conversion function.
func ToSliceInt(s SliceAny) SliceInt {
	r := make(SliceInt, len(s))
	for i, x := range s {
		r[i] = x.(int)
	}
	return r
}
