package collections

func (m MapT0T1) FlatMapT2(f func(PairT0T1) SliceT2) SliceT2 {
	r := make(SliceT2, 0, len(m)) // optimizing for speed vs space
	for k, v := range m {
		r = append(r, f(PairT0T1{k, v})...)
	}
	return r
}

func (m MapT0T1) MapT2(f func(PairT0T1) T2) SliceT2 {
	r := make(SliceT2, 0, len(m)) // optimizing for speed vs space
	for k, v := range m {
		r = append(r, f(PairT0T1{k, v}))
	}
	return r
}

func (m MapT0T1) MapKeysT2(f func(PairT0T1) T2) MapT2T1 {
	r := make(MapT2T1)
	for k, v := range m {
		r[f(PairT0T1{k, v})] = v
	}
	return r
}

func (m MapT0T1) MapValuesT2(f func(PairT0T1) T2) MapT0T2 {
	r := make(MapT2T1)
	for k, v := range m {
		r[k] = f(PairT0T1{k, v})
	}
	return r
}
