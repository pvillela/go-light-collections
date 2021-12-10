// Code generated -- DO NOT EDIT.

/*
 * Copyright © 2021 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package collections

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

////
// Preliminaries

// Map used as input to functions below.
func mBase() Mapintstring {
	return Mapintstring{1: "w1", 22: "w22", 333: "w333", 4444: "w4444"}
}

////
// TestMaps

func TestMap_Copy(t *testing.T) {
	cases := []struct {
		msg      string
		receiver Mapintstring
	}{
		{"Copy: nonempty map", mBase()},
		{"Copy: empty map", Mapintstring{}},
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
		receiver Mapintstring
		want     []PairMpintstring
	}{
		{"Entries: nonempty map", mBase(), []PairMpintstring{{1, "w1"}, {22, "w22"}, {333, "w333"},
			{4444, "w4444"}}},
		{"Entries: empty map", Mapintstring{}, []PairMpintstring{}},
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
		receiver Mapintstring
		want     []int
	}{
		{"Keys: nonempty map", mBase(), []int{1, 22, 333, 4444}},
		{"Keys: empty map", Mapintstring{}, []int{}},
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
		receiver Mapintstring
		want     []string
	}{
		{"Values: nonempty map", mBase(), []string{"w1", "w22", "w333", "w4444"}},
		{"Values: empty map", Mapintstring{}, []string{}},
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
		receiver Mapintstring
		want     int
	}{
		{"Length and Size: nonempty map", mBase(), 4},
		{"Length and Size: empty map", Mapintstring{}, 0},
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
		receiver Mapintstring
		arg      int
		want     bool
	}{
		{"CotainsKey: present", mBase(), 22, true},
		{"ContainsKey: absent", mBase(), 0, false},
		{"ContainsKey: empty map", Mapintstring{}, 22, false},
		{"ContainsKey: nil map", nil, 22, false},
	}

	for _, cs := range cases {
		got := cs.receiver.ContainsKey(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_ContainsValue(t *testing.T) {
	cases := []struct {
		msg      string
		receiver Mapintstring
		arg      string
		want     bool
	}{
		{"ContainsValue: present", mBase(), "w22", true},
		{"ContainsValue: absent", mBase(), "w0", false},
		{"ContainsValue: empty map", Mapintstring{}, "w22", false},
		{"ContainsValue: nil map", nil, "w22", false},
	}

	for _, cs := range cases {
		got := cs.receiver.ContainsValue(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_Count(t *testing.T) {
	pred1 := func(a PairMpintstring) bool { return toInt(a.X1) > 0 }
	pred2 := func(a PairMpintstring) bool { return toInt(a.X1)%2 == 0 }
	pred3 := func(a PairMpintstring) bool { return toInt(a.X1) < 0 }

	cases := []struct {
		msg      string
		receiver Mapintstring
		arg      func(PairMpintstring) bool
		want     int
	}{
		{"Count: pred matches all", mBase(), pred1, len(mBase())},
		{"Count: pred matches some", mBase(), pred2, 2},
		{"Count: pred matches none", mBase(), pred3, 0},
		{"Count: empty receiver", Mapintstring{}, pred2, 0},
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
		receiver Mapintstring
		arg      int
		want     string
		wok      bool
	}{
		{"Get: key exists", mBase(), 22, mBase()[22], true},
		{"Get: key doesn't exist", mBase(), 0, mBase()[0], false},
		{"Get: empty map", Mapintstring{}, 22, "", false},
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
		receiver Mapintstring
		want     bool
	}{
		{"IsEmpty: nonempty", mBase(), false},
		{"IsEmpty: empty", Mapintstring{}, true},
		{"IsEmpty: nil", nil, true},
	}

	for _, cs := range cases {
		got := cs.receiver.IsEmpty()
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_All(t *testing.T) {
	pred1 := func(a PairMpintstring) bool { return toInt(a.X1) > 0 }
	pred2 := func(a PairMpintstring) bool { return toInt(a.X1)%2 == 0 }
	pred3 := func(a PairMpintstring) bool { return toInt(a.X1) < 0 }

	cases := []struct {
		msg      string
		receiver Mapintstring
		arg      func(PairMpintstring) bool
		want     bool
	}{
		{"All: pred matches all", mBase(), pred1, true},
		{"All: pred matches some", mBase(), pred2, false},
		{"All: pred matches none", mBase(), pred3, false},
		{"All: empty receiver", Mapintstring{}, pred2, true},
		{"All: nil receiver", nil, pred2, true},
	}

	for _, cs := range cases {
		got := cs.receiver.All(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_Any(t *testing.T) {
	pred1 := func(a PairMpintstring) bool { return toInt(a.X1) > 0 }
	pred2 := func(a PairMpintstring) bool { return toInt(a.X1)%2 == 0 }
	pred3 := func(a PairMpintstring) bool { return toInt(a.X1) < 0 }

	cases := []struct {
		msg      string
		receiver Mapintstring
		arg      func(PairMpintstring) bool
		want     bool
	}{
		{"Any: pred matches all", mBase(), pred1, true},
		{"Any: pred matches some", mBase(), pred2, true},
		{"Any: pred matches none", mBase(), pred3, false},
		{"Any: empty receiver", Mapintstring{}, pred2, false},
		{"Any: nil receiver", nil, pred2, false},
	}

	for _, cs := range cases {
		got := cs.receiver.Any(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_ToSlice(t *testing.T) {
	pairsBase := []PairMpintstring{{1, "w1"}, {22, "w22"}, {333, "w333"}, {4444, "w4444"}}

	cases := []struct {
		msg      string
		receiver Mapintstring
		want     []PairMpintstring
	}{
		{"ToSlice: nonempty", mBase(), pairsBase},
		{"ToSlice: empty", Mapintstring{}, []PairMpintstring{}},
		{"ToSlice: nil", nil, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.ToSlice()
		assert.ElementsMatch(t, cs.want, got, cs.msg)
	}
}

func TestMap_Filter(t *testing.T) {
	pred1 := func(a PairMpintstring) bool { return toInt(a.X1) > 0 }
	pred2 := func(a PairMpintstring) bool { return toInt(a.X1)%2 == 0 }
	pred3 := func(a PairMpintstring) bool { return toInt(a.X1) < 0 }

	cases := []struct {
		msg      string
		receiver Mapintstring
		arg      func(PairMpintstring) bool
		want     Mapintstring
	}{
		{"Filter: pred matches all", mBase(), pred1, mBase()},
		{"Filter: pred matches some", mBase(), pred2, Mapintstring{22: "w22", 4444: "w4444"}},
		{"Filter: pred matches none", mBase(), pred3, Mapintstring{}},
		{"Filter: empty receiver", Mapintstring{}, pred2, Mapintstring{}},
		{"Filter: nil receiver", nil, pred2, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.Filter(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_FilterKeys(t *testing.T) {
	pred1 := func(a int) bool { return toInt(a) > 0 }
	pred2 := func(a int) bool { return toInt(a)%2 == 0 }
	pred3 := func(a int) bool { return toInt(a) < 0 }

	cases := []struct {
		msg      string
		receiver Mapintstring
		arg      func(int) bool
		want     Mapintstring
	}{
		{"FilterKeys: pred matches all", mBase(), pred1, mBase()},
		{"FilterKeys: pred matches some", mBase(), pred2, Mapintstring{22: "w22", 4444: "w4444"}},
		{"FilterKeys: pred matches none", mBase(), pred3, Mapintstring{}},
		{"FilterKeys: empty receiver", Mapintstring{}, pred2, Mapintstring{}},
		{"FilterKeys: nil receiver", nil, pred2, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.FilterKeys(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_FilterNot(t *testing.T) {
	pred1 := func(a PairMpintstring) bool { return toInt(a.X1) > 0 }
	pred2 := func(a PairMpintstring) bool { return toInt(a.X1)%2 == 1 }
	pred3 := func(a PairMpintstring) bool { return toInt(a.X1) < 0 }

	cases := []struct {
		msg      string
		receiver Mapintstring
		arg      func(PairMpintstring) bool
		want     Mapintstring
	}{
		{"FilterNot: pred matches all", mBase(), pred1, Mapintstring{}},
		{"FilterNot: pred matches some", mBase(), pred2, Mapintstring{22: "w22", 4444: "w4444"}},
		{"FilterNot: pred matches none", mBase(), pred3, mBase()},
		{"FilterNot: empty receiver", Mapintstring{}, pred2, Mapintstring{}},
		{"FilterNot: nil receiver", nil, pred2, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.FilterNot(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_FilterValues(t *testing.T) {
	pred1 := func(a string) bool { return true }
	pred2 := func(a string) bool { return len(toString(a))%2 == 1 }
	pred3 := func(a string) bool { return false }

	cases := []struct {
		msg      string
		receiver Mapintstring
		arg      func(string) bool
		want     Mapintstring
	}{
		{"FilterValues: pred matches all", mBase(), pred1, mBase()},
		{"FilterValues: pred matches some", mBase(), pred2, Mapintstring{22: "w22", 4444: "w4444"}},
		{"FilterValues: pred matches none", mBase(), pred3, Mapintstring{}},
		{"FilterValues: empty receiver", Mapintstring{}, pred2, Mapintstring{}},
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
		receiver Mapintstring
		want     map[Any]bool
	}{
		{"ForEach: nonempty receiver", mBase(),
			map[Any]bool{1: true, 22: true, 333: true, 4444: true}},
		{"ForEach: empty receiver", Mapintstring{}, map[Any]bool{}},
		{"ForEach: nil receiver", nil, map[Any]bool{}},
	}

	for _, cs := range cases {
		got := map[Any]bool{}
		f := func(a PairMpintstring) {
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
		receiver Mapintstring
		arg      int
		want     string
	}{
		{"GetOrElse: key prsent", mBase(), 22, "w22"},
		{"GetOrElse: key absent", mBase(), 9, "9"},
		{"GetOrElse: empty", Mapintstring{}, 22, "22"},
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
		receiver Mapintstring
		want     bool
	}{
		{"IsNotEmpty: nonempty", mBase(), true},
		{"IsNotEmpty: empty", Mapintstring{}, false},
		{"IsNotEmpty: nil", nil, false},
	}

	for _, cs := range cases {
		got := cs.receiver.IsNotEmpty()
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_MaxWith(t *testing.T) {
	comp := func(a1 PairMpintstring, a2 PairMpintstring) int { return toInt(a1.X1) - toInt(a2.X1) }

	cases := []struct {
		msg      string
		receiver Mapintstring
		arg      func(PairMpintstring, PairMpintstring) int
		want     PairMpintstring
		werr     error
	}{
		{"MaxWith: nonempty receiver", mBase(), comp, PairMpintstring{4444, "w4444"}, nil},
		{"MaxWith: empty receiver", Mapintstring{}, comp, PairMpintstring{}, errors.New("empty or nil map")},
		{"MaxWith: nil receiver", nil, comp, PairMpintstring{}, errors.New("empty or nil map")},
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
		receiver Mapintstring
		arg      int
		want     Mapintstring
	}{
		{"MinusKey: present", mBase(), 22, Mapintstring{1: "w1", 333: "w333", 4444: "w4444"}},
		{"MinusKey: absent", mBase(), 9, mBase()},
		{"MinusKey: empty map", Mapintstring{}, 22, Mapintstring{}},
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
		receiver Mapintstring
		arg      []int
		want     Mapintstring
	}{
		{"MinusKeys: subset", mBase(), []int{22, 333}, Mapintstring{1: "w1", 4444: "w4444"}},
		{"MinusKeys: intersects", mBase(), []int{0, 22, 9, 333},
			Mapintstring{1: "w1", 4444: "w4444"}},
		{"MinusKeys: disjoint", mBase(), []int{0, 9, 42}, mBase()},
		{"MinusKeys: empty slice", Mapintstring{}, []int{22, 333}, Mapintstring{}},
		{"MinusKeys: nil slice", nil, []int{22, 333}, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.MinusKeys(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_MinWith(t *testing.T) {
	comp := func(a1 PairMpintstring, a2 PairMpintstring) int { return -(toInt(a1.X1) - toInt(a2.X1)) }

	cases := []struct {
		msg      string
		receiver Mapintstring
		arg      func(PairMpintstring, PairMpintstring) int
		want     PairMpintstring
		werr     error
	}{
		{"MinWith: nonempty receiver", mBase(), comp, PairMpintstring{4444, "w4444"}, nil},
		{"MinWith: empty receiver", Mapintstring{}, comp, PairMpintstring{}, errors.New("empty or nil map")},
		{"MinWith: nil receiver", nil, comp, PairMpintstring{}, errors.New("empty or nil map")},
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
		receiver Mapintstring
		arg      PairMpintstring
		want     Mapintstring
	}{
		{"PlusEntry: nonempty", Mapintstring{1: "w1", 22: "w22", 4444: "w4444"}, PairMpintstring{333, "w333"},
			mBase()},
		{"PlusEntry: empty", Mapintstring{}, PairMpintstring{333, "w333"}, Mapintstring{333: "w333"}},
		{"PlusEntry: nil", nil, PairMpintstring{333, "w333"}, Mapintstring{333: "w333"}},
	}

	for _, cs := range cases {
		got := cs.receiver.PlusEntry(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_PlusMap(t *testing.T) {
	cases := []struct {
		msg      string
		receiver Mapintstring
		arg      Mapintstring
		want     Mapintstring
	}{
		{"PlusMap: nonempty + nonempty", mBase(), Mapintstring{9: "x9", 333: "x3"},
			Mapintstring{1: "w1", 9: "x9", 22: "w22", 333: "x3", 4444: "w4444"}},
		{"PlusMap: nonempty + empty", mBase(), Mapintstring{}, mBase()},
		{"PlusMap: nonempty + nil", mBase(), nil, mBase()},
		{"PlusMap: empty + nonempty", Mapintstring{}, mBase(), mBase()},
		{"PlusMap: nil + nonempty", nil, mBase(), mBase()},
		{"PlusMap: empty + empty", Mapintstring{}, Mapintstring{}, Mapintstring{}},
		{"PlusMap: empty + nil", Mapintstring{}, nil, Mapintstring{}},
		{"PlusMap: nil + empty", nil, Mapintstring{}, Mapintstring{}},
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
		receiver Mapintstring
		arg      []PairMpintstring
		want     Mapintstring
	}{
		{"PlusSlice: nonempty + nonempty", mBase(), []PairMpintstring{{9, "x9"}, {333, "x3"}},
			Mapintstring{1: "w1", 9: "x9", 22: "w22", 333: "x3", 4444: "w4444"}},
		{"PlusSlice: nonempty + empty", mBase(), []PairMpintstring{}, mBase()},
		{"PlusSlice: empty + nonempty", Mapintstring{}, []PairMpintstring{{9, "x9"}, {333, "x3"}},
			Mapintstring{9: "x9", 333: "x3"}},
		{"PlusSlice: nil + nonempty", nil, []PairMpintstring{{9, "x9"}, {333, "x3"}},
			Mapintstring{9: "x9", 333: "x3"}},
		{"PlusSlice: empty + empty", Mapintstring{}, []PairMpintstring{}, Mapintstring{}},
		{"PlusSlice: empty + nil", Mapintstring{}, nil, Mapintstring{}},
		{"PlusSlice: nil + empty", nil, []PairMpintstring{}, Mapintstring{}},
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
		receiver Mapintstring
		arg1     int
		arg2     string
		want     Mapintstring
	}{
		{"Add: key present", mBase(), 333, "x3",
			Mapintstring{1: "w1", 22: "w22", 333: "x3", 4444: "w4444"}},
		{"Add: key absent", mBase(), 9, "x9",
			Mapintstring{1: "w1", 9: "x9", 22: "w22", 333: "w333", 4444: "w4444"}},
		{"Add: empty", Mapintstring{}, 333, "w333", Mapintstring{333: "w333"}},
		{"Add: nil", nil, 333, "w333", Mapintstring{333: "w333"}},
	}

	for _, cs := range cases {
		got := cs.receiver.Add(cs.arg1, cs.arg2)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}
