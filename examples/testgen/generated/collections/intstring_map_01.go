// Code generated -- DO NOT EDIT.

/*
 * Copyright © 2021 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package collections

import "errors"

// Mapintstring is a type wrapper, implements IMap interfaces.
type Mapintstring map[int]string

// PairMpintstring is a type alias used only in Map methods to avoid code generation issues.
type PairMpintstring = struct {
	X1 int
	X2 string
}

func (m Mapintstring) Copy() Mapintstring {
	if m == nil {
		return nil
	}
	m1 := make(Mapintstring, len(m))
	for k, v := range m {
		m1[k] = v
	}
	return m1
}

func (m Mapintstring) Entries() []PairMpintstring {
	if m == nil {
		return nil
	}
	entries := make([]PairMpintstring, len(m))
	i := 0
	for k, v := range m {
		entries[i] = PairMpintstring{k, v}
		i++
	}
	return entries
}

func (m Mapintstring) Keys() []int {
	if m == nil {
		return nil
	}
	keys := make([]int, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

func (m Mapintstring) Values() []string {
	if m == nil {
		return nil
	}
	values := make([]string, len(m))
	i := 0
	for _, v := range m {
		values[i] = v
		i++
	}
	return values
}

// Length returns the number of items in the receiver.
func (m Mapintstring) Length() int {
	return len(m)
}

// Size returns the number of items in the receiver. Same as Length.
func (m Mapintstring) Size() int {
	return len(m)
}

func (m Mapintstring) ContainsKey(k int) bool {
	_, ok := m[k]
	return ok
}

func (m Mapintstring) ContainsValue(v string) bool {
	for _, v1 := range m {
		if v1 == v {
			return true
		}
	}
	return false
}

// Count returns the number of entries in the receiver that satisfy the predicate.
func (m Mapintstring) Count(pred func(PairMpintstring) bool) int {
	count := 0
	for k, v := range m {
		if pred(PairMpintstring{k, v}) {
			count++
		}
	}
	return count
}

func (m Mapintstring) Get(k int) (string, bool) {
	v, ok := m[k]
	return v, ok
}

func (m Mapintstring) IsEmpty() bool {
	return len(m) == 0
}

func (m Mapintstring) All(pred func(PairMpintstring) bool) bool {
	for k, v := range m {
		if !pred(PairMpintstring{k, v}) {
			return false
		}
	}
	return true
}

func (m Mapintstring) Any(pred func(PairMpintstring) bool) bool {
	for k, v := range m {
		if pred(PairMpintstring{k, v}) {
			return true
		}
	}
	return false
}

func (m Mapintstring) ToSlice() []PairMpintstring {
	if m == nil {
		return nil
	}
	s := make([]PairMpintstring, len(m))
	i := 0
	for k, v := range m {
		s[i] = PairMpintstring{k, v}
		i++
	}
	return s
}

func (m Mapintstring) Filter(pred func(PairMpintstring) bool) Mapintstring {
	if m == nil {
		return nil
	}
	m1 := Mapintstring{}
	for k, v := range m {
		if pred(PairMpintstring{k, v}) {
			m1[k] = v
		}
	}
	return m1
}

func (m Mapintstring) FilterKeys(pred func(int) bool) Mapintstring {
	if m == nil {
		return nil
	}
	m1 := Mapintstring{}
	for k, v := range m {
		if pred(k) {
			m1[k] = v
		}
	}
	return m1
}

func (m Mapintstring) FilterNot(pred func(PairMpintstring) bool) Mapintstring {
	if m == nil {
		return nil
	}
	m1 := Mapintstring{}
	for k, v := range m {
		if !pred(PairMpintstring{k, v}) {
			m1[k] = v
		}
	}
	return m1
}

func (m Mapintstring) FilterValues(pred func(string) bool) Mapintstring {
	if m == nil {
		return nil
	}
	m1 := Mapintstring{}
	for k, v := range m {
		if pred(v) {
			m1[k] = v
		}
	}
	return m1
}

func (m Mapintstring) ForEach(f func(PairMpintstring)) {
	for k, v := range m {
		f(PairMpintstring{k, v})
	}
}

func (m Mapintstring) GetOrElse(k int, f func(int) string) string {
	if v, ok := m[k]; ok {
		return v
	}
	return f(k)
}

func (m Mapintstring) IsNotEmpty() bool {
	return len(m) > 0
}

// MaxWith returns an entry in the map with maximum value, using a comparator function.
// Returns an error if the map is empty.
func (m Mapintstring) MaxWith(comparator func(PairMpintstring, PairMpintstring) int) (PairMpintstring, error) {
	var max PairMpintstring

	if len(m) == 0 {
		return max, errors.New("empty or nil map")
	}

	first := true
	for k, v := range m {
		if first {
			max = PairMpintstring{k, v}
			first = false
			continue
		}
		if pair := (PairMpintstring{k, v}); comparator(max, pair) < 0 {
			max = pair
		}
	}
	return max, nil
}

// MinusKey returns a new Mapintstring without the entry associated with the given key. If the
// key is not in the receiver then it returns a copy of the receiver.
func (m Mapintstring) MinusKey(k int) Mapintstring {
	m1 := m.Copy()
	delete(m1, k)
	return m1
}

func (m Mapintstring) MinusKeys(s []int) Mapintstring {
	m1 := m.Copy()
	for _, k := range s {
		delete(m1, k)
	}
	return m1
}

func (m Mapintstring) MinWith(comparator func(PairMpintstring, PairMpintstring) int) (PairMpintstring, error) {
	reverseComp := func(p1 PairMpintstring, p2 PairMpintstring) int { return -comparator(p1, p2) }
	return m.MaxWith(reverseComp)
}

func (m Mapintstring) PlusEntry(entry PairMpintstring) Mapintstring {
	m1 := m.Copy()
	if m1 == nil {
		m1 = Mapintstring{}
	}
	m1[entry.X1] = entry.X2
	return m1
}

func (m Mapintstring) PlusMap(other Mapintstring) Mapintstring {
	var m1 Mapintstring
	switch {
	case m == nil && other == nil:
		return nil
	case m == nil:
		m1 = Mapintstring{}
	default:
		m1 = m.Copy()
	}

	for k, v := range other {
		m1[k] = v
	}
	return m1
}

func (m Mapintstring) PlusSlice(s []PairMpintstring) Mapintstring {
	var m1 Mapintstring
	switch {
	case m == nil && s == nil:
		return nil
	case m == nil:
		m1 = Mapintstring{}
	default:
		m1 = m.Copy()
	}

	for _, pair := range s {
		m1[pair.X1] = pair.X2
	}
	return m1
}

func (m Mapintstring) Add(k int, v string) Mapintstring {
	m1 := m.Copy()
	if m1 == nil {
		m1 = Mapintstring{}
	}
	m1[k] = v
	return m1
}
