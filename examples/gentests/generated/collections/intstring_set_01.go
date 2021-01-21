// Code generated -- DO NOT EDIT.

package collections

// FlatMapstring returns the set obtained by applying the argument function to each item in the
// receiver and taking the union of the results.
func (s Setint) FlatMapstring(f func(int) Setstring) Setstring {
	if s == nil {
		return nil
	}
	r := make(Setstring, len(s)) // optimizing for speed vs space
	for x := range s {
		for e := range f(x) {
			r[e] = true
		}
	}
	return r
}

// GroupBystring returns a map whose keys are outputs of the keySelector function applied to
// the elements in the receiver and whose values are sets containing the elements in the
// receiver that correspond to each key obtained with the keySelector function.
func (s Setint) GroupBystring(keySelector func(int) string) MapstringSetint {
	if s == nil {
		return nil
	}
	m := make(MapstringSetint, len(s)/2) // optimizing for speed vs space
	for x := range s {
		k := keySelector(x)
		set, ok := m[k]
		if !ok {
			set = make(Setint, 1)
		}
		set[x] = true
		m[k] = set
	}
	return m
}

// Mapstring returns a new set resulting from the application of a given function to
// each element of a given set.
func (s Setint) Mapstring(f func(int) string) Setstring {
	if s == nil {
		return nil
	}
	r := make(Setstring, len(s))
	for a := range s {
		r[f(a)] = true
	}
	return r
}

// ToMap returns a map whose keys are the first components in the elements of the receiver and
// whose values are the corresonding second components in the elements of the receiver.
// If multiple elements in the receiver have the same first component, the corresponding
// value in the resulting map will be picked from one of them.
func (s SetOfPairintstring) ToMap() Mapintstring {
	if s == nil {
		return nil
	}
	m := make(map[int]string, len(s))
	for p := range s {
		m[p.X1] = p.X2
	}
	return m
}
