// Code generated -- DO NOT EDIT.

package coll

import "errors"

func (m MapPersonint) Copy() MapPersonint {
	if m == nil {
		return nil
	}
	m1 := make(MapPersonint, len(m))
	for k, v := range m {
		m1[k] = v
	}
	return m1
}

func (m MapPersonint) Entries() SetOfPairPersonint {
	entries := make(SetOfPairPersonint, len(m))
	for k, v := range m {
		entries[PairPersonint{k, v}] = true
	}
	return entries
}

func (m MapPersonint) Keys() SetPerson {
	keys := make(map[Person]bool, len(m))
	for k := range m {
		keys[k] = true
	}
	return keys
}

// Length returns the number of items in the receiver.
func (m MapPersonint) Length() int {
	return len(m)
}

// Size returns the number of items in the receiver. Same as Length.
func (m MapPersonint) Size() int {
	return len(m)
}

func (m MapPersonint) Values() Setint {
	values := make(map[int]bool, len(m))
	for _, v := range m {
		values[v] = true
	}
	return values
}

func (m MapPersonint) ContainsKey(k Person) bool {
	_, ok := m[k]
	return ok
}

func (m MapPersonint) ContainsValue(v int) bool {
	for _, v1 := range m {
		if v1 == v {
			return true
		}
	}
	return false
}

// Count returns the number of entries in the receiver that satisfy the predicate.
func (m MapPersonint) Count(pred func(PairPersonint) bool) int {
	count := 0
	for k, v := range m {
		if pred(PairPersonint{k, v}) {
			count++
		}
	}
	return count
}

func (m MapPersonint) Get(k Person) (int, bool) {
	v, ok := m[k]
	return v, ok
}

func (m MapPersonint) IsEmpty() bool {
	return len(m) == 0
}

func (m MapPersonint) All(pred func(PairPersonint) bool) bool {
	for k, v := range m {
		if !pred(PairPersonint{k, v}) {
			return false
		}
	}
	return true
}

func (m MapPersonint) Any(pred func(PairPersonint) bool) bool {
	for k, v := range m {
		if pred(PairPersonint{k, v}) {
			return true
		}
	}
	return false
}

func (m MapPersonint) ToSlice() SliceOfPairPersonint {
	if m == nil {
		return nil
	}
	s := make(SliceOfPairPersonint, len(m))
	i := 0
	for k, v := range m {
		s[i] = PairPersonint{k, v}
		i++
	}
	return s
}

func (m MapPersonint) Filter(pred func(PairPersonint) bool) MapPersonint {
	if m == nil {
		return nil
	}
	m1 := MapPersonint{}
	for k, v := range m {
		if pred(PairPersonint{k, v}) {
			m1[k] = v
		}
	}
	return m1
}

func (m MapPersonint) FilterKeys(pred func(Person) bool) MapPersonint {
	if m == nil {
		return nil
	}
	m1 := MapPersonint{}
	for k, v := range m {
		if pred(k) {
			m1[k] = v
		}
	}
	return m1
}

func (m MapPersonint) FilterNot(pred func(PairPersonint) bool) MapPersonint {
	if m == nil {
		return nil
	}
	m1 := MapPersonint{}
	for k, v := range m {
		if !pred(PairPersonint{k, v}) {
			m1[k] = v
		}
	}
	return m1
}

func (m MapPersonint) FilterValues(pred func(int) bool) MapPersonint {
	if m == nil {
		return nil
	}
	m1 := MapPersonint{}
	for k, v := range m {
		if pred(v) {
			m1[k] = v
		}
	}
	return m1
}

func (m MapPersonint) ForEach(f func(PairPersonint)) {
	for k, v := range m {
		f(PairPersonint{k, v})
	}
}

func (m MapPersonint) GetOrElse(k Person, f func(Person) int) int {
	if v, ok := m[k]; ok {
		return v
	}
	return f(k)
}

func (m MapPersonint) IsNotEmpty() bool {
	return len(m) > 0
}

// MaxWith returns an entry in the map with maximum value, using a comparator function.
// Returns an error if the map is empty.
func (m MapPersonint) MaxWith(comparator func(PairPersonint, PairPersonint) int) (PairPersonint, error) {
	var max PairPersonint

	if len(m) == 0 {
		return max, errors.New("empty or nil map")
	}

	first := true
	for k, v := range m {
		if first {
			max = PairPersonint{k, v}
			first = false
			continue
		}
		if pair := (PairPersonint{k, v}); comparator(max, pair) < 0 {
			max = pair
		}
	}
	return max, nil
}

// MinusKey returns a new MapPersonint without the entry associated with the given key. If the
// key is not in the receiver then it returns a copy of the receiver.
func (m MapPersonint) MinusKey(k Person) MapPersonint {
	m1 := m.Copy()
	delete(m1, k)
	return m1
}

func (m MapPersonint) MinusKeys(s SlicePerson) MapPersonint {
	m1 := m.Copy()
	for _, k := range s {
		delete(m1, k)
	}
	return m1
}

func (m MapPersonint) MinWith(comparator func(PairPersonint, PairPersonint) int) (PairPersonint, error) {
	reverseComp := func(p1 PairPersonint, p2 PairPersonint) int { return -comparator(p1, p2) }
	return m.MaxWith(reverseComp)
}

func (m MapPersonint) PlusEntry(entry PairPersonint) MapPersonint {
	m1 := m.Copy()
	if m1 == nil {
		m1 = MapPersonint{}
	}
	m1[entry.X1] = entry.X2
	return m1
}

func (m MapPersonint) PlusMap(other MapPersonint) MapPersonint {
	var m1 MapPersonint
	switch {
	case m == nil && other == nil:
		return nil
	case m == nil:
		m1 = MapPersonint{}
	default:
		m1 = m.Copy()
	}

	if m1 == nil {
		m1 = MapPersonint{}
	}
	for k, v := range other {
		m1[k] = v
	}
	return m1
}

func (m MapPersonint) PlusSlice(s SliceOfPairPersonint) MapPersonint {
	var m1 MapPersonint
	switch {
	case m == nil && s == nil:
		return nil
	case m == nil:
		m1 = MapPersonint{}
	default:
		m1 = m.Copy()
	}

	if m1 == nil {
		m1 = MapPersonint{}
	}
	for _, pair := range s {
		m1[pair.X1] = pair.X2
	}
	return m1
}

func (m MapPersonint) Add(k Person, v int) MapPersonint {
	m1 := m.Copy()
	if m1 == nil {
		m1 = MapPersonint{}
	}
	m1[k] = v
	return m1
}
