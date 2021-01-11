package internal

import c "github.com/pvillela/go-light-collections/pkg/collections"

/////////////////////
// This file is used to define the intended methods to be implemented by certain
// c.MapX types and to check that the implementations conform to the intention.
// Nothing here is exported.

// IMapT0T1 defines the methods to be implemented by the concrete type c.MapT0T1 that only
// depend on types c.T0 and c.T1.
type IMapT0T1 interface {
	Entries() c.SetOfPairT0T1
	Keys() c.SetT0
	Count() int
	Values() c.SetT0
	ContainsKey(c.T0) bool
	ContainsValue(c.T1) bool
	Get(k c.T0) c.T1
	IsEmpty() bool
	All(pred func(c.T0) bool) bool
	Any(pred func(c.T0) bool) bool
	ToSlice() c.SliceOfPairT0T1
	Filter(func(c.PairT0T1) bool) c.MapT0T1
	FilterKeys(func(c.T0) bool) c.MapT0T1
	FilterNot(func(c.PairT0T1) bool) c.MapT0T1
	FilterValues(func(c.T1) bool) c.MapT0T1
	ForEach(func(c.PairT0T1))
	GetOrElse(func() c.T1) c.T1
	IsNotEmpty() bool
	MaxWith(func(c.PairT0T1) int) c.MapT0T1
	MinusKey(c.T0) c.MapT0T1
	Minus(c.SliceT0) c.MapT0T1
	MinWith(func(c.PairT0T1) int) c.MapT0T1
	PlusEntry(c.PairT0T1) c.MapT0T1
	Plus(c.MapT0T1) c.MapT0T1
	PlusSlice(c.SliceOfPairT0T1) c.MapT0T1
	Put(k c.T0, v c.T0)
}

// IMapT0T1T2 defines the methods to be implemented by the concrete type c.MapT0T1 that also
// depend on type c.T2.
type IMapT0T1T2 interface {
	FlatMap(func(c.PairT0T1) c.SliceT2) c.SliceT2
	Map(func(c.PairT0T1) c.T2) c.SliceT2
	MapKeys(func(c.T0) c.T2) c.MapT2T1
	MapValues(func(c.T1) c.T2) c.MapT0T2
}
