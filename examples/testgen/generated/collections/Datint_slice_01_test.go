// Code generated -- DO NOT EDIT.

/*
 * Copyright © 2021 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlice_FlatMapint(t *testing.T) {
	f := func(a Dat) []int {
		n := toDat(a).V1 % 10
		s := make([]int, n)
		for i := range s {
			s[i] = n
		}
		return s
	}

	cases := []struct {
		msg      string
		receiver SliceDat
		arg      func(Dat) []int
		want     []int
	}{
		{"FlatMapint: nonempty receiver", sDat(), f, []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 2, 2}},
		{"FlatMapint: empty receiver", SliceDat{}, f, []int{}},
		{"FlatMapint: nil receiver", nil, f, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.FlatMapint(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_Foldint(t *testing.T) {
	op := func(z int, a Dat) int { return Any(z).(int) + toDat(a).V1 }

	cases := []struct {
		msg      string
		receiver SliceDat
		arg1     int
		arg2     func(z int, a Dat) int
		want     int
	}{
		{"Foldint: nonempty receiver", sDat(), 1, op, 1 + 1 + 22 + 333 + 4444 + 22},
		{"Foldint: empty receiver", SliceDat{}, 42, op, 42},
		{"Foldint: nil receiver", nil, 42, op, 42},
	}

	for _, cs := range cases {
		got := cs.receiver.Foldint(cs.arg1, cs.arg2)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_Mapint(t *testing.T) {
	f := func(a Dat) int { return toDat(a).V1 + 1 }

	cases := []struct {
		msg      string
		receiver SliceDat
		arg      func(Dat) int
		want     []int
	}{
		{"Mapint: nonempty receiver", sDat(), f, []int{2, 23, 334, 4445, 23}},
		{"Mapint: empty receiver", SliceDat{}, f, []int{}},
		{"Mapint: nil receiver", nil, f, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.Mapint(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_Zipint(t *testing.T) {
	shorterOther := []int{1, 2, 3}
	longerOther := []int{1, 2, 3, 4, 5, 6, 7}

	cases := []struct {
		msg      string
		receiver SliceDat
		arg      []int
		want     []PairSlDatint
	}{
		{"Zipint: nonempty receiver, shorter other", sDat(), shorterOther,
			[]PairSlDatint{{Dat{1, "w1"}, 1}, {Dat{22, "w22"}, 2}, {Dat{333, "w333"}, 3}}},
		{"Zipint: nonempty receiver, longer other", sDat(), longerOther,
			[]PairSlDatint{{Dat{1, "w1"}, 1}, {Dat{22, "w22"}, 2}, {Dat{333, "w333"}, 3},
				{Dat{4444, "w4444"}, 4}, {Dat{22, "w22"}, 5}}},
		{"Zipint: nonempty receiver, empty other", sDat(), []int{}, []PairSlDatint{}},
		{"Zipint: nonempty receiver, nil other", sDat(), []int{}, []PairSlDatint{}},
		{"Zipint: empty receiver, nonempty other", SliceDat{}, shorterOther, []PairSlDatint{}},
		{"Zipint: empty receiver, empty other", SliceDat{}, []int{}, []PairSlDatint{}},
		{"Zipint: empty receiver, nil other", SliceDat{}, []int{}, []PairSlDatint{}},
		{"Zipint: nil receiver, nonempty other", nil, shorterOther, nil},
		{"Zipint: nil receiver, empty other", nil, []int{}, nil},
		{"Zipint: nil receiver, nil other", nil, nil, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.Zipint(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}
