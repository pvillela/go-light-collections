package glc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlice_GroupByT1(t *testing.T) {
	f := func(a T0) T1 { return toDat(a).V1 % 2 }

	cases := []struct {
		msg      string
		receiver SliceT0
		arg      func(T0) T1
		want     map[T1]SliceT0
	}{
		{"GroupByT1: nonempty receiver", sDat(), f, map[T1]SliceT0{
			0: {Dat{22, "w22"}, Dat{4444, "w4444"}, Dat{22, "w22"}},
			1: {Dat{1, "w1"}, Dat{333, "w333"}},
		}},
		{"GroupByT1: empty receiver", SliceT0{}, f, map[T1]SliceT0{}},
		{"GroupByT1: nil receiver", nil, f, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.GroupByT1(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}
