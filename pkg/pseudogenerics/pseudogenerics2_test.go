package pseudogenerics_test

import (
	"fmt"

	pg "github.com/pvillela/GoSimpleCollections/pkg/pseudogenerics"
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
func (s SliceFoo) ToSliceAny() pg.SliceAny {
	r := make(pg.SliceAny, len(s))
	for i, x := range s {
		r[i] = x
	}
	return r
}

// ToSliceFoo is a conversion function.
func ToSliceFoo(s pg.SliceAny) SliceFoo {
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
func (s SlicePFoo) ToSliceAny() pg.SliceAny {
	r := make(pg.SliceAny, len(s))
	for i, x := range s {
		r[i] = x
	}
	return r
}

// ToSlicePFoo is a conversion function.
func ToSlicePFoo(s pg.SliceAny) SlicePFoo {
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
func (s SliceBar) ToSliceAny() pg.SliceAny {
	r := make(pg.SliceAny, len(s))
	for i, x := range s {
		r[i] = x
	}
	return r
}

// ToSliceBar is a conversion function.
func ToSliceBar(s pg.SliceAny) SliceBar {
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
func (s SlicePBar) ToSliceAny() pg.SliceAny {
	r := make(pg.SliceAny, len(s))
	for i, x := range s {
		r[i] = x
	}
	return r
}

// ToSlicePBar is a conversion function.
func ToSlicePBar(s pg.SliceAny) SlicePBar {
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

var xin = []Foo{{1, "w1"}, {22, "w22"}, {333, "w333"}}

var xinP = SliceFoo(xin).toP()

// Definition and use of specific map function from Foo to int.

func mapSliceFooToInt(s []Foo, f func(Foo) int) []int {
	sa := SliceFoo(s)
	fa := func(a pg.AnyT0) pg.AnyT0 { return f(a.(Foo)) }
	ra := pg.MapSlice(sa, fa)
	return pg.ToSliceInt(ra)
}

func example_MapSlice_fooToInt() {
	f := func(i Foo) int { return i.v + len(i.w) }
	rslt := mapSliceFooToInt(xin, f)

	want := []int{3, 25, 337}
	assert.Equal(rslt, want)
}

// Definition and use of specific map function from Foo to Bar.

func mapSliceFooToBar(s []Foo, f func(Foo) Bar) []Bar {
	sa := SliceFoo(s)
	fa := func(a pg.AnyT0) pg.AnyT0 { return f(a.(Foo)) }
	ra := pg.MapSlice(sa, fa)
	return ToSliceBar(ra)
}

func example_MapSlice_fooToBar() {
	f := func(i Foo) Bar { return Bar{i.v + 1} }
	rslt := mapSliceFooToBar(xin, f)

	want := []Bar{{2}, {23}, {334}}
	assert.Equal(rslt, want)
}

// Definition and use of specific map function from *Foo to *Bar.

func mapSlicePFooToPBar(s []*Foo, f func(*Foo) *Bar) []*Bar {
	sa := SlicePFoo(s)
	fa := func(a pg.AnyT0) pg.AnyT0 { return f(a.(*Foo)) }
	ra := pg.MapSlice(sa, fa)
	return ToSlicePBar(ra)
}

func example_MapSlice_fooToBar_withPointers() {
	f := func(p *Foo) *Bar { return &Bar{(*p).v + 1} }
	rslt := mapSlicePFooToPBar(xinP, f)

	rsltV := SlicePBar(rslt).toV()

	want := []Bar{{2}, {23}, {334}}
	assert.Equal(rsltV, want)
}

// Definition and use of specific filter function.

func filterSliceFoo(s []Foo, pred func(Foo) bool) []Foo {
	sa := SliceFoo(s)
	preda := func(a pg.AnyT0) bool { return pred(a.(Foo)) }
	ra := pg.FilterSlice(sa, preda)
	return ToSliceFoo(ra)
}

func example_FilterSlice_foo() {
	pred := func(i Foo) bool { return i.v%2 != 0 }
	rslt := filterSliceFoo(xin, pred)

	want := []Foo{{1, "w1"}, {333, "w333"}}
	assert.Equal(rslt, want)
}

// Definition and use of specific fold function.

func foldLeftSliceIntFoo(s []Foo, z int, op func(int, Foo) int) int {
	sa := SliceFoo(s)
	opa := func(z pg.AnyT0, a pg.AnyT0) pg.AnyT0 { return op(z.(int), a.(Foo)) }
	ra := pg.FoldLeftSlice(sa, z, opa)
	return ra.(int)
}

func exampleFoldLeftSlice_intFoo() {
	op := func(z int, x Foo) int { return z + x.v }
	rslt := foldLeftSliceIntFoo(xin, 0, op)

	want := 356
	assert.Equal(rslt, want)
}

// Run examples.
func Example() {
	example_MapSlice_fooToInt()
	example_MapSlice_fooToBar()
	example_MapSlice_fooToBar_withPointers()
	example_FilterSlice_foo()
	exampleFoldLeftSlice_intFoo()
	fmt.Println("End of Example")
	// Output: End of Example
}
