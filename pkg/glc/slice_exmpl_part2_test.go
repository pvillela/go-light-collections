/*
 * Copyright © 2021 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package glc_test

import (
	"fmt"

	c "github.com/pvillela/go-light-collections/pkg/glc"
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
	f := func(i Foo) int { return i.v1 + len(i.v2) }
	rslt := xin.MapInt(f)

	want := []int{3, 25, 337}
	assert.Equal(want, rslt)
}

// Definition and use of specific map function from Foo to Bar.

func (s SliceFoo) MapBar(f func(Foo) Bar) []Bar {
	sa := s.ToSliceAny()
	fa := func(a c.Any) c.Any { return f(a.(Foo)) }
	ra := sa.Map(fa)
	return ToSliceBar(ra)
}

func example_SliceFoo_MapBar() {
	f := func(x Foo) Bar { return Bar{x.v1 + 1, []string{x.v2}} }
	rslt := xin.MapBar(f)

	want := []Bar{{2, []string{"w1"}}, {23, []string{"w22"}}, {334, []string{"w333"}}}
	assert.Equal(want, rslt)
}

// Definition and use of specific map function from *Foo to *Bar.

func (s SlicePFoo) MapPBar(f func(*Foo) *Bar) []*Bar {
	sa := s.ToSliceAny()
	fa := func(a c.Any) c.Any { return f(a.(*Foo)) }
	ra := sa.Map(fa)
	return ToSlicePBar(ra)
}

func example_SlicePFoo_MapPBar() {
	f := func(p *Foo) *Bar { return &Bar{(*p).v1 + 1, []string{(*p).v2}} }
	rslt := xinP.MapPBar(f)

	rsltV := SlicePBar(rslt).toV()

	want := []Bar{{2, []string{"w1"}}, {23, []string{"w22"}}, {334, []string{"w333"}}}
	assert.Equal(want, rsltV)
}

// Definition and use of specific filter function.

func (s SliceFoo) Filter(pred func(Foo) bool) SliceFoo {
	sa := s.ToSliceAny()
	preda := func(a c.Any) bool { return pred(a.(Foo)) }
	ra := sa.Filter(preda)
	return ToSliceFoo(ra)
}

func example_SliceFoo_Filter() {
	pred := func(i Foo) bool { return i.v1%2 != 0 }
	rslt := xin.Filter(pred)

	var want SliceFoo = []Foo{{1, "w1"}, {333, "w333"}}
	assert.Equal(want, rslt)
}

// Definition and use of specific fold function.

func (s SliceFoo) FoldInt(z int, op func(int, Foo) int) int {
	sa := s.ToSliceAny()
	opa := func(z c.Any, a c.Any) c.Any { return op(z.(int), a.(Foo)) }
	ra := sa.Fold(z, opa)
	return ra.(int)
}

func example_SliceFoo_FoldInt() {
	op := func(z int, x Foo) int { return z + x.v1 }
	rslt := xin.FoldInt(0, op)

	want := 356
	assert.Equal(want, rslt)
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
