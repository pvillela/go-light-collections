package collections

// ToMap returns a map whose keys are the first components in the items of the receiver and
// whose values are the corresonding second components in the items of the receiver.
// If multiple items in the receiver have the same first component, the corresponding
// value in the resulting map will be taken from the last such item in the receiver.
func (s SliceOfPairT0T1) ToMap() MapT0T1 {
	m := make(map[T0]T1, len(s))
	for _, p := range s {
		m[p.X1] = p.X2
	}
	return m
}
