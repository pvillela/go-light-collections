package glc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapT1_FooBar(t *testing.T) {
	f := func(a T0) T1 { return Bar{a.(Dat).V1 + 1, []string{a.(Dat).V2}} }

	cases := []struct {
		msg      string
		receiver SliceT0
		arg      func(T0) T1
		want     []T1
	}{
		{"MapT1: nonempty receiver", sDat(), f, []T1{Bar{2, []string{"w1"}}, Bar{23, []string{"w22"}}, Bar{334, []string{"w333"}}, Bar{4445, []string{"w4444"}}, Bar{23, []string{"w22"}}}},
		{"MapT1: empty receiver", SliceT0{}, f, []T1{}},
	}

	for _, cs := range cases {
		got := cs.receiver.MapT1(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}
