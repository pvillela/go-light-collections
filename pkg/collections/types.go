package collections

// AnyT is a type alias.
type AnyT = interface{}

// AnyT0 is a type alias.
// Used to clarify method signatures and facilitate replacement for code generation.
type AnyT0 = interface{}

// AnyT1 is a type alias.
// Used to clarify method signatures and facilitate replacement for code generation.
type AnyT1 = interface{}

// SliceT0 is a type wrapper, implements List interface.
// Used to clarify method signatures and facilitate replacement for code generation.
type SliceT0 []AnyT0

// SliceT is a type alias.
type SliceT = SliceT0

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

// PairT is a type alias.
type PairT = PairT01

// SliceTPair01 is a type wrapper.
type SliceTPair01 []PairT01

// MapT01 is a type wrapper, implements Map interface.
type MapT01 map[AnyT0]AnyT1

// MapT is a type alias.
type MapT = MapT01

// ConvertibleToSliceAny is used as a type for slice parameters in pseudo-polymorphic functions.
type ConvertibleToSliceAny interface {
	ToSliceAny() SliceT0
}

// FuncAnyAny is a type alias.
type FuncAnyAny = func(AnyT0) AnyT0

// FuncAnyBool is a type alias.
type FuncAnyBool = func(AnyT0) bool

// FuncAnyAnyAny is a type alias.
type FuncAnyAnyAny = func(AnyT0, AnyT0) AnyT0

// FuncAnyVoid is a type alias.
type FuncAnyVoid = func(AnyT0)
