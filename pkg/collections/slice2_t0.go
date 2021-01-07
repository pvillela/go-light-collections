package collections

func (s Slice2T0) Flatten() SliceT0 {
	r := make([]AnyT0, 0)
	for _, x := range s {
		r = append(r, x)
	}
	return r
}
