/*
 *  Copyright © 2021 Paulo Villela. All rights reserved.
 *  Use of this source code is governed by the Apache 2.0 license
 *  that can be found in the LICENSE file.
 */

package interfaces

import (
	"github.com/pvillela/go-light-collections/pkg/glc"
)

/////////////////////
// This file is used to define the intended methods to be implemented by certain
// SliceX types and to check that the implementations conform to the intention.
// Nothing here is exported.

// ISlice[T any] defines the methods to be implemented by the concrete type c.Slice[T any] that only
// depend on type c.T0.
type ISlice[T any] interface {
	Copy() glc.Slice[T]
	Length() int
	Size() int
	Contains(elem T) bool
	ContainsAll(elems glc.Slice[T]) bool
	Get(index int) (T, bool)
	IndexOf(elem T) int
	IsEmpty() bool
	LastIndexOf(elem T) int
	SubSlice(fromIndex int, toIndex int) glc.Slice[T]
	All(pred func(T) bool) bool
	Any(pred func(T) bool) bool
	Count(pred func(T) bool) int
	Drop(n int) glc.Slice[T]
	DropLast(n int) glc.Slice[T]
	DropLastWhile(pred func(T) bool) glc.Slice[T]
	DropWhile(pred func(T) bool) glc.Slice[T]
	Filter(pred func(T) bool) glc.Slice[T]
	FilterNot(pred func(T) bool) glc.Slice[T]
	First() (T, error)
	ForEach(f func(T))
	IndexOfFirst(pred func(T) bool) int
	IndexOfLast(pred func(T) bool) int
	IsNotEmpty() bool
	Last() (T, error)
	MaxWith(comparator func(T, T) int) (T, error)
	MinusSlice(other glc.Slice[T]) glc.Slice[T]
	MinusElement(elem T) glc.Slice[T]
	MinWith(comparator func(T, T) int) (T, error)
	Partition(pred func(T) bool) (glc.Slice[T], glc.Slice[T])
	PlusElement(elem T) glc.Slice[T]
	PlusSlice(other glc.Slice[T]) glc.Slice[T]
	Reduce(op func(T, T) T) (T, error)
	Reversed() glc.Slice[T]
	SortedWith(comparator func(T, T) int) glc.Slice[T]
	Take(n int) glc.Slice[T]
	TakeLast(n int) glc.Slice[T]
	TakeLastWhile(pred func(T) bool) glc.Slice[T]
	TakeWhile(pred func(T) bool) glc.Slice[T]
}

// Check that the concrete type satisfies the interfaces.
func validateListInterface[T any](s glc.Slice[T]) {
	f := func(itf ISlice[T]) {}
	f(s)
}
