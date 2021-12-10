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
// This file is used to define the intended methods to be implemented by glc.Map.

// IMap[T0, T1] defines the methods to be implemented by a Map collection.
type IMap[T0 comparable, T1 any] interface {
	Copy() glc.Map[T0, T1]
	Entries() []glc.Pair[T0, T1]
	Keys() []T0
	Length() int
	Size() int
	Values() []T1
	ContainsKey(T0) bool
	Count(func(glc.Pair[T0, T1]) bool) int
	Get(k T0) (T1, bool)
	IsEmpty() bool
	All(func(glc.Pair[T0, T1]) bool) bool
	Any(func(glc.Pair[T0, T1]) bool) bool
	ToSlice() []glc.Pair[T0, T1]
	Filter(func(glc.Pair[T0, T1]) bool) glc.Map[T0, T1]
	FilterKeys(func(T0) bool) glc.Map[T0, T1]
	FilterNot(func(glc.Pair[T0, T1]) bool) glc.Map[T0, T1]
	FilterValues(func(T1) bool) glc.Map[T0, T1]
	ForEach(func(glc.Pair[T0, T1]))
	GetOrElse(T0, func(T0) T1) T1
	IsNotEmpty() bool
	MaxWith(func(glc.Pair[T0, T1], glc.Pair[T0, T1]) int) (glc.Pair[T0, T1], error)
	MinusKey(T0) glc.Map[T0, T1]
	MinusKeys([]T0) glc.Map[T0, T1]
	MinWith(func(glc.Pair[T0, T1], glc.Pair[T0, T1]) int) (glc.Pair[T0, T1], error)
	PlusEntry(glc.Pair[T0, T1]) glc.Map[T0, T1]
	PlusMap(glc.Map[T0, T1]) glc.Map[T0, T1]
	PlusSlice([]glc.Pair[T0, T1]) glc.Map[T0, T1]
	Add(k T0, v T1) glc.Map[T0, T1]
}

// Check that the concrete type satisfies the interfaces.
func validateMapInterface[T0 comparable, T1 any](m glc.Map[T0, T1]) {
	f := func(itf IMap[T0, T1]) {}
	f(m)
}
