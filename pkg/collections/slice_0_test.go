package collections_test

import (
	"errors"
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

func sBar() c.SliceAny {
	var s SliceBar = []Bar{{1, []string{"w1"}}, {22, []string{"w22"}}, {333, []string{"w333"}}, {4444, []string{"w4444"}}, {22, []string{"w22"}}}
	return s.ToSliceAny()
}

func sInt() c.SliceAny {
	var s c.SliceAny = []c.Any{1, 22, 333, 4444, 22}
	return s
}

////
// Tests

func TestSliceCopy(t *testing.T) {
	cases := []struct {
		msg      string
		receiver c.SliceAny
	}{
		{"Copy: non-empty slice", sFoo()},
		{"Copy: empty slice", sEmpty()},
		{"Copy: nil slice", nil},
	}

	for _, cs := range cases {
		got := cs.receiver.Copy()
		assert.Equal(t, cs.receiver, got, cs.msg)
		assert.True(t, &cs.receiver != &got, cs.msg)
	}
}

func TestLengthSize(t *testing.T) {
	cases := []struct {
		msg      string
		receiver c.SliceAny
		want     int
	}{
		{"Length and Size: non-empty slice", sFoo(), 5},
		{"Length and Size: empty slice", sEmpty(), 0},
	}

	for _, cs := range cases {
		got1 := cs.receiver.Length()
		assert.Equal(t, cs.want, got1, cs.msg)
		got2 := cs.receiver.Size()
		assert.Equal(t, cs.want, got2, cs.msg)
	}
}

func TestContains(t *testing.T) {
	cases := []struct {
		msg      string
		receiver c.SliceAny
		arg      Bar
		want     bool
	}{
		{"Cotains: present", sBar(), Bar{22, []string{"w22"}}, true},
		{"Contains: absent", sBar(), Bar{22, []string{"xyz"}}, false},
		{"Contains: empty slice", sEmpty(), Bar{22, []string{"w22"}}, false},
	}

	for _, cs := range cases {
		got := cs.receiver.Contains(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestContainsAll(t *testing.T) {
	cases := []struct {
		msg      string
		receiver c.SliceAny
		arg      c.SliceAny
		want     bool
	}{
		{"ContainsSlice: subset", sBar(), append(sBar()[2:3], sBar()[1]), true},
		{"ContainsSlice: intersects", sBar(), append(sBar()[1:2], Bar{22, []string{"xyz"}}), false},
		{"ContainsSlice: disjoint", sBar(), append(sBar()[:0], Bar{22, []string{"xyz"}}, Bar{0, []string{"abc"}}), false},
		{"ContainsSlice: empty slice", sEmpty(), append(sBar()[2:3], sBar()[1]), false},
	}

	for _, cs := range cases {
		got := cs.receiver.ContainsSlice(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestGet(t *testing.T) {
	size := len(sFoo())
	cases := []struct {
		msg      string
		receiver c.SliceAny
		arg      int
		want     c.Any
		wok      bool
	}{
		{"Get: from middle", sFoo(), 2, sFoo()[2], true},
		{"Get: from beginning", sFoo(), 0, sFoo()[0], true},
		{"Get: from end", sFoo(), size - 1, sFoo()[size-1], true},
		{"Get: outside range", sFoo(), size, nil, false},
		{"Get: empty slice", sEmpty(), 0, nil, false},
	}

	for _, cs := range cases {
		got, ok := cs.receiver.Get(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
		assert.Equal(t, cs.wok, ok, cs.msg)
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
		{"SubSlice: nonempty - empty sub-slice", sFoo(), 2, 2, sEmpty()},
		{"SubSlice: empty - empty sub-slice", sEmpty(), 0, 0, sEmpty()},
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

func TestFirst(t *testing.T) {
	cases := []struct {
		msg      string
		receiver c.SliceAny
		want     c.Any
		werr     error
	}{
		{"First: non-empty", sFoo(), Foo{1, "w1"}, nil},
		{"First: empty", sEmpty(), nil, errors.New("empty slice")},
	}

	for _, cs := range cases {
		got, err := cs.receiver.First()
		assert.Equal(t, cs.want, got, cs.msg)
		assert.Equal(t, cs.werr, err, cs.msg)
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

func TestIndexOfFirst(t *testing.T) {
	pred1 := func(a c.Any) bool { return a.(Foo).v1 > 0 }
	pred2 := func(a c.Any) bool { return a.(Foo).v1%2 == 0 }
	pred3 := func(a c.Any) bool { return a.(Foo).v1 < 0 }

	cases := []struct {
		msg      string
		receiver c.SliceAny
		arg      func(c.Any) bool
		want     int
	}{
		{"IndexOfFirst: match all", sFoo(), pred1, 0},
		{"IndexOfFirst: match some", sFoo(), pred2, 1},
		{"IndexOfFirst: match none", sFoo(), pred3, -1},
		{"IndexOfFirst: empty", sEmpty(), pred1, -1},
	}

	for _, cs := range cases {
		got := cs.receiver.IndexOfFirst(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestIndexOfLast(t *testing.T) {
	pred1 := func(a c.Any) bool { return a.(Foo).v1 > 0 }
	pred2 := func(a c.Any) bool { return a.(Foo).v1%2 == 1 }
	pred3 := func(a c.Any) bool { return a.(Foo).v1 < 0 }

	cases := []struct {
		msg      string
		receiver c.SliceAny
		arg      func(c.Any) bool
		want     int
	}{
		{"IndexOfLast: match all", sFoo(), pred1, 4},
		{"IndexOfLast: match some", sFoo(), pred2, 2},
		{"IndexOfLast: match none", sFoo(), pred3, -1},
		{"IndexOfLast: empty", sEmpty(), pred1, -1},
	}

	for _, cs := range cases {
		got := cs.receiver.IndexOfLast(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestIsNotEmpty(t *testing.T) {
	cases := []struct {
		msg      string
		receiver c.SliceAny
		want     bool
	}{
		{"IsNotEmpty: non-empty", sFoo(), true},
		{"IsNotEmpty: empty", sEmpty(), false},
	}

	for _, cs := range cases {
		got := cs.receiver.IsNotEmpty()
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestLast(t *testing.T) {
	cases := []struct {
		msg      string
		receiver c.SliceAny
		want     c.Any
		werr     error
	}{
		{"Last: non-empty", sFoo(), Foo{22, "w22"}, nil},
		{"Last: empty", sEmpty(), nil, errors.New("empty slice")},
	}

	for _, cs := range cases {
		got, err := cs.receiver.Last()
		assert.Equal(t, cs.want, got, cs.msg)
		assert.Equal(t, cs.werr, err, cs.msg)
	}
}

func TestMaxWith(t *testing.T) {
	comp := func(a1 c.Any, a2 c.Any) int { return a1.(Foo).v1 - a2.(Foo).v1 }

	cases := []struct {
		msg      string
		receiver c.SliceAny
		arg      func(c.Any, c.Any) int
		want     c.Any
		werr     error
	}{
		{"MaxWith: non-empty receiver", sFoo(), comp, Foo{4444, "w4444"}, nil},
		{"MaxWith: empty receiver", sEmpty(), comp, nil, errors.New("empty slice")},
	}

	for _, cs := range cases {
		got, err := cs.receiver.MaxWith(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
		assert.Equal(t, cs.werr, err, cs.msg)
	}
}

// []Foo{{1, "w1"}, {22, "w22"}, {333, "w333"}, {4444, "w4444"}, {22, "w22"}}
func TestMinus(t *testing.T) {
	cases := []struct {
		msg      string
		receiver c.SliceAny
		arg      c.SliceAny
		want     c.SliceAny
	}{
		{"Minus: subset", sBar(), append(sBar()[3:4], sBar()[1]), append(sBar()[0:1], sBar()[2])},
		{"Minus: intersects", sBar(), append(sBar()[1:2], Bar{22, []string{"xyz"}}), append(sBar()[0:1], sBar()[2], sBar()[3])},
		{"Minus: disjoint", sBar(), append(sBar()[:0], Bar{22, []string{"xyz"}}, Bar{0, []string{"abc"}}), sBar()},
		{"Minus: empty slice", sEmpty(), append(sBar()[2:2], sBar()[1]), sEmpty()},
	}

	for _, cs := range cases {
		got := cs.receiver.Minus(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMinusElement(t *testing.T) {
	cases := []struct {
		msg      string
		receiver c.SliceAny
		arg      Bar
		want     c.SliceAny
	}{
		{"MinusElement: present", sBar(), Bar{22, []string{"w22"}}, append(sBar()[0:1], sBar()[2:]...)},
		{"MinusElement: absent", sBar(), Bar{22, []string{"xyz"}}, sBar()},
		{"MinusElement: empty slice", sEmpty(), Bar{22, []string{"xyz"}}, sEmpty()},
	}

	for _, cs := range cases {
		got := cs.receiver.MinusElement(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMinWith(t *testing.T) {
	comp := func(a1 c.Any, a2 c.Any) int { return -(a1.(Foo).v1 - a2.(Foo).v1) }

	cases := []struct {
		msg      string
		receiver c.SliceAny
		arg      func(c.Any, c.Any) int
		want     c.Any
		werr     error
	}{
		{"MinWith: non-empty receiver", sFoo(), comp, Foo{4444, "w4444"}, nil},
		{"MinWith: empty receiver", sEmpty(), comp, nil, errors.New("empty slice")},
	}

	for _, cs := range cases {
		got, err := cs.receiver.MinWith(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
		assert.Equal(t, cs.werr, err, cs.msg)
	}
}

func TestPartition(t *testing.T) {
	pred1 := func(a c.Any) bool { return a.(Foo).v1 > 0 }
	pred2 := func(a c.Any) bool { return a.(Foo).v1%2 == 0 }
	pred3 := func(a c.Any) bool { return a.(Foo).v1 < 0 }

	cases := []struct {
		msg      string
		receiver c.SliceAny
		arg      func(c.Any) bool
		want1    c.SliceAny
		want2    c.SliceAny
	}{
		{"Partition: match all", sFoo(), pred1, sFoo(), sEmpty()},
		{"Partition: match some", sFoo(), pred2, append(sFoo()[1:2], sFoo()[3], sFoo()[4]), append(sFoo()[0:1], sFoo()[2])},
		{"Partition: match none", sFoo(), pred3, sEmpty(), sFoo()},
		{"Partition: empty", sEmpty(), pred1, sEmpty(), sEmpty()},
	}

	for _, cs := range cases {
		got1, got2 := cs.receiver.Partition(cs.arg)
		assert.Equal(t, cs.want1, got1, cs.msg)
		assert.Equal(t, cs.want2, got2, cs.msg)
	}
}

func TestPlus(t *testing.T) {
	cases := []struct {
		msg      string
		receiver c.SliceAny
		arg      c.SliceAny
		want     c.SliceAny
	}{
		{"PlusMap: non-empty + non-empty", sFoo()[:3], sFoo()[3:], sFoo()},
		{"PlusMap: non-empty + empty", sFoo()[:3], sEmpty(), sFoo()[:3]},
		{"PlusMap: empty + non-empty", sEmpty(), sFoo()[3:], sFoo()[3:]},
		{"PlusMap: empty + empty", sEmpty(), sEmpty(), sEmpty()},
	}

	for _, cs := range cases {
		got := cs.receiver.PlusSlice(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestPlusElement(t *testing.T) {
	cases := []struct {
		msg      string
		receiver c.SliceAny
		arg      c.Any
		want     c.SliceAny
	}{
		{"PlusElement: non-empty", sFoo()[:4], sFoo()[4], sFoo()},
		{"PlusElement: empty", sEmpty(), sFoo()[4], sFoo()[4:5]},
	}

	for _, cs := range cases {
		got := cs.receiver.PlusElement(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestReduce(t *testing.T) {
	op := func(a1 c.Any, a2 c.Any) c.Any { return Foo{a1.(Foo).v1 + a2.(Foo).v1, a1.(Foo).v2 + a2.(Foo).v2} }

	cases := []struct {
		msg      string
		receiver c.SliceAny
		arg      func(c.Any, c.Any) c.Any
		want     c.Any
		werr     error
	}{
		{"Reduce: receiver length > 1", sFoo(), op,
			Foo{1 + 22 + 333 + 4444 + 22, "w1w22w333w4444w22"}, nil},
		{"Reduce: receiver length = 1", sFoo()[2:3], op, sFoo()[2], nil},
		{"Reduce: empty receiver", sEmpty(), op, nil, errors.New("empty slice")},
	}

	for _, cs := range cases {
		got, err := cs.receiver.Reduce(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
		assert.Equal(t, cs.werr, err, cs.msg)
	}
}

func TestReversed(t *testing.T) {
	rev := []Foo{{22, "w22"}, {4444, "w4444"}, {333, "w333"}, {22, "w22"}, {1, "w1"}}

	cases := []struct {
		msg      string
		receiver c.SliceAny
		want     c.SliceAny
	}{
		{"Reversed: non-empty slice", sFoo(), SliceFoo(rev).ToSliceAny()},
		{"Reversed: empty slice", sEmpty(), sEmpty()},
	}

	for _, cs := range cases {
		got := cs.receiver.Reversed()
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSortedWith(t *testing.T) {
	comp := func(a1 c.Any, a2 c.Any) int { return -(a1.(Foo).v1 - a2.(Foo).v1) }

	var sorted SliceFoo = []Foo{{4444, "w4444"}, {333, "w333"}, {22, "w22"}, {22, "w22"}, {1, "w1"}}

	cases := []struct {
		msg      string
		receiver c.SliceAny
		arg      func(c.Any, c.Any) int
		want     c.SliceAny
	}{
		{"SortedWith: non-empty receiver", sFoo(), comp, sorted.ToSliceAny()},
		{"SortedWith: empty receiver", sEmpty(), comp, sEmpty()},
	}

	for _, cs := range cases {
		got := cs.receiver.SortedWith(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestTake(t *testing.T) {
	size := len(sFoo())
	cases := []struct {
		msg      string
		receiver c.SliceAny
		arg      int
		want     c.SliceAny
	}{
		{"Take: some", sFoo(), 2, sFoo()[:2]},
		{"Take: all", sFoo(), size, sFoo()},
		{"Take: none", sFoo(), 0, sEmpty()},
		{"Take: more than length", sFoo(), size + 5, sFoo()},
		{"Take: empty receiver", sEmpty(), 1, sEmpty()},
	}

	for _, cs := range cases {
		got := cs.receiver.Take(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestTakeLast(t *testing.T) {
	size := len(sFoo())
	cases := []struct {
		msg      string
		receiver c.SliceAny
		arg      int
		want     c.SliceAny
	}{
		{"TakeLast: some", sFoo(), 2, sFoo()[size-2:]},
		{"TakeLast: all", sFoo(), size, sFoo()},
		{"TakeLast: none", sFoo(), 0, sEmpty()},
		{"TakeLast: more than length", sFoo(), size + 5, sFoo()},
		{"TakeLast: empty receiver", sEmpty(), 1, sEmpty()},
	}

	for _, cs := range cases {
		got := cs.receiver.TakeLast(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestTakeLastWhile(t *testing.T) {
	pred1 := func(a c.Any) bool { return a.(Foo).v1 > 0 }
	pred2 := func(a c.Any) bool { return a.(Foo).v1%2 == 0 }
	pred3 := func(a c.Any) bool { return a.(Foo).v1 < 0 }

	cases := []struct {
		msg      string
		receiver c.SliceAny
		arg      func(c.Any) bool
		want     c.SliceAny
	}{
		{"TakeLastWhile: pred matches all", sFoo(), pred1, sFoo()},
		{"TakeLastWhile: pred matches some", sFoo(), pred2, SliceFoo([]Foo{{4444, "w4444"}, {22, "w22"}}).ToSliceAny()},
		{"TakeLastWhile: pred matches none", sFoo(), pred3, sEmpty()},
		{"TakeLastWhile: empty receiver", sEmpty(), pred2, sEmpty()},
	}

	for _, cs := range cases {
		got := cs.receiver.TakeLastWhile(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestTakeWhile(t *testing.T) {
	pred1 := func(a c.Any) bool { return a.(Foo).v1 > 0 }
	pred2 := func(a c.Any) bool { return a.(Foo).v1%2 == 1 }
	pred3 := func(a c.Any) bool { return a.(Foo).v1 < 0 }

	cases := []struct {
		msg      string
		receiver c.SliceAny
		arg      func(c.Any) bool
		want     c.SliceAny
	}{
		{"TakeWhile: pred matches all", sFoo(), pred1, sFoo()},
		{"TakeWhile: pred matches some", sFoo(), pred2, SliceFoo([]Foo{{1, "w1"}}).ToSliceAny()},
		{"TakeWhile: pred matches none", sFoo(), pred3, sEmpty()},
		{"TakeWhile: empty receiver", sEmpty(), pred2, sEmpty()},
	}

	for _, cs := range cases {
		got := cs.receiver.TakeWhile(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestToSlice(t *testing.T) {
	cases := []struct {
		msg      string
		receiver c.SliceAny
		want     c.SliceAny
	}{
		{"ToSlice: non-empty", sFoo(), sFoo()},
		{"ToSlice: empty", sEmpty(), []c.Any{}},
	}

	for _, cs := range cases {
		got := cs.receiver.ToSlice()
		assert.Equal(t, cs.want, got, cs.msg)
	}
}
