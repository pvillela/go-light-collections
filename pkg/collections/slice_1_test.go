package collections_test

import (
	"testing"

	c "github.com/pvillela/go-light-collections/pkg/collections"
	"github.com/stretchr/testify/assert"
)

func TestFlatMapT1(t *testing.T) {
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
		{"FlatMapT1: non-empty receiver", sInt.ToSliceAny(), f, []c.Any{1, 2, 2, 3, 3, 3}},
		{"FlatMapT1: empty receiver", sEmpty(), f, []c.Any{}},
	}

	for _, cs := range cases {
		got := cs.receiver.FlatMapT1(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestFoldT1(t *testing.T) {
	op := func(z c.Any, a c.Any) c.Any { return z.(int) + a.(Foo).v1 }

	cases := []struct {
		msg      string
		receiver c.Slice
		arg1     int
		arg2     func(z c.Any, a c.Any) c.Any
		want     c.Any
	}{
		{"FoldT1: non-empty receiver", sFoo(), 1, op, 1 + 1 + 22 + 333 + 4444 + 22},
		{"FoldT1: empty receiver", sEmpty(), 42, op, 42},
	}

	for _, cs := range cases {
		got := cs.receiver.FoldT1(cs.arg1, cs.arg2)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestGroupByT1(t *testing.T) {
	f := func(a c.Any) c.Any { return a.(Foo).v1 % 2 }

	cases := []struct {
		msg      string
		receiver c.Slice
		arg      func(c.Any) c.Any
		want     map[c.Any][]c.Any
	}{
		{"GroupByT1: non-empty receiver", sFoo(), f, map[c.Any][]c.Any{
			0: {Foo{22, "w22"}, Foo{4444, "w4444"}, Foo{22, "w22"}},
			1: {Foo{1, "w1"}, Foo{333, "w333"}},
		}},
		{"GroupByT1: empty receiver", sEmpty(), f, map[c.Any][]c.Any{}},
	}

	for _, cs := range cases {
		got := cs.receiver.GroupByT1(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMapT1(t *testing.T) {
	f := func(a c.Any) c.Any { return Bar{a.(Foo).v1 + 1, []string{a.(Foo).v2}} }

	cases := []struct {
		msg      string
		receiver c.Slice
		arg      func(c.Any) c.Any
		want     []c.Any
	}{
		{"MapT1: non-empty receiver", sFoo(), f, []c.Any{Bar{2, []string{"w1"}}, Bar{23, []string{"w22"}}, Bar{334, []string{"w333"}}, Bar{4445, []string{"w4444"}}, Bar{23, []string{"w22"}}}},
		{"MapT1: empty receiver", sEmpty(), f, []c.Any{}},
	}

	for _, cs := range cases {
		got := cs.receiver.MapT1(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestZipT1(t *testing.T) {
	var shorterOther c.SliceInt = []int{1, 2, 3}
	var longerOther c.SliceInt = []int{1, 2, 3, 4, 5, 6, 7}

	cases := []struct {
		msg      string
		receiver c.Slice
		arg      c.Slice
		want     []c.Pair
	}{
		{"ZipT1: non-empty receiver, shorter other", sFoo(), shorterOther.ToSliceAny(),
			[]c.Pair{{Foo{1, "w1"}, 1}, {Foo{22, "w22"}, 2}, {Foo{333, "w333"}, 3}}},
		{"ZipT1: non-empty receiver, longer other", sFoo(), longerOther.ToSliceAny(),
			[]c.Pair{{Foo{1, "w1"}, 1}, {Foo{22, "w22"}, 2}, {Foo{333, "w333"}, 3},
				{Foo{4444, "w4444"}, 4}, {Foo{22, "w22"}, 5}}},
		{"ZipT1: non-empty receiver, empty other", sFoo(), sEmpty(), []c.Pair{}},
		{"ZipT1: empty receiver, non-empty other", sEmpty(), sFoo(), []c.Pair{}},
		{"ZipT1: empty receiver, empty other", sEmpty(), sEmpty(), []c.Pair{}},
	}

	for _, cs := range cases {
		got := cs.receiver.ZipT1(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}
