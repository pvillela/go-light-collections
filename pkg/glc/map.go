/*
 *  Copyright © 2021 Paulo Villela. All rights reserved.
 *  Use of this source code is governed by the Apache 2.0 license
 *  that can be found in the LICENSE file.
 */

package glc

import (
	"errors"
)

// Map[T0, T1] is a type wrapper, implements IMap interfaces.
type Map[T0 comparable, T1 any] map[T0]T1

func (m Map[T0, T1]) Copy() Map[T0, T1] {
	if m == nil {
		return nil
	}
	m1 := make(Map[T0, T1], len(m))
	for k, v := range m {
		m1[k] = v
	}
	return m1
}

func (m Map[T0, T1]) Entries() []Pair[T0, T1] {
	if m == nil {
		return nil
	}
	entries := make([]Pair[T0, T1], len(m))
	i := 0
	for k, v := range m {
		entries[i] = Pair[T0, T1]{k, v}
		i++
	}
	return entries
}

func (m Map[T0, T1]) Keys() []T0 {
	if m == nil {
		return nil
	}
	keys := make([]T0, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

func (m Map[T0, T1]) Values() []T1 {
	if m == nil {
		return nil
	}
	values := make([]T1, len(m))
	i := 0
	for _, v := range m {
		values[i] = v
		i++
	}
	return values
}

// Length returns the number of items in the receiver.
func (m Map[T0, T1]) Length() int {
	return len(m)
}

// Size returns the number of items in the receiver. Same as Length.
func (m Map[T0, T1]) Size() int {
	return len(m)
}

func (m Map[T0, T1]) ContainsKey(k T0) bool {
	_, ok := m[k]
	return ok
}

func MapContainsValue[T0 comparable, T1 comparable](m Map[T0, T1], v T1) bool {
	for _, v1 := range m {
		if v1 == v {
			return true
		}
	}
	return false
}

// Count returns the number of entries in the receiver that satisfy the predicate.
func (m Map[T0, T1]) Count(pred func(Pair[T0, T1]) bool) int {
	count := 0
	for k, v := range m {
		if pred(Pair[T0, T1]{k, v}) {
			count++
		}
	}
	return count
}

func (m Map[T0, T1]) Get(k T0) (T1, bool) {
	v, ok := m[k]
	return v, ok
}

func (m Map[T0, T1]) IsEmpty() bool {
	return len(m) == 0
}

func (m Map[T0, T1]) All(pred func(Pair[T0, T1]) bool) bool {
	for k, v := range m {
		if !pred(Pair[T0, T1]{k, v}) {
			return false
		}
	}
	return true
}

func (m Map[T0, T1]) Any(pred func(Pair[T0, T1]) bool) bool {
	for k, v := range m {
		if pred(Pair[T0, T1]{k, v}) {
			return true
		}
	}
	return false
}

func (m Map[T0, T1]) ToSlice() []Pair[T0, T1] {
	if m == nil {
		return nil
	}
	s := make([]Pair[T0, T1], len(m))
	i := 0
	for k, v := range m {
		s[i] = Pair[T0, T1]{k, v}
		i++
	}
	return s
}

func (m Map[T0, T1]) Filter(pred func(Pair[T0, T1]) bool) Map[T0, T1] {
	if m == nil {
		return nil
	}
	m1 := Map[T0, T1]{}
	for k, v := range m {
		if pred(Pair[T0, T1]{k, v}) {
			m1[k] = v
		}
	}
	return m1
}

func (m Map[T0, T1]) FilterKeys(pred func(T0) bool) Map[T0, T1] {
	if m == nil {
		return nil
	}
	m1 := Map[T0, T1]{}
	for k, v := range m {
		if pred(k) {
			m1[k] = v
		}
	}
	return m1
}

func (m Map[T0, T1]) FilterNot(pred func(Pair[T0, T1]) bool) Map[T0, T1] {
	if m == nil {
		return nil
	}
	m1 := Map[T0, T1]{}
	for k, v := range m {
		if !pred(Pair[T0, T1]{k, v}) {
			m1[k] = v
		}
	}
	return m1
}

func (m Map[T0, T1]) FilterValues(pred func(T1) bool) Map[T0, T1] {
	if m == nil {
		return nil
	}
	m1 := Map[T0, T1]{}
	for k, v := range m {
		if pred(v) {
			m1[k] = v
		}
	}
	return m1
}

func (m Map[T0, T1]) ForEach(f func(Pair[T0, T1])) {
	for k, v := range m {
		f(Pair[T0, T1]{k, v})
	}
}

func (m Map[T0, T1]) GetOrElse(k T0, f func(T0) T1) T1 {
	if v, ok := m[k]; ok {
		return v
	}
	return f(k)
}

func (m Map[T0, T1]) IsNotEmpty() bool {
	return len(m) > 0
}

// MaxWith returns an entry in the map with maximum value, using a comparator function.
// Returns an error if the map is empty.
func (m Map[T0, T1]) MaxWith(comparator func(Pair[T0, T1], Pair[T0, T1]) int) (Pair[T0, T1], error) {
	var max Pair[T0, T1]

	if len(m) == 0 {
		return max, errors.New("empty or nil map")
	}

	first := true
	for k, v := range m {
		if first {
			max = Pair[T0, T1]{k, v}
			first = false
			continue
		}
		if pair := (Pair[T0, T1]{k, v}); comparator(max, pair) < 0 {
			max = pair
		}
	}
	return max, nil
}

// MinusKey returns a new Map[T0, T1] without the entry associated with the given key. If the
// key is not in the receiver then it returns a copy of the receiver.
func (m Map[T0, T1]) MinusKey(k T0) Map[T0, T1] {
	m1 := m.Copy()
	delete(m1, k)
	return m1
}

func (m Map[T0, T1]) MinusKeys(s []T0) Map[T0, T1] {
	m1 := m.Copy()
	for _, k := range s {
		delete(m1, k)
	}
	return m1
}

func (m Map[T0, T1]) MinWith(comparator func(Pair[T0, T1], Pair[T0, T1]) int) (Pair[T0, T1], error) {
	reverseComp := func(p1 Pair[T0, T1], p2 Pair[T0, T1]) int { return -comparator(p1, p2) }
	return m.MaxWith(reverseComp)
}

func (m Map[T0, T1]) PlusEntry(entry Pair[T0, T1]) Map[T0, T1] {
	m1 := m.Copy()
	if m1 == nil {
		m1 = Map[T0, T1]{}
	}
	m1[entry.X1] = entry.X2
	return m1
}

func (m Map[T0, T1]) PlusMap(other Map[T0, T1]) Map[T0, T1] {
	var m1 Map[T0, T1]
	switch {
	case m == nil && other == nil:
		return nil
	case m == nil:
		m1 = Map[T0, T1]{}
	default:
		m1 = m.Copy()
	}

	for k, v := range other {
		m1[k] = v
	}
	return m1
}

func (m Map[T0, T1]) PlusSlice(s []Pair[T0, T1]) Map[T0, T1] {
	var m1 Map[T0, T1]
	switch {
	case m == nil && s == nil:
		return nil
	case m == nil:
		m1 = Map[T0, T1]{}
	default:
		m1 = m.Copy()
	}

	for _, pair := range s {
		m1[pair.X1] = pair.X2
	}
	return m1
}

func (m Map[T0, T1]) Add(k T0, v T1) Map[T0, T1] {
	m1 := m.Copy()
	if m1 == nil {
		m1 = Map[T0, T1]{}
	}
	m1[k] = v
	return m1
}

func MapFlatMap[T0 comparable, T1 any, T2 any](m Map[T0, T1], f func(Pair[T0, T1]) []T2) []T2 {
	if m == nil {
		return nil
	}
	r := make([]T2, 0, len(m)) // optimizing for speed vs space
	for k, v := range m {
		r = append(r, f(Pair[T0, T1]{k, v})...)
	}
	return r
}

func MapMap[T0 comparable, T1 any, T2 any](m Map[T0, T1], f func(Pair[T0, T1]) T2) []T2 {
	if m == nil {
		return nil
	}
	r := make([]T2, 0, len(m)) // optimizing for speed vs space
	for k, v := range m {
		r = append(r, f(Pair[T0, T1]{k, v}))
	}
	return r
}

func MapMapValues[T0 comparable, T1 any, T2 any](m Map[T0, T1], f func(Pair[T0, T1]) T2) map[T0]T2 {
	if m == nil {
		return nil
	}
	r := make(map[T0]T2)
	for k, v := range m {
		r[k] = f(Pair[T0, T1]{k, v})
	}
	return r
}

func MapMapKeys[T0 comparable, T1 any, T2 comparable](m Map[T0, T1], f func(Pair[T0, T1]) T2) map[T2]T1 {
	if m == nil {
		return nil
	}
	r := make(map[T2]T1)
	for k, v := range m {
		r[f(Pair[T0, T1]{k, v})] = v
	}
	return r
}
