/*
 * Copyright © 2021 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package glc

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Hack
func xmap_01_test[T any]() {}

////
// Preliminaries

// Map used as input to functions below.
func mBase() Map[int, string] {
	return Map[int, string]{1: "w1", 22: "w22", 333: "w333", 4444: "w4444"}
}

////
// TestMaps

func TestMap_Copy(t *testing.T) {
	cases := []struct {
		msg      string
		receiver Map[int, string]
	}{
		{"Copy: nonempty map", mBase()},
		{"Copy: empty map", Map[int, string]{}},
		{"Copy: nil map", nil},
	}

	for _, cs := range cases {
		got := cs.receiver.Copy()
		assert.Equal(t, cs.receiver, got, cs.msg)
		assert.True(t, &cs.receiver != &got, cs.msg)
	}
}

func TestMap_Entries(t *testing.T) {
	cases := []struct {
		msg      string
		receiver Map[int, string]
		want     []Pair[int, string]
	}{
		{"Entries: nonempty map", mBase(), []Pair[int, string]{{1, "w1"}, {22, "w22"}, {333, "w333"},
			{4444, "w4444"}}},
		{"Entries: empty map", Map[int, string]{}, []Pair[int, string]{}},
		{"Entries: nil map", nil, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.Entries()
		assert.ElementsMatch(t, cs.want, got, cs.msg)
	}
}

func TestMap_Keys(t *testing.T) {
	cases := []struct {
		msg      string
		receiver Map[int, string]
		want     []int
	}{
		{"Keys: nonempty map", mBase(), []int{1, 22, 333, 4444}},
		{"Keys: empty map", Map[int, string]{}, []int{}},
		{"Keys: nil map", nil, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.Keys()
		assert.ElementsMatch(t, cs.want, got, cs.msg)
	}
}

func TestMap_Values(t *testing.T) {
	cases := []struct {
		msg      string
		receiver Map[int, string]
		want     []string
	}{
		{"Values: nonempty map", mBase(), []string{"w1", "w22", "w333", "w4444"}},
		{"Values: empty map", Map[int, string]{}, []string{}},
		{"Values: nil map", nil, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.Values()
		assert.ElementsMatch(t, cs.want, got, cs.msg)
	}
}

func TestMap_LengthSize(t *testing.T) {
	cases := []struct {
		msg      string
		receiver Map[int, string]
		want     int
	}{
		{"Length and Size: nonempty map", mBase(), 4},
		{"Length and Size: empty map", Map[int, string]{}, 0},
		{"Length and Size: nil map", nil, 0},
	}

	for _, cs := range cases {
		got1 := cs.receiver.Length()
		assert.Equal(t, cs.want, got1, cs.msg)
		got2 := cs.receiver.Size()
		assert.Equal(t, cs.want, got2, cs.msg)
	}
}

func TestMap_ContainsKey(t *testing.T) {
	cases := []struct {
		msg      string
		receiver Map[int, string]
		arg      int
		want     bool
	}{
		{"CotainsKey: present", mBase(), 22, true},
		{"ContainsKey: absent", mBase(), 0, false},
		{"ContainsKey: empty map", Map[int, string]{}, 22, false},
		{"ContainsKey: nil map", nil, 22, false},
	}

	for _, cs := range cases {
		got := cs.receiver.ContainsKey(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_ContainsValue(t *testing.T) {
	cases := []struct {
		msg  string
		arg0 Map[int, string]
		arg1 string
		want bool
	}{
		{"ContainsValue: present", mBase(), "w22", true},
		{"ContainsValue: absent", mBase(), "w0", false},
		{"ContainsValue: empty map", Map[int, string]{}, "w22", false},
		{"ContainsValue: nil map", nil, "w22", false},
	}

	for _, cs := range cases {
		got := MapContainsValue(cs.arg0, cs.arg1)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_Count(t *testing.T) {
	pred1 := func(a Pair[int, string]) bool { return a.X1 > 0 }
	pred2 := func(a Pair[int, string]) bool { return a.X1%2 == 0 }
	pred3 := func(a Pair[int, string]) bool { return a.X1 < 0 }

	cases := []struct {
		msg      string
		receiver Map[int, string]
		arg      func(Pair[int, string]) bool
		want     int
	}{
		{"Count: pred matches all", mBase(), pred1, len(mBase())},
		{"Count: pred matches some", mBase(), pred2, 2},
		{"Count: pred matches none", mBase(), pred3, 0},
		{"Count: empty receiver", Map[int, string]{}, pred2, 0},
		{"Count: nil receiver", nil, pred2, 0},
	}

	for _, cs := range cases {
		got := cs.receiver.Count(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_Get(t *testing.T) {
	cases := []struct {
		msg      string
		receiver Map[int, string]
		arg      int
		want     string
		wok      bool
	}{
		{"Get: key exists", mBase(), 22, mBase()[22], true},
		{"Get: key doesn't exist", mBase(), 0, mBase()[0], false},
		{"Get: empty map", Map[int, string]{}, 22, "", false},
		{"Get: nil map", nil, 2, "", false},
	}

	for _, cs := range cases {
		got, ok := cs.receiver.Get(cs.arg)
		assert.Equal(t, cs.wok, ok, cs.msg)
		if ok {
			assert.Equal(t, cs.want, got, cs.msg)
		}
	}
}

func TestMap_IsEmpty(t *testing.T) {
	cases := []struct {
		msg      string
		receiver Map[int, string]
		want     bool
	}{
		{"IsEmpty: nonempty", mBase(), false},
		{"IsEmpty: empty", Map[int, string]{}, true},
		{"IsEmpty: nil", nil, true},
	}

	for _, cs := range cases {
		got := cs.receiver.IsEmpty()
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_All(t *testing.T) {
	pred1 := func(a Pair[int, string]) bool { return a.X1 > 0 }
	pred2 := func(a Pair[int, string]) bool { return a.X1%2 == 0 }
	pred3 := func(a Pair[int, string]) bool { return a.X1 < 0 }

	cases := []struct {
		msg      string
		receiver Map[int, string]
		arg      func(Pair[int, string]) bool
		want     bool
	}{
		{"All: pred matches all", mBase(), pred1, true},
		{"All: pred matches some", mBase(), pred2, false},
		{"All: pred matches none", mBase(), pred3, false},
		{"All: empty receiver", Map[int, string]{}, pred2, true},
		{"All: nil receiver", nil, pred2, true},
	}

	for _, cs := range cases {
		got := cs.receiver.All(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_Any(t *testing.T) {
	pred1 := func(a Pair[int, string]) bool { return a.X1 > 0 }
	pred2 := func(a Pair[int, string]) bool { return a.X1%2 == 0 }
	pred3 := func(a Pair[int, string]) bool { return a.X1 < 0 }

	cases := []struct {
		msg      string
		receiver Map[int, string]
		arg      func(Pair[int, string]) bool
		want     bool
	}{
		{"Any: pred matches all", mBase(), pred1, true},
		{"Any: pred matches some", mBase(), pred2, true},
		{"Any: pred matches none", mBase(), pred3, false},
		{"Any: empty receiver", Map[int, string]{}, pred2, false},
		{"Any: nil receiver", nil, pred2, false},
	}

	for _, cs := range cases {
		got := cs.receiver.Any(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_ToSlice(t *testing.T) {
	pairsBase := []Pair[int, string]{{1, "w1"}, {22, "w22"}, {333, "w333"}, {4444, "w4444"}}

	cases := []struct {
		msg      string
		receiver Map[int, string]
		want     []Pair[int, string]
	}{
		{"ToSlice: nonempty", mBase(), pairsBase},
		{"ToSlice: empty", Map[int, string]{}, []Pair[int, string]{}},
		{"ToSlice: nil", nil, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.ToSlice()
		assert.ElementsMatch(t, cs.want, got, cs.msg)
	}
}

func TestMap_Filter(t *testing.T) {
	pred1 := func(a Pair[int, string]) bool { return a.X1 > 0 }
	pred2 := func(a Pair[int, string]) bool { return a.X1%2 == 0 }
	pred3 := func(a Pair[int, string]) bool { return a.X1 < 0 }

	cases := []struct {
		msg      string
		receiver Map[int, string]
		arg      func(Pair[int, string]) bool
		want     Map[int, string]
	}{
		{"Filter: pred matches all", mBase(), pred1, mBase()},
		{"Filter: pred matches some", mBase(), pred2, Map[int, string]{22: "w22", 4444: "w4444"}},
		{"Filter: pred matches none", mBase(), pred3, Map[int, string]{}},
		{"Filter: empty receiver", Map[int, string]{}, pred2, Map[int, string]{}},
		{"Filter: nil receiver", nil, pred2, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.Filter(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_FilterKeys(t *testing.T) {
	pred1 := func(a int) bool { return a > 0 }
	pred2 := func(a int) bool { return a%2 == 0 }
	pred3 := func(a int) bool { return a < 0 }

	cases := []struct {
		msg      string
		receiver Map[int, string]
		arg      func(int) bool
		want     Map[int, string]
	}{
		{"FilterKeys: pred matches all", mBase(), pred1, mBase()},
		{"FilterKeys: pred matches some", mBase(), pred2, Map[int, string]{22: "w22", 4444: "w4444"}},
		{"FilterKeys: pred matches none", mBase(), pred3, Map[int, string]{}},
		{"FilterKeys: empty receiver", Map[int, string]{}, pred2, Map[int, string]{}},
		{"FilterKeys: nil receiver", nil, pred2, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.FilterKeys(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_FilterNot(t *testing.T) {
	pred1 := func(a Pair[int, string]) bool { return a.X1 > 0 }
	pred2 := func(a Pair[int, string]) bool { return a.X1%2 == 1 }
	pred3 := func(a Pair[int, string]) bool { return a.X1 < 0 }

	cases := []struct {
		msg      string
		receiver Map[int, string]
		arg      func(Pair[int, string]) bool
		want     Map[int, string]
	}{
		{"FilterNot: pred matches all", mBase(), pred1, Map[int, string]{}},
		{"FilterNot: pred matches some", mBase(), pred2, Map[int, string]{22: "w22", 4444: "w4444"}},
		{"FilterNot: pred matches none", mBase(), pred3, mBase()},
		{"FilterNot: empty receiver", Map[int, string]{}, pred2, Map[int, string]{}},
		{"FilterNot: nil receiver", nil, pred2, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.FilterNot(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_FilterValues(t *testing.T) {
	pred1 := func(a string) bool { return true }
	pred2 := func(a string) bool { return len(a)%2 == 1 }
	pred3 := func(a string) bool { return false }

	cases := []struct {
		msg      string
		receiver Map[int, string]
		arg      func(string) bool
		want     Map[int, string]
	}{
		{"FilterValues: pred matches all", mBase(), pred1, mBase()},
		{"FilterValues: pred matches some", mBase(), pred2, Map[int, string]{22: "w22", 4444: "w4444"}},
		{"FilterValues: pred matches none", mBase(), pred3, Map[int, string]{}},
		{"FilterValues: empty receiver", Map[int, string]{}, pred2, Map[int, string]{}},
		{"FilterValues: nil receiver", nil, pred2, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.FilterValues(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_ForEach(t *testing.T) {
	cases := []struct {
		msg      string
		receiver Map[int, string]
		want     map[int]bool
	}{
		{"ForEach: nonempty receiver", mBase(),
			map[int]bool{1: true, 22: true, 333: true, 4444: true}},
		{"ForEach: empty receiver", Map[int, string]{}, map[int]bool{}},
		{"ForEach: nil receiver", nil, map[int]bool{}},
	}

	for _, cs := range cases {
		got := map[int]bool{}
		f := func(a Pair[int, string]) {
			got[a.X1] = true
		}

		cs.receiver.ForEach(f)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_GetOrElse(t *testing.T) {
	f := func(a int) string { return fmt.Sprintf("%v", a) }

	cases := []struct {
		msg      string
		receiver Map[int, string]
		arg      int
		want     string
	}{
		{"GetOrElse: key prsent", mBase(), 22, "w22"},
		{"GetOrElse: key absent", mBase(), 9, "9"},
		{"GetOrElse: empty", Map[int, string]{}, 22, "22"},
		{"GetOrElse: nil", nil, 22, "22"},
	}

	for _, cs := range cases {
		got := cs.receiver.GetOrElse(cs.arg, f)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_IsNotEmpty(t *testing.T) {
	cases := []struct {
		msg      string
		receiver Map[int, string]
		want     bool
	}{
		{"IsNotEmpty: nonempty", mBase(), true},
		{"IsNotEmpty: empty", Map[int, string]{}, false},
		{"IsNotEmpty: nil", nil, false},
	}

	for _, cs := range cases {
		got := cs.receiver.IsNotEmpty()
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_MaxWith(t *testing.T) {
	comp := func(a1 Pair[int, string], a2 Pair[int, string]) int { return a1.X1 - a2.X1 }

	cases := []struct {
		msg      string
		receiver Map[int, string]
		arg      func(Pair[int, string], Pair[int, string]) int
		want     Pair[int, string]
		werr     error
	}{
		{"MaxWith: nonempty receiver", mBase(), comp, Pair[int, string]{4444, "w4444"}, nil},
		{"MaxWith: empty receiver", Map[int, string]{}, comp, Pair[int, string]{}, errors.New("empty or nil map")},
		{"MaxWith: nil receiver", nil, comp, Pair[int, string]{}, errors.New("empty or nil map")},
	}

	for _, cs := range cases {
		got, err := cs.receiver.MaxWith(cs.arg)
		assert.Equal(t, cs.werr, err, cs.msg)
		if err == nil {
			assert.Equal(t, cs.want, got, cs.msg)
		}
	}
}

func TestMap_MinusKey(t *testing.T) {
	cases := []struct {
		msg      string
		receiver Map[int, string]
		arg      int
		want     Map[int, string]
	}{
		{"MinusKey: present", mBase(), 22, Map[int, string]{1: "w1", 333: "w333", 4444: "w4444"}},
		{"MinusKey: absent", mBase(), 9, mBase()},
		{"MinusKey: empty map", Map[int, string]{}, 22, Map[int, string]{}},
		{"MinusKey: nil map", nil, 22, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.MinusKey(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_MinusKeys(t *testing.T) {
	cases := []struct {
		msg      string
		receiver Map[int, string]
		arg      []int
		want     Map[int, string]
	}{
		{"MinusKeys: subset", mBase(), []int{22, 333}, Map[int, string]{1: "w1", 4444: "w4444"}},
		{"MinusKeys: intersects", mBase(), []int{0, 22, 9, 333},
			Map[int, string]{1: "w1", 4444: "w4444"}},
		{"MinusKeys: disjoint", mBase(), []int{0, 9, 42}, mBase()},
		{"MinusKeys: empty slice", Map[int, string]{}, []int{22, 333}, Map[int, string]{}},
		{"MinusKeys: nil slice", nil, []int{22, 333}, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.MinusKeys(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_MinWith(t *testing.T) {
	comp := func(a1 Pair[int, string], a2 Pair[int, string]) int { return -(a1.X1 - a2.X1) }

	cases := []struct {
		msg      string
		receiver Map[int, string]
		arg      func(Pair[int, string], Pair[int, string]) int
		want     Pair[int, string]
		werr     error
	}{
		{"MinWith: nonempty receiver", mBase(), comp, Pair[int, string]{4444, "w4444"}, nil},
		{"MinWith: empty receiver", Map[int, string]{}, comp, Pair[int, string]{}, errors.New("empty or nil map")},
		{"MinWith: nil receiver", nil, comp, Pair[int, string]{}, errors.New("empty or nil map")},
	}

	for _, cs := range cases {
		got, err := cs.receiver.MinWith(cs.arg)
		assert.Equal(t, cs.werr, err, cs.msg)
		if err == nil {
			assert.Equal(t, cs.want, got, cs.msg)
		}
	}
}

func TestMap_PlusEntry(t *testing.T) {
	cases := []struct {
		msg      string
		receiver Map[int, string]
		arg      Pair[int, string]
		want     Map[int, string]
	}{
		{"PlusEntry: nonempty", Map[int, string]{1: "w1", 22: "w22", 4444: "w4444"}, Pair[int, string]{333, "w333"},
			mBase()},
		{"PlusEntry: empty", Map[int, string]{}, Pair[int, string]{333, "w333"}, Map[int, string]{333: "w333"}},
		{"PlusEntry: nil", nil, Pair[int, string]{333, "w333"}, Map[int, string]{333: "w333"}},
	}

	for _, cs := range cases {
		got := cs.receiver.PlusEntry(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_PlusMap(t *testing.T) {
	cases := []struct {
		msg      string
		receiver Map[int, string]
		arg      Map[int, string]
		want     Map[int, string]
	}{
		{"PlusMap: nonempty + nonempty", mBase(), Map[int, string]{9: "x9", 333: "x3"},
			Map[int, string]{1: "w1", 9: "x9", 22: "w22", 333: "x3", 4444: "w4444"}},
		{"PlusMap: nonempty + empty", mBase(), Map[int, string]{}, mBase()},
		{"PlusMap: nonempty + nil", mBase(), nil, mBase()},
		{"PlusMap: empty + nonempty", Map[int, string]{}, mBase(), mBase()},
		{"PlusMap: nil + nonempty", nil, mBase(), mBase()},
		{"PlusMap: empty + empty", Map[int, string]{}, Map[int, string]{}, Map[int, string]{}},
		{"PlusMap: empty + nil", Map[int, string]{}, nil, Map[int, string]{}},
		{"PlusMap: nil + empty", nil, Map[int, string]{}, Map[int, string]{}},
		{"PlusMap: nil + nil", nil, nil, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.PlusMap(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_PlusSlice(t *testing.T) {
	cases := []struct {
		msg      string
		receiver Map[int, string]
		arg      []Pair[int, string]
		want     Map[int, string]
	}{
		{"PlusSlice: nonempty + nonempty", mBase(), []Pair[int, string]{{9, "x9"}, {333, "x3"}},
			Map[int, string]{1: "w1", 9: "x9", 22: "w22", 333: "x3", 4444: "w4444"}},
		{"PlusSlice: nonempty + empty", mBase(), []Pair[int, string]{}, mBase()},
		{"PlusSlice: empty + nonempty", Map[int, string]{}, []Pair[int, string]{{9, "x9"}, {333, "x3"}},
			Map[int, string]{9: "x9", 333: "x3"}},
		{"PlusSlice: nil + nonempty", nil, []Pair[int, string]{{9, "x9"}, {333, "x3"}},
			Map[int, string]{9: "x9", 333: "x3"}},
		{"PlusSlice: empty + empty", Map[int, string]{}, []Pair[int, string]{}, Map[int, string]{}},
		{"PlusSlice: empty + nil", Map[int, string]{}, nil, Map[int, string]{}},
		{"PlusSlice: nil + empty", nil, []Pair[int, string]{}, Map[int, string]{}},
		{"PlusSlice: nil + nil", nil, nil, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.PlusSlice(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_Add(t *testing.T) {
	cases := []struct {
		msg      string
		receiver Map[int, string]
		arg1     int
		arg2     string
		want     Map[int, string]
	}{
		{"Add: key present", mBase(), 333, "x3",
			Map[int, string]{1: "w1", 22: "w22", 333: "x3", 4444: "w4444"}},
		{"Add: key absent", mBase(), 9, "x9",
			Map[int, string]{1: "w1", 9: "x9", 22: "w22", 333: "w333", 4444: "w4444"}},
		{"Add: empty", Map[int, string]{}, 333, "w333", Map[int, string]{333: "w333"}},
		{"Add: nil", nil, 333, "w333", Map[int, string]{333: "w333"}},
	}

	for _, cs := range cases {
		got := cs.receiver.Add(cs.arg1, cs.arg2)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_FlatMap(t *testing.T) {
	f := func(a Pair[int, string]) []int {
		n := a.X1 % 10
		s := make([]int, n)
		for i := range s {
			s[i] = n
		}
		return s
	}

	cases := []struct {
		msg      string
		receiver Map[int, string]
		arg      func(Pair[int, string]) []int
		want     []int
	}{
		{"FlatMap: nonempty receiver", mBase(), f, []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}},
		{"FlatMap: empty receiver", Map[int, string]{}, f, []int{}},
		{"FlatMap: nil receiver", nil, f, nil},
	}

	for _, cs := range cases {
		got := MapFlatMap(cs.receiver, cs.arg)
		assert.ElementsMatch(t, cs.want, got, cs.msg)
	}
}

func TestMap_Map(t *testing.T) {
	f := func(a Pair[int, string]) int { return a.X1 + 1 }

	cases := []struct {
		msg      string
		receiver Map[int, string]
		arg      func(Pair[int, string]) int
		want     []int
	}{
		{"Map: nonempty receiver", mBase(), f, []int{2, 23, 334, 4445}},
		{"Map: empty receiver", Map[int, string]{}, f, []int{}},
		{"Map: nil receiver", nil, f, nil},
	}

	for _, cs := range cases {
		got := MapMap(cs.receiver, cs.arg)
		assert.ElementsMatch(t, cs.want, got, cs.msg)
	}
}

func TestMap_MapValues(t *testing.T) {
	f := func(a Pair[int, string]) int { return a.X1 + len(a.X2) }

	cases := []struct {
		msg      string
		receiver Map[int, string]
		arg      func(Pair[int, string]) int
		want     map[int]int
	}{
		{"MapValues: nonempty receiver", mBase(), f, map[int]int{1: 3, 22: 25, 333: 337,
			4444: 4449}},
		{"MapValues: empty receiver", Map[int, string]{}, f, map[int]int{}},
		{"MapValues: nil receiver", nil, f, nil},
	}

	for _, cs := range cases {
		got := MapMapValues(cs.receiver, cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_MapKeysint(t *testing.T) {
	f := func(a Pair[int, string]) int { return a.X1 + 1 }

	cases := []struct {
		msg      string
		receiver Map[int, string]
		arg      func(Pair[int, string]) int
		want     map[int]string
	}{
		{"MapKeysint: nonempty receiver", mBase(), f, map[int]string{2: "w1", 23: "w22", 334: "w333",
			4445: "w4444"}},
		{"MapKeysint: empty receiver", Map[int, string]{}, f, map[int]string{}},
		{"MapKeysint: nil receiver", nil, f, nil},
	}

	for _, cs := range cases {
		got := MapMapKeys(cs.receiver, cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}
