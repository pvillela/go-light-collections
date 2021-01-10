package collections

func (s SliceOfPairT0T1) ToMap() MapT0T1 {
	m := make(map[T0]T1)
	for _, p := range s {
		m[p.X1] = p.X2
	}
	return m
}
