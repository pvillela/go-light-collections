package glc

// PairT0T1 defines a pseudo-generic pair.
type PairT0T1 struct {
	X1 T0
	X2 T1
}

// MapT0T1 is a type wrapper, implements IMap interfaces.
type MapT0T1 map[T0]T1

// SliceOfPairT0T1 is a type wrapper.
type SliceOfPairT0T1 []PairT0T1

// SetOfPairT0T1 is a type wrapper, implements ISetOfPairT0T1 interface.
type SetOfPairT0T1 map[PairT0T1]bool
