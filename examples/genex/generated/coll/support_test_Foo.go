// Code generated -- DO NOT EDIT.

package coll

// Foo is an example data structure.
type Foo struct {
	V1 int
	V2 string
}

// SliceFoo is a wrapper type to enable extension methods.
// Used with all pseudo-generic functions for slices.
type SliceFoo []Foo

// Bar is another example data structure.
type Bar struct {
	W1 int
	W2 []string
}

// SliceBar is a wrapper type to enable extension methods.
// Used with all pseudo-generic functions for slices.
type SliceBar []Bar

////
// Slices used in tests. Cloned each time to avoid nasty side-effects.

func sFoo() SliceFoo {
	return SliceFoo{Foo{1, "w1"}, Foo{22, "w22"}, Foo{333, "w333"}, Foo{4444, "w4444"},
		Foo{22, "w22"}}
}
