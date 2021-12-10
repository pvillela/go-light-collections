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
// c.SetX types and to check that the implementations conform to the intention.
// Nothing here is exported.

// ISetT0 defines the methods to be implemented by the concrete type c.SetT0 that only
// depend on type c.T0.
type ISetT0 interface {
	Copy() glc.SetT0
	Length() int
	Size() int
	All(func(glc.T0) bool) bool
	Any(func(glc.T0) bool) bool
	Contains(glc.T0) bool
	ContainsSet(glc.SetT0) bool
	ContainsSlice([]glc.T0) bool
	Count(func(glc.T0) bool) int
	Filter(func(glc.T0) bool) glc.SetT0
	FilterNot(func(glc.T0) bool) glc.SetT0
	ForEach(func(glc.T0))
	Intersect(glc.SetT0) glc.SetT0
	IsEmpty() bool
	IsNotEmpty() bool
	MaxWith(func(glc.T0, glc.T0) int) (glc.T0, error)
	MinusElement(glc.T0) glc.SetT0
	MinusSet(glc.SetT0) glc.SetT0
	MinWith(func(glc.T0, glc.T0) int) (glc.T0, error)
	Partition(pred func(glc.T0) bool) (glc.SetT0, glc.SetT0)
	PlusElement(glc.T0) glc.SetT0
	PlusSet(glc.SetT0) glc.SetT0
	PlusSlice([]glc.T0) glc.SetT0
	ToSlice() []glc.T0
	Put(glc.T0)
}

// ISetT0T1 defines the methods to be implemented by the concrete type c.SetT0 that also
// depend on type c.T1.
type ISetT0T1 interface {
	FlatMapT1(func(glc.T0) map[glc.T1]bool) map[glc.T1]bool
	GroupByT1(keySelector func(glc.T0) glc.T1) map[glc.T1]glc.SetT0
	MapT1(f func(glc.T0) glc.T1) map[glc.T1]bool
}

// ISetOfPairT0T1 defines the methods to be implemented by the concrete type ISliceOfPairT0T1.
type ISetOfPairT0T1 interface {
	ToMap() glc.MapT0T1
}

// Check that the concrete type satisfies the interfaces.
func validateSetInterface(s glc.SetT0) {
	f := func(itf ISetT0) {}
	f(s)
	g := func(itf ISetT0T1) {}
	g(s)
}

// Check that the concrete type satisfies the interface.
func validateSetOfPairInterface(s glc.SetOfPairT0T1) {
	f := func(itf ISetOfPairT0T1) {}
	f(s)
}
