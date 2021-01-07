// Package pseudogenerics contains types and functions to facilitate functional-style
// programming in the absence of Go generics.
//
// Rather than rely on reflection, the approach supported by this package is based on:
//
// - Pseudo-generic functions that operate on type Any = interface{} and container and functional
// types composed from Any.  These include higher-order functions like map, filter, fold, etc.
//
// - Conversions from normal types, data structures, and functions to the types built on
// type Any, and vice-versa.
//
// This way, the only reflection construct used is type assertions (which are very fast) and
// the performance diadvantages of reflection are avoided.  On the other hand, some additional
// simple coding is required to create the aformentioned conversions.
//
// This approach is roughly as performant as recoding a function such as filter for slices
// for each underlying type ([]int, []string, []Foo, etc.), but the core algorith is coded
// only once (as a pseudo-generic function) and only simple conversion functions need to be
// coded for different underlying types.  The additional performancce overhead is mostly
// associated with copying slices, which in most cases will have low impact.  In cases of
// large slices with elements that are large data structures, using slices of pointers
// may be considered to minimize the overhead.
//
// The pattern exemplified in this package can be applied more broadly, beyond the functions
// and examples herein.
package pseudogenerics
