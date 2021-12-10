/*
 * Copyright © 2021 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package internal

import (
	"github.com/pvillela/go-light-collections/legacy/pkg/glc"
)

/////////////////////
// This file is used to define the intended methods to be implemented by certain
// SliceX types and to check that the implementations conform to the intention.
// Nothing here is exported.

// ISliceT0 defines the methods to be implemented by the concrete type c.SliceT0 that only
// depend on type c.T0.
type ISliceT0 interface {
	Copy() glc.SliceT0
	Length() int
	Size() int
	Contains(elem glc.T0) bool
	ContainsAll(elems glc.SliceT0) bool
	Get(index int) (glc.T0, bool)
	IndexOf(elem glc.T0) int
	IsEmpty() bool
	LastIndexOf(elem glc.T0) int
	SubSlice(fromIndex int, toIndex int) glc.SliceT0
	All(pred func(glc.T0) bool) bool
	Any(pred func(glc.T0) bool) bool
	Count(pred func(glc.T0) bool) int
	Drop(n int) glc.SliceT0
	DropLast(n int) glc.SliceT0
	DropLastWhile(pred func(glc.T0) bool) glc.SliceT0
	DropWhile(pred func(glc.T0) bool) glc.SliceT0
	Filter(pred func(glc.T0) bool) glc.SliceT0
	FilterNot(pred func(glc.T0) bool) glc.SliceT0
	First() (glc.T0, error)
	ForEach(f func(glc.T0))
	IndexOfFirst(pred func(glc.T0) bool) int
	IndexOfLast(pred func(glc.T0) bool) int
	IsNotEmpty() bool
	Last() (glc.T0, error)
	MaxWith(comparator func(glc.T0, glc.T0) int) (glc.T0, error)
	MinusSlice(other glc.SliceT0) glc.SliceT0
	MinusElement(elem glc.T0) glc.SliceT0
	MinWith(comparator func(glc.T0, glc.T0) int) (glc.T0, error)
	Partition(pred func(glc.T0) bool) (glc.SliceT0, glc.SliceT0)
	PlusElement(elem glc.T0) glc.SliceT0
	PlusSlice(other glc.SliceT0) glc.SliceT0
	Reduce(op func(glc.T0, glc.T0) glc.T0) (glc.T0, error)
	Reversed() glc.SliceT0
	SortedWith(comparator func(glc.T0, glc.T0) int) glc.SliceT0
	Take(n int) glc.SliceT0
	TakeLast(n int) glc.SliceT0
	TakeLastWhile(pred func(glc.T0) bool) glc.SliceT0
	TakeWhile(pred func(glc.T0) bool) glc.SliceT0
	ToSet() map[glc.T0]bool
}

// ISliceT0T1 defines the methods to be implemented by the concrete type c.SliceT0 that
// depend on type T1.
type ISliceT0T1 interface {
	FlatMapT1(func(glc.T0) []glc.T1) []glc.T1
	FoldT1(z glc.T1, op func(glc.T1, glc.T0) glc.T1) glc.T1
	GroupByT1(keySelector func(glc.T0) glc.T1) map[glc.T1]glc.SliceT0
	MapT1(f func(glc.T0) glc.T1) []glc.T1
	ZipT1(other []glc.T1) []glc.PairSlT0T1
}

// ISlice2T0 defines the methods to be implemented by the concrete type Slice2T0.
type ISlice2T0 interface {
	Flatten() glc.SliceT0
}

// ISliceOfPairT0T1 defines the methods to be implemented by the concrete type ISliceOfPairT0T1.
type ISliceOfPairT0T1 interface {
	ToMap() map[glc.T0]glc.T1
}

// Check that the concrete type satisfies the interfaces.
func validateListInterface(s glc.SliceT0) {
	f := func(itf ISliceT0) {}
	f(s)
	g := func(itf ISliceT0T1) {}
	g(s)
}

// Check that the concrete type satisfies the interface.
func validateListOfListInterface(s glc.Slice2T0) {
	f := func(itf ISlice2T0) {}
	f(s)
}

// Check that the concrete type satisfies the interface.
func validateListOfPairInterface(s glc.SliceOfPairT0T1) {
	f := func(itf ISliceOfPairT0T1) {}
	f(s)
}
