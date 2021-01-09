package collections_test

import (
	"testing"

	c "github.com/pvillela/go-light-collections/pkg/collections"
	"github.com/stretchr/testify/assert"
)

////
// Slices used as inputs to functions below. Cloned each time to avoid nasty side-effects
// in tests.

func sEmptyAny() c.SliceAny {
	return []c.Any{}
}

func sFooAny() c.SliceAny {
	var sFoo SliceFoo = []Foo{{1, "w1"}, {22, "w22"}, {333, "w333"}, {4444, "w4444"}, {22, "w22"}}
	return sFoo.ToSliceAny()
}

func sFooAny1() c.SliceAny {
	var sOtherFoo SliceFoo = []Foo{{0, "w"}, {22, "w22"}, {55555, "w55555"}}
	return sOtherFoo.ToSliceAny()
}

////
// Tests

func TestLength(t *testing.T) {
	cases := []struct {
		msg      string
		receiver c.SliceAny
		want     int
	}{
		{"Length: non-empty slice", sFooAny(), 5},
		{"Length: empty slice", sEmptyAny(), 0},
	}

	for _, cs := range cases {
		got := cs.receiver.Length()
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestContains(t *testing.T) {
	cases := []struct {
		msg  string
		arg  Foo
		want bool
	}{
		{"Cotains: present", Foo{22, "w22"}, true},
		{"Contains: absent", Foo{22, "xyz"}, false},
	}

	for _, cs := range cases {
		got := sFooAny().Contains(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestContainsAll(t *testing.T) {
	cases := []struct {
		msg  string
		arg  c.SliceAny
		want bool
	}{
		{"ContainsAll: subset", append(sFooAny()[2:2], sFooAny()[1]), true},
		{"ContainsAll: not subset", append(sFooAny()[1:1], Foo{22, "xyz"}), false},
	}

	for _, cs := range cases {
		got := sFooAny().ContainsAll(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestGet(t *testing.T) {
	size := len(sFooAny())
	cases := []struct {
		msg  string
		arg  int
		want c.Any
	}{
		{"Get: from middle", 2, sFooAny()[2]},
		{"Get: from beginning", 0, sFooAny()[0]},
		{"Get: from end", size - 1, sFooAny()[size-1]},
	}

	for _, cs := range cases {
		got := sFooAny().Get(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestIndexOf(t *testing.T) {
	cases := []struct {
		msg      string
		receiver c.SliceAny
		arg      Foo
		want     int
	}{
		{"IndexOf: non-empty, present", sFooAny(), Foo{22, "w22"}, 1},
		{"IndexOf: non-empty, absent", sFooAny(), Foo{0, "xyz"}, -1},
		{"IndexOf: empty", sEmptyAny(), Foo{0, "xyz"}, -1},
	}

	for _, cs := range cases {
		got := cs.receiver.IndexOf(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestIsEmpty(t *testing.T) {
	cases := []struct {
		msg      string
		receiver c.SliceAny
		want     bool
	}{
		{"IsEmpty: non-empty", sFooAny(), false},
		{"IsEmpty: empty", sEmptyAny(), true},
	}

	for _, cs := range cases {
		got := cs.receiver.IsEmpty()
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestLastIndexOf(t *testing.T) {
	cases := []struct {
		msg      string
		receiver c.SliceAny
		arg      Foo
		want     int
	}{
		{"LastIndexOf: non-empty, present", sFooAny(), Foo{22, "w22"}, 4},
		{"LastIndexOf: non-empty, absent", sFooAny(), Foo{0, "xyz"}, -1},
		{"LastIndexOf: empty", sEmptyAny(), Foo{0, "xyz"}, -1},
	}

	for _, cs := range cases {
		got := cs.receiver.LastIndexOf(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSubSlice(t *testing.T) {
	size := len(sFooAny())
	cases := []struct {
		msg      string
		receiver c.SliceAny
		arg1     int
		arg2     int
		want     c.SliceAny
	}{
		{"SubSlice: nonempty - from beginning", sFooAny(), 0, 2, sFooAny()[:2]},
		{"SubSlice: nonempty - from middle", sFooAny(), 1, 3, sFooAny()[1:3]},
		{"SubSlice: nonempty - from end", sFooAny(), size - 3, size, sFooAny()[size-3:]},
		{"SubSlice: nonempty - empty sub-slice", sFooAny(), 2, 2, sFooAny()[2:2]},
		{"SubSlice: empty - empty sub-slice", sEmptyAny(), 0, 0, sEmptyAny()[0:0]},
	}

	for _, cs := range cases {
		got := cs.receiver.SubSlice(cs.arg1, cs.arg2)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestAll(t *testing.T) {
	pred1 := func(a c.Any) bool { return a.(Foo).v1 > 0 }
	pred2 := func(a c.Any) bool { return a.(Foo).v1%2 == 0 }
	pred3 := func(a c.Any) bool { return a.(Foo).v1 < 0 }

	cases := []struct {
		msg      string
		receiver c.SliceAny
		arg      func(c.Any) bool
		want     bool
	}{
		{"All: pred matches all", sFooAny(), pred1, true},
		{"All: pred matches some", sFooAny(), pred2, false},
		{"All: pred matches none", sFooAny(), pred3, false},
		{"All: empty receiver", sEmptyAny(), pred2, true},
	}

	for _, cs := range cases {
		got := cs.receiver.All(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestAny(t *testing.T) {
	pred1 := func(a c.Any) bool { return a.(Foo).v1 > 0 }
	pred2 := func(a c.Any) bool { return a.(Foo).v1%2 == 0 }
	pred3 := func(a c.Any) bool { return a.(Foo).v1 < 0 }

	cases := []struct {
		msg      string
		receiver c.SliceAny
		arg      func(c.Any) bool
		want     bool
	}{
		{"Any: pred matches all", sFooAny(), pred1, true},
		{"Any: pred matches some", sFooAny(), pred2, true},
		{"Any: pred matches none", sFooAny(), pred3, false},
		{"Any: empty receiver", sEmptyAny(), pred2, false},
	}

	for _, cs := range cases {
		got := cs.receiver.All(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestCount(t *testing.T) {
	pred1 := func(a c.Any) bool { return a.(Foo).v1 > 0 }
	pred2 := func(a c.Any) bool { return a.(Foo).v1%2 == 0 }
	pred3 := func(a c.Any) bool { return a.(Foo).v1 < 0 }

	cases := []struct {
		msg      string
		receiver c.SliceAny
		arg      func(c.Any) bool
		want     int
	}{
		{"Count: pred matches all", sFooAny(), pred1, len(sFooAny())},
		{"Count: pred matches some", sFooAny(), pred2, 3},
		{"Count: pred matches none", sFooAny(), pred3, 0},
		{"Count: empty receiver", sEmptyAny(), pred2, 0},
	}

	for _, cs := range cases {
		got := cs.receiver.All(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestDrop(t *testing.T) {
	size := len(sFooAny())
	cases := []struct {
		msg      string
		receiver c.SliceAny
		arg      int
		want     c.SliceAny
	}{
		{"Drop: some", sFooAny(), 2, sFooAny()[2:]},
		{"Drop: all", sFooAny(), size, sEmptyAny()},
		{"Drop: none", sFooAny(), 0, sFooAny()},
		{"Drop: more than length", sFooAny(), size + 5, sEmptyAny()},
		{"Drop: empty receiver", sEmptyAny(), 1, sEmptyAny()},
	}

	for _, cs := range cases {
		got := cs.receiver.Drop(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestDropLast(t *testing.T) {
	size := len(sFooAny())
	cases := []struct {
		msg      string
		receiver c.SliceAny
		arg      int
		want     c.SliceAny
	}{
		{"DropLast: some", sFooAny(), 2, sFooAny()[:size-2]},
		{"DropLast: all", sFooAny(), size, sEmptyAny()},
		{"DropLast: none", sFooAny(), 0, sFooAny()},
		{"DropLast: more than length", sFooAny(), size + 5, sEmptyAny()},
		{"DropLast: empty receiver", sEmptyAny(), 1, sEmptyAny()},
	}

	for _, cs := range cases {
		got := cs.receiver.DropLast(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestDropLastWhile(t *testing.T) {
	pred1 := func(a c.Any) bool { return a.(Foo).v1 > 0 }
	pred2 := func(a c.Any) bool { return a.(Foo).v1%2 == 0 }
	pred3 := func(a c.Any) bool { return a.(Foo).v1 < 0 }

	cases := []struct {
		msg      string
		receiver c.SliceAny
		arg      func(c.Any) bool
		want     c.SliceAny
	}{
		{"DropLastWhile: pred matches all", sFooAny(), pred1, sEmptyAny()},
		{"DropLastWhile: pred matches some", sFooAny(), pred2, SliceFoo([]Foo{{1, "w1"}, {22, "w22"}, {333, "w333"}}).ToSliceAny()},
		{"DropLastWhile: pred matches none", sFooAny(), pred3, sFooAny()},
		{"DropLastWhile: empty receiver", sEmptyAny(), pred2, sEmptyAny()},
	}

	for _, cs := range cases {
		got := cs.receiver.DropLastWhile(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestDropWhile(t *testing.T) {
	pred1 := func(a c.Any) bool { return a.(Foo).v1 > 0 }
	pred2 := func(a c.Any) bool { return a.(Foo).v1%2 == 1 }
	pred3 := func(a c.Any) bool { return a.(Foo).v1 < 0 }

	cases := []struct {
		msg      string
		receiver c.SliceAny
		arg      func(c.Any) bool
		want     c.SliceAny
	}{
		{"DropWhile: pred matches all", sFooAny(), pred1, sEmptyAny()},
		{"DropWhile: pred matches some", sFooAny(), pred2, SliceFoo([]Foo{{22, "w22"}, {333, "w333"}, {4444, "w4444"}, {22, "w22"}}).ToSliceAny()},
		{"DropWhile: pred matches none", sFooAny(), pred3, sFooAny()},
		{"DropWhile: empty receiver", sEmptyAny(), pred2, sEmptyAny()},
	}

	for _, cs := range cases {
		got := cs.receiver.DropWhile(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestFilter(t *testing.T) {
	pred1 := func(a c.Any) bool { return a.(Foo).v1 > 0 }
	pred2 := func(a c.Any) bool { return a.(Foo).v1%2 == 0 }
	pred3 := func(a c.Any) bool { return a.(Foo).v1 < 0 }

	cases := []struct {
		msg      string
		receiver c.SliceAny
		arg      func(c.Any) bool
		want     c.SliceAny
	}{
		{"Filter: pred matches all", sFooAny(), pred1, sFooAny()},
		{"Filter: pred matches some", sFooAny(), pred2, SliceFoo([]Foo{{22, "w22"}, {4444, "w4444"}, {22, "w22"}}).ToSliceAny()},
		{"Filter: pred matches none", sFooAny(), pred3, sEmptyAny()},
		{"Filter: empty receiver", sEmptyAny(), pred2, sEmptyAny()},
	}

	for _, cs := range cases {
		got := cs.receiver.Filter(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestFilterNot(t *testing.T) {
	pred1 := func(a c.Any) bool { return a.(Foo).v1 > 0 }
	pred2 := func(a c.Any) bool { return a.(Foo).v1%2 == 1 }
	pred3 := func(a c.Any) bool { return a.(Foo).v1 < 0 }

	cases := []struct {
		msg      string
		receiver c.SliceAny
		arg      func(c.Any) bool
		want     c.SliceAny
	}{
		{"FilterNot: pred matches all", sFooAny(), pred1, sFooAny()},
		{"FilterNot: pred matches some", sFooAny(), pred2, SliceFoo([]Foo{{22, "w22"}, {4444, "w4444"}, {22, "w22"}}).ToSliceAny()},
		{"FilterNot: pred matches none", sFooAny(), pred3, sEmptyAny()},
		{"FilterNot: empty receiver", sEmptyAny(), pred2, sEmptyAny()},
	}

	for _, cs := range cases {
		got := cs.receiver.FilterNot(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestFind(t *testing.T) {
	cases := []struct {
		msg      string
		receiver c.SliceAny
		arg      c.Any
		want     c.Any
	}{
		{"Find: present", sFooAny(), Foo{22, "w222"}, Foo{22, "w222"}},
		{"Find: absent", sFooAny(), Foo{22, "xyz"}, nil},
		{"Find: empty receiver", sEmptyAny(), Foo{22, "w222"}, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.Find(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestFirst(t *testing.T) {
	cases := []struct {
		msg      string
		receiver c.SliceAny
		want     c.Any
	}{
		{"First: non-empty", sFooAny(), Foo{1, "w1"}},
		{"First: empty", sEmptyAny(), nil},
	}

	for _, cs := range cases {
		if !cs.receiver.IsEmpty() {
			got := cs.receiver.First()
			assert.Equal(t, cs.want, got, cs.msg)
		} else {
			var ptf assert.PanicTestFunc = func() { cs.receiver.First() }
			assert.Panics(t, ptf, cs.msg)
		}
	}
}

func TestFlatMap(t *testing.T) {

	// f := func(a )
}

func TestFold(t *testing.T) {

}

func TestForEach(t *testing.T) {

}

func TestGroupBy(t *testing.T) {

}

func TestIndexOfFirst(t *testing.T) {

}

func TestIndexOfLast(t *testing.T) {

}

func TestIsNotEmpty(t *testing.T) {

}

func TestLast(t *testing.T) {

}

func TestMap(t *testing.T) {

}

func TestMaxWithOrNil(t *testing.T) {

}

func TestMinus(t *testing.T) {

}

func TestMinusElement(t *testing.T) {

}

func TestMinWithOrNil(t *testing.T) {

}

func TestPartition(t *testing.T) {

}

func TestPlus(t *testing.T) {

}

func TestPlusElement(t *testing.T) {

}

func TestReduceOrNil(t *testing.T) {

}

func TestReversed(t *testing.T) {

}

func TestSortedWith(t *testing.T) {

}

func TestTake(t *testing.T) {

}

func TestTakeLast(t *testing.T) {

}

func TestTakeLastWhile(t *testing.T) {

}

func TestTakeWhile(t *testing.T) {

}

func TestToSlice(t *testing.T) {

}

func TestZip(t *testing.T) {

}
