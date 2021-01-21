// Code generated -- DO NOT EDIT.

package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet_FlatMapint(t *testing.T) {
	f := func(a int) Setint {
		n := toInt(a)
		s := make(Setint, n%10)
		for i := 0; i < n%10; i++ {
			s[n+i] = true
		}
		return s
	}

	cases := []struct {
		msg      string
		receiver Setint
		arg      func(int) Setint
		want     Setint
	}{
		{"FlatMapint: nonempty receiver", sBase(), f, Setint{1: true, 22: true, 23: true,
			333: true, 334: true, 335: true, 4444: true, 4445: true, 4446: true, 4447: true}},
		{"FlatMapint: empty receiver", Setint{}, f, Setint{}},
		{"FlatMapint: nil receiver", nil, f, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.FlatMapint(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSet_GroupByint(t *testing.T) {
	f := func(a int) int { return toInt(a) % 2 }

	cases := []struct {
		msg      string
		receiver Setint
		arg      func(int) int
		want     MapintSetint
	}{
		{"GroupByint: nonempty receiver", sBase(), f, MapintSetint{
			0: {22: true, 4444: true},
			1: {1: true, 333: true},
		}},
		{"GroupByint: empty receiver", Setint{}, f, MapintSetint{}},
		{"GroupByint: nil receiver", nil, f, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.GroupByint(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSet_Mapint(t *testing.T) {
	f := func(a int) int { return toInt(a) + 1 }

	cases := []struct {
		msg      string
		receiver Setint
		arg      func(int) int
		want     Setint
	}{
		{"Mapint: nonempty receiver", sBase(), f, Setint{2: true, 23: true, 334: true, 4445: true}},
		{"Mapint: empty receiver", Setint{}, f, Setint{}},
		{"Mapint: nil receiver", nil, f, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.Mapint(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSet_ToMap(t *testing.T) {
	data := SetOfPairintint{{1, 10}: true, {22, 42}: true, {1, 9}: true}

	cases := []struct {
		msg      string
		receiver SetOfPairintint
		want     Mapintint
	}{
		{"ToMap: nonempty receiver", data, Mapintint{1: 9, 22: 42}},
		{"ToMap: empty receiver", SetOfPairintint{}, Mapintint{}},
		{"ToMap: nil receiver", nil, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.ToMap()
		assert.Equal(t, cs.want, got, cs.msg)
	}
}
