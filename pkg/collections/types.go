package collections

// T0 is a type wrapper.
// Used to clarify method signatures and facilitate replacement for code generation.
type T0 = interface{}

// Any is a type alias.
type Any = T0

// T1 is a type alias.
// Used to clarify method signatures and facilitate replacement for code generation.
type T1 = Any

// T2 is a type alias.
// Used to clarify method signatures and facilitate replacement for code generation.
type T2 = Any

// SliceAny is a type alias.
type SliceAny = SliceT0

// SliceT1 is a type alias.
// Used to clarify method signatures and facilitate replacement for code generation.
type SliceT1 = SliceT0

// SliceT2 is a type alias.
// Used to clarify method signatures and facilitate replacement for code generation.
type SliceT2 = SliceT0

// Slice2T0 is a type wrapper.
type Slice2T0 []SliceT0

// PairT0T1 defines a pseudo-generic pair.
type PairT0T1 struct {
	X1 T0
	X2 T1
}

// PairAnyAny is a type alias.
type PairAnyAny = PairT0T1

// SliceOfPairT0T1 is a type wrapper.
type SliceOfPairT0T1 []PairT0T1

// SliceOfPairAnyAny is a type alias.
type SliceOfPairAnyAny = SliceOfPairT0T1

// MapT0T1 is a type wrapper, implements Map interface.
type MapT0T1 map[T0]T1

// MapAnyAny is a type alias.
type MapAnyAny = MapT0T1

// MapT0T2 is a type wrapper, implements Map interface.
type MapT0T2 = MapT0T1

// MapT2T1 is a alias.
type MapT2T1 = MapT0T1

// MapT0SliceT1 is a type wrapper.
type MapT0SliceT1 map[T0]SliceT1

// MapAnySliceAny is a type alias.
type MapAnySliceAny = MapT0SliceT1

// SetT0 is a type wrapper, implements Set interface.
type SetT0 map[T0]bool

// SetAny is a type alias.
type SetAny = SetT0

// SetOfPairT0T1 is a type wrapper, implements Set interface.
type SetOfPairT0T1 map[PairT0T1]bool

// SetOfPairAnyAny is a type alias.
type SetOfPairAnyAny = SetOfPairT0T1
