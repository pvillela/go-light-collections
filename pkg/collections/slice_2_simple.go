package collections

/////////////////////
// Simple names of methods without mangling.

func (s SliceT0) Zip(other SliceT1) SliceTPair01 {
	size := len(s)
	if size > len(other) {
		size = len(other)
	}
	r := make([]PairT01, size)
	for i := 0; i < size; i++ {
		r[i] = PairT01{s[i], other[i]}
	}
	return r
}
