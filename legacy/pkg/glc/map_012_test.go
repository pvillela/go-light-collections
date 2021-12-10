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

////
// Tests

func TestMap_FlatMapT2(t *testing.T) {
	f := func(a PairMpT0T1) []T2 {
		n := toInt(a.X1) % 10
		s := make([]T2, n)
		for i := range s {
			s[i] = n
		}
		return s
	}

	cases := []struct {
		msg      string
		receiver MapT0T1
		arg      func(PairMpT0T1) []T2
		want     []T2
	}{
		{"FlatMapT2: nonempty receiver", mBase(), f, []T2{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}},
		{"FlatMapT2: empty receiver", MapT0T1{}, f, []T2{}},
		{"FlatMapT2: nil receiver", nil, f, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.FlatMapT2(cs.arg)
		assert.ElementsMatch(t, cs.want, got, cs.msg)
	}
}

func TestMap_MapT2(t *testing.T) {
	f := func(a PairMpT0T1) T2 { return toInt(a.X1) + 1 }

	cases := []struct {
		msg      string
		receiver MapT0T1
		arg      func(PairMpT0T1) T2
		want     []T2
	}{
		{"MapT2: nonempty receiver", mBase(), f, []T2{2, 23, 334, 4445}},
		{"MapT2: empty receiver", MapT0T1{}, f, []T2{}},
		{"MapT2: nil receiver", nil, f, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.MapT2(cs.arg)
		assert.ElementsMatch(t, cs.want, got, cs.msg)
	}
}

func TestMap_MapValuesT2(t *testing.T) {
	f := func(a PairMpT0T1) T2 { return toInt(a.X1) + len(toString(a.X2)) }

	cases := []struct {
		msg      string
		receiver MapT0T1
		arg      func(PairMpT0T1) T2
		want     map[T0]T2
	}{
		{"MapValuesT2: nonempty receiver", mBase(), f, map[T0]T2{1: 3, 22: 25, 333: 337,
			4444: 4449}},
		{"MapValuesT2: empty receiver", MapT0T1{}, f, map[T0]T2{}},
		{"MapValuesT2: nil receiver", nil, f, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.MapValuesT2(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}
