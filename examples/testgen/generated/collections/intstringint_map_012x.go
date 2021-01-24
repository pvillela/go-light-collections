// Code generated -- DO NOT EDIT.

package collections

func (m Mapintstring) MapKeysint(f func(PairMpintstring) int) map[int]string {
	if m == nil {
		return nil
	}
	r := make(map[int]string)
	for k, v := range m {
		r[f(PairMpintstring{k, v})] = v
	}
	return r
}
