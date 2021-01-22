// Code generated -- DO NOT EDIT.

package collections

func (m Mapintstring) FlatMapint(f func(Pairintstring) Sliceint) Sliceint {
	if m == nil {
		return nil
	}
	r := make(Sliceint, 0, len(m)) // optimizing for speed vs space
	for k, v := range m {
		r = append(r, f(Pairintstring{k, v})...)
	}
	return r
}

func (m Mapintstring) Mapint(f func(Pairintstring) int) Sliceint {
	if m == nil {
		return nil
	}
	r := make(Sliceint, 0, len(m)) // optimizing for speed vs space
	for k, v := range m {
		r = append(r, f(Pairintstring{k, v}))
	}
	return r
}

func (m Mapintstring) MapKeysint(f func(Pairintstring) int) map[int]string {
	if m == nil {
		return nil
	}
	r := make(map[int]string)
	for k, v := range m {
		r[f(Pairintstring{k, v})] = v
	}
	return r
}

func (m Mapintstring) MapValuesint(f func(Pairintstring) int) map[int]int {
	if m == nil {
		return nil
	}
	r := make(map[int]int)
	for k, v := range m {
		r[k] = f(Pairintstring{k, v})
	}
	return r
}
