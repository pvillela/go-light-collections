package glc_test

import (
	"fmt"

	c "github.com/pvillela/go-light-collections/pkg/glc"
)

/////////////////////
// Part 1 of example -- types and standard converters for data structures.

////
// Types and standard converters for our first example data structure.

// Foo is an example data structure.
type Foo struct {
	v1 int
	v2 string
}

// SliceFoo is a wrapper type to enable extension methods.
// Used with all pseudo-generic functions for slices.
type SliceFoo []Foo

// ToSliceAny converts to SliceAny.
func (s SliceFoo) ToSliceAny() c.SliceAny {
	r := make(c.SliceAny, len(s))
	for i, x := range s {
		r[i] = x
	}
	return r
}

// ToSliceFoo is a conversion function.
func ToSliceFoo(s c.SliceAny) SliceFoo {
	r := make(SliceFoo, len(s))
	for i, x := range s {
		r[i] = x.(Foo)
	}
	return r
}

// SlicePFoo is a wrapper type to enable extension methods.
// Used with pseudo-generic functions for slices of pointers.
type SlicePFoo []*Foo

// ToSliceAny converts to SliceAny.
func (s SlicePFoo) ToSliceAny() c.SliceAny {
	r := make(c.SliceAny, len(s))
	for i, x := range s {
		r[i] = x
	}
	return r
}

////
// Types and standard converters for our second example data structure.

// Bar is another example data structure.
type Bar struct {
	w1 int
	w2 []string
}

// SliceBar is a wrapper type to enable extension methods.
// Used with all pseudo-generic functions for slices.
type SliceBar []Bar

// ToSliceAny converts to SliceAny.
func (s SliceBar) ToSliceAny() c.SliceAny {
	r := make(c.SliceAny, len(s))
	for i, x := range s {
		r[i] = x
	}
	return r
}

// ToSliceBar is a conversion function.
func ToSliceBar(s c.SliceAny) SliceBar {
	r := make(SliceBar, len(s))
	for i, x := range s {
		r[i] = x.(Bar)
	}
	return r
}

// SlicePBar is a wrapper type to enable extension methods.
// Used with pseudo-generic functions for slices of pointers.
type SlicePBar []*Bar

// ToSliceAny converts to SliceAny.
func (s SlicePBar) ToSliceAny() c.SliceAny {
	r := make(c.SliceAny, len(s))
	for i, x := range s {
		r[i] = x
	}
	return r
}

// ToSlicePBar is a conversion function.
func ToSlicePBar(s c.SliceAny) SlicePBar {
	r := make(SlicePBar, len(s))
	for i, p := range s {
		r[i] = p.(*Bar)
	}
	return r
}

////
// Helper conversion functions for the examples below

// toP is a helper function to convert to a slice of pointers.
func (s SliceFoo) toP() []*Foo {
	r := make([]*Foo, len(s))
	for i := range s {
		r[i] = &s[i]
	}
	return r
}

// toV is a helper function to convert to a slice of values.
func (s SlicePBar) toV() []Bar {
	r := make([]Bar, len(s))
	for i := range s {
		r[i] = *s[i]
	}
	return r
}

// Run example.
func ExampleSliceAny_part1() {
	fmt.Println("End of part1")
	// Output: End of part1
}
