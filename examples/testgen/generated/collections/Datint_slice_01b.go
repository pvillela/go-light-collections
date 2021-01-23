// Code generated -- DO NOT EDIT.

package collections

// GroupByint returns a map whose keys are outputs of the keySelector function applied to
// the items in the receiver and whose values are slices containing the items in the
// receiver that correspond to each key obtained with the keySelector function.
func (s SliceDat) GroupByint(keySelector func(Dat) int) map[int]SliceDat {
	if s == nil {
		return nil
	}
	m := make(map[int]SliceDat, len(s)/2) // optimizing for speed vs space
	for _, x := range s {
		k := keySelector(x)
		lst, ok := m[k]
		if !ok {
			lst = make(SliceDat, 0, 1)
		}
		lst = append(lst, x)
		m[k] = lst
	}
	return m
}
