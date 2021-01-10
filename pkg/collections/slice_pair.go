package collections

func (s SliceTPair01) ToMap() map[AnyT0]AnyT1 {
	m := make(map[AnyT0]AnyT1)
	for _, p := range s {
		m[p.X1] = p.X2
	}
	return m
}
