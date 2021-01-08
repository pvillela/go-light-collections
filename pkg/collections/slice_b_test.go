package collections_test

import (
	"fmt"

	pg "github.com/pvillela/GoSimpleCollections/pkg/collections"
	"github.com/pvillela/GoSimpleCollections/pkg/util/assert"
)

/////////////////////
// Types and standard converters for our first example data structure.

// Foo is an example data structure.
type Foo struct {
	v int
	w string
}

// SliceFoo is a wrapper type to enable extension methods.
// Used with all pseudo-generic functions for slices.
type SliceFoo []Foo

// Und converts to the underlying type.
func (s SliceFoo) Und() []Foo { return s }

// ToSliceAny is defined to implement pg.ConvertibleToSliceAny.
func (s SliceFoo) ToSliceAny() pg.SliceT0 {
	r := make(pg.SliceT0, len(s))
	for i, x := range s {
		r[i] = x
	}
	return r
}

// ToSliceFoo is a conversion function.
func ToSliceFoo(s pg.SliceT0) SliceFoo {
	r := make(SliceFoo, len(s))
	for i, x := range s {
		r[i] = x.(Foo)
	}
	return r
}

// SlicePFoo is a wrapper type to enable extension methods.
// Used with pseudo-generic functions for slices of pointers.
type SlicePFoo []*Foo

// Und converts to the underlying type.
func (s SlicePFoo) Und() []*Foo { return s }

// ToSliceAny is defined to implement pg.ConvertibleToSliceAny.
func (s SlicePFoo) ToSliceAny() pg.SliceT0 {
	r := make(pg.SliceT0, len(s))
	for i, x := range s {
		r[i] = x
	}
	return r
}

// ToSlicePFoo is a conversion function.
func ToSlicePFoo(s pg.SliceT0) SlicePFoo {
	r := make(SlicePFoo, len(s))
	for i, p := range s {
		r[i] = p.(*Foo)
	}
	return r
}

/////////////////////
// Types and standard converters for our second example data structure.

// Bar is another example data structure.
type Bar struct {
	z int
}

// SliceBar is a wrapper type to enable extension methods.
// Used with all pseudo-generic functions for slices.
type SliceBar []Bar

// Und converts to the underlying type.
func (s SliceBar) Und() []Bar { return s }

// ToSliceAny is defined to implement pg.ConvertibleToSliceAny.
func (s SliceBar) ToSliceAny() pg.SliceT0 {
	r := make(pg.SliceT0, len(s))
	for i, x := range s {
		r[i] = x
	}
	return r
}

// ToSliceBar is a conversion function.
func ToSliceBar(s pg.SliceT0) SliceBar {
	r := make(SliceBar, len(s))
	for i, x := range s {
		r[i] = x.(Bar)
	}
	return r
}

// SlicePBar is a wrapper type to enable extension methods.
// Used with pseudo-generic functions for slices of pointers.
type SlicePBar []*Bar

// Und converts to the underlying type.
func (s SlicePBar) Und() []*Bar { return s }

// ToSliceAny is defined to implement pg.ConvertibleToSliceAny.
func (s SlicePBar) ToSliceAny() pg.SliceT0 {
	r := make(pg.SliceT0, len(s))
	for i, x := range s {
		r[i] = x
	}
	return r
}

// ToSlicePBar is a conversion function.
func ToSlicePBar(s pg.SliceT0) SlicePBar {
	r := make(SlicePBar, len(s))
	for i, p := range s {
		r[i] = p.(*Bar)
	}
	return r
}

/////////////////////
// Helper functions for the examples below

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

/////////////////////
// Usage of basic functional combinators

// vars used as inputs to functions below

var xin SliceFoo = []Foo{{1, "w1"}, {22, "w22"}, {333, "w333"}}

var xinP SlicePFoo = xin.toP()

// Definition and use of specific map function from Foo to int.

func (s SliceFoo) MapInt(f func(Foo) int) []int {
	sa := s.ToSliceAny()
	fa := func(a pg.Any) pg.Any { return f(a.(Foo)) }
	ra := sa.Map(fa)
	return pg.ToSliceInt(ra)
}

func example_SliceFoo_MapInt() {
	f := func(i Foo) int { return i.v + len(i.w) }
	rslt := xin.MapInt(f)

	want := []int{3, 25, 337}
	assert.Equal(rslt, want)
}

// Definition and use of specific map function from Foo to Bar.

func (s SliceFoo) MapBar(f func(Foo) Bar) []Bar {
	sa := s.ToSliceAny()
	fa := func(a pg.Any) pg.Any { return f(a.(Foo)) }
	ra := sa.Map(fa)
	return ToSliceBar(ra)
}

func example_SliceFoo_MapBar() {
	f := func(i Foo) Bar { return Bar{i.v + 1} }
	rslt := xin.MapBar(f)

	want := []Bar{{2}, {23}, {334}}
	assert.Equal(rslt, want)
}

// Definition and use of specific map function from *Foo to *Bar.

func (s SlicePFoo) MapPBar(f func(*Foo) *Bar) []*Bar {
	sa := s.ToSliceAny()
	fa := func(a pg.Any) pg.Any { return f(a.(*Foo)) }
	ra := sa.Map(fa)
	return ToSlicePBar(ra)
}

func example_SlicePFoo_MapPBar() {
	f := func(p *Foo) *Bar { return &Bar{(*p).v + 1} }
	rslt := xinP.MapPBar(f)

	rsltV := SlicePBar(rslt).toV()

	want := []Bar{{2}, {23}, {334}}
	assert.Equal(rsltV, want)
}

// Definition and use of specific filter function.

func (s SliceFoo) Filter(pred func(Foo) bool) SliceFoo {
	sa := s.ToSliceAny()
	preda := func(a pg.Any) bool { return pred(a.(Foo)) }
	ra := sa.Filter(preda)
	return ToSliceFoo(ra)
}

func example_SliceFoo_Filter() {
	pred := func(i Foo) bool { return i.v%2 != 0 }
	rslt := xin.Filter(pred)

	var want SliceFoo = []Foo{{1, "w1"}, {333, "w333"}}
	assert.Equal(rslt, want)
}

// Definition and use of specific fold function.

func (s SliceFoo) FoldInt(z int, op func(int, Foo) int) int {
	sa := s.ToSliceAny()
	opa := func(z pg.Any, a pg.Any) pg.Any { return op(z.(int), a.(Foo)) }
	ra := sa.Fold(z, opa)
	return ra.(int)
}

func example_SliceFoo_FoldInt() {
	op := func(z int, x Foo) int { return z + x.v }
	rslt := xin.FoldInt(0, op)

	want := 356
	assert.Equal(rslt, want)
}

// Run examples.
func Example() {
	example_SliceFoo_MapInt()
	example_SliceFoo_MapBar()
	example_SlicePFoo_MapPBar()
	example_SliceFoo_Filter()
	example_SliceFoo_FoldInt()
	fmt.Println("End of Example")
	// Output: End of Example
}
