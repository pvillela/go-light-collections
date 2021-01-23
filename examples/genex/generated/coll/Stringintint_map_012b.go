// Code generated -- DO NOT EDIT.

package coll

func (m MapStringint) MapKeysint(f func(PairMpStringint) int) map[int]int {
	if m == nil {
		return nil
	}
	r := make(map[int]int)
	for k, v := range m {
		r[f(PairMpStringint{k, v})] = v
	}
	return r
}
