package internal

import c "github.com/pvillela/go-light-collections/pkg/collections"

/////////////////////
// This file is used to define the intended methods to be implemented by certain
// SliceX types and to check that the implementations conform to the intention.
// Nothing here is exported.

// ISliceT0 defines the methods to be implemented by the concrete type c.SliceT0 that only
// depend on type c.T0.
type ISliceT0 interface {
	Copy() c.SliceT0
	Length() int
	Size() int
	Contains(elem c.T0) bool
	ContainsSlice(elems c.SliceT0) bool
	Get(index int) (c.T0, bool)
	IndexOf(elem c.T0) int
	IsEmpty() bool
	LastIndexOf(elem c.T0) int
	SubSlice(fromIndex int, toIndex int) c.SliceT0
	All(pred func(c.T0) bool) bool
	Any(pred func(c.T0) bool) bool
	Count(pred func(c.T0) bool) int
	Drop(n int) c.SliceT0
	DropLast(n int) c.SliceT0
	DropLastWhile(pred func(c.T0) bool) c.SliceT0
	DropWhile(pred func(c.T0) bool) c.SliceT0
	Filter(pred func(c.T0) bool) c.SliceT0
	FilterNot(pred func(c.T0) bool) c.SliceT0
	First() (c.T0, error)
	ForEach(f func(c.T0))
	IndexOfFirst(pred func(c.T0) bool) int
	IndexOfLast(pred func(c.T0) bool) int
	IsNotEmpty() bool
	Last() (c.T0, error)
	MaxWith(comparator func(c.T0, c.T0) int) (c.T0, error)
	MinusSlice(other c.SliceT0) c.SliceT0
	MinusElement(elem c.T0) c.SliceT0
	MinWith(comparator func(c.T0, c.T0) int) (c.T0, error)
	Partition(pred func(c.T0) bool) (c.SliceT0, c.SliceT0)
	PlusElement(elem c.T0) c.SliceT0
	PlusSlice(other c.SliceT0) c.SliceT0
	Reduce(op func(c.T0, c.T0) c.T0) (c.T0, error)
	Reversed() c.SliceT0
	SortedWith(comparator func(c.T0, c.T0) int) c.SliceT0
	Take(n int) c.SliceT0
	TakeLast(n int) c.SliceT0
	TakeLastWhile(pred func(c.T0) bool) c.SliceT0
	TakeWhile(pred func(c.T0) bool) c.SliceT0
	ToSet() c.SetT0 // implemented in set_0.go
}

// ISliceT0T1 defines the methods to be implemented by the concrete type c.SliceT0 that
// depend on type T1.
type ISliceT0T1 interface {
	FlatMapT1(func(c.T0) c.SliceT1) c.SliceT1
	FoldT1(z c.T1, op func(c.T1, c.T0) c.T1) c.T1
	GroupByT1(keySelector func(c.T0) c.T1) c.MapT1SliceT0
	MapT1(f func(c.T0) c.T1) c.SliceT1
	ZipT1(other c.SliceT1) c.SliceOfPairT0T1
}

// ISlice2T0 defines the methods to be implemented by the concrete type Slice2T0.
type ISlice2T0 interface {
	Flatten() c.SliceT0
}

// ISliceOfPairT0T1 defines the methods to be implemented by the concrete type ISliceOfPairT0T1.
type ISliceOfPairT0T1 interface {
	ToMap() c.MapT0T1
}

// Check that the concrete type satisfies the interfaces.
func validateListInterface(s c.SliceT0) {
	f := func(itf ISliceT0) {}
	f(s)
	g := func(itf ISliceT0T1) {}
	g(s)
}

// Check that the concrete type satisfies the interface.
func validateListOfListInterface(s c.Slice2T0) {
	f := func(itf ISlice2T0) {}
	f(s)
}

// Check that the concrete type satisfies the interface.
func validateListOfPairInterface(s c.SliceOfPairT0T1) {
	f := func(itf ISliceOfPairT0T1) {}
	f(s)
}
