// Code generated -- DO NOT EDIT.

package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
