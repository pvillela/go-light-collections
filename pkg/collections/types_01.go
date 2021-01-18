package collections

// PairT0T1 defines a pseudo-generic pair.
type PairT0T1 struct {
	X1 T0
	X2 T1
}

// SliceOfPairT0T1 is a type wrapper.
type SliceOfPairT0T1 []PairT0T1

// MapT0T1 is a type wrapper, implements Map interface.
type MapT0T1 map[T0]T1

// MapT1SliceT0 is a type wrapper.
type MapT1SliceT0 map[T1]SliceT0

// MapT1SetT0 is a type wrapper.
type MapT1SetT0 map[T1]SetT0

// SetOfPairT0T1 is a type wrapper, implements Set interface.
type SetOfPairT0T1 map[PairT0T1]bool
