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
// c.MapX types and to check that the implementations conform to the intention.
// Nothing here is exported.

// IMapT0T1 defines the methods to be implemented by the concrete type c.MapT0T1 that only
// depend on types c.T0 and c.T1.
type IMapT0T1 interface {
	Copy() glc.MapT0T1
	Entries() []glc.PairMpT0T1
	Keys() []glc.T0
	Length() int
	Size() int
	Values() []glc.T1
	ContainsKey(glc.T0) bool
	ContainsValue(glc.T1) bool
	Count(func(glc.PairMpT0T1) bool) int
	Get(k glc.T0) (glc.T1, bool)
	IsEmpty() bool
	All(func(glc.PairMpT0T1) bool) bool
	Any(func(glc.PairMpT0T1) bool) bool
	ToSlice() []glc.PairMpT0T1
	Filter(func(glc.PairMpT0T1) bool) glc.MapT0T1
	FilterKeys(func(glc.T0) bool) glc.MapT0T1
	FilterNot(func(glc.PairMpT0T1) bool) glc.MapT0T1
	FilterValues(func(glc.T1) bool) glc.MapT0T1
	ForEach(func(glc.PairMpT0T1))
	GetOrElse(glc.T0, func(glc.T0) glc.T1) glc.T1
	IsNotEmpty() bool
	MaxWith(func(glc.PairMpT0T1, glc.PairMpT0T1) int) (glc.PairMpT0T1, error)
	MinusKey(glc.T0) glc.MapT0T1
	MinusKeys([]glc.T0) glc.MapT0T1
	MinWith(func(glc.PairMpT0T1, glc.PairMpT0T1) int) (glc.PairMpT0T1, error)
	PlusEntry(glc.PairMpT0T1) glc.MapT0T1
	PlusMap(glc.MapT0T1) glc.MapT0T1
	PlusSlice([]glc.PairMpT0T1) glc.MapT0T1
	Add(k glc.T0, v glc.T1) glc.MapT0T1
}

// IMapT0T1T2 defines the methods to be implemented by the concrete type c.MapT0T1 that also
// depend on type c.T2.
type IMapT0T1T2 interface {
	FlatMapT2(func(glc.PairMpT0T1) []glc.T2) []glc.T2
	MapT2(func(glc.PairMpT0T1) glc.T2) []glc.T2
	MapKeysT2(func(glc.PairMpT0T1) glc.T2) map[glc.T2]glc.T1
	MapValuesT2(func(glc.PairMpT0T1) glc.T2) map[glc.T0]glc.T2
}

// Check that the concrete type satisfies the interfaces.
func validateMapInterface(m glc.MapT0T1) {
	f := func(itf IMapT0T1) {}
	f(m)
	g := func(itf IMapT0T1T2) {}
	g(m)
}
