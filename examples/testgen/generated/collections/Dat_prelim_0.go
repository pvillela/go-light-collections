// Code generated -- DO NOT EDIT.

package collections

////
// Preliminaries

// Dat is an example data structure.
type Dat struct {
	V1 int
	V2 string
}

// Slice used in tests. Cloned each time to avoid nasty side-effects.

func sDat() SliceDat {
	return SliceDat{Dat{1, "w1"}, Dat{22, "w22"}, Dat{333, "w333"}, Dat{4444, "w4444"},
		Dat{22, "w22"}}
}

// Type conversion functions.

func toDat(a Any) Dat { return Any(a).(Dat) }

func toInt(a Any) int { return Any(a).(int) }

func toString(a Any) string { return Any(a).(string) }