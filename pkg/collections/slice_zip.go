package collections

func (s SliceT0) ZipT1(other SliceT1) SliceOfPairT0T1 {
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
