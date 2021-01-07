package collections

func (s SliceTPair01) ToMap() map[AnyT0]AnyT1 {
	m := make(map[AnyT0]AnyT1)
	for _, p := range s {
		m[p.x1] = p.x2
	}
	return m
}
