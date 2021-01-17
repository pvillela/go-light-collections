// Code generated -- DO NOT EDIT.

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
		receiver SliceDat
	}{
		{"Copy: non-empty slice", sDat()},
		{"Copy: empty slice", SliceDat{}},
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
		receiver SliceDat
		want     int
	}{
		{"Length and Size: non-empty slice", sDat(), 5},
		{"Length and Size: empty slice", SliceDat{}, 0},
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
		receiver SliceDat
		arg      Dat
		want     bool
	}{
		{"Cotains: present", sDat(), Dat{22, "w22"}, true},
		{"Contains: absent", sDat(), Dat{22, "xyz"}, false},
		{"Contains: empty slice", SliceDat{}, Dat{22, "w22"}, false},
	}

	for _, cs := range cases {
		got := cs.receiver.Contains(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestContainsAll(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SliceDat
		arg      SliceDat
		want     bool
	}{
		{"ContainsSlice: subset", sDat(), append(sDat()[2:3], sDat()[1]), true},
		{"ContainsSlice: intersects", sDat(), append(sDat()[1:2], Dat{22, "xyz"}), false},
		{"ContainsSlice: disjoint", sDat(), append(sDat()[:0], Dat{22, "xyz"}, Dat{0, "abc"}),
			false},
		{"ContainsSlice: empty slice", SliceDat{}, append(sDat()[2:3], sDat()[1]), false},
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
		receiver SliceDat
		arg      int
		want     Dat
		wok      bool
	}{
		{"Get: from middle", sDat(), 2, sDat()[2], true},
		{"Get: from beginning", sDat(), 0, sDat()[0], true},
		{"Get: from end", sDat(), size - 1, sDat()[size-1], true},
		{"Get: outside range", sDat(), size, Dat{}, false},
		{"Get: empty slice", SliceDat{}, 0, Dat{}, false},
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
		receiver SliceDat
		arg      Dat
		want     int
	}{
		{"IndexOf: non-empty, present", sDat(), Dat{22, "w22"}, 1},
		{"IndexOf: non-empty, absent", sDat(), Dat{0, "xyz"}, -1},
		{"IndexOf: empty", SliceDat{}, Dat{0, "xyz"}, -1},
	}

	for _, cs := range cases {
		got := cs.receiver.IndexOf(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestIsEmpty(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SliceDat
		want     bool
	}{
		{"IsEmpty: non-empty", sDat(), false},
		{"IsEmpty: empty", SliceDat{}, true},
	}

	for _, cs := range cases {
		got := cs.receiver.IsEmpty()
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestLastIndexOf(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SliceDat
		arg      Dat
		want     int
	}{
		{"LastIndexOf: non-empty, present", sDat(), Dat{22, "w22"}, 4},
		{"LastIndexOf: non-empty, absent", sDat(), Dat{0, "xyz"}, -1},
		{"LastIndexOf: empty", SliceDat{}, Dat{0, "xyz"}, -1},
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
		receiver SliceDat
		arg1     int
		arg2     int
		want     SliceDat
	}{
		{"SubSlice: nonempty - from beginning", sDat(), 0, 2, sDat()[:2]},
		{"SubSlice: nonempty - from middle", sDat(), 1, 3, sDat()[1:3]},
		{"SubSlice: nonempty - from end", sDat(), size - 3, size, sDat()[size-3:]},
		{"SubSlice: nonempty - empty sub-slice", sDat(), 2, 2, SliceDat{}},
		{"SubSlice: empty - empty sub-slice", SliceDat{}, 0, 0, SliceDat{}},
	}

	for _, cs := range cases {
		got := cs.receiver.SubSlice(cs.arg1, cs.arg2)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestAll(t *testing.T) {
	pred1 := func(a Dat) bool { return Any(a).(Dat).V1 > 0 }
	pred2 := func(a Dat) bool { return Any(a).(Dat).V1%2 == 0 }
	pred3 := func(a Dat) bool { return Any(a).(Dat).V1 < 0 }

	cases := []struct {
		msg      string
		receiver SliceDat
		arg      func(Dat) bool
		want     bool
	}{
		{"All: pred matches all", sDat(), pred1, true},
		{"All: pred matches some", sDat(), pred2, false},
		{"All: pred matches none", sDat(), pred3, false},
		{"All: empty receiver", SliceDat{}, pred2, true},
	}

	for _, cs := range cases {
		got := cs.receiver.All(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestAny(t *testing.T) {
	pred1 := func(a Dat) bool { return Any(a).(Dat).V1 > 0 }
	pred2 := func(a Dat) bool { return Any(a).(Dat).V1%2 == 0 }
	pred3 := func(a Dat) bool { return Any(a).(Dat).V1 < 0 }

	cases := []struct {
		msg      string
		receiver SliceDat
		arg      func(Dat) bool
		want     bool
	}{
		{"Any: pred matches all", sDat(), pred1, true},
		{"Any: pred matches some", sDat(), pred2, true},
		{"Any: pred matches none", sDat(), pred3, false},
		{"Any: empty receiver", SliceDat{}, pred2, false},
	}

	for _, cs := range cases {
		got := cs.receiver.Any(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestCount(t *testing.T) {
	pred1 := func(a Dat) bool { return Any(a).(Dat).V1 > 0 }
	pred2 := func(a Dat) bool { return Any(a).(Dat).V1%2 == 0 }
	pred3 := func(a Dat) bool { return Any(a).(Dat).V1 < 0 }

	cases := []struct {
		msg      string
		receiver SliceDat
		arg      func(Dat) bool
		want     int
	}{
		{"Count: pred matches all", sDat(), pred1, len(sDat())},
		{"Count: pred matches some", sDat(), pred2, 3},
		{"Count: pred matches none", sDat(), pred3, 0},
		{"Count: empty receiver", SliceDat{}, pred2, 0},
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
		receiver SliceDat
		arg      int
		want     SliceDat
	}{
		{"Drop: some", sDat(), 2, sDat()[2:]},
		{"Drop: all", sDat(), size, SliceDat{}},
		{"Drop: none", sDat(), 0, sDat()},
		{"Drop: more than length", sDat(), size + 5, SliceDat{}},
		{"Drop: empty receiver", SliceDat{}, 1, SliceDat{}},
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
		receiver SliceDat
		arg      int
		want     SliceDat
	}{
		{"DropLast: some", sDat(), 2, sDat()[:size-2]},
		{"DropLast: all", sDat(), size, SliceDat{}},
		{"DropLast: none", sDat(), 0, sDat()},
		{"DropLast: more than length", sDat(), size + 5, SliceDat{}},
		{"DropLast: empty receiver", SliceDat{}, 1, SliceDat{}},
	}

	for _, cs := range cases {
		got := cs.receiver.DropLast(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestDropLastWhile(t *testing.T) {
	pred1 := func(a Dat) bool { return Any(a).(Dat).V1 > 0 }
	pred2 := func(a Dat) bool { return Any(a).(Dat).V1%2 == 0 }
	pred3 := func(a Dat) bool { return Any(a).(Dat).V1 < 0 }

	cases := []struct {
		msg      string
		receiver SliceDat
		arg      func(Dat) bool
		want     SliceDat
	}{
		{"DropLastWhile: pred matches all", sDat(), pred1, SliceDat{}},
		{"DropLastWhile: pred matches some", sDat(), pred2,
			SliceDat{Dat{1, "w1"}, Dat{22, "w22"}, Dat{333, "w333"}}},
		{"DropLastWhile: pred matches none", sDat(), pred3, sDat()},
		{"DropLastWhile: empty receiver", SliceDat{}, pred2, SliceDat{}},
	}

	for _, cs := range cases {
		got := cs.receiver.DropLastWhile(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestDropWhile(t *testing.T) {
	pred1 := func(a Dat) bool { return Any(a).(Dat).V1 > 0 }
	pred2 := func(a Dat) bool { return Any(a).(Dat).V1%2 == 1 }
	pred3 := func(a Dat) bool { return Any(a).(Dat).V1 < 0 }

	cases := []struct {
		msg      string
		receiver SliceDat
		arg      func(Dat) bool
		want     SliceDat
	}{
		{"DropWhile: pred matches all", sDat(), pred1, SliceDat{}},
		{"DropWhile: pred matches some", sDat(), pred2,
			SliceDat{Dat{22, "w22"}, Dat{333, "w333"}, Dat{4444, "w4444"}, Dat{22, "w22"}}},
		{"DropWhile: pred matches none", sDat(), pred3, sDat()},
		{"DropWhile: empty receiver", SliceDat{}, pred2, SliceDat{}},
	}

	for _, cs := range cases {
		got := cs.receiver.DropWhile(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestFilter(t *testing.T) {
	pred1 := func(a Dat) bool { return Any(a).(Dat).V1 > 0 }
	pred2 := func(a Dat) bool { return Any(a).(Dat).V1%2 == 0 }
	pred3 := func(a Dat) bool { return Any(a).(Dat).V1 < 0 }

	cases := []struct {
		msg      string
		receiver SliceDat
		arg      func(Dat) bool
		want     SliceDat
	}{
		{"Filter: pred matches all", sDat(), pred1, sDat()},
		{"Filter: pred matches some", sDat(), pred2,
			SliceDat{Dat{22, "w22"}, Dat{4444, "w4444"}, Dat{22, "w22"}}},
		{"Filter: pred matches none", sDat(), pred3, SliceDat{}},
		{"Filter: empty receiver", SliceDat{}, pred2, SliceDat{}},
	}

	for _, cs := range cases {
		got := cs.receiver.Filter(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestFilterNot(t *testing.T) {
	pred1 := func(a Dat) bool { return Any(a).(Dat).V1 > 0 }
	pred2 := func(a Dat) bool { return Any(a).(Dat).V1%2 == 1 }
	pred3 := func(a Dat) bool { return Any(a).(Dat).V1 < 0 }

	cases := []struct {
		msg      string
		receiver SliceDat
		arg      func(Dat) bool
		want     SliceDat
	}{
		{"FilterNot: pred matches all", sDat(), pred1, SliceDat{}},
		{"FilterNot: pred matches some", sDat(), pred2,
			SliceDat{Dat{22, "w22"}, Dat{4444, "w4444"}, Dat{22, "w22"}}},
		{"FilterNot: pred matches none", sDat(), pred3, sDat()},
		{"FilterNot: empty receiver", SliceDat{}, pred2, SliceDat{}},
	}

	for _, cs := range cases {
		got := cs.receiver.FilterNot(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestFirst(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SliceDat
		want     Dat
		werr     error
	}{
		{"First: non-empty", sDat(), Dat{1, "w1"}, nil},
		{"First: empty", SliceDat{}, Dat{}, errors.New("empty slice")},
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
		receiver SliceDat
		want     []int
	}{
		{"ForEach: non-empty receiver", sDat(), []int{1, 22, 333, 4444, 22}},
		{"ForEach: empty receiver", SliceDat{}, []int{}},
	}

	for _, cs := range cases {
		got := []int{}
		f := func(a Dat) {
			got = append(got, Any(a).(Dat).V1)
		}

		cs.receiver.ForEach(f)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestIndexOfFirst(t *testing.T) {
	pred1 := func(a Dat) bool { return Any(a).(Dat).V1 > 0 }
	pred2 := func(a Dat) bool { return Any(a).(Dat).V1%2 == 0 }
	pred3 := func(a Dat) bool { return Any(a).(Dat).V1 < 0 }

	cases := []struct {
		msg      string
		receiver SliceDat
		arg      func(Dat) bool
		want     int
	}{
		{"IndexOfFirst: match all", sDat(), pred1, 0},
		{"IndexOfFirst: match some", sDat(), pred2, 1},
		{"IndexOfFirst: match none", sDat(), pred3, -1},
		{"IndexOfFirst: empty", SliceDat{}, pred1, -1},
	}

	for _, cs := range cases {
		got := cs.receiver.IndexOfFirst(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestIndexOfLast(t *testing.T) {
	pred1 := func(a Dat) bool { return Any(a).(Dat).V1 > 0 }
	pred2 := func(a Dat) bool { return Any(a).(Dat).V1%2 == 1 }
	pred3 := func(a Dat) bool { return Any(a).(Dat).V1 < 0 }

	cases := []struct {
		msg      string
		receiver SliceDat
		arg      func(Dat) bool
		want     int
	}{
		{"IndexOfLast: match all", sDat(), pred1, 4},
		{"IndexOfLast: match some", sDat(), pred2, 2},
		{"IndexOfLast: match none", sDat(), pred3, -1},
		{"IndexOfLast: empty", SliceDat{}, pred1, -1},
	}

	for _, cs := range cases {
		got := cs.receiver.IndexOfLast(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestIsNotEmpty(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SliceDat
		want     bool
	}{
		{"IsNotEmpty: non-empty", sDat(), true},
		{"IsNotEmpty: empty", SliceDat{}, false},
	}

	for _, cs := range cases {
		got := cs.receiver.IsNotEmpty()
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestLast(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SliceDat
		want     Dat
		werr     error
	}{
		{"Last: non-empty", sDat(), Dat{22, "w22"}, nil},
		{"Last: empty", SliceDat{}, Dat{}, errors.New("empty slice")},
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
	comp := func(a1 Dat, a2 Dat) int { return Any(a1).(Dat).V1 - Any(a2).(Dat).V1 }

	cases := []struct {
		msg      string
		receiver SliceDat
		arg      func(Dat, Dat) int
		want     Dat
		werr     error
	}{
		{"MaxWith: non-empty receiver", sDat(), comp, Dat{4444, "w4444"}, nil},
		{"MaxWith: empty receiver", SliceDat{}, comp, Dat{}, errors.New("empty slice")},
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
		receiver SliceDat
		arg      SliceDat
		want     SliceDat
	}{
		{"MinusSlice: subset", sDat(), append(sDat()[3:4], sDat()[1]), append(sDat()[0:1], sDat()[2])},
		{"MinusSlice: intersects", sDat(), append(sDat()[1:2], Dat{22, "xyz"}), append(sDat()[0:1], sDat()[2], sDat()[3])},
		{"MinusSlice: disjoint", sDat(), append(sDat()[:0], Dat{22, "xyz"}, Dat{0, "abc"}), sDat()},
		{"MinusSlice: empty slice", SliceDat{}, append(sDat()[2:2], sDat()[1]), SliceDat{}},
	}

	for _, cs := range cases {
		got := cs.receiver.MinusSlice(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMinusElement(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SliceDat
		arg      Dat
		want     SliceDat
	}{
		{"MinusElement: present", sDat(), Dat{22, "w22"}, append(sDat()[0:1], sDat()[2:]...)},
		{"MinusElement: absent", sDat(), Dat{22, "xyz"}, sDat()},
		{"MinusElement: empty slice", SliceDat{}, Dat{22, "xyz"}, SliceDat{}},
	}

	for _, cs := range cases {
		got := cs.receiver.MinusElement(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMinWith(t *testing.T) {
	comp := func(a1 Dat, a2 Dat) int { return -(Any(a1).(Dat).V1 - Any(a2).(Dat).V1) }

	cases := []struct {
		msg      string
		receiver SliceDat
		arg      func(Dat, Dat) int
		want     Dat
		werr     error
	}{
		{"MinWith: non-empty receiver", sDat(), comp, Dat{4444, "w4444"}, nil},
		{"MinWith: empty receiver", SliceDat{}, comp, Dat{}, errors.New("empty slice")},
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
	pred1 := func(a Dat) bool { return Any(a).(Dat).V1 > 0 }
	pred2 := func(a Dat) bool { return Any(a).(Dat).V1%2 == 0 }
	pred3 := func(a Dat) bool { return Any(a).(Dat).V1 < 0 }

	cases := []struct {
		msg      string
		receiver SliceDat
		arg      func(Dat) bool
		want1    SliceDat
		want2    SliceDat
	}{
		{"Partition: match all", sDat(), pred1, sDat(), SliceDat{}},
		{"Partition: match some", sDat(), pred2, append(sDat()[1:2], sDat()[3], sDat()[4]), append(sDat()[0:1], sDat()[2])},
		{"Partition: match none", sDat(), pred3, SliceDat{}, sDat()},
		{"Partition: empty", SliceDat{}, pred1, SliceDat{}, SliceDat{}},
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
		receiver SliceDat
		arg      SliceDat
		want     SliceDat
	}{
		{"PlusMap: non-empty + non-empty", sDat()[:3], sDat()[3:], sDat()},
		{"PlusMap: non-empty + empty", sDat()[:3], SliceDat{}, sDat()[:3]},
		{"PlusMap: empty + non-empty", SliceDat{}, sDat()[3:], sDat()[3:]},
		{"PlusMap: empty + empty", SliceDat{}, SliceDat{}, SliceDat{}},
	}

	for _, cs := range cases {
		got := cs.receiver.PlusSlice(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestPlusElement(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SliceDat
		arg      Dat
		want     SliceDat
	}{
		{"PlusElement: non-empty", sDat()[:4], sDat()[4], sDat()},
		{"PlusElement: empty", SliceDat{}, sDat()[4], sDat()[4:5]},
	}

	for _, cs := range cases {
		got := cs.receiver.PlusElement(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestReduce(t *testing.T) {
	op := func(a1 Dat, a2 Dat) Dat {
		foo1 := Any(a1).(Dat)
		foo2 := Any(a2).(Dat)
		return Dat{foo1.V1 + foo2.V1, foo1.V2 + foo2.V2}
	}

	cases := []struct {
		msg      string
		receiver SliceDat
		arg      func(Dat, Dat) Dat
		want     Dat
		werr     error
	}{
		{"Reduce: receiver length > 1", sDat(), op,
			Dat{1 + 22 + 333 + 4444 + 22, "w1w22w333w4444w22"}, nil},
		{"Reduce: receiver length = 1", sDat()[2:3], op, sDat()[2], nil},
		{"Reduce: empty receiver", SliceDat{}, op, Dat{}, errors.New("empty slice")},
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
	rev := SliceDat{Dat{22, "w22"}, Dat{4444, "w4444"}, Dat{333, "w333"}, Dat{22, "w22"},
		Dat{1, "w1"}}

	cases := []struct {
		msg      string
		receiver SliceDat
		want     SliceDat
	}{
		{"Reversed: non-empty slice", sDat(), rev},
		{"Reversed: empty slice", SliceDat{}, SliceDat{}},
	}

	for _, cs := range cases {
		got := cs.receiver.Reversed()
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSortedWith(t *testing.T) {
	comp := func(a1 Dat, a2 Dat) int { return -(Any(a1).(Dat).V1 - Any(a2).(Dat).V1) }

	sorted := SliceDat{Dat{4444, "w4444"}, Dat{333, "w333"}, Dat{22, "w22"},
		Dat{22, "w22"}, Dat{1, "w1"}}

	cases := []struct {
		msg      string
		receiver SliceDat
		arg      func(Dat, Dat) int
		want     SliceDat
	}{
		{"SortedWith: non-empty receiver", sDat(), comp, sorted},
		{"SortedWith: empty receiver", SliceDat{}, comp, SliceDat{}},
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
		receiver SliceDat
		arg      int
		want     SliceDat
	}{
		{"Take: some", sDat(), 2, sDat()[:2]},
		{"Take: all", sDat(), size, sDat()},
		{"Take: none", sDat(), 0, SliceDat{}},
		{"Take: more than length", sDat(), size + 5, sDat()},
		{"Take: empty receiver", SliceDat{}, 1, SliceDat{}},
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
		receiver SliceDat
		arg      int
		want     SliceDat
	}{
		{"TakeLast: some", sDat(), 2, sDat()[size-2:]},
		{"TakeLast: all", sDat(), size, sDat()},
		{"TakeLast: none", sDat(), 0, SliceDat{}},
		{"TakeLast: more than length", sDat(), size + 5, sDat()},
		{"TakeLast: empty receiver", SliceDat{}, 1, SliceDat{}},
	}

	for _, cs := range cases {
		got := cs.receiver.TakeLast(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestTakeLastWhile(t *testing.T) {
	pred1 := func(a Dat) bool { return Any(a).(Dat).V1 > 0 }
	pred2 := func(a Dat) bool { return Any(a).(Dat).V1%2 == 0 }
	pred3 := func(a Dat) bool { return Any(a).(Dat).V1 < 0 }

	cases := []struct {
		msg      string
		receiver SliceDat
		arg      func(Dat) bool
		want     SliceDat
	}{
		{"TakeLastWhile: pred matches all", sDat(), pred1, sDat()},
		{"TakeLastWhile: pred matches some", sDat(), pred2,
			SliceDat{Dat{4444, "w4444"}, Dat{22, "w22"}}},
		{"TakeLastWhile: pred matches none", sDat(), pred3, SliceDat{}},
		{"TakeLastWhile: empty receiver", SliceDat{}, pred2, SliceDat{}},
	}

	for _, cs := range cases {
		got := cs.receiver.TakeLastWhile(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestTakeWhile(t *testing.T) {
	pred1 := func(a Dat) bool { return Any(a).(Dat).V1 > 0 }
	pred2 := func(a Dat) bool { return Any(a).(Dat).V1%2 == 1 }
	pred3 := func(a Dat) bool { return Any(a).(Dat).V1 < 0 }

	cases := []struct {
		msg      string
		receiver SliceDat
		arg      func(Dat) bool
		want     SliceDat
	}{
		{"TakeWhile: pred matches all", sDat(), pred1, sDat()},
		{"TakeWhile: pred matches some", sDat(), pred2, SliceDat{Dat{1, "w1"}}},
		{"TakeWhile: pred matches none", sDat(), pred3, SliceDat{}},
		{"TakeWhile: empty receiver", SliceDat{}, pred2, SliceDat{}},
	}

	for _, cs := range cases {
		got := cs.receiver.TakeWhile(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}
