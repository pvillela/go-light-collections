// Code generated -- DO NOT EDIT.

package collections

// SliceOfPairDatint is a type wrapper.
type SliceOfPairDatint []struct {
	X1 Dat
	X2 int
}

// ToMap returns a map whose keys are the first components in the items of the receiver and
// whose values are the corresonding second components in the items of the receiver.
// If multiple items in the receiver have the same first component, the corresponding
// value in the resulting map will be taken from the last such item in the receiver.
func (s SliceOfPairDatint) ToMap() map[Dat]int {
	if s == nil {
		return nil
	}
	m := make(map[Dat]int, len(s))
	for _, p := range s {
		m[p.X1] = p.X2
	}
	return m
}
