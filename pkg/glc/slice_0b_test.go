package glc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlice_ToSet(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SliceT0
		want     map[T0]bool
	}{
		{"ToSet: nonempty receiver", sDat(), map[T0]bool{
			1: true, 22: true, 333: true, 4444: true}},
		{"ToSet: empty receiver", SliceT0{}, map[T0]bool{}},
		{"ToSet: nil receiver", nil, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.ToSet()
		assert.Equal(t, cs.want, got, cs.msg)
	}
}
