package glc

// FlatMapT1 returns the set obtained by applying the argument function to each item in the
// receiver and taking the union of the results.
func (s SetT0) FlatMapT1(f func(T0) SetT1) SetT1 {
	if s == nil {
		return nil
	}
	r := make(SetT1, len(s)) // optimizing for speed vs space
	for x := range s {
		for e := range f(x) {
			r[e] = true
		}
	}
	return r
}

// GroupByT1 returns a map whose keys are outputs of the keySelector function applied to
// the elements in the receiver and whose values are sets containing the elements in the
// receiver that correspond to each key obtained with the keySelector function.
func (s SetT0) GroupByT1(keySelector func(T0) T1) map[T1]SetT0 {
	if s == nil {
		return nil
	}
	m := make(map[T1]SetT0, len(s)/2) // optimizing for speed vs space
	for x := range s {
		k := keySelector(x)
		set, ok := m[k]
		if !ok {
			set = make(SetT0, 1)
		}
		set[x] = true
		m[k] = set
	}
	return m
}

// MapT1 returns a new set resulting from the application of a given function to
// each element of a given set.
func (s SetT0) MapT1(f func(T0) T1) SetT1 {
	if s == nil {
		return nil
	}
	r := make(SetT1, len(s))
	for a := range s {
		r[f(a)] = true
	}
	return r
}
