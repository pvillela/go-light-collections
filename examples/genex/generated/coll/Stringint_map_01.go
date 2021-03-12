// Code generated -- DO NOT EDIT.

/*
 * Copyright © 2021 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package coll

import "errors"

// MapStringint is a type wrapper, implements IMap interfaces.
type MapStringint map[String]int

// PairMpStringint is a type alias used only in Map methods to avoid code generation issues.
type PairMpStringint = struct {
	X1 String
	X2 int
}

func (m MapStringint) Copy() MapStringint {
	if m == nil {
		return nil
	}
	m1 := make(MapStringint, len(m))
	for k, v := range m {
		m1[k] = v
	}
	return m1
}

func (m MapStringint) Entries() []PairMpStringint {
	if m == nil {
		return nil
	}
	entries := make([]PairMpStringint, len(m))
	i := 0
	for k, v := range m {
		entries[i] = PairMpStringint{k, v}
		i++
	}
	return entries
}

func (m MapStringint) Keys() []String {
	if m == nil {
		return nil
	}
	keys := make([]String, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

func (m MapStringint) Values() []int {
	if m == nil {
		return nil
	}
	values := make([]int, len(m))
	i := 0
	for _, v := range m {
		values[i] = v
		i++
	}
	return values
}

// Length returns the number of items in the receiver.
func (m MapStringint) Length() int {
	return len(m)
}

// Size returns the number of items in the receiver. Same as Length.
func (m MapStringint) Size() int {
	return len(m)
}

func (m MapStringint) ContainsKey(k String) bool {
	_, ok := m[k]
	return ok
}

func (m MapStringint) ContainsValue(v int) bool {
	for _, v1 := range m {
		if v1 == v {
			return true
		}
	}
	return false
}

// Count returns the number of entries in the receiver that satisfy the predicate.
func (m MapStringint) Count(pred func(PairMpStringint) bool) int {
	count := 0
	for k, v := range m {
		if pred(PairMpStringint{k, v}) {
			count++
		}
	}
	return count
}

func (m MapStringint) Get(k String) (int, bool) {
	v, ok := m[k]
	return v, ok
}

func (m MapStringint) IsEmpty() bool {
	return len(m) == 0
}

func (m MapStringint) All(pred func(PairMpStringint) bool) bool {
	for k, v := range m {
		if !pred(PairMpStringint{k, v}) {
			return false
		}
	}
	return true
}

func (m MapStringint) Any(pred func(PairMpStringint) bool) bool {
	for k, v := range m {
		if pred(PairMpStringint{k, v}) {
			return true
		}
	}
	return false
}

func (m MapStringint) ToSlice() []PairMpStringint {
	if m == nil {
		return nil
	}
	s := make([]PairMpStringint, len(m))
	i := 0
	for k, v := range m {
		s[i] = PairMpStringint{k, v}
		i++
	}
	return s
}

func (m MapStringint) Filter(pred func(PairMpStringint) bool) MapStringint {
	if m == nil {
		return nil
	}
	m1 := MapStringint{}
	for k, v := range m {
		if pred(PairMpStringint{k, v}) {
			m1[k] = v
		}
	}
	return m1
}

func (m MapStringint) FilterKeys(pred func(String) bool) MapStringint {
	if m == nil {
		return nil
	}
	m1 := MapStringint{}
	for k, v := range m {
		if pred(k) {
			m1[k] = v
		}
	}
	return m1
}

func (m MapStringint) FilterNot(pred func(PairMpStringint) bool) MapStringint {
	if m == nil {
		return nil
	}
	m1 := MapStringint{}
	for k, v := range m {
		if !pred(PairMpStringint{k, v}) {
			m1[k] = v
		}
	}
	return m1
}

func (m MapStringint) FilterValues(pred func(int) bool) MapStringint {
	if m == nil {
		return nil
	}
	m1 := MapStringint{}
	for k, v := range m {
		if pred(v) {
			m1[k] = v
		}
	}
	return m1
}

func (m MapStringint) ForEach(f func(PairMpStringint)) {
	for k, v := range m {
		f(PairMpStringint{k, v})
	}
}

func (m MapStringint) GetOrElse(k String, f func(String) int) int {
	if v, ok := m[k]; ok {
		return v
	}
	return f(k)
}

func (m MapStringint) IsNotEmpty() bool {
	return len(m) > 0
}

// MaxWith returns an entry in the map with maximum value, using a comparator function.
// Returns an error if the map is empty.
func (m MapStringint) MaxWith(comparator func(PairMpStringint, PairMpStringint) int) (PairMpStringint, error) {
	var max PairMpStringint

	if len(m) == 0 {
		return max, errors.New("empty or nil map")
	}

	first := true
	for k, v := range m {
		if first {
			max = PairMpStringint{k, v}
			first = false
			continue
		}
		if pair := (PairMpStringint{k, v}); comparator(max, pair) < 0 {
			max = pair
		}
	}
	return max, nil
}

// MinusKey returns a new MapStringint without the entry associated with the given key. If the
// key is not in the receiver then it returns a copy of the receiver.
func (m MapStringint) MinusKey(k String) MapStringint {
	m1 := m.Copy()
	delete(m1, k)
	return m1
}

func (m MapStringint) MinusKeys(s []String) MapStringint {
	m1 := m.Copy()
	for _, k := range s {
		delete(m1, k)
	}
	return m1
}

func (m MapStringint) MinWith(comparator func(PairMpStringint, PairMpStringint) int) (PairMpStringint, error) {
	reverseComp := func(p1 PairMpStringint, p2 PairMpStringint) int { return -comparator(p1, p2) }
	return m.MaxWith(reverseComp)
}

func (m MapStringint) PlusEntry(entry PairMpStringint) MapStringint {
	m1 := m.Copy()
	if m1 == nil {
		m1 = MapStringint{}
	}
	m1[entry.X1] = entry.X2
	return m1
}

func (m MapStringint) PlusMap(other MapStringint) MapStringint {
	var m1 MapStringint
	switch {
	case m == nil && other == nil:
		return nil
	case m == nil:
		m1 = MapStringint{}
	default:
		m1 = m.Copy()
	}

	for k, v := range other {
		m1[k] = v
	}
	return m1
}

func (m MapStringint) PlusSlice(s []PairMpStringint) MapStringint {
	var m1 MapStringint
	switch {
	case m == nil && s == nil:
		return nil
	case m == nil:
		m1 = MapStringint{}
	default:
		m1 = m.Copy()
	}

	for _, pair := range s {
		m1[pair.X1] = pair.X2
	}
	return m1
}

func (m MapStringint) Add(k String, v int) MapStringint {
	m1 := m.Copy()
	if m1 == nil {
		m1 = MapStringint{}
	}
	m1[k] = v
	return m1
}
