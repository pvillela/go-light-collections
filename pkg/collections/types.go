package collections

// Any is a type alias.
type Any = interface{}

// AnyT0 is a type alias.
// Used to clarify method signatures and facilitate replacement for code generation.
type AnyT0 = interface{}

// AnyT1 is a type alias.
// Used to clarify method signatures and facilitate replacement for code generation.
type AnyT1 = interface{}

// SliceT0 is a type wrapper, implements List interface.
type SliceT0 []AnyT0

// SliceAny is a type alias.
type SliceAny = SliceT0

// SliceT1 is a type alias.
// Used to clarify method signatures and facilitate replacement for code generation.
type SliceT1 = SliceT0

// Slice2T0 is a type wrapper.
type Slice2T0 [][]AnyT0

// PairT01 defines a pseudo-generic pair.
type PairT01 struct {
	x1 AnyT0
	x2 AnyT1
}

// PairAnyAny is a type alias.
type PairAnyAny = PairT01

// SliceTPair01 is a type wrapper.
type SliceTPair01 []PairT01

// MapT01 is a type wrapper, implements Map interface.
type MapT01 map[AnyT0]AnyT1

// MapAnyAny is a type alias.
type MapAnyAny = MapT01

// SetT0 is a type wrapper, implements Set interface.
type SetT0 map[AnyT0]bool

// SetAny is a type alias.
type SetAny = SetT0
