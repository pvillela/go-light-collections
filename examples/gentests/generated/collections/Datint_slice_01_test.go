// Code generated -- DO NOT EDIT.

package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlice_FlatMapint(t *testing.T) {
	f := func(a Dat) Sliceint {
		n := toDat(a).V1 % 10
		s := make(Sliceint, n)
		for i := range s {
			s[i] = n
		}
		return s
	}

	cases := []struct {
		msg      string
		receiver SliceDat
		arg      func(Dat) Sliceint
		want     Sliceint
	}{
		{"FlatMapint: nonempty receiver", sDat(), f, Sliceint{1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 2, 2}},
		{"FlatMapint: empty receiver", SliceDat{}, f, Sliceint{}},
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

func TestSlice_GroupByint(t *testing.T) {
	f := func(a Dat) int { return toDat(a).V1 % 2 }

	cases := []struct {
		msg      string
		receiver SliceDat
		arg      func(Dat) int
		want     MapintSliceDat
	}{
		{"GroupByint: nonempty receiver", sDat(), f, MapintSliceDat{
			0: {Dat{22, "w22"}, Dat{4444, "w4444"}, Dat{22, "w22"}},
			1: {Dat{1, "w1"}, Dat{333, "w333"}},
		}},
		{"GroupByint: empty receiver", SliceDat{}, f, MapintSliceDat{}},
		{"GroupByint: nil receiver", nil, f, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.GroupByint(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_Mapint(t *testing.T) {
	f := func(a Dat) int { return toDat(a).V1 + 1 }

	cases := []struct {
		msg      string
		receiver SliceDat
		arg      func(Dat) int
		want     Sliceint
	}{
		{"Mapint: nonempty receiver", sDat(), f, Sliceint{2, 23, 334, 4445, 23}},
		{"Mapint: empty receiver", SliceDat{}, f, Sliceint{}},
		{"Mapint: nil receiver", nil, f, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.Mapint(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_Zipint(t *testing.T) {
	shorterOther := Sliceint{1, 2, 3}
	longerOther := Sliceint{1, 2, 3, 4, 5, 6, 7}

	cases := []struct {
		msg      string
		receiver SliceDat
		arg      Sliceint
		want     SliceOfPairDatint
	}{
		{"Zipint: nonempty receiver, shorter other", sDat(), shorterOther,
			SliceOfPairDatint{{Dat{1, "w1"}, 1}, {Dat{22, "w22"}, 2}, {Dat{333, "w333"}, 3}}},
		{"Zipint: nonempty receiver, longer other", sDat(), longerOther,
			SliceOfPairDatint{{Dat{1, "w1"}, 1}, {Dat{22, "w22"}, 2}, {Dat{333, "w333"}, 3},
				{Dat{4444, "w4444"}, 4}, {Dat{22, "w22"}, 5}}},
		{"Zipint: nonempty receiver, empty other", sDat(), Sliceint{}, SliceOfPairDatint{}},
		{"Zipint: nonempty receiver, nil other", sDat(), Sliceint{}, SliceOfPairDatint{}},
		{"Zipint: empty receiver, nonempty other", SliceDat{}, shorterOther, SliceOfPairDatint{}},
		{"Zipint: empty receiver, empty other", SliceDat{}, Sliceint{}, SliceOfPairDatint{}},
		{"Zipint: empty receiver, nil other", SliceDat{}, Sliceint{}, SliceOfPairDatint{}},
		{"Zipint: nil receiver, nonempty other", nil, shorterOther, nil},
		{"Zipint: nil receiver, empty other", nil, Sliceint{}, nil},
		{"Zipint: nil receiver, nil other", nil, nil, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.Zipint(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_ToMap(t *testing.T) {
	data := SliceOfPairDatint{{Dat{1, "w1"}, 10}, {Dat{22, "w22"}, 42}, {Dat{1, "w1"}, 9}}

	cases := []struct {
		msg      string
		receiver SliceOfPairDatint
		want     MapDatint
	}{
		{"ToMap: nonempty receiver", data, MapDatint{Dat{1, "w1"}: 9, Dat{22, "w22"}: 42}},
		{"ToMap: empty receiver", SliceOfPairDatint{}, MapDatint{}},
		{"ToMap: nil receiver", nil, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.ToMap()
		assert.Equal(t, cs.want, got, cs.msg)
	}
}
