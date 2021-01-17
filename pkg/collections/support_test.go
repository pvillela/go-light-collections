package collections

// SliceFoo is a wrapper type to enable extension methods.
// Used with all pseudo-generic functions for slices.
type SliceDat []Dat

// Bar is another example data structure.
type Bar struct {
	W1 int
	W2 []string
}

// SliceBar is a wrapper type to enable extension methods.
// Used with all pseudo-generic functions for slices.
type SliceBar []Bar
