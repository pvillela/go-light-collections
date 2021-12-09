/*
 * Copyright © 2021 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package g2lc

import (
	"errors"
	"github.com/pvillela/go-light-collections/pkg/g2lc/m"
	"github.com/pvillela/go-light-collections/pkg/g2lc/pair"
	"github.com/pvillela/go-light-collections/pkg/g2lc/set"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Hack
func xset_0_test[T any]() {}

////
// Preliminaries

// Map used as input to functions below.
func sBase() set.Set[int] {
	return set.Set[int]{1: true, 22: true, 333: true, 4444: true}
}

func sliceBase() []int {
	return []int{1, 22, 333, 4444}
}

////
// TestMaps

func TestSet_Copy(t *testing.T) {
	cases := []struct {
		msg      string
		receiver set.Set[int]
	}{
		{"Copy: nonempty set", sBase()},
		{"Copy: empty set", set.Set[int]{}},
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
		receiver set.Set[int]
		want     int
	}{
		{"Length and Size: nonempty set", sBase(), 4},
		{"Length and Size: empty set", set.Set[int]{}, 0},
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
	pred1 := func(a int) bool { return a > 0 }
	pred2 := func(a int) bool { return a%2 == 0 }
	pred3 := func(a int) bool { return a < 0 }

	cases := []struct {
		msg      string
		receiver set.Set[int]
		arg      func(int) bool
		want     bool
	}{
		{"All: pred matches all", sBase(), pred1, true},
		{"All: pred matches some", sBase(), pred2, false},
		{"All: pred matches none", sBase(), pred3, false},
		{"All: empty receiver", set.Set[int]{}, pred2, true},
		{"All: nil receiver", nil, pred2, true},
	}

	for _, cs := range cases {
		got := cs.receiver.All(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSet_Any(t *testing.T) {
	pred1 := func(a int) bool { return a > 0 }
	pred2 := func(a int) bool { return a%2 == 0 }
	pred3 := func(a int) bool { return a < 0 }

	cases := []struct {
		msg      string
		receiver set.Set[int]
		arg      func(int) bool
		want     bool
	}{
		{"Any: pred matches all", sBase(), pred1, true},
		{"Any: pred matches some", sBase(), pred2, true},
		{"Any: pred matches none", sBase(), pred3, false},
		{"Any: empty receiver", set.Set[int]{}, pred2, false},
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
		receiver set.Set[int]
		arg      int
		want     bool
	}{
		{"CotainsKey: present", sBase(), 22, true},
		{"Contains: absent", sBase(), 0, false},
		{"Contains: empty set", set.Set[int]{}, 22, false},
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
		receiver set.Set[int]
		arg      set.Set[int]
		want     bool
	}{
		{"ContainsSet: nonempty receiver, other subset", sBase(),
			set.Set[int]{22: true, 333: true}, true},
		{"ContainsSet: nonempty receiver, other intersects", sBase(),
			set.Set[int]{22: true, 25: true, 333: true}, false},
		{"ContainsSet: nonempty receiver, other disjoint", sBase(),
			set.Set[int]{11: true, 25: true, 33: true}, false},
		{"ContainsSet: nonempty receiver, other empty", sBase(), set.Set[int]{}, true},
		{"ContainsSet: nonempty receiver, other nil", sBase(), nil, true},
		{"ContainsSet: empty receiver, other nonempty", set.Set[int]{},
			set.Set[int]{22: true, 333: true}, false},
		{"ContainsSet: empty receiver, other empty", set.Set[int]{}, set.Set[int]{}, true},
		{"ContainsSet: empty receiver, other nil", set.Set[int]{}, nil, true},
		{"ContainsSet: nil receiver, other nonempty", nil,
			set.Set[int]{22: true, 333: true}, false},
		{"ContainsSet: nil receiver, other empty", nil, set.Set[int]{}, true},
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
		receiver set.Set[int]
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
		{"ContainsSlice: empty receiver, elems nonempty", set.Set[int]{},
			[]int{22, 333, 22}, false},
		{"ContainsSlice: empty receiver, elems empty", set.Set[int]{}, []int{}, true},
		{"ContainsSlice: empty receiver, elems nil", set.Set[int]{}, nil, true},
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
	pred1 := func(a int) bool { return a > 0 }
	pred2 := func(a int) bool { return a%2 == 0 }
	pred3 := func(a int) bool { return a < 0 }

	cases := []struct {
		msg      string
		receiver set.Set[int]
		arg      func(int) bool
		want     int
	}{
		{"Count: pred matches all", sBase(), pred1, len(sBase())},
		{"Count: pred matches some", sBase(), pred2, 2},
		{"Count: pred matches none", sBase(), pred3, 0},
		{"Count: empty receiver", set.Set[int]{}, pred2, 0},
		{"Count: nil receiver", nil, pred2, 0},
	}

	for _, cs := range cases {
		got := cs.receiver.Count(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSet_Filter(t *testing.T) {
	pred1 := func(a int) bool { return a > 0 }
	pred2 := func(a int) bool { return a%2 == 0 }
	pred3 := func(a int) bool { return a < 0 }

	cases := []struct {
		msg      string
		receiver set.Set[int]
		arg      func(int) bool
		want     set.Set[int]
	}{
		{"Filter: pred matches all", sBase(), pred1, sBase()},
		{"Filter: pred matches some", sBase(), pred2, set.Set[int]{22: true, 4444: true}},
		{"Filter: pred matches none", sBase(), pred3, set.Set[int]{}},
		{"Filter: empty receiver", set.Set[int]{}, pred2, set.Set[int]{}},
		{"Filter: nil receiver", nil, pred2, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.Filter(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSet_FilterNot(t *testing.T) {
	pred1 := func(a int) bool { return a > 0 }
	pred2 := func(a int) bool { return a%2 == 1 }
	pred3 := func(a int) bool { return a < 0 }

	cases := []struct {
		msg      string
		receiver set.Set[int]
		arg      func(int) bool
		want     set.Set[int]
	}{
		{"FilterNot: pred matches all", sBase(), pred1, set.Set[int]{}},
		{"FilterNot: pred matches some", sBase(), pred2, set.Set[int]{22: true, 4444: true}},
		{"FilterNot: pred matches none", sBase(), pred3, sBase()},
		{"FilterNot: empty receiver", set.Set[int]{}, pred2, set.Set[int]{}},
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
		receiver set.Set[int]
		want     set.Set[int]
	}{
		{"ForEach: nonempty receiver", sBase(),
			set.Set[int]{1: true, 22: true, 333: true, 4444: true}},
		{"ForEach: empty receiver", set.Set[int]{}, set.Set[int]{}},
		{"ForEach: nil receiver", nil, set.Set[int]{}},
	}

	for _, cs := range cases {
		got := set.Set[int]{}
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
		receiver set.Set[int]
		arg      set.Set[int]
		want     set.Set[int]
	}{
		{"Intersect: nonempty receiver, other subset", sBase(),
			set.Set[int]{22: true, 333: true}, set.Set[int]{22: true, 333: true}},
		{"Intersect: nonempty receiver, other intersects", sBase(),
			set.Set[int]{22: true, 25: true, 333: true}, set.Set[int]{22: true, 333: true}},
		{"Intersect: nonempty receiver, other disjoint", sBase(),
			set.Set[int]{11: true, 25: true, 33: true}, set.Set[int]{}},
		{"Intersect: nonempty receiver, other empty", sBase(), set.Set[int]{}, set.Set[int]{}},
		{"Intersect: nonempty receiver, other nil", sBase(), nil, set.Set[int]{}},
		{"Intersect: empty receiver, other nonempty", set.Set[int]{},
			set.Set[int]{22: true, 333: true}, set.Set[int]{}},
		{"Intersect: empty receiver, other empty", set.Set[int]{}, set.Set[int]{}, set.Set[int]{}},
		{"Intersect: empty receiver, other nil", set.Set[int]{}, nil, set.Set[int]{}},
		{"Intersect: nil receiver, other nonempty", nil,
			set.Set[int]{22: true, 333: true}, nil},
		{"Intersect: nil receiver, other empty", nil, set.Set[int]{}, nil},
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
		receiver set.Set[int]
		want     bool
	}{
		{"IsEmpty: nonempty", sBase(), false},
		{"IsEmpty: empty", set.Set[int]{}, true},
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
		receiver set.Set[int]
		want     bool
	}{
		{"IsNotEmpty: nonempty", sBase(), true},
		{"IsNotEmpty: empty", set.Set[int]{}, false},
		{"IsNotEmpty: nil", nil, false},
	}

	for _, cs := range cases {
		got := cs.receiver.IsNotEmpty()
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSet_MaxWith(t *testing.T) {
	comp := func(a1 int, a2 int) int { return a1 - a2 }
	var zero int

	cases := []struct {
		msg      string
		receiver set.Set[int]
		arg      func(int, int) int
		want     int
		werr     error
	}{
		{"MaxWith: nonempty receiver", sBase(), comp, 4444, nil},
		{"MaxWith: empty receiver", set.Set[int]{}, comp, zero, errors.New("empty or nil set")},
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
		receiver set.Set[int]
		arg      int
		want     set.Set[int]
	}{
		{"MinusElement: present", sBase(), 22, set.Set[int]{1: true, 333: true, 4444: true}},
		{"MinusElement: absent", sBase(), 9, sBase()},
		{"MinusElement: empty set", set.Set[int]{}, 22, set.Set[int]{}},
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
		receiver set.Set[int]
		arg      set.Set[int]
		want     set.Set[int]
	}{
		{"MinusSet: subset", sBase(), set.Set[int]{22: true, 333: true}, set.Set[int]{1: true, 4444: true}},
		{"MinusSet: intersects", sBase(), set.Set[int]{0: true, 22: true, 9: true, 333: true},
			set.Set[int]{1: true, 4444: true}},
		{"MinusSet: disjoint", sBase(), set.Set[int]{0: true, 9: true, 42: true}, sBase()},
		{"MinusSet: empty slice", set.Set[int]{}, set.Set[int]{22: true, 333: true}, set.Set[int]{}},
		{"MinusSet: nil slice", nil, set.Set[int]{22: true, 333: true}, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.MinusSet(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSet_MinusSlice(t *testing.T) {
	cases := []struct {
		msg      string
		receiver set.Set[int]
		arg      []int
		want     set.Set[int]
	}{
		{"MinusSlice: subset", sBase(), []int{22, 333}, set.Set[int]{1: true, 4444: true}},
		{"MinusSlice: intersects", sBase(), []int{0, 22, 9, 333}, set.Set[int]{1: true, 4444: true}},
		{"MinusSlice: disjoint", sBase(), []int{0, 9, 42}, sBase()},
		{"MinusSlice: empty slice", set.Set[int]{}, []int{22, 333}, set.Set[int]{}},
		{"MinusSlice: nil slice", nil, []int{22, 333}, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.MinusSlice(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSet_MinWith(t *testing.T) {
	comp := func(a1 int, a2 int) int { return -(a1 - a2) }
	var zero int

	cases := []struct {
		msg      string
		receiver set.Set[int]
		arg      func(int, int) int
		want     int
		werr     error
	}{
		{"MinWith: nonempty receiver", sBase(), comp, 4444, nil},
		{"MinWith: empty receiver", set.Set[int]{}, comp, zero, errors.New("empty or nil set")},
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
	pred1 := func(a int) bool { return a > 0 }
	pred2 := func(a int) bool { return a%2 == 0 }
	pred3 := func(a int) bool { return a < 0 }

	cases := []struct {
		msg      string
		receiver set.Set[int]
		arg      func(int) bool
		want1    set.Set[int]
		want2    set.Set[int]
	}{
		{"Partition: match all", sBase(), pred1, sBase(), set.Set[int]{}},
		{"Partition: match some", sBase(), pred2,
			set.Set[int]{22: true, 4444: true}, set.Set[int]{1: true, 333: true}},
		{"Partition: match none", sBase(), pred3, set.Set[int]{}, sBase()},
		{"Partition: empty", set.Set[int]{}, pred1, set.Set[int]{}, set.Set[int]{}},
		{"Partition: nil", nil, pred1, set.Set[int]{}, set.Set[int]{}},
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
		receiver set.Set[int]
		arg      int
		want     set.Set[int]
	}{
		{"PlusElement: nonempty, absent", set.Set[int]{1: true, 22: true, 4444: true}, 333, sBase()},
		{"PlusElement: nonempty, present", sBase(), 333, sBase()},
		{"PlusElement: empty", set.Set[int]{}, 333, set.Set[int]{333: true}},
		{"PlusElement: nil", nil, 333, set.Set[int]{333: true}},
	}

	for _, cs := range cases {
		got := cs.receiver.PlusElement(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSet_PlusSet(t *testing.T) {
	cases := []struct {
		msg      string
		receiver set.Set[int]
		arg      set.Set[int]
		want     set.Set[int]
	}{
		{"PlusSet: nonempty + nonempty", sBase(), set.Set[int]{9: true, 333: true},
			set.Set[int]{1: true, 9: true, 22: true, 333: true, 4444: true}},
		{"PlusSet: nonempty + empty", sBase(), set.Set[int]{}, sBase()},
		{"PlusSet: nonempty + nil", sBase(), nil, sBase()},
		{"PlusSet: empty + nonempty", set.Set[int]{}, sBase(), sBase()},
		{"PlusSet: nil + nonempty", nil, sBase(), sBase()},
		{"PlusSet: empty + empty", set.Set[int]{}, set.Set[int]{}, set.Set[int]{}},
		{"PlusSet: empty + nil", set.Set[int]{}, nil, set.Set[int]{}},
		{"PlusSet: nil + empty", nil, set.Set[int]{}, set.Set[int]{}},
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
		receiver set.Set[int]
		arg      []int
		want     set.Set[int]
	}{
		{"PlusSlice: nonempty + nonempty", sBase(), []int{9, 333},
			set.Set[int]{1: true, 9: true, 22: true, 333: true, 4444: true}},
		{"PlusSlice: nonempty + empty", sBase(), []int{}, sBase()},
		{"PlusSlice: nonempty + nil", sBase(), nil, sBase()},
		{"PlusSlice: empty + nonempty", set.Set[int]{}, sliceBase(), sBase()},
		{"PlusSlice: nil + nonempty", nil, sliceBase(), sBase()},
		{"PlusSlice: empty + empty", set.Set[int]{}, []int{}, set.Set[int]{}},
		{"PlusSlice: empty + nil", set.Set[int]{}, nil, set.Set[int]{}},
		{"PlusSlice: nil + empty", nil, []int{}, set.Set[int]{}},
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
		receiver set.Set[int]
		want     []int
	}{
		{"ToSlice: nonempty", sBase(), sliceBase()},
		{"ToSlice: empty", set.Set[int]{}, []int{}},
		{"ToSlice: nil", nil, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.ToSlice()
		assert.Equal(t, set.Set[int]{}.PlusSlice(cs.want), set.Set[int]{}.PlusSlice(got), cs.msg)
	}
}

func TestSet_Put(t *testing.T) {
	cases := []struct {
		msg      string
		receiver set.Set[int]
		arg      int
		want     set.Set[int]
	}{
		{"Put: nonempty absent", set.Set[int]{1: true, 22: true, 4444: true}, 333, sBase()},
		{"Put: nonempty present", sBase(), 333, sBase()},
		{"Put: empty", set.Set[int]{}, 333, set.Set[int]{333: true}},
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

func TestSet_FlatMap(t *testing.T) {
	f := func(a int) map[string]bool {
		n := a
		s := make(map[string]bool, n%10)
		for i := 0; i < n%10; i++ {
			s[strconv.Itoa(n+i)] = true
		}
		return s
	}

	cases := []struct {
		msg      string
		receiver set.Set[int]
		arg      func(int) map[string]bool
		want     map[string]bool
	}{
		{"FlatMap: nonempty receiver", sBase(), f, map[string]bool{"1": true, "22": true, "23": true,
			"333": true, "334": true, "335": true, "4444": true, "4445": true, "4446": true,
			"4447": true}},
		{"FlatMap: empty receiver", set.Set[int]{}, f, map[string]bool{}},
		{"FlatMap: nil receiver", nil, f, nil},
	}

	for _, cs := range cases {
		got := set.SetFlatMap(cs.receiver, cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSet_GroupBy(t *testing.T) {
	f := func(a int) string { return strconv.Itoa(a % 2) }

	cases := []struct {
		msg      string
		receiver set.Set[int]
		arg      func(int) string
		want     map[string]set.Set[int]
	}{
		{"GroupBy: nonempty receiver", sBase(), f, map[string]set.Set[int]{
			"0": {22: true, 4444: true},
			"1": {1: true, 333: true},
		}},
		{"GroupBy: empty receiver", set.Set[int]{}, f, map[string]set.Set[int]{}},
		{"GroupBy: nil receiver", nil, f, nil},
	}

	for _, cs := range cases {
		got := set.SetGroupBy(cs.receiver, cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSet_Map(t *testing.T) {
	f := func(a int) string { return strconv.Itoa(a + 1) }

	cases := []struct {
		msg      string
		receiver set.Set[int]
		arg      func(int) string
		want     map[string]bool
	}{
		{"Map: nonempty receiver", sBase(), f, map[string]bool{"2": true, "23": true, "334": true,
			"4445": true}},
		{"Map: empty receiver", set.Set[int]{}, f, map[string]bool{}},
		{"Map: nil receiver", nil, f, nil},
	}

	for _, cs := range cases {
		got := set.SetMap(cs.receiver, cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSet_ToMap(t *testing.T) {
	data := set.Set[pair.Pair[int, string]]{{22, "42"}: true, {1, "9"}: true}

	cases := []struct {
		msg      string
		receiver Set[Pair[int, string]]
		want     m.Map[int, string]
	}{
		{"ToMap: nonempty receiver", data, m.Map[int, string]{1: "9", 22: "42"}},
		{"ToMap: empty receiver", Set[Pair[int, string]]{}, m.Map[int, string]{}},
		{"ToMap: nil receiver", nil, nil},
	}

	for _, cs := range cases {
		got := set.SetToMap(cs.receiver)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}
