package collections_test

import (
	"testing"

	c "github.com/pvillela/go-light-collections/pkg/collections"
	"github.com/stretchr/testify/assert"
)

////
// Slices used as inputs to functions below. Cloned each time to avoid nasty side-effects
// in tests.

func sEmpty() c.SliceAny {
	return []c.Any{}
}

func sFoo() c.SliceAny {
	var s SliceFoo = []Foo{{1, "w1"}, {22, "w22"}, {333, "w333"}, {4444, "w4444"}, {22, "w22"}}
	return s.ToSliceAny()
}

func sFooOther() c.SliceAny {
	var s SliceFoo = []Foo{{0, "w"}, {22, "w22"}, {55555, "w55555"}}
	return s.ToSliceAny()
}

////
// Tests

func TestLength(t *testing.T) {
	cases := []struct {
		msg      string
		receiver c.SliceAny
		want     int
	}{
		{"Length: non-empty slice", sFoo(), 5},
		{"Length: empty slice", sEmpty(), 0},
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
		got := sFoo().Contains(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestContainsAll(t *testing.T) {
	cases := []struct {
		msg  string
		arg  c.SliceAny
		want bool
	}{
		{"ContainsAll: subset", append(sFoo()[2:2], sFoo()[1]), true},
		{"ContainsAll: not subset", append(sFoo()[1:1], Foo{22, "xyz"}), false},
	}

	for _, cs := range cases {
		got := sFoo().ContainsAll(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestGet(t *testing.T) {
	size := len(sFoo())
	cases := []struct {
		msg  string
		arg  int
		want c.Any
	}{
		{"Get: from middle", 2, sFoo()[2]},
		{"Get: from beginning", 0, sFoo()[0]},
		{"Get: from end", size - 1, sFoo()[size-1]},
	}

	for _, cs := range cases {
		got := sFoo().Get(cs.arg)
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
		{"IndexOf: non-empty, present", sFoo(), Foo{22, "w22"}, 1},
		{"IndexOf: non-empty, absent", sFoo(), Foo{0, "xyz"}, -1},
		{"IndexOf: empty", sEmpty(), Foo{0, "xyz"}, -1},
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
		{"IsEmpty: non-empty", sFoo(), false},
		{"IsEmpty: empty", sEmpty(), true},
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
		{"LastIndexOf: non-empty, present", sFoo(), Foo{22, "w22"}, 4},
		{"LastIndexOf: non-empty, absent", sFoo(), Foo{0, "xyz"}, -1},
		{"LastIndexOf: empty", sEmpty(), Foo{0, "xyz"}, -1},
	}

	for _, cs := range cases {
		got := cs.receiver.LastIndexOf(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSubSlice(t *testing.T) {
	size := len(sFoo())
	cases := []struct {
		msg      string
		receiver c.SliceAny
		arg1     int
		arg2     int
		want     c.SliceAny
	}{
		{"SubSlice: nonempty - from beginning", sFoo(), 0, 2, sFoo()[:2]},
		{"SubSlice: nonempty - from middle", sFoo(), 1, 3, sFoo()[1:3]},
		{"SubSlice: nonempty - from end", sFoo(), size - 3, size, sFoo()[size-3:]},
		{"SubSlice: nonempty - empty sub-slice", sFoo(), 2, 2, sFoo()[2:2]},
		{"SubSlice: empty - empty sub-slice", sEmpty(), 0, 0, sEmpty()[0:0]},
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
		{"All: pred matches all", sFoo(), pred1, true},
		{"All: pred matches some", sFoo(), pred2, false},
		{"All: pred matches none", sFoo(), pred3, false},
		{"All: empty receiver", sEmpty(), pred2, true},
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
		{"Any: pred matches all", sFoo(), pred1, true},
		{"Any: pred matches some", sFoo(), pred2, true},
		{"Any: pred matches none", sFoo(), pred3, false},
		{"Any: empty receiver", sEmpty(), pred2, false},
	}

	for _, cs := range cases {
		got := cs.receiver.Any(cs.arg)
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
		{"Count: pred matches all", sFoo(), pred1, len(sFoo())},
		{"Count: pred matches some", sFoo(), pred2, 3},
		{"Count: pred matches none", sFoo(), pred3, 0},
		{"Count: empty receiver", sEmpty(), pred2, 0},
	}

	for _, cs := range cases {
		got := cs.receiver.Count(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestDrop(t *testing.T) {
	size := len(sFoo())
	cases := []struct {
		msg      string
		receiver c.SliceAny
		arg      int
		want     c.SliceAny
	}{
		{"Drop: some", sFoo(), 2, sFoo()[2:]},
		{"Drop: all", sFoo(), size, sEmpty()},
		{"Drop: none", sFoo(), 0, sFoo()},
		{"Drop: more than length", sFoo(), size + 5, sEmpty()},
		{"Drop: empty receiver", sEmpty(), 1, sEmpty()},
	}

	for _, cs := range cases {
		got := cs.receiver.Drop(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestDropLast(t *testing.T) {
	size := len(sFoo())
	cases := []struct {
		msg      string
		receiver c.SliceAny
		arg      int
		want     c.SliceAny
	}{
		{"DropLast: some", sFoo(), 2, sFoo()[:size-2]},
		{"DropLast: all", sFoo(), size, sEmpty()},
		{"DropLast: none", sFoo(), 0, sFoo()},
		{"DropLast: more than length", sFoo(), size + 5, sEmpty()},
		{"DropLast: empty receiver", sEmpty(), 1, sEmpty()},
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
		{"DropLastWhile: pred matches all", sFoo(), pred1, sEmpty()},
		{"DropLastWhile: pred matches some", sFoo(), pred2, SliceFoo([]Foo{{1, "w1"}, {22, "w22"}, {333, "w333"}}).ToSliceAny()},
		{"DropLastWhile: pred matches none", sFoo(), pred3, sFoo()},
		{"DropLastWhile: empty receiver", sEmpty(), pred2, sEmpty()},
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
		{"DropWhile: pred matches all", sFoo(), pred1, sEmpty()},
		{"DropWhile: pred matches some", sFoo(), pred2, SliceFoo([]Foo{{22, "w22"}, {333, "w333"}, {4444, "w4444"}, {22, "w22"}}).ToSliceAny()},
		{"DropWhile: pred matches none", sFoo(), pred3, sFoo()},
		{"DropWhile: empty receiver", sEmpty(), pred2, sEmpty()},
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
		{"Filter: pred matches all", sFoo(), pred1, sFoo()},
		{"Filter: pred matches some", sFoo(), pred2, SliceFoo([]Foo{{22, "w22"}, {4444, "w4444"}, {22, "w22"}}).ToSliceAny()},
		{"Filter: pred matches none", sFoo(), pred3, sEmpty()},
		{"Filter: empty receiver", sEmpty(), pred2, sEmpty()},
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
		{"FilterNot: pred matches all", sFoo(), pred1, sEmpty()},
		{"FilterNot: pred matches some", sFoo(), pred2, SliceFoo([]Foo{{22, "w22"}, {4444, "w4444"}, {22, "w22"}}).ToSliceAny()},
		{"FilterNot: pred matches none", sFoo(), pred3, sFoo()},
		{"FilterNot: empty receiver", sEmpty(), pred2, sEmpty()},
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
		{"Find: present", sFoo(), Foo{22, "w22"}, Foo{22, "w22"}},
		{"Find: absent", sFoo(), Foo{22, "xyz"}, nil},
		{"Find: empty receiver", sEmpty(), Foo{22, "w222"}, nil},
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
		{"First: non-empty", sFoo(), Foo{1, "w1"}},
		{"First: empty", sEmpty(), nil},
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
	var sInt c.SliceInt = []int{1, 2, 3}
	f := func(a c.Any) []c.Any {
		n := a.(int)
		s := make([]c.Any, n)
		for i := range s {
			s[i] = n
		}
		return s
	}

	cases := []struct {
		msg      string
		receiver c.SliceAny
		arg      func(c.Any) []c.Any
		want     []c.Any
	}{
		{"FlatMap: non-empty receiver", sInt.ToSliceAny(), f, []c.Any{1, 2, 2, 3, 3, 3}},
		{"FlatMap: empty receiver", sEmpty(), f, []c.Any{}},
	}

	for _, cs := range cases {
		got := cs.receiver.FlatMap(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestFold(t *testing.T) {
	op := func(z c.Any, a c.Any) c.Any { return z.(int) + a.(Foo).v1 }

	cases := []struct {
		msg      string
		receiver c.SliceAny
		arg1     int
		arg2     func(z c.Any, a c.Any) c.Any
		want     c.Any
	}{
		{"Fold: non-empty receiver", sFoo(), 1, op, 1 + 1 + 22 + 333 + 4444 + 22},
		{"Fold: empty receiver", sEmpty(), 42, op, 42},
	}

	for _, cs := range cases {
		got := cs.receiver.Fold(cs.arg1, cs.arg2)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestForEach(t *testing.T) {
	cases := []struct {
		msg      string
		receiver c.SliceAny
		want     []c.Any
	}{
		{"ForEach: non-empty receiver", sFoo(), []c.Any{1, 22, 333, 4444, 22}},
		{"ForEach: empty receiver", sEmpty(), []c.Any{}},
	}

	for _, cs := range cases {
		got := []c.Any{}
		f := func(a c.Any) {
			got = append(got, a.(Foo).v1)
		}

		cs.receiver.ForEach(f)
		assert.Equal(t, cs.want, got, cs.msg)
	}
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
