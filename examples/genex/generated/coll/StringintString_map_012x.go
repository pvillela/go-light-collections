// Code generated -- DO NOT EDIT.

package coll

func (m MapStringint) MapKeysString(f func(PairMpStringint) String) map[String]int {
	if m == nil {
		return nil
	}
	r := make(map[String]int)
	for k, v := range m {
		r[f(PairMpStringint{k, v})] = v
	}
	return r
}
