package collections_test

import (
	"testing"

	c "github.com/pvillela/go-light-collections/pkg/collections"
	"github.com/stretchr/testify/assert"
)

func TestFlatMapAnyT1(t *testing.T) {
	var sInt c.SliceInt = []int{1, 2, 3}
	f := func(a c.Any) []c.Any {
		n := a.(int)
		s := make([]c.Any, n)
		for i := range s {
			s[i] = n
		}
		return s
	}

	cases := []struct {
		msg      string
		receiver c.Slice
		arg      func(c.Any) []c.Any
		want     []c.Any
	}{
		{"FlatMapAnyT1: non-empty receiver", sInt.ToSliceAny(), f, []c.Any{1, 2, 2, 3, 3, 3}},
		{"FlatMapAnyT1: empty receiver", sEmpty(), f, []c.Any{}},
	}

	for _, cs := range cases {
		got := cs.receiver.FlatMapAnyT1(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestFoldAnyT1(t *testing.T) {
	op := func(z c.Any, a c.Any) c.Any { return z.(int) + a.(Foo).v1 }

	cases := []struct {
		msg      string
		receiver c.Slice
		arg1     int
		arg2     func(z c.Any, a c.Any) c.Any
		want     c.Any
	}{
		{"FoldAnyT1: non-empty receiver", sFoo(), 1, op, 1 + 1 + 22 + 333 + 4444 + 22},
		{"FoldAnyT1: empty receiver", sEmpty(), 42, op, 42},
	}

	for _, cs := range cases {
		got := cs.receiver.FoldAnyT1(cs.arg1, cs.arg2)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestGroupByAnyT1(t *testing.T) {
	f := func(a c.Any) c.Any { return a.(Foo).v1 % 2 }

	cases := []struct {
		msg      string
		receiver c.Slice
		arg      func(c.Any) c.Any
		want     map[c.Any][]c.Any
	}{
		{"GroupByAnyT1: non-empty receiver", sFoo(), f, map[c.Any][]c.Any{
			0: {Foo{22, "w22"}, Foo{4444, "w4444"}, Foo{22, "w22"}},
			1: {Foo{1, "w1"}, Foo{333, "w333"}},
		}},
		{"GroupByAnyT1: empty receiver", sEmpty(), f, map[c.Any][]c.Any{}},
	}

	for _, cs := range cases {
		got := cs.receiver.GroupByAnyT1(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMapAnyT1(t *testing.T) {
	f := func(a c.Any) c.Any { return Bar{a.(Foo).v1 + 1, []string{a.(Foo).v2}} }

	cases := []struct {
		msg      string
		receiver c.Slice
		arg      func(c.Any) c.Any
		want     []c.Any
	}{
		{"MapAnyT1: non-empty receiver", sFoo(), f, []c.Any{Bar{2, []string{"w1"}}, Bar{23, []string{"w22"}}, Bar{334, []string{"w333"}}, Bar{4445, []string{"w4444"}}, Bar{23, []string{"w22"}}}},
		{"MapAnyT1: empty receiver", sEmpty(), f, []c.Any{}},
	}

	for _, cs := range cases {
		got := cs.receiver.MapAnyT1(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestZipAnyT1(t *testing.T) {
	var shorterOther c.SliceInt = []int{1, 2, 3}
	var longerOther c.SliceInt = []int{1, 2, 3, 4, 5, 6, 7}

	cases := []struct {
		msg      string
		receiver c.Slice
		arg      c.Slice
		want     []c.Pair
	}{
		{"ZipAnyT1: non-empty receiver, shorter other", sFoo(), shorterOther.ToSliceAny(),
			[]c.Pair{{Foo{1, "w1"}, 1}, {Foo{22, "w22"}, 2}, {Foo{333, "w333"}, 3}}},
		{"ZipAnyT1: non-empty receiver, longer other", sFoo(), longerOther.ToSliceAny(),
			[]c.Pair{{Foo{1, "w1"}, 1}, {Foo{22, "w22"}, 2}, {Foo{333, "w333"}, 3},
				{Foo{4444, "w4444"}, 4}, {Foo{22, "w22"}, 5}}},
		{"ZipAnyT1: non-empty receiver, empty other", sFoo(), sEmpty(), []c.Pair{}},
		{"ZipAnyT1: empty receiver, non-empty other", sEmpty(), sFoo(), []c.Pair{}},
		{"ZipAnyT1: empty receiver, empty other", sEmpty(), sEmpty(), []c.Pair{}},
	}

	for _, cs := range cases {
		got := cs.receiver.ZipAnyT1(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}
