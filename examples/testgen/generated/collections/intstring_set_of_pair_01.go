// Code generated -- DO NOT EDIT.

package collections

// SetOfPairintstring is a type wrapper, implements ISetOfPairintstring interface.
type SetOfPairintstring map[struct {
	X1 int
	X2 string
}]bool

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
