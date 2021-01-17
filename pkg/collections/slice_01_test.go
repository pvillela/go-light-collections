package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlatMapT1(t *testing.T) {
	var sInt SliceInt = []int{1, 2, 3}
	f := func(a T0) SliceT1 {
		n := a.(int)
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
		{"FlatMapT1: non-empty receiver", sInt.ToSliceAny(), f, SliceT1{1, 2, 2, 3, 3, 3}},
		{"FlatMapT1: empty receiver", SliceT0{}, f, SliceT1{}},
	}

	for _, cs := range cases {
		got := cs.receiver.FlatMapT1(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestFoldT1(t *testing.T) {
	op := func(z T1, a T0) T1 { return z.(int) + a.(Foo).V1 }

	cases := []struct {
		msg      string
		receiver SliceT0
		arg1     int
		arg2     func(z T1, a T0) T1
		want     T1
	}{
		{"FoldT1: non-empty receiver", sFoo(), 1, op, 1 + 1 + 22 + 333 + 4444 + 22},
		{"FoldT1: empty receiver", SliceT0{}, 42, op, 42},
	}

	for _, cs := range cases {
		got := cs.receiver.FoldT1(cs.arg1, cs.arg2)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestGroupByT1(t *testing.T) {
	f := func(a T0) T1 { return a.(Foo).V1 % 2 }

	cases := []struct {
		msg      string
		receiver SliceT0
		arg      func(T0) T1
		want     MapT1SliceT0
	}{
		{"GroupByT1: non-empty receiver", sFoo(), f, MapT1SliceT0{
			0: {Foo{22, "w22"}, Foo{4444, "w4444"}, Foo{22, "w22"}},
			1: {Foo{1, "w1"}, Foo{333, "w333"}},
		}},
		{"GroupByT1: empty receiver", SliceT0{}, f, MapT1SliceT0{}},
	}

	for _, cs := range cases {
		got := cs.receiver.GroupByT1(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMapT1(t *testing.T) {
	f := func(a T0) T1 { return Bar{a.(Foo).V1 + 1, []string{a.(Foo).V2}} }

	cases := []struct {
		msg      string
		receiver SliceT0
		arg      func(T0) T1
		want     SliceT1
	}{
		{"MapT1: non-empty receiver", sFoo(), f, SliceT1{Bar{2, []string{"w1"}}, Bar{23, []string{"w22"}}, Bar{334, []string{"w333"}}, Bar{4445, []string{"w4444"}}, Bar{23, []string{"w22"}}}},
		{"MapT1: empty receiver", SliceT0{}, f, SliceT1{}},
	}

	for _, cs := range cases {
		got := cs.receiver.MapT1(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestZipT1(t *testing.T) {
	shorterOther := SliceT1{1, 2, 3}
	longerOther := SliceT1{1, 2, 3, 4, 5, 6, 7}

	cases := []struct {
		msg      string
		receiver SliceT0
		arg      SliceT1
		want     SliceOfPairT0T1
	}{
		{"ZipT1: non-empty receiver, shorter other", sFoo(), shorterOther,
			SliceOfPairT0T1{{Foo{1, "w1"}, 1}, {Foo{22, "w22"}, 2}, {Foo{333, "w333"}, 3}}},
		{"ZipT1: non-empty receiver, longer other", sFoo(), longerOther,
			SliceOfPairT0T1{{Foo{1, "w1"}, 1}, {Foo{22, "w22"}, 2}, {Foo{333, "w333"}, 3},
				{Foo{4444, "w4444"}, 4}, {Foo{22, "w22"}, 5}}},
		{"ZipT1: non-empty receiver, empty other", sFoo(), SliceT0{}, SliceOfPairT0T1{}},
		{"ZipT1: empty receiver, non-empty other", SliceT0{}, sFoo(), SliceOfPairT0T1{}},
		{"ZipT1: empty receiver, empty other", SliceT0{}, SliceT0{}, SliceOfPairT0T1{}},
	}

	for _, cs := range cases {
		got := cs.receiver.ZipT1(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}
