package collections

func (s Slice2T0) Flatten() SliceT0 {
	r := make([]T0, 0, len(s)) // optimizing for speed vs space
	for _, x := range s {
		r = append(r, x)
	}
	return r
}
