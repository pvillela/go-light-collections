// Code generated -- DO NOT EDIT.

package coll

func (m MapStringint) FlatMapint(f func(PairMpStringint) []int) []int {
	if m == nil {
		return nil
	}
	r := make([]int, 0, len(m)) // optimizing for speed vs space
	for k, v := range m {
		r = append(r, f(PairMpStringint{k, v})...)
	}
	return r
}

func (m MapStringint) Mapint(f func(PairMpStringint) int) []int {
	if m == nil {
		return nil
	}
	r := make([]int, 0, len(m)) // optimizing for speed vs space
	for k, v := range m {
		r = append(r, f(PairMpStringint{k, v}))
	}
	return r
}

func (m MapStringint) MapValuesint(f func(PairMpStringint) int) map[String]int {
	if m == nil {
		return nil
	}
	r := make(map[String]int)
	for k, v := range m {
		r[k] = f(PairMpStringint{k, v})
	}
	return r
}
