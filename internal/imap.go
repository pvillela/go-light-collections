package internal

import . "github.com/pvillela/go-light-collections/pkg/collections"

/////////////////////
// This file is used to define the intended methods to be implemented by certain
// MapX types and to check that the implementations conform to the intention.
// Nothing here is exported.

// IMapT0T1 defines the methods to be implemented by the concrete type MapT0T1 that only
// depend on types T0 and T1.
type IMapT0T1 interface {
	Entries() SetOfPairT0T1
	Keys() SetT0
	Count() int
	Values() SetT0
	ContainsKey(T0) bool
	ContainsValue(T1) bool
	Get(k T0) T1
	IsEmpty() bool
	All(pred func(T0) bool) bool
	Any(pred func(T0) bool) bool
	ToSlice() SliceOfPairT0T1
	Filter(func(PairT0T1) bool) MapT0T1
	FilterKeys(func(T0) bool) MapT0T1
	FilterNot(func(PairT0T1) bool) MapT0T1
	FilterValues(func(T1) bool) MapT0T1
	ForEach(func(PairT0T1))
	GetOrElse(func() T1) T1
	IsNotEmpty() bool
	MaxWith(func(PairT0T1) int) MapT0T1
	MinusKey(T0) MapT0T1
	Minus(SliceT0) MapT0T1
	MinWith(func(PairT0T1) int) MapT0T1
	PlusEntry(PairT0T1) MapT0T1
	Plus(MapT0T1) MapT0T1
	PlusSlice(SliceOfPairT0T1) MapT0T1
	Put(k T0, v T0)
}

// IMapT0T1T2 defines the methods to be implemented by the concrete type MapT0T1 that also
// depend on type T2.
type IMapT0T1T2 interface {
	FlatMap(func(PairT0T1) SliceT2) SliceT2
	Map(func(PairT0T1) T2) SliceT2
	MapKeys(func(T0) T2) MapT2T1
	MapValues(func(T1) T2) MapT0T2
}
