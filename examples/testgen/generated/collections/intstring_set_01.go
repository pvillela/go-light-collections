// Code generated -- DO NOT EDIT.

package collections

// FlatMapstring returns the set obtained by applying the argument function to each item in the
// receiver and taking the union of the results.
func (s Setint) FlatMapstring(f func(int) map[string]bool) map[string]bool {
	if s == nil {
		return nil
	}
	r := make(map[string]bool, len(s)) // optimizing for speed vs space
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
func (s Setint) GroupBystring(keySelector func(int) string) map[string]Setint {
	if s == nil {
		return nil
	}
	m := make(map[string]Setint, len(s)/2) // optimizing for speed vs space
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
func (s Setint) Mapstring(f func(int) string) map[string]bool {
	if s == nil {
		return nil
	}
	r := make(map[string]bool, len(s))
	for a := range s {
		r[f(a)] = true
	}
	return r
}
