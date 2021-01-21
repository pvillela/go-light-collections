package glc

// PairT0T1 defines a pseudo-generic pair.
type PairT0T1 struct {
	X1 T0
	X2 T1
}

// SliceOfPairT0T1 is a type wrapper.
type SliceOfPairT0T1 []PairT0T1

// MapT0T1 is a type wrapper, implements Map interface.
type MapT0T1 map[T0]T1

// MapT0SliceT1 is a type wrapper.
type MapT0SliceT1 map[T0]SliceT1

// MapT0SetT1 is a type wrapper.
type MapT0SetT1 map[T0]SetT1

// SetOfPairT0T1 is a type wrapper, implements Set interface.
type SetOfPairT0T1 map[PairT0T1]bool
