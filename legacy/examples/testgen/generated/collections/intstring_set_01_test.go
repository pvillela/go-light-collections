// Code generated -- DO NOT EDIT.

/*
 * Copyright © 2021 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package collections

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet_FlatMapstring(t *testing.T) {
	f := func(a int) map[string]bool {
		n := toInt(a)
		s := make(map[string]bool, n%10)
		for i := 0; i < n%10; i++ {
			s[strconv.Itoa(n+i)] = true
		}
		return s
	}

	cases := []struct {
		msg      string
		receiver Setint
		arg      func(int) map[string]bool
		want     map[string]bool
	}{
		{"FlatMapstring: nonempty receiver", sBase(), f, map[string]bool{"1": true, "22": true, "23": true,
			"333": true, "334": true, "335": true, "4444": true, "4445": true, "4446": true,
			"4447": true}},
		{"FlatMapstring: empty receiver", Setint{}, f, map[string]bool{}},
		{"FlatMapstring: nil receiver", nil, f, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.FlatMapstring(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSet_GroupBystring(t *testing.T) {
	f := func(a int) string { return strconv.Itoa(toInt(a) % 2) }

	cases := []struct {
		msg      string
		receiver Setint
		arg      func(int) string
		want     map[string]Setint
	}{
		{"GroupBystring: nonempty receiver", sBase(), f, map[string]Setint{
			"0": {22: true, 4444: true},
			"1": {1: true, 333: true},
		}},
		{"GroupBystring: empty receiver", Setint{}, f, map[string]Setint{}},
		{"GroupBystring: nil receiver", nil, f, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.GroupBystring(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSet_Mapstring(t *testing.T) {
	f := func(a int) string { return strconv.Itoa(toInt(a) + 1) }

	cases := []struct {
		msg      string
		receiver Setint
		arg      func(int) string
		want     map[string]bool
	}{
		{"Mapstring: nonempty receiver", sBase(), f, map[string]bool{"2": true, "23": true, "334": true,
			"4445": true}},
		{"Mapstring: empty receiver", Setint{}, f, map[string]bool{}},
		{"Mapstring: nil receiver", nil, f, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.Mapstring(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}
