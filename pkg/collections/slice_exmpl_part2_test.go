package collections_test

import (
	"fmt"

	c "github.com/pvillela/go-light-collections/pkg/collections"
	"github.com/pvillela/go-light-collections/pkg/util/assert"
)

/////////////////////
// Part 2 of example -- usage of basic functional combinators

// vars used as inputs to functions below

var xin SliceFoo = []Foo{{1, "w1"}, {22, "w22"}, {333, "w333"}}

var xinP SlicePFoo = xin.toP()

// Definition and use of specific map function from Foo to int.

func (s SliceFoo) MapInt(f func(Foo) int) []int {
	sa := s.ToSliceAny()
	fa := func(a c.Any) c.Any { return f(a.(Foo)) }
	ra := sa.Map(fa)
	return c.ToSliceInt(ra)
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
	fa := func(a c.Any) c.Any { return f(a.(Foo)) }
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
	fa := func(a c.Any) c.Any { return f(a.(*Foo)) }
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
	preda := func(a c.Any) bool { return pred(a.(Foo)) }
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
	opa := func(z c.Any, a c.Any) c.Any { return op(z.(int), a.(Foo)) }
	ra := sa.Fold(z, opa)
	return ra.(int)
}

func example_SliceFoo_FoldInt() {
	op := func(z int, x Foo) int { return z + x.v }
	rslt := xin.FoldInt(0, op)

	want := 356
	assert.Equal(rslt, want)
}

// Run example.
func ExampleSliceAny_part2() {
	example_SliceFoo_MapInt()
	example_SliceFoo_MapBar()
	example_SlicePFoo_MapPBar()
	example_SliceFoo_Filter()
	example_SliceFoo_FoldInt()
	fmt.Println("End of Example")
	// Output: End of Example
}
