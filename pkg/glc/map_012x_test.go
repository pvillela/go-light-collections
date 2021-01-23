package glc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

////
// Tests

func TestMap_MapKeysT2(t *testing.T) {
	f := func(a PairMpT0T1) T2 { return toInt(a.X1) + 1 }

	cases := []struct {
		msg      string
		receiver MapT0T1
		arg      func(PairMpT0T1) T2
		want     map[T2]T1
	}{
		{"MapKeysT2: nonempty receiver", mBase(), f, map[T2]T1{2: "w1", 23: "w22", 334: "w333",
			4445: "w4444"}},
		{"MapKeysT2: empty receiver", MapT0T1{}, f, map[T2]T1{}},
		{"MapKeysT2: nil receiver", nil, f, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.MapKeysT2(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}
