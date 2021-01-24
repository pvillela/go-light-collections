package glc

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

////
// Preliminaries

// Map used as input to functions below.
func sBase() SetT0 {
	return SetT0{1: true, 22: true, 333: true, 4444: true}
}

func sliceBase() SliceT0 {
	return SliceT0{1, 22, 333, 4444}
}

////
// Test sets

func TestSlice_ToSet(t *testing.T) {
	slice := SliceT0{1, 22, 333, 4444}

	cases := []struct {
		msg      string
		receiver SliceT0
		want     SetT0
	}{
		{"ToSet: nonempty receiver", slice, sBase()},
		{"ToSet: empty receiver", SliceT0{}, SetT0{}},
		{"ToSet: nil receiver", nil, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.ToSet()
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSet_Copy(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SetT0
	}{
		{"Copy: nonempty set", sBase()},
		{"Copy: empty set", SetT0{}},
		{"Copy: nil set", nil},
	}

	for _, cs := range cases {
		got := cs.receiver.Copy()
		assert.Equal(t, cs.receiver, got, cs.msg)
		assert.True(t, &cs.receiver != &got, cs.msg)
	}
}

func TestSet_LengthSize(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SetT0
		want     int
	}{
		{"Length and Size: nonempty set", sBase(), 4},
		{"Length and Size: empty set", SetT0{}, 0},
		{"Length and Size: nil set", nil, 0},
	}

	for _, cs := range cases {
		got1 := cs.receiver.Length()
		assert.Equal(t, cs.want, got1, cs.msg)
		got2 := cs.receiver.Size()
		assert.Equal(t, cs.want, got2, cs.msg)
	}
}

func TestSet_All(t *testing.T) {
	pred1 := func(a T0) bool { return toInt(a) > 0 }
	pred2 := func(a T0) bool { return toInt(a)%2 == 0 }
	pred3 := func(a T0) bool { return toInt(a) < 0 }

	cases := []struct {
		msg      string
		receiver SetT0
		arg      func(T0) bool
		want     bool
	}{
		{"All: pred matches all", sBase(), pred1, true},
		{"All: pred matches some", sBase(), pred2, false},
		{"All: pred matches none", sBase(), pred3, false},
		{"All: empty receiver", SetT0{}, pred2, true},
		{"All: nil receiver", nil, pred2, true},
	}

	for _, cs := range cases {
		got := cs.receiver.All(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSet_Any(t *testing.T) {
	pred1 := func(a T0) bool { return toInt(a) > 0 }
	pred2 := func(a T0) bool { return toInt(a)%2 == 0 }
	pred3 := func(a T0) bool { return toInt(a) < 0 }

	cases := []struct {
		msg      string
		receiver SetT0
		arg      func(T0) bool
		want     bool
	}{
		{"Any: pred matches all", sBase(), pred1, true},
		{"Any: pred matches some", sBase(), pred2, true},
		{"Any: pred matches none", sBase(), pred3, false},
		{"Any: empty receiver", SetT0{}, pred2, false},
		{"Any: nil receiver", nil, pred2, false},
	}

	for _, cs := range cases {
		got := cs.receiver.Any(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSet_Contains(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SetT0
		arg      int
		want     bool
	}{
		{"CotainsKey: present", sBase(), 22, true},
		{"Contains: absent", sBase(), 0, false},
		{"Contains: empty set", SetT0{}, 22, false},
		{"Contains: nil set", nil, 22, false},
	}

	for _, cs := range cases {
		got := cs.receiver.Contains(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSet_ContainsSet(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SetT0
		arg      SetT0
		want     bool
	}{
		{"ContainsSet: nonempty receiver, other subset", sBase(),
			SetT0{22: true, 333: true}, true},
		{"ContainsSet: nonempty receiver, other intersects", sBase(),
			SetT0{22: true, 25: true, 333: true}, false},
		{"ContainsSet: nonempty receiver, other disjoint", sBase(),
			SetT0{11: true, 25: true, 33: true}, false},
		{"ContainsSet: nonempty receiver, other empty", sBase(), SetT0{}, true},
		{"ContainsSet: nonempty receiver, other nil", sBase(), nil, true},
		{"ContainsSet: empty receiver, other nonempty", SetT0{},
			SetT0{22: true, 333: true}, false},
		{"ContainsSet: empty receiver, other empty", SetT0{}, SetT0{}, true},
		{"ContainsSet: empty receiver, other nil", SetT0{}, nil, true},
		{"ContainsSet: nil receiver, other nonempty", nil,
			SetT0{22: true, 333: true}, false},
		{"ContainsSet: nil receiver, other empty", nil, SetT0{}, true},
		{"ContainsSet: nil receiver, other nil", nil, nil, true},
	}

	for _, cs := range cases {
		got := cs.receiver.ContainsSet(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSet_ContainsSlice(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SetT0
		arg      SliceT0
		want     bool
	}{
		{"ContainsSlice: nonempty receiver, elems subset", sBase(),
			SliceT0{22, 333, 22}, true},
		{"ContainsSlice: nonempty receiver, elems intersects", sBase(),
			SliceT0{22, 25, 333}, false},
		{"ContainsSlice: nonempty receiver, elems disjoint", sBase(),
			SliceT0{11, 25, 33}, false},
		{"ContainsSlice: nonempty receiver, elems empty", sBase(), SliceT0{}, true},
		{"ContainsSlice: nonempty receiver, elems nil", sBase(), nil, true},
		{"ContainsSlice: empty receiver, elems nonempty", SetT0{},
			SliceT0{22, 333, 22}, false},
		{"ContainsSlice: empty receiver, elems empty", SetT0{}, SliceT0{}, true},
		{"ContainsSlice: empty receiver, elems nil", SetT0{}, nil, true},
		{"ContainsSlice: nil receiver, elems nonempty", nil,
			SliceT0{22, 333, 22}, false},
		{"ContainsSlice: nil receiver, elems empty", nil, SliceT0{}, true},
		{"ContainsSlice: nil receiver, elems nil", nil, nil, true},
	}

	for _, cs := range cases {
		got := cs.receiver.ContainsSlice(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSet_Count(t *testing.T) {
	pred1 := func(a T0) bool { return toInt(a) > 0 }
	pred2 := func(a T0) bool { return toInt(a)%2 == 0 }
	pred3 := func(a T0) bool { return toInt(a) < 0 }

	cases := []struct {
		msg      string
		receiver SetT0
		arg      func(T0) bool
		want     int
	}{
		{"Count: pred matches all", sBase(), pred1, len(sBase())},
		{"Count: pred matches some", sBase(), pred2, 2},
		{"Count: pred matches none", sBase(), pred3, 0},
		{"Count: empty receiver", SetT0{}, pred2, 0},
		{"Count: nil receiver", nil, pred2, 0},
	}

	for _, cs := range cases {
		got := cs.receiver.Count(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSet_Filter(t *testing.T) {
	pred1 := func(a T0) bool { return toInt(a) > 0 }
	pred2 := func(a T0) bool { return toInt(a)%2 == 0 }
	pred3 := func(a T0) bool { return toInt(a) < 0 }

	cases := []struct {
		msg      string
		receiver SetT0
		arg      func(T0) bool
		want     SetT0
	}{
		{"Filter: pred matches all", sBase(), pred1, sBase()},
		{"Filter: pred matches some", sBase(), pred2, SetT0{22: true, 4444: true}},
		{"Filter: pred matches none", sBase(), pred3, SetT0{}},
		{"Filter: empty receiver", SetT0{}, pred2, SetT0{}},
		{"Filter: nil receiver", nil, pred2, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.Filter(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSet_FilterNot(t *testing.T) {
	pred1 := func(a T0) bool { return toInt(a) > 0 }
	pred2 := func(a T0) bool { return toInt(a)%2 == 1 }
	pred3 := func(a T0) bool { return toInt(a) < 0 }

	cases := []struct {
		msg      string
		receiver SetT0
		arg      func(T0) bool
		want     SetT0
	}{
		{"FilterNot: pred matches all", sBase(), pred1, SetT0{}},
		{"FilterNot: pred matches some", sBase(), pred2, SetT0{22: true, 4444: true}},
		{"FilterNot: pred matches none", sBase(), pred3, sBase()},
		{"FilterNot: empty receiver", SetT0{}, pred2, SetT0{}},
		{"FilterNot: nil receiver", nil, pred2, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.FilterNot(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSet_ForEach(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SetT0
		want     SetAny
	}{
		{"ForEach: nonempty receiver", sBase(),
			SetAny{1: true, 22: true, 333: true, 4444: true}},
		{"ForEach: empty receiver", SetT0{}, SetAny{}},
		{"ForEach: nil receiver", nil, SetAny{}},
	}

	for _, cs := range cases {
		got := SetAny{}
		f := func(a T0) {
			got[a] = true
		}

		cs.receiver.ForEach(f)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSet_Intersect(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SetT0
		arg      SetT0
		want     SetT0
	}{
		{"Intersect: nonempty receiver, other subset", sBase(),
			SetT0{22: true, 333: true}, SetT0{22: true, 333: true}},
		{"Intersect: nonempty receiver, other intersects", sBase(),
			SetT0{22: true, 25: true, 333: true}, SetT0{22: true, 333: true}},
		{"Intersect: nonempty receiver, other disjoint", sBase(),
			SetT0{11: true, 25: true, 33: true}, SetT0{}},
		{"Intersect: nonempty receiver, other empty", sBase(), SetT0{}, SetT0{}},
		{"Intersect: nonempty receiver, other nil", sBase(), nil, SetT0{}},
		{"Intersect: empty receiver, other nonempty", SetT0{},
			SetT0{22: true, 333: true}, SetT0{}},
		{"Intersect: empty receiver, other empty", SetT0{}, SetT0{}, SetT0{}},
		{"Intersect: empty receiver, other nil", SetT0{}, nil, SetT0{}},
		{"Intersect: nil receiver, other nonempty", nil,
			SetT0{22: true, 333: true}, nil},
		{"Intersect: nil receiver, other empty", nil, SetT0{}, nil},
		{"Intersect: nil receiver, other nil", nil, nil, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.Intersect(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSet_IsEmpty(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SetT0
		want     bool
	}{
		{"IsEmpty: nonempty", sBase(), false},
		{"IsEmpty: empty", SetT0{}, true},
		{"IsEmpty: nil", nil, true},
	}

	for _, cs := range cases {
		got := cs.receiver.IsEmpty()
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSet_IsNotEmpty(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SetT0
		want     bool
	}{
		{"IsNotEmpty: nonempty", sBase(), true},
		{"IsNotEmpty: empty", SetT0{}, false},
		{"IsNotEmpty: nil", nil, false},
	}

	for _, cs := range cases {
		got := cs.receiver.IsNotEmpty()
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSet_MaxWith(t *testing.T) {
	comp := func(a1 T0, a2 T0) int { return toInt(a1) - toInt(a2) }
	var zero T0

	cases := []struct {
		msg      string
		receiver SetT0
		arg      func(T0, T0) int
		want     T0
		werr     error
	}{
		{"MaxWith: nonempty receiver", sBase(), comp, 4444, nil},
		{"MaxWith: empty receiver", SetT0{}, comp, zero, errors.New("empty or nil set")},
		{"MaxWith: nil receiver", nil, comp, zero, errors.New("empty or nil set")},
	}

	for _, cs := range cases {
		got, err := cs.receiver.MaxWith(cs.arg)
		assert.Equal(t, cs.werr, err, cs.msg)
		if err == nil {
			assert.Equal(t, cs.want, got, cs.msg)
		}
	}
}

func TestSet_MinusElement(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SetT0
		arg      T0
		want     SetT0
	}{
		{"MinusElement: present", sBase(), 22, SetT0{1: true, 333: true, 4444: true}},
		{"MinusElement: absent", sBase(), 9, sBase()},
		{"MinusElement: empty set", SetT0{}, 22, SetT0{}},
		{"MinusElement: nil set", nil, 22, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.MinusElement(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSet_MinusSet(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SetT0
		arg      SetT0
		want     SetT0
	}{
		{"MinusSet: subset", sBase(), SetT0{22: true, 333: true}, SetT0{1: true, 4444: true}},
		{"MinusSet: intersects", sBase(), SetT0{0: true, 22: true, 9: true, 333: true},
			SetT0{1: true, 4444: true}},
		{"MinusSet: disjoint", sBase(), SetT0{0: true, 9: true, 42: true}, sBase()},
		{"MinusSet: empty slice", SetT0{}, SetT0{22: true, 333: true}, SetT0{}},
		{"MinusSet: nil slice", nil, SetT0{22: true, 333: true}, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.MinusSet(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSet_MinusSlice(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SetT0
		arg      SliceT0
		want     SetT0
	}{
		{"MinusSlice: subset", sBase(), SliceT0{22, 333}, SetT0{1: true, 4444: true}},
		{"MinusSlice: intersects", sBase(), SliceT0{0, 22, 9, 333}, SetT0{1: true, 4444: true}},
		{"MinusSlice: disjoint", sBase(), SliceT0{0, 9, 42}, sBase()},
		{"MinusSlice: empty slice", SetT0{}, SliceT0{22, 333}, SetT0{}},
		{"MinusSlice: nil slice", nil, SliceT0{22, 333}, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.MinusSlice(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSet_MinWith(t *testing.T) {
	comp := func(a1 T0, a2 T0) int { return -(toInt(a1) - toInt(a2)) }
	var zero T0

	cases := []struct {
		msg      string
		receiver SetT0
		arg      func(T0, T0) int
		want     T0
		werr     error
	}{
		{"MinWith: nonempty receiver", sBase(), comp, 4444, nil},
		{"MinWith: empty receiver", SetT0{}, comp, zero, errors.New("empty or nil set")},
		{"MinWith: nil receiver", nil, comp, zero, errors.New("empty or nil set")},
	}

	for _, cs := range cases {
		got, err := cs.receiver.MinWith(cs.arg)
		assert.Equal(t, cs.werr, err, cs.msg)
		if err == nil {
			assert.Equal(t, cs.want, got, cs.msg)
		}
	}
}

func TestSet_Partition(t *testing.T) {
	pred1 := func(a T0) bool { return toInt(a) > 0 }
	pred2 := func(a T0) bool { return toInt(a)%2 == 0 }
	pred3 := func(a T0) bool { return toInt(a) < 0 }

	cases := []struct {
		msg      string
		receiver SetT0
		arg      func(T0) bool
		want1    SetT0
		want2    SetT0
	}{
		{"Partition: match all", sBase(), pred1, sBase(), SetT0{}},
		{"Partition: match some", sBase(), pred2,
			SetT0{22: true, 4444: true}, SetT0{1: true, 333: true}},
		{"Partition: match none", sBase(), pred3, SetT0{}, sBase()},
		{"Partition: empty", SetT0{}, pred1, SetT0{}, SetT0{}},
		{"Partition: nil", nil, pred1, SetT0{}, SetT0{}},
	}

	for _, cs := range cases {
		got1, got2 := cs.receiver.Partition(cs.arg)
		assert.Equal(t, cs.want1, got1, cs.msg)
		assert.Equal(t, cs.want2, got2, cs.msg)
	}
}

func TestSet_PlusElement(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SetT0
		arg      T0
		want     SetT0
	}{
		{"PlusElement: nonempty, absent", SetT0{1: true, 22: true, 4444: true}, 333, sBase()},
		{"PlusElement: nonempty, present", sBase(), 333, sBase()},
		{"PlusElement: empty", SetT0{}, 333, SetT0{333: true}},
		{"PlusElement: nil", nil, 333, SetT0{333: true}},
	}

	for _, cs := range cases {
		got := cs.receiver.PlusElement(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSet_PlusSet(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SetT0
		arg      SetT0
		want     SetT0
	}{
		{"PlusSet: nonempty + nonempty", sBase(), SetT0{9: true, 333: true},
			SetT0{1: true, 9: true, 22: true, 333: true, 4444: true}},
		{"PlusSet: nonempty + empty", sBase(), SetT0{}, sBase()},
		{"PlusSet: nonempty + nil", sBase(), nil, sBase()},
		{"PlusSet: empty + nonempty", SetT0{}, sBase(), sBase()},
		{"PlusSet: nil + nonempty", nil, sBase(), sBase()},
		{"PlusSet: empty + empty", SetT0{}, SetT0{}, SetT0{}},
		{"PlusSet: empty + nil", SetT0{}, nil, SetT0{}},
		{"PlusSet: nil + empty", nil, SetT0{}, SetT0{}},
		{"PlusSet: nil + nil", nil, nil, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.PlusSet(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSet_PlusSlice(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SetT0
		arg      SliceT0
		want     SetT0
	}{
		{"PlusSlice: nonempty + nonempty", sBase(), SliceT0{9, 333},
			SetT0{1: true, 9: true, 22: true, 333: true, 4444: true}},
		{"PlusSlice: nonempty + empty", sBase(), SliceT0{}, sBase()},
		{"PlusSlice: nonempty + nil", sBase(), nil, sBase()},
		{"PlusSlice: empty + nonempty", SetT0{}, sliceBase(), sBase()},
		{"PlusSlice: nil + nonempty", nil, sliceBase(), sBase()},
		{"PlusSlice: empty + empty", SetT0{}, SliceT0{}, SetT0{}},
		{"PlusSlice: empty + nil", SetT0{}, nil, SetT0{}},
		{"PlusSlice: nil + empty", nil, SliceT0{}, SetT0{}},
		{"PlusSlice: nil + nil", nil, nil, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.PlusSlice(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSet_ToSlice(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SetT0
		want     SliceT0
	}{
		{"ToSlice: nonempty", sBase(), sliceBase()},
		{"ToSlice: empty", SetT0{}, SliceT0{}},
		{"ToSlice: nil", nil, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.ToSlice()
		assert.Equal(t, cs.want.ToSet(), got.ToSet(), cs.msg)
	}
}

func TestSet_Put(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SetT0
		arg      T0
		want     SetT0
	}{
		{"Put: nonempty absent", SetT0{1: true, 22: true, 4444: true}, 333, sBase()},
		{"Put: nonempty present", sBase(), 333, sBase()},
		{"Put: empty", SetT0{}, 333, SetT0{333: true}},
		{"Put: nil", nil, 333, nil},
	}

	for _, cs := range cases {
		if cs.receiver != nil {
			got := cs.receiver
			got.Put(cs.arg)
			assert.Equal(t, cs.want, got, cs.msg)
		} else {
			var ptf assert.PanicTestFunc = func() { cs.receiver.Put(cs.arg) }
			assert.Panics(t, ptf, cs.msg)
		}

	}
}
