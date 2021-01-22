// Code generated -- DO NOT EDIT.

package coll

import "errors"

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

func (m MapStringint) Entries() SetOfPairStringint {
	entries := make(SetOfPairStringint, len(m))
	for k, v := range m {
		entries[PairStringint{k, v}] = true
	}
	return entries
}

func (m MapStringint) Keys() SetString {
	keys := make(map[String]bool, len(m))
	for k := range m {
		keys[k] = true
	}
	return keys
}

// Length returns the number of items in the receiver.
func (m MapStringint) Length() int {
	return len(m)
}

// Size returns the number of items in the receiver. Same as Length.
func (m MapStringint) Size() int {
	return len(m)
}

func (m MapStringint) Values() Setint {
	values := make(map[int]bool, len(m))
	for _, v := range m {
		values[v] = true
	}
	return values
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
func (m MapStringint) Count(pred func(PairStringint) bool) int {
	count := 0
	for k, v := range m {
		if pred(PairStringint{k, v}) {
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

func (m MapStringint) All(pred func(PairStringint) bool) bool {
	for k, v := range m {
		if !pred(PairStringint{k, v}) {
			return false
		}
	}
	return true
}

func (m MapStringint) Any(pred func(PairStringint) bool) bool {
	for k, v := range m {
		if pred(PairStringint{k, v}) {
			return true
		}
	}
	return false
}

func (m MapStringint) ToSlice() SliceOfPairStringint {
	if m == nil {
		return nil
	}
	s := make(SliceOfPairStringint, len(m))
	i := 0
	for k, v := range m {
		s[i] = PairStringint{k, v}
		i++
	}
	return s
}

func (m MapStringint) Filter(pred func(PairStringint) bool) MapStringint {
	if m == nil {
		return nil
	}
	m1 := MapStringint{}
	for k, v := range m {
		if pred(PairStringint{k, v}) {
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

func (m MapStringint) FilterNot(pred func(PairStringint) bool) MapStringint {
	if m == nil {
		return nil
	}
	m1 := MapStringint{}
	for k, v := range m {
		if !pred(PairStringint{k, v}) {
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

func (m MapStringint) ForEach(f func(PairStringint)) {
	for k, v := range m {
		f(PairStringint{k, v})
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
func (m MapStringint) MaxWith(comparator func(PairStringint, PairStringint) int) (PairStringint, error) {
	var max PairStringint

	if len(m) == 0 {
		return max, errors.New("empty or nil map")
	}

	first := true
	for k, v := range m {
		if first {
			max = PairStringint{k, v}
			first = false
			continue
		}
		if pair := (PairStringint{k, v}); comparator(max, pair) < 0 {
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

func (m MapStringint) MinusKeys(s SliceString) MapStringint {
	m1 := m.Copy()
	for _, k := range s {
		delete(m1, k)
	}
	return m1
}

func (m MapStringint) MinWith(comparator func(PairStringint, PairStringint) int) (PairStringint, error) {
	reverseComp := func(p1 PairStringint, p2 PairStringint) int { return -comparator(p1, p2) }
	return m.MaxWith(reverseComp)
}

func (m MapStringint) PlusEntry(entry PairStringint) MapStringint {
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

func (m MapStringint) PlusSlice(s SliceOfPairStringint) MapStringint {
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
