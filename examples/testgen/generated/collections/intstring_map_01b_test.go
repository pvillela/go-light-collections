// Code generated -- DO NOT EDIT.

package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap_Entries(t *testing.T) {
	cases := []struct {
		msg      string
		receiver Mapintstring
		want     map[PairMpintstring]bool
	}{
		{"Entries: nonempty map", mBase(), map[PairMpintstring]bool{
			{1, "w1"}: true, {22, "w22"}: true, {333, "w333"}: true,
			{4444, "w4444"}: true}},
		{"Entries: empty map", Mapintstring{}, map[PairMpintstring]bool{}},
		{"Entries: nil map", nil, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.Entries()
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_Values(t *testing.T) {
	cases := []struct {
		msg      string
		receiver Mapintstring
		want     map[string]bool
	}{
		{"Values: nonempty map", mBase(), map[string]bool{"w1": true, "w22": true, "w333": true,
			"w4444": true}},
		{"Values: empty map", Mapintstring{}, map[string]bool{}},
		{"Values: nil map", nil, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.Values()
		assert.Equal(t, cs.want, got, cs.msg)
	}
}
