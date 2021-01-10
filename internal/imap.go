package internal

import . "github.com/pvillela/go-light-collections/pkg/collections"

/////////////////////
// This file is used to define the intended methods to be implemented by certain
// MapX types and to check that the implementations conform to the intention.
// Nothing here is exported.

// IMap is the interface for IMap operations.
type IMap interface {
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
	FlatMap(func(PairT0T1) SliceT2) SliceT2
	ForEach(func(PairT0T1))
	GetOrElse(func() T1) T1
	IsNotEmpty() bool
	Map(func(PairT0T1) T2) SliceT2
	MapKeys(func(T0) T2) MapT2T1
	MapValues(func(T1) T2) MapT0T2
	MaxWith(func(PairT0T1) int) MapT0T1
	MinusKey(T0) MapT0T1
	Minus(SliceT0) MapT0T1
	MinWith(func(PairT0T1) int) MapT0T1
	PlusEntry(PairT0T1) MapT0T1
	Plus(MapT0T1) MapT0T1
	PlusSlice(SliceOfPairT0T1) MapT0T1
	Put(k T0, v T0)
}
