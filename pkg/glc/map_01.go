package glc

import "errors"

func (m MapT0T1) Copy() MapT0T1 {
	if m == nil {
		return nil
	}
	m1 := make(MapT0T1, len(m))
	for k, v := range m {
		m1[k] = v
	}
	return m1
}

func (m MapT0T1) Entries() SetOfPairT0T1 {
	entries := make(SetOfPairT0T1, len(m))
	for k, v := range m {
		entries[PairT0T1{k, v}] = true
	}
	return entries
}

func (m MapT0T1) Keys() SetT0 {
	keys := make(map[T0]bool, len(m))
	for k := range m {
		keys[k] = true
	}
	return keys
}

// Length returns the number of items in the receiver.
func (m MapT0T1) Length() int {
	return len(m)
}

// Size returns the number of items in the receiver. Same as Length.
func (m MapT0T1) Size() int {
	return len(m)
}

func (m MapT0T1) Values() SetT1 {
	values := make(map[T1]bool, len(m))
	for _, v := range m {
		values[v] = true
	}
	return values
}

func (m MapT0T1) ContainsKey(k T0) bool {
	_, ok := m[k]
	return ok
}

func (m MapT0T1) ContainsValue(v T1) bool {
	for _, v1 := range m {
		if v1 == v {
			return true
		}
	}
	return false
}

// Count returns the number of entries in the receiver that satisfy the predicate.
func (m MapT0T1) Count(pred func(PairT0T1) bool) int {
	count := 0
	for k, v := range m {
		if pred(PairT0T1{k, v}) {
			count++
		}
	}
	return count
}

func (m MapT0T1) Get(k T0) (T1, bool) {
	v, ok := m[k]
	return v, ok
}

func (m MapT0T1) IsEmpty() bool {
	return len(m) == 0
}

func (m MapT0T1) All(pred func(PairT0T1) bool) bool {
	for k, v := range m {
		if !pred(PairT0T1{k, v}) {
			return false
		}
	}
	return true
}

func (m MapT0T1) Any(pred func(PairT0T1) bool) bool {
	for k, v := range m {
		if pred(PairT0T1{k, v}) {
			return true
		}
	}
	return false
}

func (m MapT0T1) ToSlice() SliceOfPairT0T1 {
	if m == nil {
		return nil
	}
	s := make(SliceOfPairT0T1, len(m))
	i := 0
	for k, v := range m {
		s[i] = PairT0T1{k, v}
		i++
	}
	return s
}

func (m MapT0T1) Filter(pred func(PairT0T1) bool) MapT0T1 {
	if m == nil {
		return nil
	}
	m1 := MapT0T1{}
	for k, v := range m {
		if pred(PairT0T1{k, v}) {
			m1[k] = v
		}
	}
	return m1
}

func (m MapT0T1) FilterKeys(pred func(T0) bool) MapT0T1 {
	if m == nil {
		return nil
	}
	m1 := MapT0T1{}
	for k, v := range m {
		if pred(k) {
			m1[k] = v
		}
	}
	return m1
}

func (m MapT0T1) FilterNot(pred func(PairT0T1) bool) MapT0T1 {
	if m == nil {
		return nil
	}
	m1 := MapT0T1{}
	for k, v := range m {
		if !pred(PairT0T1{k, v}) {
			m1[k] = v
		}
	}
	return m1
}

func (m MapT0T1) FilterValues(pred func(T1) bool) MapT0T1 {
	if m == nil {
		return nil
	}
	m1 := MapT0T1{}
	for k, v := range m {
		if pred(v) {
			m1[k] = v
		}
	}
	return m1
}

func (m MapT0T1) ForEach(f func(PairT0T1)) {
	for k, v := range m {
		f(PairT0T1{k, v})
	}
}

func (m MapT0T1) GetOrElse(k T0, f func(T0) T1) T1 {
	if v, ok := m[k]; ok {
		return v
	}
	return f(k)
}

func (m MapT0T1) IsNotEmpty() bool {
	return len(m) > 0
}

// MaxWith returns an entry in the map with maximum value, using a comparator function.
// Returns an error if the map is empty.
func (m MapT0T1) MaxWith(comparator func(PairT0T1, PairT0T1) int) (PairT0T1, error) {
	var max PairT0T1

	if len(m) == 0 {
		return max, errors.New("empty or nil map")
	}

	first := true
	for k, v := range m {
		if first {
			max = PairT0T1{k, v}
			first = false
			continue
		}
		if pair := (PairT0T1{k, v}); comparator(max, pair) < 0 {
			max = pair
		}
	}
	return max, nil
}

// MinusKey returns a new MapT0T1 without the entry associated with the given key. If the
// key is not in the receiver then it returns a copy of the receiver.
func (m MapT0T1) MinusKey(k T0) MapT0T1 {
	m1 := m.Copy()
	delete(m1, k)
	return m1
}

func (m MapT0T1) MinusKeys(s SliceT0) MapT0T1 {
	m1 := m.Copy()
	for _, k := range s {
		delete(m1, k)
	}
	return m1
}

func (m MapT0T1) MinWith(comparator func(PairT0T1, PairT0T1) int) (PairT0T1, error) {
	reverseComp := func(p1 PairT0T1, p2 PairT0T1) int { return -comparator(p1, p2) }
	return m.MaxWith(reverseComp)
}

func (m MapT0T1) PlusEntry(entry PairT0T1) MapT0T1 {
	m1 := m.Copy()
	if m1 == nil {
		m1 = MapT0T1{}
	}
	m1[entry.X1] = entry.X2
	return m1
}

func (m MapT0T1) PlusMap(other MapT0T1) MapT0T1 {
	var m1 MapT0T1
	switch {
	case m == nil && other == nil:
		return nil
	case m == nil:
		m1 = MapT0T1{}
	default:
		m1 = m.Copy()
	}

	for k, v := range other {
		m1[k] = v
	}
	return m1
}

func (m MapT0T1) PlusSlice(s SliceOfPairT0T1) MapT0T1 {
	var m1 MapT0T1
	switch {
	case m == nil && s == nil:
		return nil
	case m == nil:
		m1 = MapT0T1{}
	default:
		m1 = m.Copy()
	}

	for _, pair := range s {
		m1[pair.X1] = pair.X2
	}
	return m1
}

func (m MapT0T1) Add(k T0, v T1) MapT0T1 {
	m1 := m.Copy()
	if m1 == nil {
		m1 = MapT0T1{}
	}
	m1[k] = v
	return m1
}
