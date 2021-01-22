// Code generated -- DO NOT EDIT.

package coll

func (m MapStringint) FlatMapint(f func(PairStringint) Sliceint) Sliceint {
	if m == nil {
		return nil
	}
	r := make(Sliceint, 0, len(m)) // optimizing for speed vs space
	for k, v := range m {
		r = append(r, f(PairStringint{k, v})...)
	}
	return r
}

func (m MapStringint) Mapint(f func(PairStringint) int) Sliceint {
	if m == nil {
		return nil
	}
	r := make(Sliceint, 0, len(m)) // optimizing for speed vs space
	for k, v := range m {
		r = append(r, f(PairStringint{k, v}))
	}
	return r
}

func (m MapStringint) MapKeysint(f func(PairStringint) int) Mapintint {
	if m == nil {
		return nil
	}
	r := make(Mapintint)
	for k, v := range m {
		r[f(PairStringint{k, v})] = v
	}
	return r
}

func (m MapStringint) MapValuesint(f func(PairStringint) int) MapStringint {
	if m == nil {
		return nil
	}
	r := make(MapStringint)
	for k, v := range m {
		r[k] = f(PairStringint{k, v})
	}
	return r
}
