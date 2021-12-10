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
// c.SetX types and to check that the implementations conform to the intention.
// Nothing here is exported.

// ISet[T] defines the methods to be implemented by the concrete type c.Set[T] that only
// depend on type c.T0.
type ISet[T comparable] interface {
	Copy() glc.Set[T]
	Length() int
	Size() int
	All(func(T) bool) bool
	Any(func(T) bool) bool
	Contains(T) bool
	ContainsSet(glc.Set[T]) bool
	ContainsSlice([]T) bool
	Count(func(T) bool) int
	Filter(func(T) bool) glc.Set[T]
	FilterNot(func(T) bool) glc.Set[T]
	ForEach(func(T))
	Intersect(glc.Set[T]) glc.Set[T]
	IsEmpty() bool
	IsNotEmpty() bool
	MaxWith(func(T, T) int) (T, error)
	MinusElement(T) glc.Set[T]
	MinusSet(glc.Set[T]) glc.Set[T]
	MinWith(func(T, T) int) (T, error)
	Partition(pred func(T) bool) (glc.Set[T], glc.Set[T])
	PlusElement(T) glc.Set[T]
	PlusSet(glc.Set[T]) glc.Set[T]
	PlusSlice([]T) glc.Set[T]
	ToSlice() []T
	Put(T)
}

// Check that the concrete type satisfies the interfaces.
func validateSetInterface[T comparable](s glc.Set[T]) {
	f := func(itf ISet[T]) {}
	f(s)
}
