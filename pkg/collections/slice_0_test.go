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
		{"Copy: non-empty slice", sDat()},
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
		{"Length and Size: non-empty slice", sDat(), 5},
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
		arg      Dat
		want     bool
	}{
		{"Cotains: present", sDat(), Dat{22, "w22"}, true},
		{"Contains: absent", sDat(), Dat{22, "xyz"}, false},
		{"Contains: empty slice", SliceT0{}, Dat{22, "w22"}, false},
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
		{"ContainsSlice: subset", sDat(), append(sDat()[2:3], sDat()[1]), true},
		{"ContainsSlice: intersects", sDat(), append(sDat()[1:2], Dat{22, "xyz"}), false},
		{"ContainsSlice: disjoint", sDat(), append(sDat()[:0], Dat{22, "xyz"}, Dat{0, "abc"}),
			false},
		{"ContainsSlice: empty slice", SliceT0{}, append(sDat()[2:3], sDat()[1]), false},
	}

	for _, cs := range cases {
		got := cs.receiver.ContainsSlice(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestGet(t *testing.T) {
	size := len(sDat())
	cases := []struct {
		msg      string
		receiver SliceT0
		arg      int
		want     T0
		wok      bool
	}{
		{"Get: from middle", sDat(), 2, sDat()[2], true},
		{"Get: from beginning", sDat(), 0, sDat()[0], true},
		{"Get: from end", sDat(), size - 1, sDat()[size-1], true},
		{"Get: outside range", sDat(), size, Dat{}, false},
		{"Get: empty slice", SliceT0{}, 0, Dat{}, false},
	}

	for _, cs := range cases {
		got, ok := cs.receiver.Get(cs.arg)
		assert.Equal(t, cs.wok, ok, cs.msg)
		if ok {
			assert.Equal(t, cs.want, got, cs.msg)
		}
	}
}

func TestIndexOf(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SliceT0
		arg      Dat
		want     int
	}{
		{"IndexOf: non-empty, present", sDat(), Dat{22, "w22"}, 1},
		{"IndexOf: non-empty, absent", sDat(), Dat{0, "xyz"}, -1},
		{"IndexOf: empty", SliceT0{}, Dat{0, "xyz"}, -1},
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
		{"IsEmpty: non-empty", sDat(), false},
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
		arg      Dat
		want     int
	}{
		{"LastIndexOf: non-empty, present", sDat(), Dat{22, "w22"}, 4},
		{"LastIndexOf: non-empty, absent", sDat(), Dat{0, "xyz"}, -1},
		{"LastIndexOf: empty", SliceT0{}, Dat{0, "xyz"}, -1},
	}

	for _, cs := range cases {
		got := cs.receiver.LastIndexOf(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSubSlice(t *testing.T) {
	size := len(sDat())
	cases := []struct {
		msg      string
		receiver SliceT0
		arg1     int
		arg2     int
		want     SliceT0
	}{
		{"SubSlice: nonempty - from beginning", sDat(), 0, 2, sDat()[:2]},
		{"SubSlice: nonempty - from middle", sDat(), 1, 3, sDat()[1:3]},
		{"SubSlice: nonempty - from end", sDat(), size - 3, size, sDat()[size-3:]},
		{"SubSlice: nonempty - empty sub-slice", sDat(), 2, 2, SliceT0{}},
		{"SubSlice: empty - empty sub-slice", SliceT0{}, 0, 0, SliceT0{}},
	}

	for _, cs := range cases {
		got := cs.receiver.SubSlice(cs.arg1, cs.arg2)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestAll(t *testing.T) {
	pred1 := func(a T0) bool { return Any(a).(Dat).V1 > 0 }
	pred2 := func(a T0) bool { return Any(a).(Dat).V1%2 == 0 }
	pred3 := func(a T0) bool { return Any(a).(Dat).V1 < 0 }

	cases := []struct {
		msg      string
		receiver SliceT0
		arg      func(T0) bool
		want     bool
	}{
		{"All: pred matches all", sDat(), pred1, true},
		{"All: pred matches some", sDat(), pred2, false},
		{"All: pred matches none", sDat(), pred3, false},
		{"All: empty receiver", SliceT0{}, pred2, true},
	}

	for _, cs := range cases {
		got := cs.receiver.All(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestAny(t *testing.T) {
	pred1 := func(a T0) bool { return Any(a).(Dat).V1 > 0 }
	pred2 := func(a T0) bool { return Any(a).(Dat).V1%2 == 0 }
	pred3 := func(a T0) bool { return Any(a).(Dat).V1 < 0 }

	cases := []struct {
		msg      string
		receiver SliceT0
		arg      func(T0) bool
		want     bool
	}{
		{"Any: pred matches all", sDat(), pred1, true},
		{"Any: pred matches some", sDat(), pred2, true},
		{"Any: pred matches none", sDat(), pred3, false},
		{"Any: empty receiver", SliceT0{}, pred2, false},
	}

	for _, cs := range cases {
		got := cs.receiver.Any(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestCount(t *testing.T) {
	pred1 := func(a T0) bool { return Any(a).(Dat).V1 > 0 }
	pred2 := func(a T0) bool { return Any(a).(Dat).V1%2 == 0 }
	pred3 := func(a T0) bool { return Any(a).(Dat).V1 < 0 }

	cases := []struct {
		msg      string
		receiver SliceT0
		arg      func(T0) bool
		want     int
	}{
		{"Count: pred matches all", sDat(), pred1, len(sDat())},
		{"Count: pred matches some", sDat(), pred2, 3},
		{"Count: pred matches none", sDat(), pred3, 0},
		{"Count: empty receiver", SliceT0{}, pred2, 0},
	}

	for _, cs := range cases {
		got := cs.receiver.Count(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestDrop(t *testing.T) {
	size := len(sDat())
	cases := []struct {
		msg      string
		receiver SliceT0
		arg      int
		want     SliceT0
	}{
		{"Drop: some", sDat(), 2, sDat()[2:]},
		{"Drop: all", sDat(), size, SliceT0{}},
		{"Drop: none", sDat(), 0, sDat()},
		{"Drop: more than length", sDat(), size + 5, SliceT0{}},
		{"Drop: empty receiver", SliceT0{}, 1, SliceT0{}},
	}

	for _, cs := range cases {
		got := cs.receiver.Drop(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestDropLast(t *testing.T) {
	size := len(sDat())
	cases := []struct {
		msg      string
		receiver SliceT0
		arg      int
		want     SliceT0
	}{
		{"DropLast: some", sDat(), 2, sDat()[:size-2]},
		{"DropLast: all", sDat(), size, SliceT0{}},
		{"DropLast: none", sDat(), 0, sDat()},
		{"DropLast: more than length", sDat(), size + 5, SliceT0{}},
		{"DropLast: empty receiver", SliceT0{}, 1, SliceT0{}},
	}

	for _, cs := range cases {
		got := cs.receiver.DropLast(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestDropLastWhile(t *testing.T) {
	pred1 := func(a T0) bool { return Any(a).(Dat).V1 > 0 }
	pred2 := func(a T0) bool { return Any(a).(Dat).V1%2 == 0 }
	pred3 := func(a T0) bool { return Any(a).(Dat).V1 < 0 }

	cases := []struct {
		msg      string
		receiver SliceT0
		arg      func(T0) bool
		want     SliceT0
	}{
		{"DropLastWhile: pred matches all", sDat(), pred1, SliceT0{}},
		{"DropLastWhile: pred matches some", sDat(), pred2,
			SliceT0{Dat{1, "w1"}, Dat{22, "w22"}, Dat{333, "w333"}}},
		{"DropLastWhile: pred matches none", sDat(), pred3, sDat()},
		{"DropLastWhile: empty receiver", SliceT0{}, pred2, SliceT0{}},
	}

	for _, cs := range cases {
		got := cs.receiver.DropLastWhile(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestDropWhile(t *testing.T) {
	pred1 := func(a T0) bool { return Any(a).(Dat).V1 > 0 }
	pred2 := func(a T0) bool { return Any(a).(Dat).V1%2 == 1 }
	pred3 := func(a T0) bool { return Any(a).(Dat).V1 < 0 }

	cases := []struct {
		msg      string
		receiver SliceT0
		arg      func(T0) bool
		want     SliceT0
	}{
		{"DropWhile: pred matches all", sDat(), pred1, SliceT0{}},
		{"DropWhile: pred matches some", sDat(), pred2,
			SliceT0{Dat{22, "w22"}, Dat{333, "w333"}, Dat{4444, "w4444"}, Dat{22, "w22"}}},
		{"DropWhile: pred matches none", sDat(), pred3, sDat()},
		{"DropWhile: empty receiver", SliceT0{}, pred2, SliceT0{}},
	}

	for _, cs := range cases {
		got := cs.receiver.DropWhile(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestFilter(t *testing.T) {
	pred1 := func(a T0) bool { return Any(a).(Dat).V1 > 0 }
	pred2 := func(a T0) bool { return Any(a).(Dat).V1%2 == 0 }
	pred3 := func(a T0) bool { return Any(a).(Dat).V1 < 0 }

	cases := []struct {
		msg      string
		receiver SliceT0
		arg      func(T0) bool
		want     SliceT0
	}{
		{"Filter: pred matches all", sDat(), pred1, sDat()},
		{"Filter: pred matches some", sDat(), pred2,
			SliceT0{Dat{22, "w22"}, Dat{4444, "w4444"}, Dat{22, "w22"}}},
		{"Filter: pred matches none", sDat(), pred3, SliceT0{}},
		{"Filter: empty receiver", SliceT0{}, pred2, SliceT0{}},
	}

	for _, cs := range cases {
		got := cs.receiver.Filter(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestFilterNot(t *testing.T) {
	pred1 := func(a T0) bool { return Any(a).(Dat).V1 > 0 }
	pred2 := func(a T0) bool { return Any(a).(Dat).V1%2 == 1 }
	pred3 := func(a T0) bool { return Any(a).(Dat).V1 < 0 }

	cases := []struct {
		msg      string
		receiver SliceT0
		arg      func(T0) bool
		want     SliceT0
	}{
		{"FilterNot: pred matches all", sDat(), pred1, SliceT0{}},
		{"FilterNot: pred matches some", sDat(), pred2,
			SliceT0{Dat{22, "w22"}, Dat{4444, "w4444"}, Dat{22, "w22"}}},
		{"FilterNot: pred matches none", sDat(), pred3, sDat()},
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
		{"First: non-empty", sDat(), Dat{1, "w1"}, nil},
		{"First: empty", SliceT0{}, Dat{}, errors.New("empty slice")},
	}

	for _, cs := range cases {
		got, err := cs.receiver.First()
		assert.Equal(t, cs.werr, err, cs.msg)
		if err == nil {
			assert.Equal(t, cs.want, got, cs.msg)
		}
	}
}

func TestForEach(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SliceT0
		want     []int
	}{
		{"ForEach: non-empty receiver", sDat(), []int{1, 22, 333, 4444, 22}},
		{"ForEach: empty receiver", SliceT0{}, []int{}},
	}

	for _, cs := range cases {
		got := []int{}
		f := func(a T0) {
			got = append(got, Any(a).(Dat).V1)
		}

		cs.receiver.ForEach(f)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestIndexOfFirst(t *testing.T) {
	pred1 := func(a T0) bool { return Any(a).(Dat).V1 > 0 }
	pred2 := func(a T0) bool { return Any(a).(Dat).V1%2 == 0 }
	pred3 := func(a T0) bool { return Any(a).(Dat).V1 < 0 }

	cases := []struct {
		msg      string
		receiver SliceT0
		arg      func(T0) bool
		want     int
	}{
		{"IndexOfFirst: match all", sDat(), pred1, 0},
		{"IndexOfFirst: match some", sDat(), pred2, 1},
		{"IndexOfFirst: match none", sDat(), pred3, -1},
		{"IndexOfFirst: empty", SliceT0{}, pred1, -1},
	}

	for _, cs := range cases {
		got := cs.receiver.IndexOfFirst(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestIndexOfLast(t *testing.T) {
	pred1 := func(a T0) bool { return Any(a).(Dat).V1 > 0 }
	pred2 := func(a T0) bool { return Any(a).(Dat).V1%2 == 1 }
	pred3 := func(a T0) bool { return Any(a).(Dat).V1 < 0 }

	cases := []struct {
		msg      string
		receiver SliceT0
		arg      func(T0) bool
		want     int
	}{
		{"IndexOfLast: match all", sDat(), pred1, 4},
		{"IndexOfLast: match some", sDat(), pred2, 2},
		{"IndexOfLast: match none", sDat(), pred3, -1},
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
		{"IsNotEmpty: non-empty", sDat(), true},
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
		{"Last: non-empty", sDat(), Dat{22, "w22"}, nil},
		{"Last: empty", SliceT0{}, Dat{}, errors.New("empty slice")},
	}

	for _, cs := range cases {
		got, err := cs.receiver.Last()
		assert.Equal(t, cs.werr, err, cs.msg)
		if err == nil {
			assert.Equal(t, cs.want, got, cs.msg)
		}
	}
}

func TestMaxWith(t *testing.T) {
	comp := func(a1 T0, a2 T0) int { return Any(a1).(Dat).V1 - Any(a2).(Dat).V1 }

	cases := []struct {
		msg      string
		receiver SliceT0
		arg      func(T0, T0) int
		want     T0
		werr     error
	}{
		{"MaxWith: non-empty receiver", sDat(), comp, Dat{4444, "w4444"}, nil},
		{"MaxWith: empty receiver", SliceT0{}, comp, Dat{}, errors.New("empty slice")},
	}

	for _, cs := range cases {
		got, err := cs.receiver.MaxWith(cs.arg)
		assert.Equal(t, cs.werr, err, cs.msg)
		if err == nil {
			assert.Equal(t, cs.want, got, cs.msg)
		}
	}
}

func TestMinus(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SliceT0
		arg      SliceT0
		want     SliceT0
	}{
		{"MinusSlice: subset", sDat(), append(sDat()[3:4], sDat()[1]), append(sDat()[0:1], sDat()[2])},
		{"MinusSlice: intersects", sDat(), append(sDat()[1:2], Dat{22, "xyz"}), append(sDat()[0:1], sDat()[2], sDat()[3])},
		{"MinusSlice: disjoint", sDat(), append(sDat()[:0], Dat{22, "xyz"}, Dat{0, "abc"}), sDat()},
		{"MinusSlice: empty slice", SliceT0{}, append(sDat()[2:2], sDat()[1]), SliceT0{}},
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
		arg      Dat
		want     SliceT0
	}{
		{"MinusElement: present", sDat(), Dat{22, "w22"}, append(sDat()[0:1], sDat()[2:]...)},
		{"MinusElement: absent", sDat(), Dat{22, "xyz"}, sDat()},
		{"MinusElement: empty slice", SliceT0{}, Dat{22, "xyz"}, SliceT0{}},
	}

	for _, cs := range cases {
		got := cs.receiver.MinusElement(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMinWith(t *testing.T) {
	comp := func(a1 T0, a2 T0) int { return -(Any(a1).(Dat).V1 - Any(a2).(Dat).V1) }

	cases := []struct {
		msg      string
		receiver SliceT0
		arg      func(T0, T0) int
		want     T0
		werr     error
	}{
		{"MinWith: non-empty receiver", sDat(), comp, Dat{4444, "w4444"}, nil},
		{"MinWith: empty receiver", SliceT0{}, comp, Dat{}, errors.New("empty slice")},
	}

	for _, cs := range cases {
		got, err := cs.receiver.MinWith(cs.arg)
		assert.Equal(t, cs.werr, err, cs.msg)
		if err == nil {
			assert.Equal(t, cs.want, got, cs.msg)
		}
	}
}

func TestPartition(t *testing.T) {
	pred1 := func(a T0) bool { return Any(a).(Dat).V1 > 0 }
	pred2 := func(a T0) bool { return Any(a).(Dat).V1%2 == 0 }
	pred3 := func(a T0) bool { return Any(a).(Dat).V1 < 0 }

	cases := []struct {
		msg      string
		receiver SliceT0
		arg      func(T0) bool
		want1    SliceT0
		want2    SliceT0
	}{
		{"Partition: match all", sDat(), pred1, sDat(), SliceT0{}},
		{"Partition: match some", sDat(), pred2, append(sDat()[1:2], sDat()[3], sDat()[4]), append(sDat()[0:1], sDat()[2])},
		{"Partition: match none", sDat(), pred3, SliceT0{}, sDat()},
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
		{"PlusMap: non-empty + non-empty", sDat()[:3], sDat()[3:], sDat()},
		{"PlusMap: non-empty + empty", sDat()[:3], SliceT0{}, sDat()[:3]},
		{"PlusMap: empty + non-empty", SliceT0{}, sDat()[3:], sDat()[3:]},
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
		{"PlusElement: non-empty", sDat()[:4], sDat()[4], sDat()},
		{"PlusElement: empty", SliceT0{}, sDat()[4], sDat()[4:5]},
	}

	for _, cs := range cases {
		got := cs.receiver.PlusElement(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestReduce(t *testing.T) {
	op := func(a1 T0, a2 T0) T0 {
		foo1 := Any(a1).(Dat)
		foo2 := Any(a2).(Dat)
		return Dat{foo1.V1 + foo2.V1, foo1.V2 + foo2.V2}
	}

	cases := []struct {
		msg      string
		receiver SliceT0
		arg      func(T0, T0) T0
		want     T0
		werr     error
	}{
		{"Reduce: receiver length > 1", sDat(), op,
			Dat{1 + 22 + 333 + 4444 + 22, "w1w22w333w4444w22"}, nil},
		{"Reduce: receiver length = 1", sDat()[2:3], op, sDat()[2], nil},
		{"Reduce: empty receiver", SliceT0{}, op, Dat{}, errors.New("empty slice")},
	}

	for _, cs := range cases {
		got, err := cs.receiver.Reduce(cs.arg)
		assert.Equal(t, cs.werr, err, cs.msg)
		if err == nil {
			assert.Equal(t, cs.want, got, cs.msg)
		}
	}
}

func TestReversed(t *testing.T) {
	rev := SliceT0{Dat{22, "w22"}, Dat{4444, "w4444"}, Dat{333, "w333"}, Dat{22, "w22"},
		Dat{1, "w1"}}

	cases := []struct {
		msg      string
		receiver SliceT0
		want     SliceT0
	}{
		{"Reversed: non-empty slice", sDat(), rev},
		{"Reversed: empty slice", SliceT0{}, SliceT0{}},
	}

	for _, cs := range cases {
		got := cs.receiver.Reversed()
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSortedWith(t *testing.T) {
	comp := func(a1 T0, a2 T0) int { return -(Any(a1).(Dat).V1 - Any(a2).(Dat).V1) }

	sorted := SliceT0{Dat{4444, "w4444"}, Dat{333, "w333"}, Dat{22, "w22"},
		Dat{22, "w22"}, Dat{1, "w1"}}

	cases := []struct {
		msg      string
		receiver SliceT0
		arg      func(T0, T0) int
		want     SliceT0
	}{
		{"SortedWith: non-empty receiver", sDat(), comp, sorted},
		{"SortedWith: empty receiver", SliceT0{}, comp, SliceT0{}},
	}

	for _, cs := range cases {
		got := cs.receiver.SortedWith(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestTake(t *testing.T) {
	size := len(sDat())
	cases := []struct {
		msg      string
		receiver SliceT0
		arg      int
		want     SliceT0
	}{
		{"Take: some", sDat(), 2, sDat()[:2]},
		{"Take: all", sDat(), size, sDat()},
		{"Take: none", sDat(), 0, SliceT0{}},
		{"Take: more than length", sDat(), size + 5, sDat()},
		{"Take: empty receiver", SliceT0{}, 1, SliceT0{}},
	}

	for _, cs := range cases {
		got := cs.receiver.Take(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestTakeLast(t *testing.T) {
	size := len(sDat())
	cases := []struct {
		msg      string
		receiver SliceT0
		arg      int
		want     SliceT0
	}{
		{"TakeLast: some", sDat(), 2, sDat()[size-2:]},
		{"TakeLast: all", sDat(), size, sDat()},
		{"TakeLast: none", sDat(), 0, SliceT0{}},
		{"TakeLast: more than length", sDat(), size + 5, sDat()},
		{"TakeLast: empty receiver", SliceT0{}, 1, SliceT0{}},
	}

	for _, cs := range cases {
		got := cs.receiver.TakeLast(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestTakeLastWhile(t *testing.T) {
	pred1 := func(a T0) bool { return Any(a).(Dat).V1 > 0 }
	pred2 := func(a T0) bool { return Any(a).(Dat).V1%2 == 0 }
	pred3 := func(a T0) bool { return Any(a).(Dat).V1 < 0 }

	cases := []struct {
		msg      string
		receiver SliceT0
		arg      func(T0) bool
		want     SliceT0
	}{
		{"TakeLastWhile: pred matches all", sDat(), pred1, sDat()},
		{"TakeLastWhile: pred matches some", sDat(), pred2,
			SliceT0{Dat{4444, "w4444"}, Dat{22, "w22"}}},
		{"TakeLastWhile: pred matches none", sDat(), pred3, SliceT0{}},
		{"TakeLastWhile: empty receiver", SliceT0{}, pred2, SliceT0{}},
	}

	for _, cs := range cases {
		got := cs.receiver.TakeLastWhile(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestTakeWhile(t *testing.T) {
	pred1 := func(a T0) bool { return Any(a).(Dat).V1 > 0 }
	pred2 := func(a T0) bool { return Any(a).(Dat).V1%2 == 1 }
	pred3 := func(a T0) bool { return Any(a).(Dat).V1 < 0 }

	cases := []struct {
		msg      string
		receiver SliceT0
		arg      func(T0) bool
		want     SliceT0
	}{
		{"TakeWhile: pred matches all", sDat(), pred1, sDat()},
		{"TakeWhile: pred matches some", sDat(), pred2, SliceT0{Dat{1, "w1"}}},
		{"TakeWhile: pred matches none", sDat(), pred3, SliceT0{}},
		{"TakeWhile: empty receiver", SliceT0{}, pred2, SliceT0{}},
	}

	for _, cs := range cases {
		got := cs.receiver.TakeWhile(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}
