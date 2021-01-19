package collections

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlice_Copy(t *testing.T) {
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

func TestSlice_LengthSize(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SliceT0
		want     int
	}{
		{"Length and Size: non-empty slice", sDat(), 5},
		{"Length and Size: empty slice", SliceT0{}, 0},
		{"Length and Size: nil slice", nil, 0},
	}

	for _, cs := range cases {
		got1 := cs.receiver.Length()
		assert.Equal(t, cs.want, got1, cs.msg)
		got2 := cs.receiver.Size()
		assert.Equal(t, cs.want, got2, cs.msg)
	}
}

func TestSlice_Contains(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SliceT0
		arg      Dat
		want     bool
	}{
		{"Contains: present", sDat(), Dat{22, "w22"}, true},
		{"Contains: absent", sDat(), Dat{22, "xyz"}, false},
		{"Contains: empty slice", SliceT0{}, Dat{22, "w22"}, false},
		{"Contains: nil slice", nil, Dat{22, "w22"}, false},
	}

	for _, cs := range cases {
		got := cs.receiver.Contains(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_ContainsAll(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SliceT0
		arg      SliceT0
		want     bool
	}{
		{"ContainsSlice: subset", sDat(), append(sDat()[2:3], sDat()[1]), true},
		{"ContainsSlice: intersects", sDat(), append(sDat()[1:2], Dat{22, "xyz"}), false},
		{"ContainsSlice: disjoint", sDat(), SliceT0{Dat{22, "xyz"}, Dat{0, "abc"}}, false},
		{"ContainsSlice: empty slice", SliceT0{}, append(sDat()[2:3], sDat()[1]), false},
		{"ContainsSlice: nil slice", nil, append(sDat()[2:3], sDat()[1]), false},
	}

	for _, cs := range cases {
		got := cs.receiver.ContainsSlice(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_Get(t *testing.T) {
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
		{"Get: nil slice", nil, 0, Dat{}, false},
	}

	for _, cs := range cases {
		got, ok := cs.receiver.Get(cs.arg)
		assert.Equal(t, cs.wok, ok, cs.msg)
		if ok {
			assert.Equal(t, cs.want, got, cs.msg)
		}
	}
}

func TestSlice_IndexOf(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SliceT0
		arg      Dat
		want     int
	}{
		{"IndexOf: non-empty, present", sDat(), Dat{22, "w22"}, 1},
		{"IndexOf: non-empty, absent", sDat(), Dat{0, "xyz"}, -1},
		{"IndexOf: empty", SliceT0{}, Dat{0, "xyz"}, -1},
		{"IndexOf: nil", nil, Dat{0, "xyz"}, -1},
	}

	for _, cs := range cases {
		got := cs.receiver.IndexOf(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_IsEmpty(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SliceT0
		want     bool
	}{
		{"IsEmpty: non-empty", sDat(), false},
		{"IsEmpty: empty", SliceT0{}, true},
		{"IsEmpty: nil", nil, true},
	}

	for _, cs := range cases {
		got := cs.receiver.IsEmpty()
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_LastIndexOf(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SliceT0
		arg      Dat
		want     int
	}{
		{"LastIndexOf: non-empty, present", sDat(), Dat{22, "w22"}, 4},
		{"LastIndexOf: non-empty, absent", sDat(), Dat{0, "xyz"}, -1},
		{"LastIndexOf: empty", SliceT0{}, Dat{0, "xyz"}, -1},
		{"LastIndexOf: nil", nil, Dat{0, "xyz"}, -1},
	}

	for _, cs := range cases {
		got := cs.receiver.LastIndexOf(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_SubSlice(t *testing.T) {
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
		{"SubSlice: nil - empty sub-slice", nil, 0, 0, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.SubSlice(cs.arg1, cs.arg2)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_All(t *testing.T) {
	pred1 := func(a T0) bool { return toDat(a).V1 > 0 }
	pred2 := func(a T0) bool { return toDat(a).V1%2 == 0 }
	pred3 := func(a T0) bool { return toDat(a).V1 < 0 }

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
		{"All: nil receiver", nil, pred2, true},
	}

	for _, cs := range cases {
		got := cs.receiver.All(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_Any(t *testing.T) {
	pred1 := func(a T0) bool { return toDat(a).V1 > 0 }
	pred2 := func(a T0) bool { return toDat(a).V1%2 == 0 }
	pred3 := func(a T0) bool { return toDat(a).V1 < 0 }

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
		{"Any: nil receiver", nil, pred2, false},
	}

	for _, cs := range cases {
		got := cs.receiver.Any(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_Count(t *testing.T) {
	pred1 := func(a T0) bool { return toDat(a).V1 > 0 }
	pred2 := func(a T0) bool { return toDat(a).V1%2 == 0 }
	pred3 := func(a T0) bool { return toDat(a).V1 < 0 }

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
		{"Count: nil receiver", nil, pred2, 0},
	}

	for _, cs := range cases {
		got := cs.receiver.Count(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_Drop(t *testing.T) {
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
		{"Drop: nil receiver", nil, 1, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.Drop(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_DropLast(t *testing.T) {
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
		{"DropLast: nil receiver", nil, 1, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.DropLast(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_DropLastWhile(t *testing.T) {
	pred1 := func(a T0) bool { return toDat(a).V1 > 0 }
	pred2 := func(a T0) bool { return toDat(a).V1%2 == 0 }
	pred3 := func(a T0) bool { return toDat(a).V1 < 0 }

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
		{"DropLastWhile: nil receiver", nil, pred2, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.DropLastWhile(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_DropWhile(t *testing.T) {
	pred1 := func(a T0) bool { return toDat(a).V1 > 0 }
	pred2 := func(a T0) bool { return toDat(a).V1%2 == 1 }
	pred3 := func(a T0) bool { return toDat(a).V1 < 0 }

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
		{"DropWhile: nil receiver", nil, pred2, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.DropWhile(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_Filter(t *testing.T) {
	pred1 := func(a T0) bool { return toDat(a).V1 > 0 }
	pred2 := func(a T0) bool { return toDat(a).V1%2 == 0 }
	pred3 := func(a T0) bool { return toDat(a).V1 < 0 }

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
		{"Filter: nil receiver", nil, pred2, SliceT0{}},
	}

	for _, cs := range cases {
		got := cs.receiver.Filter(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_FilterNot(t *testing.T) {
	pred1 := func(a T0) bool { return toDat(a).V1 > 0 }
	pred2 := func(a T0) bool { return toDat(a).V1%2 == 1 }
	pred3 := func(a T0) bool { return toDat(a).V1 < 0 }

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
		{"FilterNot: nil receiver", nil, pred2, SliceT0{}},
	}

	for _, cs := range cases {
		got := cs.receiver.FilterNot(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_First(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SliceT0
		want     T0
		werr     error
	}{
		{"First: non-empty", sDat(), Dat{1, "w1"}, nil},
		{"First: empty", SliceT0{}, Dat{}, errors.New("empty or nil slice")},
		{"First: nil", nil, Dat{}, errors.New("empty or nil slice")},
	}

	for _, cs := range cases {
		got, err := cs.receiver.First()
		assert.Equal(t, cs.werr, err, cs.msg)
		if err == nil {
			assert.Equal(t, cs.want, got, cs.msg)
		}
	}
}

func TestSlice_ForEach(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SliceT0
		want     []int
	}{
		{"ForEach: non-empty receiver", sDat(), []int{1, 22, 333, 4444, 22}},
		{"ForEach: empty receiver", SliceT0{}, []int{}},
		{"ForEach: nil receiver", nil, []int{}},
	}

	for _, cs := range cases {
		got := []int{}
		f := func(a T0) {
			got = append(got, toDat(a).V1)
		}

		cs.receiver.ForEach(f)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_IndexOfFirst(t *testing.T) {
	pred1 := func(a T0) bool { return toDat(a).V1 > 0 }
	pred2 := func(a T0) bool { return toDat(a).V1%2 == 0 }
	pred3 := func(a T0) bool { return toDat(a).V1 < 0 }

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
		{"IndexOfFirst: nil", nil, pred1, -1},
	}

	for _, cs := range cases {
		got := cs.receiver.IndexOfFirst(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_IndexOfLast(t *testing.T) {
	pred1 := func(a T0) bool { return toDat(a).V1 > 0 }
	pred2 := func(a T0) bool { return toDat(a).V1%2 == 1 }
	pred3 := func(a T0) bool { return toDat(a).V1 < 0 }

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
		{"IndexOfLast: nil", nil, pred1, -1},
	}

	for _, cs := range cases {
		got := cs.receiver.IndexOfLast(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_IsNotEmpty(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SliceT0
		want     bool
	}{
		{"IsNotEmpty: non-empty", sDat(), true},
		{"IsNotEmpty: empty", SliceT0{}, false},
		{"IsNotEmpty: nil", nil, false},
	}

	for _, cs := range cases {
		got := cs.receiver.IsNotEmpty()
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_Last(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SliceT0
		want     T0
		werr     error
	}{
		{"Last: non-empty", sDat(), Dat{22, "w22"}, nil},
		{"Last: empty", SliceT0{}, Dat{}, errors.New("empty or nil slice")},
		{"Last: nil", nil, Dat{}, errors.New("empty or nil slice")},
	}

	for _, cs := range cases {
		got, err := cs.receiver.Last()
		assert.Equal(t, cs.werr, err, cs.msg)
		if err == nil {
			assert.Equal(t, cs.want, got, cs.msg)
		}
	}
}

func TestSlice_MaxWith(t *testing.T) {
	comp := func(a1 T0, a2 T0) int { return toDat(a1).V1 - toDat(a2).V1 }

	cases := []struct {
		msg      string
		receiver SliceT0
		arg      func(T0, T0) int
		want     T0
		werr     error
	}{
		{"MaxWith: non-empty receiver", sDat(), comp, Dat{4444, "w4444"}, nil},
		{"MaxWith: empty receiver", SliceT0{}, comp, Dat{}, errors.New("empty or nil slice")},
		{"MaxWith: nil receiver", nil, comp, Dat{}, errors.New("empty or nil slice")},
	}

	for _, cs := range cases {
		got, err := cs.receiver.MaxWith(cs.arg)
		assert.Equal(t, cs.werr, err, cs.msg)
		if err == nil {
			assert.Equal(t, cs.want, got, cs.msg)
		}
	}
}

func TestSlice_Minus(t *testing.T) {
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
		{"MinusSlice: nil slice", nil, append(sDat()[2:2], sDat()[1]), SliceT0{}},
	}

	for _, cs := range cases {
		got := cs.receiver.MinusSlice(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_MinusElement(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SliceT0
		arg      Dat
		want     SliceT0
	}{
		{"MinusElement: present", sDat(), Dat{22, "w22"}, append(sDat()[0:1], sDat()[2:]...)},
		{"MinusElement: absent", sDat(), Dat{22, "xyz"}, sDat()},
		{"MinusElement: empty slice", SliceT0{}, Dat{22, "xyz"}, SliceT0{}},
		{"MinusElement: nil slice", nil, Dat{22, "xyz"}, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.MinusElement(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_MinWith(t *testing.T) {
	comp := func(a1 T0, a2 T0) int { return -(toDat(a1).V1 - toDat(a2).V1) }

	cases := []struct {
		msg      string
		receiver SliceT0
		arg      func(T0, T0) int
		want     T0
		werr     error
	}{
		{"MinWith: non-empty receiver", sDat(), comp, Dat{4444, "w4444"}, nil},
		{"MinWith: empty receiver", SliceT0{}, comp, Dat{}, errors.New("empty or nil slice")},
		{"MinWith: nil receiver", nil, comp, Dat{}, errors.New("empty or nil slice")},
	}

	for _, cs := range cases {
		got, err := cs.receiver.MinWith(cs.arg)
		assert.Equal(t, cs.werr, err, cs.msg)
		if err == nil {
			assert.Equal(t, cs.want, got, cs.msg)
		}
	}
}

func TestSlice_Partition(t *testing.T) {
	pred1 := func(a T0) bool { return toDat(a).V1 > 0 }
	pred2 := func(a T0) bool { return toDat(a).V1%2 == 0 }
	pred3 := func(a T0) bool { return toDat(a).V1 < 0 }

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
		{"Partition: nil", nil, pred1, SliceT0{}, SliceT0{}},
	}

	for _, cs := range cases {
		got1, got2 := cs.receiver.Partition(cs.arg)
		assert.Equal(t, cs.want1, got1, cs.msg)
		assert.Equal(t, cs.want2, got2, cs.msg)
	}
}

func TestSlice_PlusSlice(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SliceT0
		arg      SliceT0
		want     SliceT0
	}{
		{"PlusSlice: non-empty + non-empty", sDat()[:3], sDat()[3:], sDat()},
		{"PlusSlice: non-empty + empty", sDat()[:3], SliceT0{}, sDat()[:3]},
		{"PlusSlice: empty + non-empty", SliceT0{}, sDat()[3:], sDat()[3:]},
		{"PlusSlice: nil + non-empty", nil, sDat()[3:], sDat()[3:]},
		{"PlusSlice: empty + empty", SliceT0{}, SliceT0{}, SliceT0{}},
		{"PlusSlice: empty + nil", SliceT0{}, nil, SliceT0{}},
		{"PlusSlice: nil + empty", nil, SliceT0{}, nil},
		{"PlusSlice: nil + nil", nil, nil, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.PlusSlice(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_PlusElement(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SliceT0
		arg      T0
		want     SliceT0
	}{
		{"PlusElement: non-empty", sDat()[:4], sDat()[4], sDat()},
		{"PlusElement: empty", SliceT0{}, sDat()[4], sDat()[4:5]},
		{"PlusElement: nil", nil, sDat()[4], sDat()[4:5]},
	}

	for _, cs := range cases {
		got := cs.receiver.PlusElement(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_Reduce(t *testing.T) {
	op := func(a1 T0, a2 T0) T0 {
		foo1 := toDat(a1)
		foo2 := toDat(a2)
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
		{"Reduce: empty receiver", SliceT0{}, op, Dat{}, errors.New("empty or nil slice")},
		{"Reduce: nil receiver", nil, op, Dat{}, errors.New("empty or nil slice")},
	}

	for _, cs := range cases {
		got, err := cs.receiver.Reduce(cs.arg)
		assert.Equal(t, cs.werr, err, cs.msg)
		if err == nil {
			assert.Equal(t, cs.want, got, cs.msg)
		}
	}
}

func TestSlice_Reversed(t *testing.T) {
	rev := SliceT0{Dat{22, "w22"}, Dat{4444, "w4444"}, Dat{333, "w333"}, Dat{22, "w22"},
		Dat{1, "w1"}}

	cases := []struct {
		msg      string
		receiver SliceT0
		want     SliceT0
	}{
		{"Reversed: non-empty slice", sDat(), rev},
		{"Reversed: empty slice", SliceT0{}, SliceT0{}},
		{"Reversed: nil slice", nil, SliceT0{}},
	}

	for _, cs := range cases {
		got := cs.receiver.Reversed()
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_SortedWith(t *testing.T) {
	comp := func(a1 T0, a2 T0) int { return -(toDat(a1).V1 - toDat(a2).V1) }

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
		{"SortedWith: nil receiver", nil, comp, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.SortedWith(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_Take(t *testing.T) {
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
		{"Take: nil receiver", nil, 1, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.Take(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_TakeLast(t *testing.T) {
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
		{"TakeLast: nil receiver", nil, 1, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.TakeLast(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_TakeLastWhile(t *testing.T) {
	pred1 := func(a T0) bool { return toDat(a).V1 > 0 }
	pred2 := func(a T0) bool { return toDat(a).V1%2 == 0 }
	pred3 := func(a T0) bool { return toDat(a).V1 < 0 }

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
		{"TakeLastWhile: nil receiver", nil, pred2, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.TakeLastWhile(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_TakeWhile(t *testing.T) {
	pred1 := func(a T0) bool { return toDat(a).V1 > 0 }
	pred2 := func(a T0) bool { return toDat(a).V1%2 == 1 }
	pred3 := func(a T0) bool { return toDat(a).V1 < 0 }

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
		{"TakeWhile: nil receiver", nil, pred2, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.TakeWhile(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}
