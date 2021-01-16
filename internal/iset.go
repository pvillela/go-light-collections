package internal

import c "github.com/pvillela/go-light-collections/pkg/collections"

/////////////////////
// This file is used to define the intended methods to be implemented by certain
// c.SetX types and to check that the implementations conform to the intention.
// Nothing here is exported.

// ISetT0 defines the methods to be implemented by the concrete type c.SetT0 that only
// depend on type c.T0.
type ISetT0 interface {
	Copy() c.SetT0
	Length() int
	Size() int
	All(func(c.T0) bool) bool
	Any(func(c.T0) bool) bool
	Contains(c.T0) bool
	ContainsSet(c.SetT0) bool
	ContainsSlice(c.SliceT0) bool
	Count(func(c.T0) bool) int
	Filter(func(c.T0) bool) c.SetT0
	FilterNot(func(c.T0) bool) c.SetT0
	ForEach(func(c.T0))
	Intersect(c.SetT0) c.SetT0
	IsEmpty() bool
	IsNotEmpty() bool
	MaxWith(func(c.T0, c.T0) int) (c.T0, error)
	MinusElement(c.T0) c.SetT0
	MinusSet(c.SetT0) c.SetT0
	MinWith(func(c.T0, c.T0) int) (c.T0, error)
	Partition(pred func(c.T0) bool) (c.SetT0, c.SetT0)
	PlusElement(c.T0) c.SetT0
	PlusSet(c.SetT0) c.SetT0
	PlusSlice(c.SliceT0) c.SetT0
	ToSlice() c.SliceT0
	Put(c.T0)
}

// ISetT0T1 defines the methods to be implemented by the concrete type c.SetT0 that also
// depend on type c.T1.
type ISetT0T1 interface {
	FlatMapT1(func(c.T0) c.SetT1) c.SetT1
	GroupByT1(keySelector func(c.T0) c.T1) c.MapT1SliceT0
	MapT1(f func(c.T0) c.T1) c.SetT1
}

// ISetOfPairT0T1 defines the methods to be implemented by the concrete type ISliceOfPairT0T1.
type ISetOfPairT0T1 interface {
	ToMap() c.MapT0T1
}

// Check that the concrete type satisfies the interfaces.
func validateSetInterface(s c.SetT0) {
	f := func(itf ISetT0) {}
	f(s)
	g := func(itf ISetT0T1) {}
	g(s)
}

// Check that the concrete type satisfies the interface.
func validateSetOfPairInterface(s c.SetOfPairT0T1) {
	f := func(itf ISetOfPairT0T1) {}
	f(s)
}
