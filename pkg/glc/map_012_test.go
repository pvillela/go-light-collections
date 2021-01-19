package glc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

////
// Preliminaries

func _map012_sliceToSet(s SliceT2) SetT2 {
	if s == nil {
		return nil
	}
	set := make(SetT2, len(s))
	for _, x := range s {
		set[x] = true
	}
	return set
}

////
// Tests

func TestMap_FlatMapT2(t *testing.T) {
	f := func(a PairT0T1) SliceT2 {
		n := toInt(a.X1) % 10
		s := make(SliceT2, n)
		for i := range s {
			s[i] = n
		}
		return s
	}

	toSet := _map012_sliceToSet

	cases := []struct {
		msg      string
		receiver MapT0T1
		arg      func(PairT0T1) SliceT2
		want     SliceT2
	}{
		{"FlatMapT2: non-empty receiver", mBase(), f, SliceT2{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}},
		{"FlatMapT2: empty receiver", MapT0T1{}, f, SliceT2{}},
		{"FlatMapT2: nil receiver", nil, f, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.FlatMapT2(cs.arg)
		assert.Equal(t, toSet(cs.want), toSet(got), cs.msg)
	}
}

func TestMap_MapT2(t *testing.T) {
	f := func(a PairT0T1) T2 { return toInt(a.X1) + 1 }

	toSet := _map012_sliceToSet

	cases := []struct {
		msg      string
		receiver MapT0T1
		arg      func(PairT0T1) T2
		want     SliceT2
	}{
		{"MapT2: non-empty receiver", mBase(), f, SliceT2{2, 23, 334, 4445}},
		{"MapT2: empty receiver", MapT0T1{}, f, SliceT2{}},
		{"MapT2: nil receiver", nil, f, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.MapT2(cs.arg)
		assert.Equal(t, toSet(cs.want), toSet(got), cs.msg)
	}
}

func TestMap_MapKeysT2(t *testing.T) {
	f := func(a PairT0T1) T2 { return toInt(a.X1) + 1 }

	cases := []struct {
		msg      string
		receiver MapT0T1
		arg      func(PairT0T1) T2
		want     MapT2T1
	}{
		{"MapKeysT2: non-empty receiver", mBase(), f, MapT2T1{2: "w1", 23: "w22", 334: "w333",
			4445: "w4444"}},
		{"MapKeysT2: empty receiver", MapT0T1{}, f, MapT2T1{}},
		{"MapKeysT2: nil receiver", nil, f, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.MapKeysT2(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_MapValuesT2(t *testing.T) {
	f := func(a PairT0T1) T2 { return toInt(a.X1) + len(toString(a.X2)) }

	cases := []struct {
		msg      string
		receiver MapT0T1
		arg      func(PairT0T1) T2
		want     MapT0T2
	}{
		{"MapValuesT2: non-empty receiver", mBase(), f, MapT0T2{1: 3, 22: 25, 333: 337,
			4444: 4449}},
		{"MapValuesT2: empty receiver", MapT0T1{}, f, MapT0T2{}},
		{"MapValuesT2: nil receiver", nil, f, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.MapValuesT2(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}
