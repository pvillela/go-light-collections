// Code generated -- DO NOT EDIT.

package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

////
// Tests

func TestMap_MapKeysint(t *testing.T) {
	f := func(a PairMpintstring) int { return toInt(a.X1) + 1 }

	cases := []struct {
		msg      string
		receiver Mapintstring
		arg      func(PairMpintstring) int
		want     map[int]string
	}{
		{"MapKeysint: nonempty receiver", mBase(), f, map[int]string{2: "w1", 23: "w22", 334: "w333",
			4445: "w4444"}},
		{"MapKeysint: empty receiver", Mapintstring{}, f, map[int]string{}},
		{"MapKeysint: nil receiver", nil, f, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.MapKeysint(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}
