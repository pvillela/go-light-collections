// Code generated -- DO NOT EDIT.

package collections

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

////
// Preliminaries

// Map used as input to functions below.
func sBase() Setint {
	return Setint{1: true, 22: true, 333: true, 4444: true}
}

func sliceBase() []int {
	return []int{1, 22, 333, 4444}
}

////
// TestMaps

func TestSet_Copy(t *testing.T) {
	cases := []struct {
		msg      string
		receiver Setint
	}{
		{"Copy: nonempty set", sBase()},
		{"Copy: empty set", Setint{}},
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
		receiver Setint
		want     int
	}{
		{"Length and Size: nonempty set", sBase(), 4},
		{"Length and Size: empty set", Setint{}, 0},
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
	pred1 := func(a int) bool { return toInt(a) > 0 }
	pred2 := func(a int) bool { return toInt(a)%2 == 0 }
	pred3 := func(a int) bool { return toInt(a) < 0 }

	cases := []struct {
		msg      string
		receiver Setint
		arg      func(int) bool
		want     bool
	}{
		{"All: pred matches all", sBase(), pred1, true},
		{"All: pred matches some", sBase(), pred2, false},
		{"All: pred matches none", sBase(), pred3, false},
		{"All: empty receiver", Setint{}, pred2, true},
		{"All: nil receiver", nil, pred2, true},
	}

	for _, cs := range cases {
		got := cs.receiver.All(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSet_Any(t *testing.T) {
	pred1 := func(a int) bool { return toInt(a) > 0 }
	pred2 := func(a int) bool { return toInt(a)%2 == 0 }
	pred3 := func(a int) bool { return toInt(a) < 0 }

	cases := []struct {
		msg      string
		receiver Setint
		arg      func(int) bool
		want     bool
	}{
		{"Any: pred matches all", sBase(), pred1, true},
		{"Any: pred matches some", sBase(), pred2, true},
		{"Any: pred matches none", sBase(), pred3, false},
		{"Any: empty receiver", Setint{}, pred2, false},
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
		receiver Setint
		arg      int
		want     bool
	}{
		{"CotainsKey: present", sBase(), 22, true},
		{"Contains: absent", sBase(), 0, false},
		{"Contains: empty set", Setint{}, 22, false},
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
		receiver Setint
		arg      Setint
		want     bool
	}{
		{"ContainsSet: nonempty receiver, other subset", sBase(),
			Setint{22: true, 333: true}, true},
		{"ContainsSet: nonempty receiver, other intersects", sBase(),
			Setint{22: true, 25: true, 333: true}, false},
		{"ContainsSet: nonempty receiver, other disjoint", sBase(),
			Setint{11: true, 25: true, 33: true}, false},
		{"ContainsSet: nonempty receiver, other empty", sBase(), Setint{}, true},
		{"ContainsSet: nonempty receiver, other nil", sBase(), nil, true},
		{"ContainsSet: empty receiver, other nonempty", Setint{},
			Setint{22: true, 333: true}, false},
		{"ContainsSet: empty receiver, other empty", Setint{}, Setint{}, true},
		{"ContainsSet: empty receiver, other nil", Setint{}, nil, true},
		{"ContainsSet: nil receiver, other nonempty", nil,
			Setint{22: true, 333: true}, false},
		{"ContainsSet: nil receiver, other empty", nil, Setint{}, true},
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
		receiver Setint
		arg      []int
		want     bool
	}{
		{"ContainsSlice: nonempty receiver, elems subset", sBase(),
			[]int{22, 333, 22}, true},
		{"ContainsSlice: nonempty receiver, elems intersects", sBase(),
			[]int{22, 25, 333}, false},
		{"ContainsSlice: nonempty receiver, elems disjoint", sBase(),
			[]int{11, 25, 33}, false},
		{"ContainsSlice: nonempty receiver, elems empty", sBase(), []int{}, true},
		{"ContainsSlice: nonempty receiver, elems nil", sBase(), nil, true},
		{"ContainsSlice: empty receiver, elems nonempty", Setint{},
			[]int{22, 333, 22}, false},
		{"ContainsSlice: empty receiver, elems empty", Setint{}, []int{}, true},
		{"ContainsSlice: empty receiver, elems nil", Setint{}, nil, true},
		{"ContainsSlice: nil receiver, elems nonempty", nil,
			[]int{22, 333, 22}, false},
		{"ContainsSlice: nil receiver, elems empty", nil, []int{}, true},
		{"ContainsSlice: nil receiver, elems nil", nil, nil, true},
	}

	for _, cs := range cases {
		got := cs.receiver.ContainsSlice(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSet_Count(t *testing.T) {
	pred1 := func(a int) bool { return toInt(a) > 0 }
	pred2 := func(a int) bool { return toInt(a)%2 == 0 }
	pred3 := func(a int) bool { return toInt(a) < 0 }

	cases := []struct {
		msg      string
		receiver Setint
		arg      func(int) bool
		want     int
	}{
		{"Count: pred matches all", sBase(), pred1, len(sBase())},
		{"Count: pred matches some", sBase(), pred2, 2},
		{"Count: pred matches none", sBase(), pred3, 0},
		{"Count: empty receiver", Setint{}, pred2, 0},
		{"Count: nil receiver", nil, pred2, 0},
	}

	for _, cs := range cases {
		got := cs.receiver.Count(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSet_Filter(t *testing.T) {
	pred1 := func(a int) bool { return toInt(a) > 0 }
	pred2 := func(a int) bool { return toInt(a)%2 == 0 }
	pred3 := func(a int) bool { return toInt(a) < 0 }

	cases := []struct {
		msg      string
		receiver Setint
		arg      func(int) bool
		want     Setint
	}{
		{"Filter: pred matches all", sBase(), pred1, sBase()},
		{"Filter: pred matches some", sBase(), pred2, Setint{22: true, 4444: true}},
		{"Filter: pred matches none", sBase(), pred3, Setint{}},
		{"Filter: empty receiver", Setint{}, pred2, Setint{}},
		{"Filter: nil receiver", nil, pred2, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.Filter(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSet_FilterNot(t *testing.T) {
	pred1 := func(a int) bool { return toInt(a) > 0 }
	pred2 := func(a int) bool { return toInt(a)%2 == 1 }
	pred3 := func(a int) bool { return toInt(a) < 0 }

	cases := []struct {
		msg      string
		receiver Setint
		arg      func(int) bool
		want     Setint
	}{
		{"FilterNot: pred matches all", sBase(), pred1, Setint{}},
		{"FilterNot: pred matches some", sBase(), pred2, Setint{22: true, 4444: true}},
		{"FilterNot: pred matches none", sBase(), pred3, sBase()},
		{"FilterNot: empty receiver", Setint{}, pred2, Setint{}},
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
		receiver Setint
		want     SetAny
	}{
		{"ForEach: nonempty receiver", sBase(),
			SetAny{1: true, 22: true, 333: true, 4444: true}},
		{"ForEach: empty receiver", Setint{}, SetAny{}},
		{"ForEach: nil receiver", nil, SetAny{}},
	}

	for _, cs := range cases {
		got := SetAny{}
		f := func(a int) {
			got[a] = true
		}

		cs.receiver.ForEach(f)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSet_Intersect(t *testing.T) {
	cases := []struct {
		msg      string
		receiver Setint
		arg      Setint
		want     Setint
	}{
		{"Intersect: nonempty receiver, other subset", sBase(),
			Setint{22: true, 333: true}, Setint{22: true, 333: true}},
		{"Intersect: nonempty receiver, other intersects", sBase(),
			Setint{22: true, 25: true, 333: true}, Setint{22: true, 333: true}},
		{"Intersect: nonempty receiver, other disjoint", sBase(),
			Setint{11: true, 25: true, 33: true}, Setint{}},
		{"Intersect: nonempty receiver, other empty", sBase(), Setint{}, Setint{}},
		{"Intersect: nonempty receiver, other nil", sBase(), nil, Setint{}},
		{"Intersect: empty receiver, other nonempty", Setint{},
			Setint{22: true, 333: true}, Setint{}},
		{"Intersect: empty receiver, other empty", Setint{}, Setint{}, Setint{}},
		{"Intersect: empty receiver, other nil", Setint{}, nil, Setint{}},
		{"Intersect: nil receiver, other nonempty", nil,
			Setint{22: true, 333: true}, nil},
		{"Intersect: nil receiver, other empty", nil, Setint{}, nil},
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
		receiver Setint
		want     bool
	}{
		{"IsEmpty: nonempty", sBase(), false},
		{"IsEmpty: empty", Setint{}, true},
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
		receiver Setint
		want     bool
	}{
		{"IsNotEmpty: nonempty", sBase(), true},
		{"IsNotEmpty: empty", Setint{}, false},
		{"IsNotEmpty: nil", nil, false},
	}

	for _, cs := range cases {
		got := cs.receiver.IsNotEmpty()
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSet_MaxWith(t *testing.T) {
	comp := func(a1 int, a2 int) int { return toInt(a1) - toInt(a2) }
	var zero int

	cases := []struct {
		msg      string
		receiver Setint
		arg      func(int, int) int
		want     int
		werr     error
	}{
		{"MaxWith: nonempty receiver", sBase(), comp, 4444, nil},
		{"MaxWith: empty receiver", Setint{}, comp, zero, errors.New("empty or nil set")},
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
		receiver Setint
		arg      int
		want     Setint
	}{
		{"MinusElement: present", sBase(), 22, Setint{1: true, 333: true, 4444: true}},
		{"MinusElement: absent", sBase(), 9, sBase()},
		{"MinusElement: empty set", Setint{}, 22, Setint{}},
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
		receiver Setint
		arg      Setint
		want     Setint
	}{
		{"MinusSet: subset", sBase(), Setint{22: true, 333: true}, Setint{1: true, 4444: true}},
		{"MinusSet: intersects", sBase(), Setint{0: true, 22: true, 9: true, 333: true},
			Setint{1: true, 4444: true}},
		{"MinusSet: disjoint", sBase(), Setint{0: true, 9: true, 42: true}, sBase()},
		{"MinusSet: empty slice", Setint{}, Setint{22: true, 333: true}, Setint{}},
		{"MinusSet: nil slice", nil, Setint{22: true, 333: true}, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.MinusSet(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSet_MinusSlice(t *testing.T) {
	cases := []struct {
		msg      string
		receiver Setint
		arg      []int
		want     Setint
	}{
		{"MinusSlice: subset", sBase(), []int{22, 333}, Setint{1: true, 4444: true}},
		{"MinusSlice: intersects", sBase(), []int{0, 22, 9, 333}, Setint{1: true, 4444: true}},
		{"MinusSlice: disjoint", sBase(), []int{0, 9, 42}, sBase()},
		{"MinusSlice: empty slice", Setint{}, []int{22, 333}, Setint{}},
		{"MinusSlice: nil slice", nil, []int{22, 333}, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.MinusSlice(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSet_MinWith(t *testing.T) {
	comp := func(a1 int, a2 int) int { return -(toInt(a1) - toInt(a2)) }
	var zero int

	cases := []struct {
		msg      string
		receiver Setint
		arg      func(int, int) int
		want     int
		werr     error
	}{
		{"MinWith: nonempty receiver", sBase(), comp, 4444, nil},
		{"MinWith: empty receiver", Setint{}, comp, zero, errors.New("empty or nil set")},
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
	pred1 := func(a int) bool { return toInt(a) > 0 }
	pred2 := func(a int) bool { return toInt(a)%2 == 0 }
	pred3 := func(a int) bool { return toInt(a) < 0 }

	cases := []struct {
		msg      string
		receiver Setint
		arg      func(int) bool
		want1    Setint
		want2    Setint
	}{
		{"Partition: match all", sBase(), pred1, sBase(), Setint{}},
		{"Partition: match some", sBase(), pred2,
			Setint{22: true, 4444: true}, Setint{1: true, 333: true}},
		{"Partition: match none", sBase(), pred3, Setint{}, sBase()},
		{"Partition: empty", Setint{}, pred1, Setint{}, Setint{}},
		{"Partition: nil", nil, pred1, Setint{}, Setint{}},
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
		receiver Setint
		arg      int
		want     Setint
	}{
		{"PlusElement: nonempty, absent", Setint{1: true, 22: true, 4444: true}, 333, sBase()},
		{"PlusElement: nonempty, present", sBase(), 333, sBase()},
		{"PlusElement: empty", Setint{}, 333, Setint{333: true}},
		{"PlusElement: nil", nil, 333, Setint{333: true}},
	}

	for _, cs := range cases {
		got := cs.receiver.PlusElement(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSet_PlusSet(t *testing.T) {
	cases := []struct {
		msg      string
		receiver Setint
		arg      Setint
		want     Setint
	}{
		{"PlusSet: nonempty + nonempty", sBase(), Setint{9: true, 333: true},
			Setint{1: true, 9: true, 22: true, 333: true, 4444: true}},
		{"PlusSet: nonempty + empty", sBase(), Setint{}, sBase()},
		{"PlusSet: nonempty + nil", sBase(), nil, sBase()},
		{"PlusSet: empty + nonempty", Setint{}, sBase(), sBase()},
		{"PlusSet: nil + nonempty", nil, sBase(), sBase()},
		{"PlusSet: empty + empty", Setint{}, Setint{}, Setint{}},
		{"PlusSet: empty + nil", Setint{}, nil, Setint{}},
		{"PlusSet: nil + empty", nil, Setint{}, Setint{}},
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
		receiver Setint
		arg      []int
		want     Setint
	}{
		{"PlusSlice: nonempty + nonempty", sBase(), []int{9, 333},
			Setint{1: true, 9: true, 22: true, 333: true, 4444: true}},
		{"PlusSlice: nonempty + empty", sBase(), []int{}, sBase()},
		{"PlusSlice: nonempty + nil", sBase(), nil, sBase()},
		{"PlusSlice: empty + nonempty", Setint{}, sliceBase(), sBase()},
		{"PlusSlice: nil + nonempty", nil, sliceBase(), sBase()},
		{"PlusSlice: empty + empty", Setint{}, []int{}, Setint{}},
		{"PlusSlice: empty + nil", Setint{}, nil, Setint{}},
		{"PlusSlice: nil + empty", nil, []int{}, Setint{}},
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
		receiver Setint
		want     []int
	}{
		{"ToSlice: nonempty", sBase(), sliceBase()},
		{"ToSlice: empty", Setint{}, []int{}},
		{"ToSlice: nil", nil, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.ToSlice()
		assert.Equal(t, Setint{}.PlusSlice(cs.want), Setint{}.PlusSlice(got), cs.msg)
	}
}

func TestSet_Put(t *testing.T) {
	cases := []struct {
		msg      string
		receiver Setint
		arg      int
		want     Setint
	}{
		{"Put: nonempty absent", Setint{1: true, 22: true, 4444: true}, 333, sBase()},
		{"Put: nonempty present", sBase(), 333, sBase()},
		{"Put: empty", Setint{}, 333, Setint{333: true}},
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
