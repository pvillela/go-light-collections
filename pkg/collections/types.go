package collections

// Any is a type alias.
type Any = interface{}

// T0 is a type alias.
// Used to clarify method signatures and facilitate replacement for code generation.
type T0 = interface{}

// T1 is a type alias.
// Used to clarify method signatures and facilitate replacement for code generation.
type T1 = interface{}

// SliceT0 is a type wrapper, implements List interface.
type SliceT0 []T0

// Slice is a type alias.
type Slice = SliceT0

// SliceT1 is a type alias.
// Used to clarify method signatures and facilitate replacement for code generation.
type SliceT1 = SliceT0

// Slice2T0 is a type wrapper.
type Slice2T0 [][]T0

// PairT0T1 defines a pseudo-generic pair.
type PairT0T1 struct {
	X1 T0
	X2 T1
}

// Pair is a type alias.
type Pair = PairT0T1

// SliceOfPairT0T1 is a type wrapper.
type SliceOfPairT0T1 []PairT0T1

// MapT0T1 is a type wrapper, implements Map interface.
type MapT0T1 map[T0]T1

// Map is a type alias.
type Map = MapT0T1

// SetT0 is a type wrapper, implements Set interface.
type SetT0 map[T0]bool

// Set is a type alias.
type Set = SetT0
