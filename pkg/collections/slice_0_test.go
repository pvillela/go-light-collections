package collections

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
		receiver SliceT0
	}{
		{"Copy: non-empty slice", sFoo()},
		{"Copy: empty slice", SliceT0{}},
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
		receiver SliceT0
		want     int
	}{
		{"Length and Size: non-empty slice", sFoo(), 5},
		{"Length and Size: empty slice", SliceT0{}, 0},
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
		receiver SliceT0
		arg      Foo
		want     bool
	}{
		{"Cotains: present", sFoo(), Foo{22, "w22"}, true},
		{"Contains: absent", sFoo(), Foo{22, "xyz"}, false},
		{"Contains: empty slice", SliceT0{}, Foo{22, "w22"}, false},
	}

	for _, cs := range cases {
		got := cs.receiver.Contains(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestContainsAll(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SliceT0
		arg      SliceT0
		want     bool
	}{
		{"ContainsSlice: subset", sFoo(), append(sFoo()[2:3], sFoo()[1]), true},
		{"ContainsSlice: intersects", sFoo(), append(sFoo()[1:2], Foo{22, "xyz"}), false},
		{"ContainsSlice: disjoint", sFoo(), append(sFoo()[:0], Foo{22, "xyz"}, Foo{0, "abc"}),
			false},
		{"ContainsSlice: empty slice", SliceT0{}, append(sFoo()[2:3], sFoo()[1]), false},
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
		receiver SliceT0
		arg      int
		want     T0
		wok      bool
	}{
		{"Get: from middle", sFoo(), 2, sFoo()[2], true},
		{"Get: from beginning", sFoo(), 0, sFoo()[0], true},
		{"Get: from end", sFoo(), size - 1, sFoo()[size-1], true},
		{"Get: outside range", sFoo(), size, Foo{}, false},
		{"Get: empty slice", SliceT0{}, 0, Foo{}, false},
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
		receiver SliceT0
		arg      Foo
		want     int
	}{
		{"IndexOf: non-empty, present", sFoo(), Foo{22, "w22"}, 1},
		{"IndexOf: non-empty, absent", sFoo(), Foo{0, "xyz"}, -1},
		{"IndexOf: empty", SliceT0{}, Foo{0, "xyz"}, -1},
	}

	for _, cs := range cases {
		got := cs.receiver.IndexOf(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestIsEmpty(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SliceT0
		want     bool
	}{
		{"IsEmpty: non-empty", sFoo(), false},
		{"IsEmpty: empty", SliceT0{}, true},
	}

	for _, cs := range cases {
		got := cs.receiver.IsEmpty()
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestLastIndexOf(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SliceT0
		arg      Foo
		want     int
	}{
		{"LastIndexOf: non-empty, present", sFoo(), Foo{22, "w22"}, 4},
		{"LastIndexOf: non-empty, absent", sFoo(), Foo{0, "xyz"}, -1},
		{"LastIndexOf: empty", SliceT0{}, Foo{0, "xyz"}, -1},
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
		receiver SliceT0
		arg1     int
		arg2     int
		want     SliceT0
	}{
		{"SubSlice: nonempty - from beginning", sFoo(), 0, 2, sFoo()[:2]},
		{"SubSlice: nonempty - from middle", sFoo(), 1, 3, sFoo()[1:3]},
		{"SubSlice: nonempty - from end", sFoo(), size - 3, size, sFoo()[size-3:]},
		{"SubSlice: nonempty - empty sub-slice", sFoo(), 2, 2, SliceT0{}},
		{"SubSlice: empty - empty sub-slice", SliceT0{}, 0, 0, SliceT0{}},
	}

	for _, cs := range cases {
		got := cs.receiver.SubSlice(cs.arg1, cs.arg2)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestAll(t *testing.T) {
	pred1 := func(a T0) bool { return Any(a).(Foo).V1 > 0 }
	pred2 := func(a T0) bool { return Any(a).(Foo).V1%2 == 0 }
	pred3 := func(a T0) bool { return Any(a).(Foo).V1 < 0 }

	cases := []struct {
		msg      string
		receiver SliceT0
		arg      func(T0) bool
		want     bool
	}{
		{"All: pred matches all", sFoo(), pred1, true},
		{"All: pred matches some", sFoo(), pred2, false},
		{"All: pred matches none", sFoo(), pred3, false},
		{"All: empty receiver", SliceT0{}, pred2, true},
	}

	for _, cs := range cases {
		got := cs.receiver.All(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestAny(t *testing.T) {
	pred1 := func(a T0) bool { return Any(a).(Foo).V1 > 0 }
	pred2 := func(a T0) bool { return Any(a).(Foo).V1%2 == 0 }
	pred3 := func(a T0) bool { return Any(a).(Foo).V1 < 0 }

	cases := []struct {
		msg      string
		receiver SliceT0
		arg      func(T0) bool
		want     bool
	}{
		{"Any: pred matches all", sFoo(), pred1, true},
		{"Any: pred matches some", sFoo(), pred2, true},
		{"Any: pred matches none", sFoo(), pred3, false},
		{"Any: empty receiver", SliceT0{}, pred2, false},
	}

	for _, cs := range cases {
		got := cs.receiver.Any(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestCount(t *testing.T) {
	pred1 := func(a T0) bool { return Any(a).(Foo).V1 > 0 }
	pred2 := func(a T0) bool { return Any(a).(Foo).V1%2 == 0 }
	pred3 := func(a T0) bool { return Any(a).(Foo).V1 < 0 }

	cases := []struct {
		msg      string
		receiver SliceT0
		arg      func(T0) bool
		want     int
	}{
		{"Count: pred matches all", sFoo(), pred1, len(sFoo())},
		{"Count: pred matches some", sFoo(), pred2, 3},
		{"Count: pred matches none", sFoo(), pred3, 0},
		{"Count: empty receiver", SliceT0{}, pred2, 0},
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
		receiver SliceT0
		arg      int
		want     SliceT0
	}{
		{"Drop: some", sFoo(), 2, sFoo()[2:]},
		{"Drop: all", sFoo(), size, SliceT0{}},
		{"Drop: none", sFoo(), 0, sFoo()},
		{"Drop: more than length", sFoo(), size + 5, SliceT0{}},
		{"Drop: empty receiver", SliceT0{}, 1, SliceT0{}},
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
		receiver SliceT0
		arg      int
		want     SliceT0
	}{
		{"DropLast: some", sFoo(), 2, sFoo()[:size-2]},
		{"DropLast: all", sFoo(), size, SliceT0{}},
		{"DropLast: none", sFoo(), 0, sFoo()},
		{"DropLast: more than length", sFoo(), size + 5, SliceT0{}},
		{"DropLast: empty receiver", SliceT0{}, 1, SliceT0{}},
	}

	for _, cs := range cases {
		got := cs.receiver.DropLast(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestDropLastWhile(t *testing.T) {
	pred1 := func(a T0) bool { return Any(a).(Foo).V1 > 0 }
	pred2 := func(a T0) bool { return Any(a).(Foo).V1%2 == 0 }
	pred3 := func(a T0) bool { return Any(a).(Foo).V1 < 0 }

	cases := []struct {
		msg      string
		receiver SliceT0
		arg      func(T0) bool
		want     SliceT0
	}{
		{"DropLastWhile: pred matches all", sFoo(), pred1, SliceT0{}},
		{"DropLastWhile: pred matches some", sFoo(), pred2,
			SliceT0{Foo{1, "w1"}, Foo{22, "w22"}, Foo{333, "w333"}}},
		{"DropLastWhile: pred matches none", sFoo(), pred3, sFoo()},
		{"DropLastWhile: empty receiver", SliceT0{}, pred2, SliceT0{}},
	}

	for _, cs := range cases {
		got := cs.receiver.DropLastWhile(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestDropWhile(t *testing.T) {
	pred1 := func(a T0) bool { return Any(a).(Foo).V1 > 0 }
	pred2 := func(a T0) bool { return Any(a).(Foo).V1%2 == 1 }
	pred3 := func(a T0) bool { return Any(a).(Foo).V1 < 0 }

	cases := []struct {
		msg      string
		receiver SliceT0
		arg      func(T0) bool
		want     SliceT0
	}{
		{"DropWhile: pred matches all", sFoo(), pred1, SliceT0{}},
		{"DropWhile: pred matches some", sFoo(), pred2,
			SliceT0{Foo{22, "w22"}, Foo{333, "w333"}, Foo{4444, "w4444"}, Foo{22, "w22"}}},
		{"DropWhile: pred matches none", sFoo(), pred3, sFoo()},
		{"DropWhile: empty receiver", SliceT0{}, pred2, SliceT0{}},
	}

	for _, cs := range cases {
		got := cs.receiver.DropWhile(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestFilter(t *testing.T) {
	pred1 := func(a T0) bool { return Any(a).(Foo).V1 > 0 }
	pred2 := func(a T0) bool { return Any(a).(Foo).V1%2 == 0 }
	pred3 := func(a T0) bool { return Any(a).(Foo).V1 < 0 }

	cases := []struct {
		msg      string
		receiver SliceT0
		arg      func(T0) bool
		want     SliceT0
	}{
		{"Filter: pred matches all", sFoo(), pred1, sFoo()},
		{"Filter: pred matches some", sFoo(), pred2,
			SliceT0{Foo{22, "w22"}, Foo{4444, "w4444"}, Foo{22, "w22"}}},
		{"Filter: pred matches none", sFoo(), pred3, SliceT0{}},
		{"Filter: empty receiver", SliceT0{}, pred2, SliceT0{}},
	}

	for _, cs := range cases {
		got := cs.receiver.Filter(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestFilterNot(t *testing.T) {
	pred1 := func(a T0) bool { return Any(a).(Foo).V1 > 0 }
	pred2 := func(a T0) bool { return Any(a).(Foo).V1%2 == 1 }
	pred3 := func(a T0) bool { return Any(a).(Foo).V1 < 0 }

	cases := []struct {
		msg      string
		receiver SliceT0
		arg      func(T0) bool
		want     SliceT0
	}{
		{"FilterNot: pred matches all", sFoo(), pred1, SliceT0{}},
		{"FilterNot: pred matches some", sFoo(), pred2,
			SliceT0{Foo{22, "w22"}, Foo{4444, "w4444"}, Foo{22, "w22"}}},
		{"FilterNot: pred matches none", sFoo(), pred3, sFoo()},
		{"FilterNot: empty receiver", SliceT0{}, pred2, SliceT0{}},
	}

	for _, cs := range cases {
		got := cs.receiver.FilterNot(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestFirst(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SliceT0
		want     T0
		werr     error
	}{
		{"First: non-empty", sFoo(), Foo{1, "w1"}, nil},
		{"First: empty", SliceT0{}, Foo{}, errors.New("empty slice")},
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
		receiver SliceT0
		want     []T0
	}{
		{"ForEach: non-empty receiver", sFoo(), []T0{1, 22, 333, 4444, 22}},
		{"ForEach: empty receiver", SliceT0{}, []T0{}},
	}

	for _, cs := range cases {
		got := []T0{}
		f := func(a T0) {
			got = append(got, Any(a).(Foo).V1)
		}

		cs.receiver.ForEach(f)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestIndexOfFirst(t *testing.T) {
	pred1 := func(a T0) bool { return Any(a).(Foo).V1 > 0 }
	pred2 := func(a T0) bool { return Any(a).(Foo).V1%2 == 0 }
	pred3 := func(a T0) bool { return Any(a).(Foo).V1 < 0 }

	cases := []struct {
		msg      string
		receiver SliceT0
		arg      func(T0) bool
		want     int
	}{
		{"IndexOfFirst: match all", sFoo(), pred1, 0},
		{"IndexOfFirst: match some", sFoo(), pred2, 1},
		{"IndexOfFirst: match none", sFoo(), pred3, -1},
		{"IndexOfFirst: empty", SliceT0{}, pred1, -1},
	}

	for _, cs := range cases {
		got := cs.receiver.IndexOfFirst(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestIndexOfLast(t *testing.T) {
	pred1 := func(a T0) bool { return Any(a).(Foo).V1 > 0 }
	pred2 := func(a T0) bool { return Any(a).(Foo).V1%2 == 1 }
	pred3 := func(a T0) bool { return Any(a).(Foo).V1 < 0 }

	cases := []struct {
		msg      string
		receiver SliceT0
		arg      func(T0) bool
		want     int
	}{
		{"IndexOfLast: match all", sFoo(), pred1, 4},
		{"IndexOfLast: match some", sFoo(), pred2, 2},
		{"IndexOfLast: match none", sFoo(), pred3, -1},
		{"IndexOfLast: empty", SliceT0{}, pred1, -1},
	}

	for _, cs := range cases {
		got := cs.receiver.IndexOfLast(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestIsNotEmpty(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SliceT0
		want     bool
	}{
		{"IsNotEmpty: non-empty", sFoo(), true},
		{"IsNotEmpty: empty", SliceT0{}, false},
	}

	for _, cs := range cases {
		got := cs.receiver.IsNotEmpty()
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestLast(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SliceT0
		want     T0
		werr     error
	}{
		{"Last: non-empty", sFoo(), Foo{22, "w22"}, nil},
		{"Last: empty", SliceT0{}, nil, errors.New("empty slice")},
	}

	for _, cs := range cases {
		got, err := cs.receiver.Last()
		assert.Equal(t, cs.want, got, cs.msg)
		assert.Equal(t, cs.werr, err, cs.msg)
	}
}

func TestMaxWith(t *testing.T) {
	comp := func(a1 T0, a2 T0) int { return a1.(Foo).V1 - a2.(Foo).V1 }

	cases := []struct {
		msg      string
		receiver SliceT0
		arg      func(T0, T0) int
		want     T0
		werr     error
	}{
		{"MaxWith: non-empty receiver", sFoo(), comp, Foo{4444, "w4444"}, nil},
		{"MaxWith: empty receiver", SliceT0{}, comp, nil, errors.New("empty slice")},
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
		receiver SliceT0
		arg      SliceT0
		want     SliceT0
	}{
		{"MinusSlice: subset", sFoo(), append(sFoo()[3:4], sFoo()[1]), append(sFoo()[0:1], sFoo()[2])},
		{"MinusSlice: intersects", sFoo(), append(sFoo()[1:2], Foo{22, "xyz"}), append(sFoo()[0:1], sFoo()[2], sFoo()[3])},
		{"MinusSlice: disjoint", sFoo(), append(sFoo()[:0], Foo{22, "xyz"}, Foo{0, "abc"}), sFoo()},
		{"MinusSlice: empty slice", SliceT0{}, append(sFoo()[2:2], sFoo()[1]), SliceT0{}},
	}

	for _, cs := range cases {
		got := cs.receiver.MinusSlice(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMinusElement(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SliceT0
		arg      Foo
		want     SliceT0
	}{
		{"MinusElement: present", sFoo(), Foo{22, "w22"}, append(sFoo()[0:1], sFoo()[2:]...)},
		{"MinusElement: absent", sFoo(), Foo{22, "xyz"}, sFoo()},
		{"MinusElement: empty slice", SliceT0{}, Foo{22, "xyz"}, SliceT0{}},
	}

	for _, cs := range cases {
		got := cs.receiver.MinusElement(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMinWith(t *testing.T) {
	comp := func(a1 T0, a2 T0) int { return -(a1.(Foo).V1 - a2.(Foo).V1) }

	cases := []struct {
		msg      string
		receiver SliceT0
		arg      func(T0, T0) int
		want     T0
		werr     error
	}{
		{"MinWith: non-empty receiver", sFoo(), comp, Foo{4444, "w4444"}, nil},
		{"MinWith: empty receiver", SliceT0{}, comp, nil, errors.New("empty slice")},
	}

	for _, cs := range cases {
		got, err := cs.receiver.MinWith(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
		assert.Equal(t, cs.werr, err, cs.msg)
	}
}

func TestPartition(t *testing.T) {
	pred1 := func(a T0) bool { return Any(a).(Foo).V1 > 0 }
	pred2 := func(a T0) bool { return Any(a).(Foo).V1%2 == 0 }
	pred3 := func(a T0) bool { return Any(a).(Foo).V1 < 0 }

	cases := []struct {
		msg      string
		receiver SliceT0
		arg      func(T0) bool
		want1    SliceT0
		want2    SliceT0
	}{
		{"Partition: match all", sFoo(), pred1, sFoo(), SliceT0{}},
		{"Partition: match some", sFoo(), pred2, append(sFoo()[1:2], sFoo()[3], sFoo()[4]), append(sFoo()[0:1], sFoo()[2])},
		{"Partition: match none", sFoo(), pred3, SliceT0{}, sFoo()},
		{"Partition: empty", SliceT0{}, pred1, SliceT0{}, SliceT0{}},
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
		receiver SliceT0
		arg      SliceT0
		want     SliceT0
	}{
		{"PlusMap: non-empty + non-empty", sFoo()[:3], sFoo()[3:], sFoo()},
		{"PlusMap: non-empty + empty", sFoo()[:3], SliceT0{}, sFoo()[:3]},
		{"PlusMap: empty + non-empty", SliceT0{}, sFoo()[3:], sFoo()[3:]},
		{"PlusMap: empty + empty", SliceT0{}, SliceT0{}, SliceT0{}},
	}

	for _, cs := range cases {
		got := cs.receiver.PlusSlice(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestPlusElement(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SliceT0
		arg      T0
		want     SliceT0
	}{
		{"PlusElement: non-empty", sFoo()[:4], sFoo()[4], sFoo()},
		{"PlusElement: empty", SliceT0{}, sFoo()[4], sFoo()[4:5]},
	}

	for _, cs := range cases {
		got := cs.receiver.PlusElement(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestReduce(t *testing.T) {
	op := func(a1 T0, a2 T0) T0 { return Foo{a1.(Foo).V1 + a2.(Foo).V1, a1.(Foo).V2 + a2.(Foo).V2} }

	cases := []struct {
		msg      string
		receiver SliceT0
		arg      func(T0, T0) T0
		want     T0
		werr     error
	}{
		{"Reduce: receiver length > 1", sFoo(), op,
			Foo{1 + 22 + 333 + 4444 + 22, "w1w22w333w4444w22"}, nil},
		{"Reduce: receiver length = 1", sFoo()[2:3], op, sFoo()[2], nil},
		{"Reduce: empty receiver", SliceT0{}, op, nil, errors.New("empty slice")},
	}

	for _, cs := range cases {
		got, err := cs.receiver.Reduce(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
		assert.Equal(t, cs.werr, err, cs.msg)
	}
}

func TestReversed(t *testing.T) {
	rev := SliceT0{Foo{22, "w22"}, Foo{4444, "w4444"}, Foo{333, "w333"}, Foo{22, "w22"},
		Foo{1, "w1"}}

	cases := []struct {
		msg      string
		receiver SliceT0
		want     SliceT0
	}{
		{"Reversed: non-empty slice", sFoo(), rev},
		{"Reversed: empty slice", SliceT0{}, SliceT0{}},
	}

	for _, cs := range cases {
		got := cs.receiver.Reversed()
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSortedWith(t *testing.T) {
	comp := func(a1 T0, a2 T0) int { return -(a1.(Foo).V1 - a2.(Foo).V1) }

	sorted := SliceT0{Foo{4444, "w4444"}, Foo{333, "w333"}, Foo{22, "w22"},
		Foo{22, "w22"}, Foo{1, "w1"}}

	cases := []struct {
		msg      string
		receiver SliceT0
		arg      func(T0, T0) int
		want     SliceT0
	}{
		{"SortedWith: non-empty receiver", sFoo(), comp, sorted},
		{"SortedWith: empty receiver", SliceT0{}, comp, SliceT0{}},
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
		receiver SliceT0
		arg      int
		want     SliceT0
	}{
		{"Take: some", sFoo(), 2, sFoo()[:2]},
		{"Take: all", sFoo(), size, sFoo()},
		{"Take: none", sFoo(), 0, SliceT0{}},
		{"Take: more than length", sFoo(), size + 5, sFoo()},
		{"Take: empty receiver", SliceT0{}, 1, SliceT0{}},
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
		receiver SliceT0
		arg      int
		want     SliceT0
	}{
		{"TakeLast: some", sFoo(), 2, sFoo()[size-2:]},
		{"TakeLast: all", sFoo(), size, sFoo()},
		{"TakeLast: none", sFoo(), 0, SliceT0{}},
		{"TakeLast: more than length", sFoo(), size + 5, sFoo()},
		{"TakeLast: empty receiver", SliceT0{}, 1, SliceT0{}},
	}

	for _, cs := range cases {
		got := cs.receiver.TakeLast(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestTakeLastWhile(t *testing.T) {
	pred1 := func(a T0) bool { return Any(a).(Foo).V1 > 0 }
	pred2 := func(a T0) bool { return Any(a).(Foo).V1%2 == 0 }
	pred3 := func(a T0) bool { return Any(a).(Foo).V1 < 0 }

	cases := []struct {
		msg      string
		receiver SliceT0
		arg      func(T0) bool
		want     SliceT0
	}{
		{"TakeLastWhile: pred matches all", sFoo(), pred1, sFoo()},
		{"TakeLastWhile: pred matches some", sFoo(), pred2,
			SliceT0{Foo{4444, "w4444"}, Foo{22, "w22"}}},
		{"TakeLastWhile: pred matches none", sFoo(), pred3, SliceT0{}},
		{"TakeLastWhile: empty receiver", SliceT0{}, pred2, SliceT0{}},
	}

	for _, cs := range cases {
		got := cs.receiver.TakeLastWhile(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestTakeWhile(t *testing.T) {
	pred1 := func(a T0) bool { return Any(a).(Foo).V1 > 0 }
	pred2 := func(a T0) bool { return Any(a).(Foo).V1%2 == 1 }
	pred3 := func(a T0) bool { return Any(a).(Foo).V1 < 0 }

	cases := []struct {
		msg      string
		receiver SliceT0
		arg      func(T0) bool
		want     SliceT0
	}{
		{"TakeWhile: pred matches all", sFoo(), pred1, sFoo()},
		{"TakeWhile: pred matches some", sFoo(), pred2, SliceT0{Foo{1, "w1"}}},
		{"TakeWhile: pred matches none", sFoo(), pred3, SliceT0{}},
		{"TakeWhile: empty receiver", SliceT0{}, pred2, SliceT0{}},
	}

	for _, cs := range cases {
		got := cs.receiver.TakeWhile(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}
