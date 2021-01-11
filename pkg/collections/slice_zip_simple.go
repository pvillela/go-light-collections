package collections

/////////////////////
// Simple names of methods without mangling.

func (s SliceT0) Zip(other SliceT1) SliceOfPairT0T1 {
	return s.ZipT1(other)
}
