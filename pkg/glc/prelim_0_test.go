package glc

////
// Preliminaries

// Dat is an example data structure.
type Dat struct {
	V1 int
	V2 string
}

// Type conversion functions.

func toDat(a Any) Dat { return Any(a).(Dat) }

func toInt(a Any) int { return Any(a).(int) }

func toString(a Any) string { return Any(a).(string) }
