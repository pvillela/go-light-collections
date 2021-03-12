/*
 * Copyright © 2021 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package g2lc

import (
	"errors"
	"github.com/pvillela/go-light-collections/pkg/g2lc/pair"
	"github.com/pvillela/go-light-collections/pkg/g2lc/slice"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Hack
func xslice_0_test[T any]() {}

////
// Preliminaries

// Dat is an example data structure.
type Dat struct {
	V1 int
	V2 string
}

// Slice used in tests. Cloned each time to avoid nasty side-effects.

func sDat() slice.Slice[Dat] {
	return slice.Slice[Dat]{Dat{1, "w1"}, Dat{22, "w22"}, Dat{333, "w333"}, Dat{4444, "w4444"},
		Dat{22, "w22"}}
}

////
// Method tests

func TestSlice_Copy(t *testing.T) {
	cases := []struct {
		msg      string
		receiver slice.Slice[Dat]
	}{
		{"Copy: nonempty slice", sDat()},
		{"Copy: empty slice", slice.Slice[Dat]{}},
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
		receiver slice.Slice[Dat]
		want     int
	}{
		{"Length and Size: nonempty slice", sDat(), 5},
		{"Length and Size: empty slice", slice.Slice[Dat]{}, 0},
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
		receiver slice.Slice[Dat]
		arg      Dat
		want     bool
	}{
		{"Contains: present", sDat(), Dat{22, "w22"}, true},
		{"Contains: absent", sDat(), Dat{22, "xyz"}, false},
		{"Contains: empty slice", slice.Slice[Dat]{}, Dat{22, "w22"}, false},
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
		receiver slice.Slice[Dat]
		arg      slice.Slice[Dat]
		want     bool
	}{
		{"ContainsAll: nonempty receiver, other subset", sDat(),
			append(sDat()[2:3], sDat()[1]), true},
		{"ContainsAll: nonempty receiver, other intersects", sDat(),
			append(sDat()[1:2], Dat{22, "xyz"}), false},
		{"ContainsAll: nonempty receiver, other disjoint", sDat(),
			slice.Slice[Dat]{Dat{22, "xyz"}, Dat{0, "abc"}}, false},
		{"ContainsAll: nonempty receiver, other empty", sDat(), slice.Slice[Dat]{}, true},
		{"ContainsAll: nonempty receiver, other nil", sDat(), nil, true},
		{"ContainsAll: empty receiver, other nonempty", slice.Slice[Dat]{}, append(sDat()[2:3], sDat()[1]), false},
		{"ContainsAll: empty receiver, other empty", slice.Slice[Dat]{}, slice.Slice[Dat]{}, true},
		{"ContainsAll: empty receiver, other nil", slice.Slice[Dat]{}, nil, true},
		{"ContainsAll: nil receiver, other nonempty", nil, append(sDat()[2:3], sDat()[1]), false},
		{"ContainsAll: nil receiver, other empty", nil, slice.Slice[Dat]{}, true},
		{"ContainsAll: nil receiver, other nil", nil, nil, true},
	}

	for _, cs := range cases {
		got := cs.receiver.ContainsAll(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_Get(t *testing.T) {
	size := len(sDat())
	cases := []struct {
		msg      string
		receiver slice.Slice[Dat]
		arg      int
		want     Dat
		wok      bool
	}{
		{"Get: from middle", sDat(), 2, sDat()[2], true},
		{"Get: from beginning", sDat(), 0, sDat()[0], true},
		{"Get: from end", sDat(), size - 1, sDat()[size-1], true},
		{"Get: outside range", sDat(), size, Dat{}, false},
		{"Get: empty slice", slice.Slice[Dat]{}, 0, Dat{}, false},
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
		receiver slice.Slice[Dat]
		arg      Dat
		want     int
	}{
		{"IndexOf: nonempty, present", sDat(), Dat{22, "w22"}, 1},
		{"IndexOf: nonempty, absent", sDat(), Dat{0, "xyz"}, -1},
		{"IndexOf: empty", slice.Slice[Dat]{}, Dat{0, "xyz"}, -1},
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
		receiver slice.Slice[Dat]
		want     bool
	}{
		{"IsEmpty: nonempty", sDat(), false},
		{"IsEmpty: empty", slice.Slice[Dat]{}, true},
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
		receiver slice.Slice[Dat]
		arg      Dat
		want     int
	}{
		{"LastIndexOf: nonempty, present", sDat(), Dat{22, "w22"}, 4},
		{"LastIndexOf: nonempty, absent", sDat(), Dat{0, "xyz"}, -1},
		{"LastIndexOf: empty", slice.Slice[Dat]{}, Dat{0, "xyz"}, -1},
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
		receiver slice.Slice[Dat]
		arg1     int
		arg2     int
		want     slice.Slice[Dat]
		succeeds bool
	}{
		{"SubSlice: nonempty - from beginning", sDat(), 0, 2, sDat()[:2], true},
		{"SubSlice: nonempty - from middle", sDat(), 1, 3, sDat()[1:3], true},
		{"SubSlice: nonempty - from end", sDat(), size - 3, size, sDat()[size-3:], true},
		{"SubSlice: nonempty - empty sub-slice", sDat(), 2, 2, slice.Slice[Dat]{}, true},
		{"SubSlice: nonempty - invalid indices", sDat(), 5, 6, slice.Slice[Dat]{}, false},
		{"SubSlice: empty - empty sub-slice", slice.Slice[Dat]{}, 0, 0, slice.Slice[Dat]{}, true},
		{"SubSlice: empty - invalid indices", slice.Slice[Dat]{}, 1, 1, slice.Slice[Dat]{}, false},
		{"SubSlice: nil - empty sub-slice", nil, 0, 0, nil, true},
		{"SubSlice: nil - invalid indices", nil, 1, 1, nil, false},
	}

	for _, cs := range cases {
		if cs.succeeds {
			got := cs.receiver.SubSlice(cs.arg1, cs.arg2)
			assert.Equal(t, cs.want, got, cs.msg)
		} else {
			var ptf assert.PanicTestFunc = func() { cs.receiver.SubSlice(cs.arg1, cs.arg2) }
			assert.Panics(t, ptf, cs.msg)
		}
	}
}

func TestSlice_All(t *testing.T) {
	pred1 := func(a Dat) bool { return a.V1 > 0 }
	pred2 := func(a Dat) bool { return a.V1%2 == 0 }
	pred3 := func(a Dat) bool { return a.V1 < 0 }

	cases := []struct {
		msg      string
		receiver slice.Slice[Dat]
		arg      func(Dat) bool
		want     bool
	}{
		{"All: pred matches all", sDat(), pred1, true},
		{"All: pred matches some", sDat(), pred2, false},
		{"All: pred matches none", sDat(), pred3, false},
		{"All: empty receiver", slice.Slice[Dat]{}, pred2, true},
		{"All: nil receiver", nil, pred2, true},
	}

	for _, cs := range cases {
		got := cs.receiver.All(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_Any(t *testing.T) {
	pred1 := func(a Dat) bool { return a.V1 > 0 }
	pred2 := func(a Dat) bool { return a.V1%2 == 0 }
	pred3 := func(a Dat) bool { return a.V1 < 0 }

	cases := []struct {
		msg      string
		receiver slice.Slice[Dat]
		arg      func(Dat) bool
		want     bool
	}{
		{"Any: pred matches all", sDat(), pred1, true},
		{"Any: pred matches some", sDat(), pred2, true},
		{"Any: pred matches none", sDat(), pred3, false},
		{"Any: empty receiver", slice.Slice[Dat]{}, pred2, false},
		{"Any: nil receiver", nil, pred2, false},
	}

	for _, cs := range cases {
		got := cs.receiver.Any(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_Count(t *testing.T) {
	pred1 := func(a Dat) bool { return a.V1 > 0 }
	pred2 := func(a Dat) bool { return a.V1%2 == 0 }
	pred3 := func(a Dat) bool { return a.V1 < 0 }

	cases := []struct {
		msg      string
		receiver slice.Slice[Dat]
		arg      func(Dat) bool
		want     int
	}{
		{"Count: pred matches all", sDat(), pred1, len(sDat())},
		{"Count: pred matches some", sDat(), pred2, 3},
		{"Count: pred matches none", sDat(), pred3, 0},
		{"Count: empty receiver", slice.Slice[Dat]{}, pred2, 0},
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
		receiver slice.Slice[Dat]
		arg      int
		want     slice.Slice[Dat]
	}{
		{"Drop: some", sDat(), 2, sDat()[2:]},
		{"Drop: all", sDat(), size, slice.Slice[Dat]{}},
		{"Drop: none", sDat(), 0, sDat()},
		{"Drop: more than length", sDat(), size + 5, slice.Slice[Dat]{}},
		{"Drop: empty receiver", slice.Slice[Dat]{}, 1, slice.Slice[Dat]{}},
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
		receiver slice.Slice[Dat]
		arg      int
		want     slice.Slice[Dat]
	}{
		{"DropLast: some", sDat(), 2, sDat()[:size-2]},
		{"DropLast: all", sDat(), size, slice.Slice[Dat]{}},
		{"DropLast: none", sDat(), 0, sDat()},
		{"DropLast: more than length", sDat(), size + 5, slice.Slice[Dat]{}},
		{"DropLast: empty receiver", slice.Slice[Dat]{}, 1, slice.Slice[Dat]{}},
		{"DropLast: nil receiver", nil, 1, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.DropLast(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_DropLastWhile(t *testing.T) {
	pred1 := func(a Dat) bool { return a.V1 > 0 }
	pred2 := func(a Dat) bool { return a.V1%2 == 0 }
	pred3 := func(a Dat) bool { return a.V1 < 0 }

	cases := []struct {
		msg      string
		receiver slice.Slice[Dat]
		arg      func(Dat) bool
		want     slice.Slice[Dat]
	}{
		{"DropLastWhile: pred matches all", sDat(), pred1, slice.Slice[Dat]{}},
		{"DropLastWhile: pred matches some", sDat(), pred2,
			slice.Slice[Dat]{Dat{1, "w1"}, Dat{22, "w22"}, Dat{333, "w333"}}},
		{"DropLastWhile: pred matches none", sDat(), pred3, sDat()},
		{"DropLastWhile: empty receiver", slice.Slice[Dat]{}, pred2, slice.Slice[Dat]{}},
		{"DropLastWhile: nil receiver", nil, pred2, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.DropLastWhile(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_DropWhile(t *testing.T) {
	pred1 := func(a Dat) bool { return a.V1 > 0 }
	pred2 := func(a Dat) bool { return a.V1%2 == 1 }
	pred3 := func(a Dat) bool { return a.V1 < 0 }

	cases := []struct {
		msg      string
		receiver slice.Slice[Dat]
		arg      func(Dat) bool
		want     slice.Slice[Dat]
	}{
		{"DropWhile: pred matches all", sDat(), pred1, slice.Slice[Dat]{}},
		{"DropWhile: pred matches some", sDat(), pred2,
			slice.Slice[Dat]{Dat{22, "w22"}, Dat{333, "w333"}, Dat{4444, "w4444"}, Dat{22, "w22"}}},
		{"DropWhile: pred matches none", sDat(), pred3, sDat()},
		{"DropWhile: empty receiver", slice.Slice[Dat]{}, pred2, slice.Slice[Dat]{}},
		{"DropWhile: nil receiver", nil, pred2, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.DropWhile(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_Filter(t *testing.T) {
	pred1 := func(a Dat) bool { return a.V1 > 0 }
	pred2 := func(a Dat) bool { return a.V1%2 == 0 }
	pred3 := func(a Dat) bool { return a.V1 < 0 }

	cases := []struct {
		msg      string
		receiver slice.Slice[Dat]
		arg      func(Dat) bool
		want     slice.Slice[Dat]
	}{
		{"Filter: pred matches all", sDat(), pred1, sDat()},
		{"Filter: pred matches some", sDat(), pred2,
			slice.Slice[Dat]{Dat{22, "w22"}, Dat{4444, "w4444"}, Dat{22, "w22"}}},
		{"Filter: pred matches none", sDat(), pred3, slice.Slice[Dat]{}},
		{"Filter: empty receiver", slice.Slice[Dat]{}, pred2, slice.Slice[Dat]{}},
		{"Filter: nil receiver", nil, pred2, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.Filter(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_FilterNot(t *testing.T) {
	pred1 := func(a Dat) bool { return a.V1 > 0 }
	pred2 := func(a Dat) bool { return a.V1%2 == 1 }
	pred3 := func(a Dat) bool { return a.V1 < 0 }

	cases := []struct {
		msg      string
		receiver slice.Slice[Dat]
		arg      func(Dat) bool
		want     slice.Slice[Dat]
	}{
		{"FilterNot: pred matches all", sDat(), pred1, slice.Slice[Dat]{}},
		{"FilterNot: pred matches some", sDat(), pred2,
			slice.Slice[Dat]{Dat{22, "w22"}, Dat{4444, "w4444"}, Dat{22, "w22"}}},
		{"FilterNot: pred matches none", sDat(), pred3, sDat()},
		{"FilterNot: empty receiver", slice.Slice[Dat]{}, pred2, slice.Slice[Dat]{}},
		{"FilterNot: nil receiver", nil, pred2, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.FilterNot(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_First(t *testing.T) {
	cases := []struct {
		msg      string
		receiver slice.Slice[Dat]
		want     Dat
		werr     error
	}{
		{"First: nonempty", sDat(), Dat{1, "w1"}, nil},
		{"First: empty", slice.Slice[Dat]{}, Dat{}, errors.New("empty or nil slice")},
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
		receiver slice.Slice[Dat]
		want     []int
	}{
		{"ForEach: nonempty receiver", sDat(), []int{1, 22, 333, 4444, 22}},
		{"ForEach: empty receiver", slice.Slice[Dat]{}, []int{}},
		{"ForEach: nil receiver", nil, []int{}},
	}

	for _, cs := range cases {
		got := []int{}
		f := func(a Dat) {
			got = append(got, a.V1)
		}

		cs.receiver.ForEach(f)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_IndexOfFirst(t *testing.T) {
	pred1 := func(a Dat) bool { return a.V1 > 0 }
	pred2 := func(a Dat) bool { return a.V1%2 == 0 }
	pred3 := func(a Dat) bool { return a.V1 < 0 }

	cases := []struct {
		msg      string
		receiver slice.Slice[Dat]
		arg      func(Dat) bool
		want     int
	}{
		{"IndexOfFirst: match all", sDat(), pred1, 0},
		{"IndexOfFirst: match some", sDat(), pred2, 1},
		{"IndexOfFirst: match none", sDat(), pred3, -1},
		{"IndexOfFirst: empty", slice.Slice[Dat]{}, pred1, -1},
		{"IndexOfFirst: nil", nil, pred1, -1},
	}

	for _, cs := range cases {
		got := cs.receiver.IndexOfFirst(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_IndexOfLast(t *testing.T) {
	pred1 := func(a Dat) bool { return a.V1 > 0 }
	pred2 := func(a Dat) bool { return a.V1%2 == 1 }
	pred3 := func(a Dat) bool { return a.V1 < 0 }

	cases := []struct {
		msg      string
		receiver slice.Slice[Dat]
		arg      func(Dat) bool
		want     int
	}{
		{"IndexOfLast: match all", sDat(), pred1, 4},
		{"IndexOfLast: match some", sDat(), pred2, 2},
		{"IndexOfLast: match none", sDat(), pred3, -1},
		{"IndexOfLast: empty", slice.Slice[Dat]{}, pred1, -1},
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
		receiver slice.Slice[Dat]
		want     bool
	}{
		{"IsNotEmpty: nonempty", sDat(), true},
		{"IsNotEmpty: empty", slice.Slice[Dat]{}, false},
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
		receiver slice.Slice[Dat]
		want     Dat
		werr     error
	}{
		{"Last: nonempty", sDat(), Dat{22, "w22"}, nil},
		{"Last: empty", slice.Slice[Dat]{}, Dat{}, errors.New("empty or nil slice")},
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
	comp := func(a1 Dat, a2 Dat) int { return a1.V1 - a2.V1 }

	cases := []struct {
		msg      string
		receiver slice.Slice[Dat]
		arg      func(Dat, Dat) int
		want     Dat
		werr     error
	}{
		{"MaxWith: nonempty receiver", sDat(), comp, Dat{4444, "w4444"}, nil},
		{"MaxWith: empty receiver", slice.Slice[Dat]{}, comp, Dat{}, errors.New("empty or nil slice")},
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

func TestSlice_MinusSlice(t *testing.T) {
	cases := []struct {
		msg      string
		receiver slice.Slice[Dat]
		arg      slice.Slice[Dat]
		want     slice.Slice[Dat]
	}{
		{"MinusSlice: nonempty receiver, other subset", sDat(),
			append(sDat()[3:4], sDat()[1]), append(sDat()[0:1], sDat()[2])},
		{"MinusSlice: nonempty receiver, other intersects", sDat(),
			append(sDat()[1:2], Dat{22, "xyz"}), append(sDat()[0:1], sDat()[2], sDat()[3])},
		{"MinusSlice: nonempty receiver, other disjoint", sDat(),
			append(sDat()[:0], Dat{22, "xyz"}, Dat{0, "abc"}), sDat()},
		{"MinusSlice: nonempty receiver, other empty", sDat(), slice.Slice[Dat]{}, sDat()},
		{"MinusSlice: nonempty receiver, other nil", sDat(), nil, sDat()},
		{"MinusSlice: empty slice", slice.Slice[Dat]{}, append(sDat()[2:2], sDat()[1]), slice.Slice[Dat]{}},
		{"MinusSlice: nil slice", nil, append(sDat()[2:2], sDat()[1]), nil},
	}

	for _, cs := range cases {
		got := cs.receiver.MinusSlice(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_MinusElement(t *testing.T) {
	cases := []struct {
		msg      string
		receiver slice.Slice[Dat]
		arg      Dat
		want     slice.Slice[Dat]
	}{
		{"MinusElement: present", sDat(), Dat{22, "w22"}, append(sDat()[0:1], sDat()[2:]...)},
		{"MinusElement: absent", sDat(), Dat{22, "xyz"}, sDat()},
		{"MinusElement: empty slice", slice.Slice[Dat]{}, Dat{22, "xyz"}, slice.Slice[Dat]{}},
		{"MinusElement: nil slice", nil, Dat{22, "xyz"}, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.MinusElement(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_MinWith(t *testing.T) {
	comp := func(a1 Dat, a2 Dat) int { return -(a1.V1 - a2.V1) }

	cases := []struct {
		msg      string
		receiver slice.Slice[Dat]
		arg      func(Dat, Dat) int
		want     Dat
		werr     error
	}{
		{"MinWith: nonempty receiver", sDat(), comp, Dat{4444, "w4444"}, nil},
		{"MinWith: empty receiver", slice.Slice[Dat]{}, comp, Dat{}, errors.New("empty or nil slice")},
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
	pred1 := func(a Dat) bool { return a.V1 > 0 }
	pred2 := func(a Dat) bool { return a.V1%2 == 0 }
	pred3 := func(a Dat) bool { return a.V1 < 0 }

	cases := []struct {
		msg      string
		receiver slice.Slice[Dat]
		arg      func(Dat) bool
		want1    slice.Slice[Dat]
		want2    slice.Slice[Dat]
	}{
		{"Partition: match all", sDat(), pred1, sDat(), slice.Slice[Dat]{}},
		{"Partition: match some", sDat(), pred2, append(sDat()[1:2], sDat()[3], sDat()[4]), append(sDat()[0:1], sDat()[2])},
		{"Partition: match none", sDat(), pred3, slice.Slice[Dat]{}, sDat()},
		{"Partition: empty", slice.Slice[Dat]{}, pred1, slice.Slice[Dat]{}, slice.Slice[Dat]{}},
		{"Partition: nil", nil, pred1, slice.Slice[Dat]{}, slice.Slice[Dat]{}},
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
		receiver slice.Slice[Dat]
		arg      slice.Slice[Dat]
		want     slice.Slice[Dat]
	}{
		{"PlusSlice: nonempty + nonempty", sDat()[:3], sDat()[3:], sDat()},
		{"PlusSlice: nonempty + empty", sDat()[:3], slice.Slice[Dat]{}, sDat()[:3]},
		{"PlusSlice: empty + nonempty", slice.Slice[Dat]{}, sDat()[3:], sDat()[3:]},
		{"PlusSlice: nil + nonempty", nil, sDat()[3:], sDat()[3:]},
		{"PlusSlice: empty + empty", slice.Slice[Dat]{}, slice.Slice[Dat]{}, slice.Slice[Dat]{}},
		{"PlusSlice: empty + nil", slice.Slice[Dat]{}, nil, slice.Slice[Dat]{}},
		{"PlusSlice: nil + empty", nil, slice.Slice[Dat]{}, nil},
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
		receiver slice.Slice[Dat]
		arg      Dat
		want     slice.Slice[Dat]
	}{
		{"PlusElement: nonempty", sDat()[:4], sDat()[4], sDat()},
		{"PlusElement: empty", slice.Slice[Dat]{}, sDat()[4], sDat()[4:5]},
		{"PlusElement: nil", nil, sDat()[4], sDat()[4:5]},
	}

	for _, cs := range cases {
		got := cs.receiver.PlusElement(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_Reduce(t *testing.T) {
	op := func(a1 Dat, a2 Dat) Dat {
		foo1 := a1
		foo2 := a2
		return Dat{foo1.V1 + foo2.V1, foo1.V2 + foo2.V2}
	}

	cases := []struct {
		msg      string
		receiver slice.Slice[Dat]
		arg      func(Dat, Dat) Dat
		want     Dat
		werr     error
	}{
		{"Reduce: receiver length > 1", sDat(), op,
			Dat{1 + 22 + 333 + 4444 + 22, "w1w22w333w4444w22"}, nil},
		{"Reduce: receiver length = 1", sDat()[2:3], op, sDat()[2], nil},
		{"Reduce: empty receiver", slice.Slice[Dat]{}, op, Dat{}, errors.New("empty or nil slice")},
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
	rev := slice.Slice[Dat]{Dat{22, "w22"}, Dat{4444, "w4444"}, Dat{333, "w333"}, Dat{22, "w22"},
		Dat{1, "w1"}}

	cases := []struct {
		msg      string
		receiver slice.Slice[Dat]
		want     slice.Slice[Dat]
	}{
		{"Reversed: nonempty slice", sDat(), rev},
		{"Reversed: empty slice", slice.Slice[Dat]{}, slice.Slice[Dat]{}},
		{"Reversed: nil slice", nil, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.Reversed()
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_SortedWith(t *testing.T) {
	comp := func(a1 Dat, a2 Dat) int { return -(a1.V1 - a2.V1) }

	sorted := slice.Slice[Dat]{Dat{4444, "w4444"}, Dat{333, "w333"}, Dat{22, "w22"},
		Dat{22, "w22"}, Dat{1, "w1"}}

	cases := []struct {
		msg      string
		receiver slice.Slice[Dat]
		arg      func(Dat, Dat) int
		want     slice.Slice[Dat]
	}{
		{"SortedWith: nonempty receiver", sDat(), comp, sorted},
		{"SortedWith: empty receiver", slice.Slice[Dat]{}, comp, slice.Slice[Dat]{}},
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
		receiver slice.Slice[Dat]
		arg      int
		want     slice.Slice[Dat]
	}{
		{"Take: some", sDat(), 2, sDat()[:2]},
		{"Take: all", sDat(), size, sDat()},
		{"Take: none", sDat(), 0, slice.Slice[Dat]{}},
		{"Take: more than length", sDat(), size + 5, sDat()},
		{"Take: empty receiver", slice.Slice[Dat]{}, 1, slice.Slice[Dat]{}},
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
		receiver slice.Slice[Dat]
		arg      int
		want     slice.Slice[Dat]
	}{
		{"TakeLast: some", sDat(), 2, sDat()[size-2:]},
		{"TakeLast: all", sDat(), size, sDat()},
		{"TakeLast: none", sDat(), 0, slice.Slice[Dat]{}},
		{"TakeLast: more than length", sDat(), size + 5, sDat()},
		{"TakeLast: empty receiver", slice.Slice[Dat]{}, 1, slice.Slice[Dat]{}},
		{"TakeLast: nil receiver", nil, 1, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.TakeLast(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_TakeLastWhile(t *testing.T) {
	pred1 := func(a Dat) bool { return a.V1 > 0 }
	pred2 := func(a Dat) bool { return a.V1%2 == 0 }
	pred3 := func(a Dat) bool { return a.V1 < 0 }

	cases := []struct {
		msg      string
		receiver slice.Slice[Dat]
		arg      func(Dat) bool
		want     slice.Slice[Dat]
	}{
		{"TakeLastWhile: pred matches all", sDat(), pred1, sDat()},
		{"TakeLastWhile: pred matches some", sDat(), pred2,
			slice.Slice[Dat]{Dat{4444, "w4444"}, Dat{22, "w22"}}},
		{"TakeLastWhile: pred matches none", sDat(), pred3, slice.Slice[Dat]{}},
		{"TakeLastWhile: empty receiver", slice.Slice[Dat]{}, pred2, slice.Slice[Dat]{}},
		{"TakeLastWhile: nil receiver", nil, pred2, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.TakeLastWhile(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_TakeWhile(t *testing.T) {
	pred1 := func(a Dat) bool { return a.V1 > 0 }
	pred2 := func(a Dat) bool { return a.V1%2 == 1 }
	pred3 := func(a Dat) bool { return a.V1 < 0 }

	cases := []struct {
		msg      string
		receiver slice.Slice[Dat]
		arg      func(Dat) bool
		want     slice.Slice[Dat]
	}{
		{"TakeWhile: pred matches all", sDat(), pred1, sDat()},
		{"TakeWhile: pred matches some", sDat(), pred2, slice.Slice[Dat]{Dat{1, "w1"}}},
		{"TakeWhile: pred matches none", sDat(), pred3, slice.Slice[Dat]{}},
		{"TakeWhile: empty receiver", slice.Slice[Dat]{}, pred2, slice.Slice[Dat]{}},
		{"TakeWhile: nil receiver", nil, pred2, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.TakeWhile(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_ToSet(t *testing.T) {
	cases := []struct {
		msg      string
		receiver slice.Slice[Dat]
		want     map[Dat]bool
	}{
		{"ToSet: nonempty receiver", sDat(), map[Dat]bool{
			{1, "w1"}: true, {22, "w22"}: true, {333, "w333"}: true,
			{4444, "w4444"}: true}},
		{"ToSet: empty receiver", slice.Slice[Dat]{}, map[Dat]bool{}},
		{"ToSet: nil receiver", nil, nil},
	}

	for _, cs := range cases {
		got := slice.SliceToSet(cs.receiver)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_Flatten(t *testing.T) {
	data := slice.Slice[slice.Slice[Dat]]{{Dat{1, "w1"}, Dat{2, "w2"}}, {Dat{22, "w22"}},
		{Dat{333, "w333"}, Dat{334, "w334"}, Dat{335, "w335"}}, {Dat{4444, "w4444"}},
		{Dat{22, "w22"}}}
	cases := []struct {
		msg      string
		receiver slice.Slice[slice.Slice[Dat]]
		want     slice.Slice[Dat]
	}{
		{"Flatten: nonempty", data, slice.Slice[Dat]{Dat{1, "w1"}, Dat{2, "w2"}, Dat{22, "w22"},
			Dat{333, "w333"}, Dat{334, "w334"}, Dat{335, "w335"}, Dat{4444, "w4444"},
			Dat{22, "w22"}}},
		{"Flatten: empty", slice.Slice[slice.Slice[Dat]]{}, slice.Slice[Dat]{}},
		{"Flatten: nil", nil, nil},
	}

	for _, cs := range cases {
		got := slice.SliceFlatten(cs.receiver)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_FlatMap(t *testing.T) {
	f := func(a Dat) []int {
		n := a.V1 % 10
		s := make([]int, n)
		for i := range s {
			s[i] = n
		}
		return s
	}

	cases := []struct {
		msg      string
		receiver slice.Slice[Dat]
		arg      func(Dat) []int
		want     []int
	}{
		{"FlatMap: nonempty receiver", sDat(), f, []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 2, 2}},
		{"FlatMap: empty receiver", slice.Slice[Dat]{}, f, []int{}},
		{"FlatMap: nil receiver", nil, f, nil},
	}

	for _, cs := range cases {
		got := slice.SliceFlatMap(cs.receiver, cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_Fold(t *testing.T) {
	op := func(z int, a Dat) int { return z + a.V1 }

	cases := []struct {
		msg      string
		receiver slice.Slice[Dat]
		arg1     int
		arg2     func(z int, a Dat) int
		want     int
	}{
		{"Fold: nonempty receiver", sDat(), 1, op, 1 + 1 + 22 + 333 + 4444 + 22},
		{"Fold: empty receiver", slice.Slice[Dat]{}, 42, op, 42},
		{"Fold: nil receiver", nil, 42, op, 42},
	}

	for _, cs := range cases {
		got := slice.SliceFold(cs.receiver, cs.arg1, cs.arg2)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_Map(t *testing.T) {
	f := func(a Dat) int { return a.V1 + 1 }

	cases := []struct {
		msg      string
		receiver slice.Slice[Dat]
		arg      func(Dat) int
		want     []int
	}{
		{"Map: nonempty receiver", sDat(), f, []int{2, 23, 334, 4445, 23}},
		{"Map: empty receiver", slice.Slice[Dat]{}, f, []int{}},
		{"Map: nil receiver", nil, f, nil},
	}

	for _, cs := range cases {
		got := slice.SliceMap(cs.receiver, cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_Zip(t *testing.T) {
	shorterOther := []int{1, 2, 3}
	longerOther := []int{1, 2, 3, 4, 5, 6, 7}

	cases := []struct {
		msg      string
		receiver slice.Slice[Dat]
		arg      []int
		want     []pair.Pair[Dat, int]
	}{
		{"Zip: nonempty receiver, shorter other", sDat(), shorterOther,
			[]pair.Pair[Dat, int]{{Dat{1, "w1"}, 1}, {Dat{22, "w22"}, 2}, {Dat{333, "w333"}, 3}}},
		{"Zip: nonempty receiver, longer other", sDat(), longerOther,
			[]pair.Pair[Dat, int]{{Dat{1, "w1"}, 1}, {Dat{22, "w22"}, 2}, {Dat{333, "w333"}, 3},
				{Dat{4444, "w4444"}, 4}, {Dat{22, "w22"}, 5}}},
		{"Zip: nonempty receiver, empty other", sDat(), []int{}, []pair.Pair[Dat, int]{}},
		{"Zip: nonempty receiver, nil other", sDat(), []int{}, []pair.Pair[Dat, int]{}},
		{"Zip: empty receiver, nonempty other", slice.Slice[Dat]{}, shorterOther, []pair.Pair[Dat, int]{}},
		{"Zip: empty receiver, empty other", slice.Slice[Dat]{}, []int{}, []pair.Pair[Dat, int]{}},
		{"Zip: empty receiver, nil other", slice.Slice[Dat]{}, []int{}, []pair.Pair[Dat, int]{}},
		{"Zip: nil receiver, nonempty other", nil, shorterOther, nil},
		{"Zip: nil receiver, empty other", nil, []int{}, nil},
		{"Zip: nil receiver, nil other", nil, nil, nil},
	}

	for _, cs := range cases {
		got := slice.SliceZip(cs.receiver, cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_GroupBy(t *testing.T) {
	f := func(a Dat) int { return a.V1 % 2 }

	cases := []struct {
		msg      string
		receiver slice.Slice[Dat]
		arg      func(Dat) int
		want     map[int]slice.Slice[Dat]
	}{
		{"GroupBy: nonempty receiver", sDat(), f, map[int]slice.Slice[Dat]{
			0: {Dat{22, "w22"}, Dat{4444, "w4444"}, Dat{22, "w22"}},
			1: {Dat{1, "w1"}, Dat{333, "w333"}},
		}},
		{"GroupBy: empty receiver", slice.Slice[Dat]{}, f, map[int]slice.Slice[Dat]{}},
		{"GroupBy: nil receiver", nil, f, nil},
	}

	for _, cs := range cases {
		got := slice.SliceGroupBy(cs.receiver, cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_ToMap(t *testing.T) {
	data := slice.Slice[pair.Pair[Dat, int]]{{Dat{1, "w1"}, 10}, {Dat{22, "w22"}, 42}, {Dat{1, "w1"}, 9}}

	cases := []struct {
		msg      string
		receiver slice.Slice[pair.Pair[Dat, int]]
		want     map[Dat]int
	}{
		{"ToMap: nonempty receiver", data, map[Dat]int{{1, "w1"}: 9, {22, "w22"}: 42}},
		{"ToMap: empty receiver", slice.Slice[pair.Pair[Dat, int]]{}, map[Dat]int{}},
		{"ToMap: nil receiver", nil, nil},
	}

	for _, cs := range cases {
		got := slice.SliceToMap(cs.receiver)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}
