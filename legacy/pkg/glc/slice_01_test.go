/*
 * Copyright © 2021 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package glc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlice_FlatMapT1(t *testing.T) {
	f := func(a T0) []T1 {
		n := toDat(a).V1 % 10
		s := make([]T1, n)
		for i := range s {
			s[i] = n
		}
		return s
	}

	cases := []struct {
		msg      string
		receiver SliceT0
		arg      func(T0) []T1
		want     []T1
	}{
		{"FlatMapT1: nonempty receiver", sDat(), f, []T1{1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 2, 2}},
		{"FlatMapT1: empty receiver", SliceT0{}, f, []T1{}},
		{"FlatMapT1: nil receiver", nil, f, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.FlatMapT1(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_FoldT1(t *testing.T) {
	op := func(z T1, a T0) T1 { return Any(z).(int) + toDat(a).V1 }

	cases := []struct {
		msg      string
		receiver SliceT0
		arg1     int
		arg2     func(z T1, a T0) T1
		want     T1
	}{
		{"FoldT1: nonempty receiver", sDat(), 1, op, 1 + 1 + 22 + 333 + 4444 + 22},
		{"FoldT1: empty receiver", SliceT0{}, 42, op, 42},
		{"FoldT1: nil receiver", nil, 42, op, 42},
	}

	for _, cs := range cases {
		got := cs.receiver.FoldT1(cs.arg1, cs.arg2)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_MapT1(t *testing.T) {
	f := func(a T0) T1 { return toDat(a).V1 + 1 }

	cases := []struct {
		msg      string
		receiver SliceT0
		arg      func(T0) T1
		want     []T1
	}{
		{"MapT1: nonempty receiver", sDat(), f, []T1{2, 23, 334, 4445, 23}},
		{"MapT1: empty receiver", SliceT0{}, f, []T1{}},
		{"MapT1: nil receiver", nil, f, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.MapT1(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_ZipT1(t *testing.T) {
	shorterOther := []T1{1, 2, 3}
	longerOther := []T1{1, 2, 3, 4, 5, 6, 7}

	cases := []struct {
		msg      string
		receiver SliceT0
		arg      []T1
		want     []PairSlT0T1
	}{
		{"ZipT1: nonempty receiver, shorter other", sDat(), shorterOther,
			[]PairSlT0T1{{Dat{1, "w1"}, 1}, {Dat{22, "w22"}, 2}, {Dat{333, "w333"}, 3}}},
		{"ZipT1: nonempty receiver, longer other", sDat(), longerOther,
			[]PairSlT0T1{{Dat{1, "w1"}, 1}, {Dat{22, "w22"}, 2}, {Dat{333, "w333"}, 3},
				{Dat{4444, "w4444"}, 4}, {Dat{22, "w22"}, 5}}},
		{"ZipT1: nonempty receiver, empty other", sDat(), []T1{}, []PairSlT0T1{}},
		{"ZipT1: nonempty receiver, nil other", sDat(), []T1{}, []PairSlT0T1{}},
		{"ZipT1: empty receiver, nonempty other", SliceT0{}, shorterOther, []PairSlT0T1{}},
		{"ZipT1: empty receiver, empty other", SliceT0{}, []T1{}, []PairSlT0T1{}},
		{"ZipT1: empty receiver, nil other", SliceT0{}, []T1{}, []PairSlT0T1{}},
		{"ZipT1: nil receiver, nonempty other", nil, shorterOther, nil},
		{"ZipT1: nil receiver, empty other", nil, []T1{}, nil},
		{"ZipT1: nil receiver, nil other", nil, nil, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.ZipT1(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}
