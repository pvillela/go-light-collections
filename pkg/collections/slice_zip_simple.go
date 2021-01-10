package collections

/////////////////////
// Simple names of methods without mangling.

func (s SliceT0) Zip(other SliceT1) []PairT0T1 {
	size := len(s)
	if size > len(other) {
		size = len(other)
	}
	r := make([]PairT0T1, size)
	for i := 0; i < size; i++ {
		r[i] = PairT0T1{s[i], other[i]}
	}
	return r
}
