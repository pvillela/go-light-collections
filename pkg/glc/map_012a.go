package glc

func (m MapT0T1) FlatMapT2(f func(PairMpT0T1) []T2) []T2 {
	if m == nil {
		return nil
	}
	r := make([]T2, 0, len(m)) // optimizing for speed vs space
	for k, v := range m {
		r = append(r, f(PairMpT0T1{k, v})...)
	}
	return r
}

func (m MapT0T1) MapT2(f func(PairMpT0T1) T2) []T2 {
	if m == nil {
		return nil
	}
	r := make([]T2, 0, len(m)) // optimizing for speed vs space
	for k, v := range m {
		r = append(r, f(PairMpT0T1{k, v}))
	}
	return r
}

func (m MapT0T1) MapValuesT2(f func(PairMpT0T1) T2) map[T0]T2 {
	if m == nil {
		return nil
	}
	r := make(map[T0]T2)
	for k, v := range m {
		r[k] = f(PairMpT0T1{k, v})
	}
	return r
}
