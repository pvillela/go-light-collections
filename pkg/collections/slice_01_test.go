package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlice_FlatMapT1(t *testing.T) {
	f := func(a T0) SliceT1 {
		n := Any(a).(Dat).V1 % 10
		s := make(SliceT1, n)
		for i := range s {
			s[i] = n
		}
		return s
	}

	cases := []struct {
		msg      string
		receiver SliceT0
		arg      func(T0) SliceT1
		want     SliceT1
	}{
		{"FlatMapT1: non-empty receiver", sDat(), f, SliceT1{1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 2, 2}},
		{"FlatMapT1: empty receiver", SliceT0{}, f, SliceT1{}},
	}

	for _, cs := range cases {
		got := cs.receiver.FlatMapT1(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_FoldT1(t *testing.T) {
	op := func(z T1, a T0) T1 { return Any(z).(int) + Any(a).(Dat).V1 }

	cases := []struct {
		msg      string
		receiver SliceT0
		arg1     int
		arg2     func(z T1, a T0) T1
		want     T1
	}{
		{"FoldT1: non-empty receiver", sDat(), 1, op, 1 + 1 + 22 + 333 + 4444 + 22},
		{"FoldT1: empty receiver", SliceT0{}, 42, op, 42},
	}

	for _, cs := range cases {
		got := cs.receiver.FoldT1(cs.arg1, cs.arg2)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_GroupByT1(t *testing.T) {
	f := func(a T0) T1 { return Any(a).(Dat).V1 % 2 }

	cases := []struct {
		msg      string
		receiver SliceT0
		arg      func(T0) T1
		want     MapT1SliceT0
	}{
		{"GroupByT1: non-empty receiver", sDat(), f, MapT1SliceT0{
			0: {Dat{22, "w22"}, Dat{4444, "w4444"}, Dat{22, "w22"}},
			1: {Dat{1, "w1"}, Dat{333, "w333"}},
		}},
		{"GroupByT1: empty receiver", SliceT0{}, f, MapT1SliceT0{}},
	}

	for _, cs := range cases {
		got := cs.receiver.GroupByT1(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_MapT1(t *testing.T) {
	f := func(a T0) T1 { return Any(a).(Dat).V1 + 1 }

	cases := []struct {
		msg      string
		receiver SliceT0
		arg      func(T0) T1
		want     SliceT1
	}{
		{"MapT1: non-empty receiver", sDat(), f, SliceT1{2, 23, 334, 4445, 23}},
		{"MapT1: empty receiver", SliceT0{}, f, SliceT1{}},
	}

	for _, cs := range cases {
		got := cs.receiver.MapT1(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSlice_ZipT1(t *testing.T) {
	shorterOther := SliceT1{1, 2, 3}
	longerOther := SliceT1{1, 2, 3, 4, 5, 6, 7}

	cases := []struct {
		msg      string
		receiver SliceT0
		arg      SliceT1
		want     SliceOfPairT0T1
	}{
		{"ZipT1: non-empty receiver, shorter other", sDat(), shorterOther,
			SliceOfPairT0T1{{Dat{1, "w1"}, 1}, {Dat{22, "w22"}, 2}, {Dat{333, "w333"}, 3}}},
		{"ZipT1: non-empty receiver, longer other", sDat(), longerOther,
			SliceOfPairT0T1{{Dat{1, "w1"}, 1}, {Dat{22, "w22"}, 2}, {Dat{333, "w333"}, 3},
				{Dat{4444, "w4444"}, 4}, {Dat{22, "w22"}, 5}}},
		{"ZipT1: non-empty receiver, empty other", sDat(), SliceT1{}, SliceOfPairT0T1{}},
		{"ZipT1: empty receiver, non-empty other", SliceT0{}, shorterOther, SliceOfPairT0T1{}},
		{"ZipT1: empty receiver, empty other", SliceT0{}, SliceT1{}, SliceOfPairT0T1{}},
	}

	for _, cs := range cases {
		got := cs.receiver.ZipT1(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}
