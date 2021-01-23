// Code generated -- DO NOT EDIT.

package collections

// ToSet returns a set containing the values in the receiver.
func (s SliceDat) ToSet() map[Dat]bool {
	if s == nil {
		return nil
	}
	set := make(map[Dat]bool, len(s)) // optimize for speed vs space
	for _, x := range s {
		set[x] = true
	}
	return set
}
