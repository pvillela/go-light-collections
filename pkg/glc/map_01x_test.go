package glc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap_Entries(t *testing.T) {
	cases := []struct {
		msg      string
		receiver MapT0T1
		want     map[PairMpT0T1]bool
	}{
		{"Entries: nonempty map", mBase(), map[PairMpT0T1]bool{
			{1, "w1"}: true, {22, "w22"}: true, {333, "w333"}: true,
			{4444, "w4444"}: true}},
		{"Entries: empty map", MapT0T1{}, map[PairMpT0T1]bool{}},
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
		receiver MapT0T1
		want     map[T1]bool
	}{
		{"Values: nonempty map", mBase(), map[T1]bool{"w1": true, "w22": true, "w333": true,
			"w4444": true}},
		{"Values: empty map", MapT0T1{}, map[T1]bool{}},
		{"Values: nil map", nil, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.Values()
		assert.Equal(t, cs.want, got, cs.msg)
	}
}
