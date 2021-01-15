package collections

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

// T3 is a type alias.
// Used to clarify method signatures and facilitate replacement for code generation.
type T3 = T0

// SliceAny is a type alias.
type SliceAny = SliceT0

// SliceT1 is a type alias.
// Used to clarify method signatures and facilitate replacement for code generation.
type SliceT1 = SliceT0

// SliceT2 is a type alias.
// Used to clarify method signatures and facilitate replacement for code generation.
type SliceT2 = SliceT0

// PairAnyAny is a type alias.
type PairAnyAny = PairT0T1

// SliceOfPairAnyAny is a type alias.
type SliceOfPairAnyAny = SliceOfPairT0T1

// MapAnyAny is a type alias.
type MapAnyAny = MapT0T1

// MapT0T2 is a type wrapper, implements Map interface.
type MapT0T2 = MapT0T1

// MapT2T1 is a alias.
type MapT2T1 = MapT0T1

// MapAnySliceAny is a type alias.
type MapAnySliceAny = MapT1SliceT0

// SetAny is a type alias.
type SetAny = SetT0

// SetT1 is a type alias.
type SetT1 = SetT0

// SetOfPairAnyAny is a type alias.
type SetOfPairAnyAny = SetOfPairT0T1
