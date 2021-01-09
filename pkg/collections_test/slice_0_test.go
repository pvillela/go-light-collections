package collections_test

import (
	"testing"

	c "github.com/pvillela/go-light-collections/pkg/collections"
	"github.com/stretchr/testify/assert"
)

////
// Slices used as inputs to functions below. Cloned each time to avoid nasty side-effects
// in tests.

func sAny() c.SliceAny {
	var sFoo SliceFoo = []Foo{{1, "w1"}, {22, "w22"}, {333, "w333"}, {4444, "w4444"}, {22, "w22"}}
	return sFoo.ToSliceAny()
}

func sOtherAny() c.SliceAny {
	var sOtherFoo SliceFoo = []Foo{{0, "w"}, {22, "w22"}, {55555, "w55555"}}
	return sOtherFoo.ToSliceAny()
}

func sEmptyAny() c.SliceAny {
	var emptySliceFoo SliceFoo = []Foo{}
	return emptySliceFoo.ToSliceAny()
}

////
// Tests

func TestLength(t *testing.T) {
	cases := []struct {
		msg      string
		receiver c.SliceAny
		want     int
	}{
		{"Length: non-empty slice", sAny(), 5},
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
		got := sAny().Contains(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestContainsAll(t *testing.T) {
	cases := []struct {
		msg  string
		arg  c.SliceAny
		want bool
	}{
		{"ContainsAll: subset", append(sAny()[2:2], sAny()[1]), true},
		{"ContainsAll: not subset", append(sAny()[1:1], Foo{22, "xyz"}), false},
	}

	for _, cs := range cases {
		got := sAny().ContainsAll(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestGet(t *testing.T) {
	size := len(sAny())
	cases := []struct {
		msg  string
		arg  int
		want c.Any
	}{
		{"Get: from middle", 2, sAny()[2]},
		{"Get: from beginning", 0, sAny()[0]},
		{"Get: from end", size - 1, sAny()[size-1]},
	}

	for _, cs := range cases {
		got := sAny().Get(cs.arg)
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
		{"IndexOf: non-empty, present", sAny(), Foo{22, "w22"}, 1},
		{"IndexOf: non-empty, absent", sAny(), Foo{0, "xyz"}, -1},
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
		{"IsEmpty: non-empty", sAny(), false},
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
		{"LastIndexOf: non-empty, present", sAny(), Foo{22, "w22"}, 4},
		{"LastIndexOf: non-empty, absent", sAny(), Foo{0, "xyz"}, -1},
		{"LastIndexOf: empty", sEmptyAny(), Foo{0, "xyz"}, -1},
	}

	for _, cs := range cases {
		got := cs.receiver.LastIndexOf(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSubSlice(t *testing.T) {
	size := len(sAny())
	cases := []struct {
		msg      string
		receiver c.SliceAny
		arg1     int
		arg2     int
		want     c.SliceAny
	}{
		{"SubSlice: nonempty - from beginning", sAny(), 0, 2, sAny()[:2]},
		{"SubSlice: nonempty - from middle", sAny(), 1, 3, sAny()[1:3]},
		{"SubSlice: nonempty - from end", sAny(), size - 3, size, sAny()[size-3:]},
		{"SubSlice: nonempty - empty sub-slice", sAny(), 2, 2, sAny()[2:2]},
		{"SubSlice: empty - empty sub-slice", sEmptyAny(), 0, 0, sEmptyAny()[0:0]},
	}

	for _, cs := range cases {
		got := cs.receiver.SubSlice(cs.arg1, cs.arg2)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestAll(t *testing.T) {
	pred1 := func(a c.Any) bool { return a.(Foo).v > 0 }
	pred2 := func(a c.Any) bool { return a.(Foo).v%2 == 0 }
	pred3 := func(a c.Any) bool { return a.(Foo).v < 0 }

	cases := []struct {
		msg      string
		receiver c.SliceAny
		arg      func(c.Any) bool
		want     bool
	}{
		{"All: pred matches all", sAny(), pred1, true},
		{"All: pred matches some", sAny(), pred2, false},
		{"All: pred matches none", sAny(), pred3, false},
		{"All: empty receiver", sEmptyAny(), pred2, true},
	}

	for _, cs := range cases {
		got := cs.receiver.All(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestAny(t *testing.T) {

}

func TestCount(t *testing.T) {

}

func TestDrop(t *testing.T) {

}

func TestDropLast(t *testing.T) {

}

func TestDropLastWhile(t *testing.T) {

}

func TestDropWhile(t *testing.T) {

}

func TestFilter(t *testing.T) {

}

func TestFilterNot(t *testing.T) {

}

func TestFind(t *testing.T) {

}

func TestFirst(t *testing.T) {

}

func TestFlatMap(t *testing.T) {

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
