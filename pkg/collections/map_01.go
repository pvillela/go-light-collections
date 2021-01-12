package collections

import "errors"

func (m MapT0T1) Entries() SetOfPairT0T1 {
	entries := map[PairT0T1]bool{}
	for k, v := range m {
		m[PairT0T1{k, v}] = true
	}
	return entries
}

func (m MapT0T1) Keys() SetT0 {
	keys := map[T0]bool{}
	for k := range m {
		keys[k] = true
	}
	return keys
}

func (m MapT0T1) Count() int {
	return len(m)
}

func (m MapT0T1) Values() SetT0 {
	values := map[T1]bool{}
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
	s := SliceOfPairT0T1{}
	for k, v := range m {
		s = append(s, PairT0T1{k, v})
	}
	return s
}

func (m MapT0T1) Filter(pred func(PairT0T1) bool) MapT0T1 {
	m1 := MapT0T1{}
	for k, v := range m {
		if pred(PairT0T1{k, v}) {
			m1[k] = v
		}
	}
	return m1
}

func (m MapT0T1) FilterKeys(pred func(T0) bool) MapT0T1 {
	m1 := MapT0T1{}
	for k, v := range m {
		if pred(k) {
			m1[k] = v
		}
	}
	return m1
}

func (m MapT0T1) FilterNot(pred func(PairT0T1) bool) MapT0T1 {
	m1 := MapT0T1{}
	for k, v := range m {
		if !pred(PairT0T1{k, v}) {
			m1[k] = v
		}
	}
	return m1
}

func (m MapT0T1) FilterValues(pred func(T1) bool) MapT0T1 {
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

func (m MapT0T1) GetOrElse(k T0, f func() T1) T1 {
	if v, ok := m[k]; ok {
		return v
	}
	return f()
}

func (m MapT0T1) IsNotEmpty() bool {
	return len(m) > 0
}

// MaxWith returns an entry in the map with maximum value, using a comparator function.
// Returns an error if the map is empty.
func (m MapT0T1) MaxWith(comparator func(PairT0T1, PairT0T1) int) (PairT0T1, error) {
	var max PairT0T1
	if len(m) == 0 {
		return max, errors.New("Empty map.")
	}
	for k, v := range m {
		if pair := (PairT0T1{k, v}); comparator(max, pair) < 0 {
			max = pair
		}
	}
	return max, nil
}

func (m MapT0T1) MinusKey(T0) MapT0T1 {

}

func (m MapT0T1) Minus(SliceT0) MapT0T1 {

}

func (m MapT0T1) MinWith(func(PairT0T1) int) MapT0T1 {

}

func (m MapT0T1) PlusEntry(PairT0T1) MapT0T1 {

}

func (m MapT0T1) Plus(MapT0T1) MapT0T1 {

}

func (m MapT0T1) PlusSlice(SliceOfPairT0T1) MapT0T1 {

}

func (m MapT0T1) Put(k T0, v T0) {

}
