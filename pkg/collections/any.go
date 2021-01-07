package collections

// AnyT0 is a type alias (a wrapper for interface{} behaves as an alias).
// Used to clarify method signatures and facilitate replacement for code generation.
type AnyT0 interface{}

// AnyT1 is a type alias (a wrapper for interface{} behaves as an alias).
// Used to clarify method signatures and facilitate replacement for code generation.
type AnyT1 interface{}

// AnyT2 is a type alias (a wrapper for interface{} behaves as an alias).
// Used to clarify method signatures and facilitate replacement for code generation.
type AnyT2 interface{}

// AnyT3 is a type alias (a wrapper for interface{} behaves as an alias).
// Used to clarify method signatures and facilitate replacement for code generation.
type AnyT3 interface{}

// AnyT4 is a type alias (a wrapper for interface{} behaves as an alias).
// Used to clarify method signatures and facilitate replacement for code generation.
type AnyT4 interface{}

// AnyT5 is a type alias (a wrapper for interface{} behaves as an alias).
// Used to clarify method signatures and facilitate replacement for code generation.
type AnyT5 interface{}

// FuncAnyAny is a type alias.
type FuncAnyAny = func(AnyT0) AnyT0

// FuncAnyBool is a type alias.
type FuncAnyBool = func(AnyT0) bool

// FuncAnyAnyAny is a type alias.
type FuncAnyAnyAny = func(AnyT0, AnyT0) AnyT0

// FuncAnyVoid is a type alias.
type FuncAnyVoid = func(AnyT0)
