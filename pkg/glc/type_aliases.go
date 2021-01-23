package glc

// T0 is a type wrapper.
// Used to clarify method signatures and facilitate replacement for code generation.
type T0 = interface{}

// Any is a type alias.
type Any = T0

// T1 is a type alias.
// Used to clarify method signatures and facilitate replacement for code generation.
type T1 = T0

// T2 is a type alias.
// Used to clarify method signatures and facilitate replacement for code generation.
type T2 = T0

// SliceAny is a type alias.
type SliceAny = SliceT0

// PairAnyAny is a type alias.
type PairAnyAny = struct {
	X1 Any
	X2 Any
}

// SliceOfPairAnyAny is a type alias.
type SliceOfPairAnyAny = SliceOfPairT0T1

// MapAnyAny is a type alias.
type MapAnyAny = MapT0T1

// SetAny is a type alias.
type SetAny = SetT0

// SetOfPairAnyAny is a type alias.
type SetOfPairAnyAny = SetOfPairT0T1
