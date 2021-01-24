package internal

import c "github.com/pvillela/go-light-collections/pkg/glc"

/////////////////////
// This file is used to define the intended methods to be implemented by certain
// c.MapX types and to check that the implementations conform to the intention.
// Nothing here is exported.

// IMapT0T1 defines the methods to be implemented by the concrete type c.MapT0T1 that only
// depend on types c.T0 and c.T1.
type IMapT0T1 interface {
	Copy() c.MapT0T1
	Entries() []c.PairMpT0T1
	Keys() []c.T0
	Length() int
	Size() int
	Values() []c.T1
	ContainsKey(c.T0) bool
	ContainsValue(c.T1) bool
	Count(func(c.PairMpT0T1) bool) int
	Get(k c.T0) (c.T1, bool)
	IsEmpty() bool
	All(func(c.PairMpT0T1) bool) bool
	Any(func(c.PairMpT0T1) bool) bool
	ToSlice() []c.PairMpT0T1
	Filter(func(c.PairMpT0T1) bool) c.MapT0T1
	FilterKeys(func(c.T0) bool) c.MapT0T1
	FilterNot(func(c.PairMpT0T1) bool) c.MapT0T1
	FilterValues(func(c.T1) bool) c.MapT0T1
	ForEach(func(c.PairMpT0T1))
	GetOrElse(c.T0, func(c.T0) c.T1) c.T1
	IsNotEmpty() bool
	MaxWith(func(c.PairMpT0T1, c.PairMpT0T1) int) (c.PairMpT0T1, error)
	MinusKey(c.T0) c.MapT0T1
	MinusKeys([]c.T0) c.MapT0T1
	MinWith(func(c.PairMpT0T1, c.PairMpT0T1) int) (c.PairMpT0T1, error)
	PlusEntry(c.PairMpT0T1) c.MapT0T1
	PlusMap(c.MapT0T1) c.MapT0T1
	PlusSlice([]c.PairMpT0T1) c.MapT0T1
	Add(k c.T0, v c.T1) c.MapT0T1
}

// IMapT0T1T2 defines the methods to be implemented by the concrete type c.MapT0T1 that also
// depend on type c.T2.
type IMapT0T1T2 interface {
	FlatMapT2(func(c.PairMpT0T1) []c.T2) []c.T2
	MapT2(func(c.PairMpT0T1) c.T2) []c.T2
	MapKeysT2(func(c.PairMpT0T1) c.T2) map[c.T2]c.T1
	MapValuesT2(func(c.PairMpT0T1) c.T2) map[c.T0]c.T2
}

// Check that the concrete type satisfies the interfaces.
func validateMapInterface(m c.MapT0T1) {
	f := func(itf IMapT0T1) {}
	f(m)
	g := func(itf IMapT0T1T2) {}
	g(m)
}
