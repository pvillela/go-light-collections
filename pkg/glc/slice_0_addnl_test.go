package glc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func sBar() SliceT0 {
	return SliceT0{Bar{1, []string{"w1"}}, Bar{22, []string{"w22"}}, Bar{333, []string{"w333"}},
		Bar{4444, []string{"w4444"}}, Bar{22, []string{"w22"}}}
}

func TestContains_Bar(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SliceT0
		arg      Bar
		want     bool
	}{
		{"Cotains: present", sBar(), Bar{22, []string{"w22"}}, true},
		{"Contains: absent", sBar(), Bar{22, []string{"xyz"}}, false},
		{"Contains: empty slice", SliceT0{}, Bar{22, []string{"w22"}}, false},
	}

	for _, cs := range cases {
		got := cs.receiver.Contains(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestContainsAll_Bar(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SliceT0
		arg      SliceT0
		want     bool
	}{
		{"ContainsAll: subset", sBar(), append(sBar()[2:3], sBar()[1]), true},
		{"ContainsAll: intersects", sBar(), append(sBar()[1:2], Bar{22, []string{"xyz"}}), false},
		{"ContainsAll: disjoint", sBar(), append(sBar()[:0], Bar{22, []string{"xyz"}}, Bar{0, []string{"abc"}}), false},
		{"ContainsAll: empty slice", SliceT0{}, append(sBar()[2:3], sBar()[1]), false},
	}

	for _, cs := range cases {
		got := cs.receiver.ContainsAll(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}
func TestMinus_Bar(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SliceT0
		arg      SliceT0
		want     SliceT0
	}{
		{"MinusSlice: subset", sBar(), append(sBar()[3:4], sBar()[1]), append(sBar()[0:1], sBar()[2])},
		{"MinusSlice: intersects", sBar(), append(sBar()[1:2], Bar{22, []string{"xyz"}}), append(sBar()[0:1], sBar()[2], sBar()[3])},
		{"MinusSlice: disjoint", sBar(), append(sBar()[:0], Bar{22, []string{"xyz"}}, Bar{0, []string{"abc"}}), sBar()},
		{"MinusSlice: empty slice", SliceT0{}, append(sBar()[2:2], sBar()[1]), SliceT0{}},
	}

	for _, cs := range cases {
		got := cs.receiver.MinusSlice(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}
func TestMinusElement_Bar(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SliceT0
		arg      Bar
		want     SliceT0
	}{
		{"MinusElement: present", sBar(), Bar{22, []string{"w22"}}, append(sBar()[0:1], sBar()[2:]...)},
		{"MinusElement: absent", sBar(), Bar{22, []string{"xyz"}}, sBar()},
		{"MinusElement: empty slice", SliceT0{}, Bar{22, []string{"xyz"}}, SliceT0{}},
	}

	for _, cs := range cases {
		got := cs.receiver.MinusElement(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}
