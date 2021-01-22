// Code generated -- DO NOT EDIT.

package collections

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func itoa(i int) string { return strconv.Itoa(i) }

func TestSet_FlatMapstring(t *testing.T) {
	f := func(a int) Setstring {
		n := toInt(a)
		s := make(Setstring, n%10)
		for i := 0; i < n%10; i++ {
			s[itoa(n+i)] = true
		}
		return s
	}

	cases := []struct {
		msg      string
		receiver Setint
		arg      func(int) Setstring
		want     Setstring
	}{
		{"FlatMapstring: nonempty receiver", sBase(), f, Setstring{"1": true, "22": true, "23": true,
			"333": true, "334": true, "335": true, "4444": true, "4445": true, "4446": true,
			"4447": true}},
		{"FlatMapstring: empty receiver", Setint{}, f, Setstring{}},
		{"FlatMapstring: nil receiver", nil, f, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.FlatMapstring(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSet_GroupBystring(t *testing.T) {
	f := func(a int) string { return itoa(toInt(a) % 2) }

	cases := []struct {
		msg      string
		receiver Setint
		arg      func(int) string
		want     MapstringSetint
	}{
		{"GroupBystring: nonempty receiver", sBase(), f, MapstringSetint{
			"0": {22: true, 4444: true},
			"1": {1: true, 333: true},
		}},
		{"GroupBystring: empty receiver", Setint{}, f, MapstringSetint{}},
		{"GroupBystring: nil receiver", nil, f, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.GroupBystring(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSet_Mapstring(t *testing.T) {
	f := func(a int) string { return itoa(toInt(a) + 1) }

	cases := []struct {
		msg      string
		receiver Setint
		arg      func(int) string
		want     Setstring
	}{
		{"Mapstring: nonempty receiver", sBase(), f, Setstring{"2": true, "23": true, "334": true,
			"4445": true}},
		{"Mapstring: empty receiver", Setint{}, f, Setstring{}},
		{"Mapstring: nil receiver", nil, f, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.Mapstring(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestSet_ToMap(t *testing.T) {
	data := SetOfPairintstring{{22, "42"}: true, {1, "9"}: true}

	cases := []struct {
		msg      string
		receiver SetOfPairintstring
		want     Mapintstring
	}{
		{"ToMap: nonempty receiver", data, Mapintstring{1: "9", 22: "42"}},
		{"ToMap: empty receiver", SetOfPairintstring{}, Mapintstring{}},
		{"ToMap: nil receiver", nil, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.ToMap()
		assert.Equal(t, cs.want, got, cs.msg)
	}
}
