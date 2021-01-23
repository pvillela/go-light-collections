// Code generated -- DO NOT EDIT.

package coll

// SliceOfPairStringint is a type wrapper.
type SliceOfPairStringint []struct {
	X1 String
	X2 int
}

// ToMap returns a map whose keys are the first components in the items of the receiver and
// whose values are the corresonding second components in the items of the receiver.
// If multiple items in the receiver have the same first component, the corresponding
// value in the resulting map will be taken from the last such item in the receiver.
func (s SliceOfPairStringint) ToMap() map[String]int {
	if s == nil {
		return nil
	}
	m := make(map[String]int, len(s))
	for _, p := range s {
		m[p.X1] = p.X2
	}
	return m
}
