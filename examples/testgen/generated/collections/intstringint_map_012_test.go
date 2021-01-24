// Code generated -- DO NOT EDIT.

package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

////
// Preliminaries

func _map012_sliceToSet(s []int) map[int]bool {
	if s == nil {
		return nil
	}
	set := make(map[int]bool, len(s))
	for _, x := range s {
		set[x] = true
	}
	return set
}

////
// Tests

func TestMap_FlatMapint(t *testing.T) {
	f := func(a PairMpintstring) Sliceint {
		n := toInt(a.X1) % 10
		s := make(Sliceint, n)
		for i := range s {
			s[i] = n
		}
		return s
	}

	toSet := _map012_sliceToSet

	cases := []struct {
		msg      string
		receiver Mapintstring
		arg      func(PairMpintstring) Sliceint
		want     Sliceint
	}{
		{"FlatMapint: nonempty receiver", mBase(), f, Sliceint{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}},
		{"FlatMapint: empty receiver", Mapintstring{}, f, Sliceint{}},
		{"FlatMapint: nil receiver", nil, f, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.FlatMapint(cs.arg)
		assert.Equal(t, toSet(cs.want), toSet(got), cs.msg)
	}
}

func TestMap_Mapint(t *testing.T) {
	f := func(a PairMpintstring) int { return toInt(a.X1) + 1 }

	toSet := _map012_sliceToSet

	cases := []struct {
		msg      string
		receiver Mapintstring
		arg      func(PairMpintstring) int
		want     Sliceint
	}{
		{"Mapint: nonempty receiver", mBase(), f, Sliceint{2, 23, 334, 4445}},
		{"Mapint: empty receiver", Mapintstring{}, f, Sliceint{}},
		{"Mapint: nil receiver", nil, f, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.Mapint(cs.arg)
		assert.Equal(t, toSet(cs.want), toSet(got), cs.msg)
	}
}

func TestMap_MapKeysint(t *testing.T) {
	f := func(a PairMpintstring) int { return toInt(a.X1) + 1 }

	cases := []struct {
		msg      string
		receiver Mapintstring
		arg      func(PairMpintstring) int
		want     map[int]string
	}{
		{"MapKeysint: nonempty receiver", mBase(), f, map[int]string{2: "w1", 23: "w22", 334: "w333",
			4445: "w4444"}},
		{"MapKeysint: empty receiver", Mapintstring{}, f, map[int]string{}},
		{"MapKeysint: nil receiver", nil, f, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.MapKeysint(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_MapValuesint(t *testing.T) {
	f := func(a PairMpintstring) int { return toInt(a.X1) + len(toString(a.X2)) }

	cases := []struct {
		msg      string
		receiver Mapintstring
		arg      func(PairMpintstring) int
		want     map[int]int
	}{
		{"MapValuesint: nonempty receiver", mBase(), f, map[int]int{1: 3, 22: 25, 333: 337,
			4444: 4449}},
		{"MapValuesint: empty receiver", Mapintstring{}, f, map[int]int{}},
		{"MapValuesint: nil receiver", nil, f, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.MapValuesint(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}
