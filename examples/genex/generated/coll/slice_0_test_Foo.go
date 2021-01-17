// Code generated -- DO NOT EDIT.

package coll

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

////
// Tests

func TestSliceCopy(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SliceFoo
	}{
		{"Copy: non-empty slice", sFoo()},
		{"Copy: empty slice", SliceFoo{}},
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
		receiver SliceFoo
		want     int
	}{
		{"Length and Size: non-empty slice", sFoo(), 5},
		{"Length and Size: empty slice", SliceFoo{}, 0},
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
		receiver SliceFoo
		arg      Foo
		want     bool
	}{
		{"Cotains: present", sFoo(), Foo{22, "w22"}, true},
		{"Contains: absent", sFoo(), Foo{22, "xyz"}, false},
		{"Contains: empty slice", SliceFoo{}, Foo{22, "w22"}, false},
	}

	for _, cs := range cases {
		got := cs.receiver.Contains(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestContainsAll(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SliceFoo
		arg      SliceFoo
		want     bool
	}{
		{"ContainsSlice: subset", sFoo(), append(sFoo()[2:3], sFoo()[1]), true},
		{"ContainsSlice: intersects", sFoo(), append(sFoo()[1:2], Foo{22, "xyz"}), false},
		{"ContainsSlice: disjoint", sFoo(), append(sFoo()[:0], Foo{22, "xyz"}, Foo{0, "abc"}),
			false},
		{"ContainsSlice: empty slice", SliceFoo{}, append(sFoo()[2:3], sFoo()[1]), false},
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
		receiver SliceFoo
		arg      int
		want     Foo
		wok      bool
	}{
		{"Get: from middle", sFoo(), 2, sFoo()[2], true},
		{"Get: from beginning", sFoo(), 0, sFoo()[0], true},
		{"Get: from end", sFoo(), size - 1, sFoo()[size-1], true},
		{"Get: outside range", sFoo(), size, Foo{}, false},
		{"Get: empty slice", SliceFoo{}, 0, Foo{}, false},
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
		receiver SliceFoo
		arg      Foo
		want     int
	}{
		{"IndexOf: non-empty, present", sFoo(), Foo{22, "w22"}, 1},
		{"IndexOf: non-empty, absent", sFoo(), Foo{0, "xyz"}, -1},
		{"IndexOf: empty", SliceFoo{}, Foo{0, "xyz"}, -1},
	}

	for _, cs := range cases {
		got := cs.receiver.IndexOf(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestIsEmpty(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SliceFoo
		want     bool
	}{
		{"IsEmpty: non-empty", sFoo(), false},
		{"IsEmpty: empty", SliceFoo{}, true},
	}

	for _, cs := range cases {
		got := cs.receiver.IsEmpty()
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestLastIndexOf(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SliceFoo
		arg      Foo
		want     int
	}{
		{"LastIndexOf: non-empty, present", sFoo(), Foo{22, "w22"}, 4},
		{"LastIndexOf: non-empty, absent", sFoo(), Foo{0, "xyz"}, -1},
		{"LastIndexOf: empty", SliceFoo{}, Foo{0, "xyz"}, -1},
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
		receiver SliceFoo
		arg1     int
		arg2     int
		want     SliceFoo
	}{
		{"SubSlice: nonempty - from beginning", sFoo(), 0, 2, sFoo()[:2]},
		{"SubSlice: nonempty - from middle", sFoo(), 1, 3, sFoo()[1:3]},
		{"SubSlice: nonempty - from end", sFoo(), size - 3, size, sFoo()[size-3:]},
		{"SubSlice: nonempty - empty sub-slice", sFoo(), 2, 2, SliceFoo{}},
		{"SubSlice: empty - empty sub-slice", SliceFoo{}, 0, 0, SliceFoo{}},
	}

	for _, cs := range cases {
		got := cs.receiver.SubSlice(cs.arg1, cs.arg2)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestAll(t *testing.T) {
	pred1 := func(a Foo) bool { return Any(a).(Foo).V1 > 0 }
	pred2 := func(a Foo) bool { return Any(a).(Foo).V1%2 == 0 }
	pred3 := func(a Foo) bool { return Any(a).(Foo).V1 < 0 }

	cases := []struct {
		msg      string
		receiver SliceFoo
		arg      func(Foo) bool
		want     bool
	}{
		{"All: pred matches all", sFoo(), pred1, true},
		{"All: pred matches some", sFoo(), pred2, false},
		{"All: pred matches none", sFoo(), pred3, false},
		{"All: empty receiver", SliceFoo{}, pred2, true},
	}

	for _, cs := range cases {
		got := cs.receiver.All(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestAny(t *testing.T) {
	pred1 := func(a Foo) bool { return Any(a).(Foo).V1 > 0 }
	pred2 := func(a Foo) bool { return Any(a).(Foo).V1%2 == 0 }
	pred3 := func(a Foo) bool { return Any(a).(Foo).V1 < 0 }

	cases := []struct {
		msg      string
		receiver SliceFoo
		arg      func(Foo) bool
		want     bool
	}{
		{"Any: pred matches all", sFoo(), pred1, true},
		{"Any: pred matches some", sFoo(), pred2, true},
		{"Any: pred matches none", sFoo(), pred3, false},
		{"Any: empty receiver", SliceFoo{}, pred2, false},
	}

	for _, cs := range cases {
		got := cs.receiver.Any(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestCount(t *testing.T) {
	pred1 := func(a Foo) bool { return Any(a).(Foo).V1 > 0 }
	pred2 := func(a Foo) bool { return Any(a).(Foo).V1%2 == 0 }
	pred3 := func(a Foo) bool { return Any(a).(Foo).V1 < 0 }

	cases := []struct {
		msg      string
		receiver SliceFoo
		arg      func(Foo) bool
		want     int
	}{
		{"Count: pred matches all", sFoo(), pred1, len(sFoo())},
		{"Count: pred matches some", sFoo(), pred2, 3},
		{"Count: pred matches none", sFoo(), pred3, 0},
		{"Count: empty receiver", SliceFoo{}, pred2, 0},
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
		receiver SliceFoo
		arg      int
		want     SliceFoo
	}{
		{"Drop: some", sFoo(), 2, sFoo()[2:]},
		{"Drop: all", sFoo(), size, SliceFoo{}},
		{"Drop: none", sFoo(), 0, sFoo()},
		{"Drop: more than length", sFoo(), size + 5, SliceFoo{}},
		{"Drop: empty receiver", SliceFoo{}, 1, SliceFoo{}},
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
		receiver SliceFoo
		arg      int
		want     SliceFoo
	}{
		{"DropLast: some", sFoo(), 2, sFoo()[:size-2]},
		{"DropLast: all", sFoo(), size, SliceFoo{}},
		{"DropLast: none", sFoo(), 0, sFoo()},
		{"DropLast: more than length", sFoo(), size + 5, SliceFoo{}},
		{"DropLast: empty receiver", SliceFoo{}, 1, SliceFoo{}},
	}

	for _, cs := range cases {
		got := cs.receiver.DropLast(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestDropLastWhile(t *testing.T) {
	pred1 := func(a Foo) bool { return Any(a).(Foo).V1 > 0 }
	pred2 := func(a Foo) bool { return Any(a).(Foo).V1%2 == 0 }
	pred3 := func(a Foo) bool { return Any(a).(Foo).V1 < 0 }

	cases := []struct {
		msg      string
		receiver SliceFoo
		arg      func(Foo) bool
		want     SliceFoo
	}{
		{"DropLastWhile: pred matches all", sFoo(), pred1, SliceFoo{}},
		{"DropLastWhile: pred matches some", sFoo(), pred2,
			SliceFoo{Foo{1, "w1"}, Foo{22, "w22"}, Foo{333, "w333"}}},
		{"DropLastWhile: pred matches none", sFoo(), pred3, sFoo()},
		{"DropLastWhile: empty receiver", SliceFoo{}, pred2, SliceFoo{}},
	}

	for _, cs := range cases {
		got := cs.receiver.DropLastWhile(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestDropWhile(t *testing.T) {
	pred1 := func(a Foo) bool { return Any(a).(Foo).V1 > 0 }
	pred2 := func(a Foo) bool { return Any(a).(Foo).V1%2 == 1 }
	pred3 := func(a Foo) bool { return Any(a).(Foo).V1 < 0 }

	cases := []struct {
		msg      string
		receiver SliceFoo
		arg      func(Foo) bool
		want     SliceFoo
	}{
		{"DropWhile: pred matches all", sFoo(), pred1, SliceFoo{}},
		{"DropWhile: pred matches some", sFoo(), pred2,
			SliceFoo{Foo{22, "w22"}, Foo{333, "w333"}, Foo{4444, "w4444"}, Foo{22, "w22"}}},
		{"DropWhile: pred matches none", sFoo(), pred3, sFoo()},
		{"DropWhile: empty receiver", SliceFoo{}, pred2, SliceFoo{}},
	}

	for _, cs := range cases {
		got := cs.receiver.DropWhile(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestFilter(t *testing.T) {
	pred1 := func(a Foo) bool { return Any(a).(Foo).V1 > 0 }
	pred2 := func(a Foo) bool { return Any(a).(Foo).V1%2 == 0 }
	pred3 := func(a Foo) bool { return Any(a).(Foo).V1 < 0 }

	cases := []struct {
		msg      string
		receiver SliceFoo
		arg      func(Foo) bool
		want     SliceFoo
	}{
		{"Filter: pred matches all", sFoo(), pred1, sFoo()},
		{"Filter: pred matches some", sFoo(), pred2,
			SliceFoo{Foo{22, "w22"}, Foo{4444, "w4444"}, Foo{22, "w22"}}},
		{"Filter: pred matches none", sFoo(), pred3, SliceFoo{}},
		{"Filter: empty receiver", SliceFoo{}, pred2, SliceFoo{}},
	}

	for _, cs := range cases {
		got := cs.receiver.Filter(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestFilterNot(t *testing.T) {
	pred1 := func(a Foo) bool { return Any(a).(Foo).V1 > 0 }
	pred2 := func(a Foo) bool { return Any(a).(Foo).V1%2 == 1 }
	pred3 := func(a Foo) bool { return Any(a).(Foo).V1 < 0 }

	cases := []struct {
		msg      string
		receiver SliceFoo
		arg      func(Foo) bool
		want     SliceFoo
	}{
		{"FilterNot: pred matches all", sFoo(), pred1, SliceFoo{}},
		{"FilterNot: pred matches some", sFoo(), pred2,
			SliceFoo{Foo{22, "w22"}, Foo{4444, "w4444"}, Foo{22, "w22"}}},
		{"FilterNot: pred matches none", sFoo(), pred3, sFoo()},
		{"FilterNot: empty receiver", SliceFoo{}, pred2, SliceFoo{}},
	}

	for _, cs := range cases {
		got := cs.receiver.FilterNot(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestFirst(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SliceFoo
		want     Foo
		werr     error
	}{
		{"First: non-empty", sFoo(), Foo{1, "w1"}, nil},
		{"First: empty", SliceFoo{}, nil, errors.New("empty slice")},
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
		receiver SliceFoo
		want     []Foo
	}{
		{"ForEach: non-empty receiver", sFoo(), []Foo{1, 22, 333, 4444, 22}},
		{"ForEach: empty receiver", SliceFoo{}, []Foo{}},
	}

	for _, cs := range cases {
		got := []Foo{}
		f := func(a Foo) {
			got = append(got, Any(a).(Foo).V1)
		}

		cs.receiver.ForEach(f)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestIndexOfFirst(t *testing.T) {
	pred1 := func(a Foo) bool { return Any(a).(Foo).V1 > 0 }
	pred2 := func(a Foo) bool { return Any(a).(Foo).V1%2 == 0 }
	pred3 := func(a Foo) bool { return Any(a).(Foo).V1 < 0 }

	cases := []struct {
		msg      string
		receiver SliceFoo
		arg      func(Foo) bool
		want     int
	}{
		{"IndexOfFirst: match all", sFoo(), pred1, 0},
		{"IndexOfFirst: match some", sFoo(), pred2, 1},
		{"IndexOfFirst: match none", sFoo(), pred3, -1},
		{"IndexOfFirst: empty", SliceFoo{}, pred1, -1},
	}

	for _, cs := range cases {
		got := cs.receiver.IndexOfFirst(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestIndexOfLast(t *testing.T) {
	pred1 := func(a Foo) bool { return Any(a).(Foo).V1 > 0 }
	pred2 := func(a Foo) bool { return Any(a).(Foo).V1%2 == 1 }
	pred3 := func(a Foo) bool { return Any(a).(Foo).V1 < 0 }

	cases := []struct {
		msg      string
		receiver SliceFoo
		arg      func(Foo) bool
		want     int
	}{
		{"IndexOfLast: match all", sFoo(), pred1, 4},
		{"IndexOfLast: match some", sFoo(), pred2, 2},
		{"IndexOfLast: match none", sFoo(), pred3, -1},
		{"IndexOfLast: empty", SliceFoo{}, pred1, -1},
	}

	for _, cs := range cases {
		got := cs.receiver.IndexOfLast(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestIsNotEmpty(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SliceFoo
		want     bool
	}{
		{"IsNotEmpty: non-empty", sFoo(), true},
		{"IsNotEmpty: empty", SliceFoo{}, false},
	}

	for _, cs := range cases {
		got := cs.receiver.IsNotEmpty()
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestLast(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SliceFoo
		want     Foo
		werr     error
	}{
		{"Last: non-empty", sFoo(), Foo{22, "w22"}, nil},
		{"Last: empty", SliceFoo{}, nil, errors.New("empty slice")},
	}

	for _, cs := range cases {
		got, err := cs.receiver.Last()
		assert.Equal(t, cs.want, got, cs.msg)
		assert.Equal(t, cs.werr, err, cs.msg)
	}
}

func TestMaxWith(t *testing.T) {
	comp := func(a1 Foo, a2 Foo) int { return a1.(Foo).V1 - a2.(Foo).V1 }

	cases := []struct {
		msg      string
		receiver SliceFoo
		arg      func(Foo, Foo) int
		want     Foo
		werr     error
	}{
		{"MaxWith: non-empty receiver", sFoo(), comp, Foo{4444, "w4444"}, nil},
		{"MaxWith: empty receiver", SliceFoo{}, comp, nil, errors.New("empty slice")},
	}

	for _, cs := range cases {
		got, err := cs.receiver.MaxWith(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
		assert.Equal(t, cs.werr, err, cs.msg)
	}
}

func TestMinus(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SliceFoo
		arg      SliceFoo
		want     SliceFoo
	}{
		{"MinusSlice: subset", sFoo(), append(sFoo()[3:4], sFoo()[1]), append(sFoo()[0:1], sFoo()[2])},
		{"MinusSlice: intersects", sFoo(), append(sFoo()[1:2], Foo{22, "xyz"}), append(sFoo()[0:1], sFoo()[2], sFoo()[3])},
		{"MinusSlice: disjoint", sFoo(), append(sFoo()[:0], Foo{22, "xyz"}, Foo{0, "abc"}), sFoo()},
		{"MinusSlice: empty slice", SliceFoo{}, append(sFoo()[2:2], sFoo()[1]), SliceFoo{}},
	}

	for _, cs := range cases {
		got := cs.receiver.MinusSlice(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMinusElement(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SliceFoo
		arg      Foo
		want     SliceFoo
	}{
		{"MinusElement: present", sFoo(), Foo{22, "w22"}, append(sFoo()[0:1], sFoo()[2:]...)},
		{"MinusElement: absent", sFoo(), Foo{22, "xyz"}, sFoo()},
		{"MinusElement: empty slice", SliceFoo{}, Foo{22, "xyz"}, SliceFoo{}},
	}

	for _, cs := range cases {
		got := cs.receiver.MinusElement(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMinWith(t *testing.T) {
	comp := func(a1 Foo, a2 Foo) int { return -(a1.(Foo).V1 - a2.(Foo).V1) }

	cases := []struct {
		msg      string
		receiver SliceFoo
		arg      func(Foo, Foo) int
		want     Foo
		werr     error
	}{
		{"MinWith: non-empty receiver", sFoo(), comp, Foo{4444, "w4444"}, nil},
		{"MinWith: empty receiver", SliceFoo{}, comp, nil, errors.New("empty slice")},
	}

	for _, cs := range cases {
		got, err := cs.receiver.MinWith(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
		assert.Equal(t, cs.werr, err, cs.msg)
	}
}

func TestPartition(t *testing.T) {
	pred1 := func(a Foo) bool { return Any(a).(Foo).V1 > 0 }
	pred2 := func(a Foo) bool { return Any(a).(Foo).V1%2 == 0 }
	pred3 := func(a Foo) bool { return Any(a).(Foo).V1 < 0 }

	cases := []struct {
		msg      string
		receiver SliceFoo
		arg      func(Foo) bool
		want1    SliceFoo
		want2    SliceFoo
	}{
		{"Partition: match all", sFoo(), pred1, sFoo(), SliceFoo{}},
		{"Partition: match some", sFoo(), pred2, append(sFoo()[1:2], sFoo()[3], sFoo()[4]), append(sFoo()[0:1], sFoo()[2])},
		{"Partition: match none", sFoo(), pred3, SliceFoo{}, sFoo()},
		{"Partition: empty", SliceFoo{}, pred1, SliceFoo{}, SliceFoo{}},
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
		receiver SliceFoo
		arg      SliceFoo
		want     SliceFoo
	}{
		{"PlusMap: non-empty + non-empty", sFoo()[:3], sFoo()[3:], sFoo()},
		{"PlusMap: non-empty + empty", sFoo()[:3], SliceFoo{}, sFoo()[:3]},
		{"PlusMap: empty + non-empty", SliceFoo{}, sFoo()[3:], sFoo()[3:]},
		{"PlusMap: empty + empty", SliceFoo{}, SliceFoo{}, SliceFoo{}},
	}

	for _, cs := range cases {
		got := cs.receiver.PlusSlice(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestPlusElement(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SliceFoo
		arg      Foo
		want     SliceFoo
	}{
		{"PlusElement: non-empty", sFoo()[:4], sFoo()[4], sFoo()},
		{"PlusElement: empty", SliceFoo{}, sFoo()[4], sFoo()[4:5]},
	}

	for _, cs := range cases {
		got := cs.receiver.PlusElement(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestReduce(t *testing.T) {
	op := func(a1 Foo, a2 Foo) Foo { return Foo{a1.(Foo).V1 + a2.(Foo).V1, a1.(Foo).V2 + a2.(Foo).V2} }

	cases := []struct {
		msg      string
		receiver SliceFoo
		arg      func(Foo, Foo) Foo
		want     Foo
		werr     error
	}{
		{"Reduce: receiver length > 1", sFoo(), op,
			Foo{1 + 22 + 333 + 4444 + 22, "w1w22w333w4444w22"}, nil},
		{"Reduce: receiver length = 1", sFoo()[2:3], op, sFoo()[2], nil},
		{"Reduce: empty receiver", SliceFoo{}, op, nil, errors.New("empty slice")},
	}

	for _, cs := range cases {
		got, err := cs.receiver.Reduce(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
		assert.Equal(t, cs.werr, err, cs.msg)
	}
}

func TestReversed(t *testing.T) {
	rev := SliceFoo{Foo{22, "w22"}, Foo{4444, "w4444"}, Foo{333, "w333"}, Foo{22, "w22"},
		Foo{1, "w1"}}

	cases := []struct {
		msg      string
		receiver SliceFoo
		want     SliceFoo
	}{
		{"Reversed: non-empty slice", sFoo(), rev},
		{"Reversed: empty slice", SliceFoo{}, SliceFoo{}},
	}

	for _, cs := range cases {
		got := cs.receiver.Reversed()
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSortedWith(t *testing.T) {
	comp := func(a1 Foo, a2 Foo) int { return -(a1.(Foo).V1 - a2.(Foo).V1) }

	sorted := SliceFoo{Foo{4444, "w4444"}, Foo{333, "w333"}, Foo{22, "w22"},
		Foo{22, "w22"}, Foo{1, "w1"}}

	cases := []struct {
		msg      string
		receiver SliceFoo
		arg      func(Foo, Foo) int
		want     SliceFoo
	}{
		{"SortedWith: non-empty receiver", sFoo(), comp, sorted},
		{"SortedWith: empty receiver", SliceFoo{}, comp, SliceFoo{}},
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
		receiver SliceFoo
		arg      int
		want     SliceFoo
	}{
		{"Take: some", sFoo(), 2, sFoo()[:2]},
		{"Take: all", sFoo(), size, sFoo()},
		{"Take: none", sFoo(), 0, SliceFoo{}},
		{"Take: more than length", sFoo(), size + 5, sFoo()},
		{"Take: empty receiver", SliceFoo{}, 1, SliceFoo{}},
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
		receiver SliceFoo
		arg      int
		want     SliceFoo
	}{
		{"TakeLast: some", sFoo(), 2, sFoo()[size-2:]},
		{"TakeLast: all", sFoo(), size, sFoo()},
		{"TakeLast: none", sFoo(), 0, SliceFoo{}},
		{"TakeLast: more than length", sFoo(), size + 5, sFoo()},
		{"TakeLast: empty receiver", SliceFoo{}, 1, SliceFoo{}},
	}

	for _, cs := range cases {
		got := cs.receiver.TakeLast(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestTakeLastWhile(t *testing.T) {
	pred1 := func(a Foo) bool { return Any(a).(Foo).V1 > 0 }
	pred2 := func(a Foo) bool { return Any(a).(Foo).V1%2 == 0 }
	pred3 := func(a Foo) bool { return Any(a).(Foo).V1 < 0 }

	cases := []struct {
		msg      string
		receiver SliceFoo
		arg      func(Foo) bool
		want     SliceFoo
	}{
		{"TakeLastWhile: pred matches all", sFoo(), pred1, sFoo()},
		{"TakeLastWhile: pred matches some", sFoo(), pred2,
			SliceFoo{Foo{4444, "w4444"}, Foo{22, "w22"}}},
		{"TakeLastWhile: pred matches none", sFoo(), pred3, SliceFoo{}},
		{"TakeLastWhile: empty receiver", SliceFoo{}, pred2, SliceFoo{}},
	}

	for _, cs := range cases {
		got := cs.receiver.TakeLastWhile(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestTakeWhile(t *testing.T) {
	pred1 := func(a Foo) bool { return Any(a).(Foo).V1 > 0 }
	pred2 := func(a Foo) bool { return Any(a).(Foo).V1%2 == 1 }
	pred3 := func(a Foo) bool { return Any(a).(Foo).V1 < 0 }

	cases := []struct {
		msg      string
		receiver SliceFoo
		arg      func(Foo) bool
		want     SliceFoo
	}{
		{"TakeWhile: pred matches all", sFoo(), pred1, sFoo()},
		{"TakeWhile: pred matches some", sFoo(), pred2, SliceFoo{Foo{1, "w1"}}},
		{"TakeWhile: pred matches none", sFoo(), pred3, SliceFoo{}},
		{"TakeWhile: empty receiver", SliceFoo{}, pred2, SliceFoo{}},
	}

	for _, cs := range cases {
		got := cs.receiver.TakeWhile(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}
