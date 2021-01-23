// Code generated -- DO NOT EDIT.

package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlice_ToSet(t *testing.T) {
	cases := []struct {
		msg      string
		receiver SliceDat
		want     map[Dat]bool
	}{
		{"ToSet: nonempty receiver", sDat(), map[Dat]bool{
			1: true, 22: true, 333: true, 4444: true}},
		{"ToSet: empty receiver", SliceDat{}, map[Dat]bool{}},
		{"ToSet: nil receiver", nil, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.ToSet()
		assert.Equal(t, cs.want, got, cs.msg)
	}
}
