// Package collections (this package) contains types and functions to facilitate
// functional-style programming in the absence of Go generics.
//
// Rather than rely on reflection, the approach supported by this package is based on
// "pseudo-generic" functions that operate on type Any = interface{} as well as
// container and types composed from Any.
// Such functions include the usual map, filter, fold, etc.
//
// There are two suggested approaches to use this framework.
//
// One approach is to create simple adapters to convert between specific types to the core framework types
// (Any, SliceAny, MapAnyAny, SetAny).
// This way, the only reflection construct used is type assertions (which are very fast) and
// the performance diadvantages of reflection are avoided.
// On the other hand, some additional simple coding is required to create the aformentioned
// conversions.
//
// This approach is not as performant as recoding or generating a function such as filter for
// slices of each underlying type ([]int, []string, []Foo, etc.), but the core algorith is coded
// only once (as a pseudo-generic function) and only simple conversion functions need to be
// coded for different concrete types.  The additional performancce overhead is mostly
// associated with copying slices to/from []interface{}, and likewise for maps and sets..
// In most cases, this will have low impact.
// In cases of large slices/maps/sets with elements that are large data structures, using
// slices/maps/sets of pointers may be considered to minimize the overhead.
//
// Use of the above first approach is shown in a detailed package example.
//
// The second approach is code generation.  This framework uses type aliases and wrappers
// to facilitate the replacement of the pseudo-generic types with specific concrete types
// in support of code generation.
package collections
