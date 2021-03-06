/*
 * Copyright © 2021 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package glc

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet_FlatMapT1(t *testing.T) {
	f := func(a T0) map[T1]bool {
		n := toInt(a)
		s := make(map[T1]bool, n%10)
		for i := 0; i < n%10; i++ {
			s[strconv.Itoa(n+i)] = true
		}
		return s
	}

	cases := []struct {
		msg      string
		receiver SetT0
		arg      func(T0) map[T1]bool
		want     map[T1]bool
	}{
		{"FlatMapT1: nonempty receiver", sBase(), f, map[T1]bool{"1": true, "22": true, "23": true,
			"333": true, "334": true, "335": true, "4444": true, "4445": true, "4446": true,
			"4447": true}},
		{"FlatMapT1: empty receiver", SetT0{}, f, map[T1]bool{}},
		{"FlatMapT1: nil receiver", nil, f, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.FlatMapT1(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSet_GroupByT1(t *testing.T) {
	f := func(a T0) T1 { return strconv.Itoa(toInt(a) % 2) }

	cases := []struct {
		msg      string
		receiver SetT0
		arg      func(T0) T1
		want     map[T1]SetT0
	}{
		{"GroupByT1: nonempty receiver", sBase(), f, map[T1]SetT0{
			"0": {22: true, 4444: true},
			"1": {1: true, 333: true},
		}},
		{"GroupByT1: empty receiver", SetT0{}, f, map[T1]SetT0{}},
		{"GroupByT1: nil receiver", nil, f, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.GroupByT1(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSet_MapT1(t *testing.T) {
	f := func(a T0) T1 { return strconv.Itoa(toInt(a) + 1) }

	cases := []struct {
		msg      string
		receiver SetT0
		arg      func(T0) T1
		want     map[T1]bool
	}{
		{"MapT1: nonempty receiver", sBase(), f, map[T1]bool{"2": true, "23": true, "334": true,
			"4445": true}},
		{"MapT1: empty receiver", SetT0{}, f, map[T1]bool{}},
		{"MapT1: nil receiver", nil, f, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.MapT1(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}
