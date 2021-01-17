package collections

// Dat is an example data structure.
type Dat struct {
	V1 int
	V2 string
}

////
// Slice used in tests. Cloned each time to avoid nasty side-effects.

func sDat() SliceT0 {
	return SliceT0{Dat{1, "w1"}, Dat{22, "w22"}, Dat{333, "w333"}, Dat{4444, "w4444"},
		Dat{22, "w22"}}
}
