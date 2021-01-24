// Code generated -- DO NOT EDIT.

package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlice_ToMap(t *testing.T) {
	data := SliceOfPairDatint{{Dat{1, "w1"}, 10}, {Dat{22, "w22"}, 42}, {Dat{1, "w1"}, 9}}

	cases := []struct {
		msg      string
		receiver SliceOfPairDatint
		want     map[Dat]int
	}{
		{"ToMap: nonempty receiver", data, map[Dat]int{Dat{1, "w1"}: 9, Dat{22, "w22"}: 42}},
		{"ToMap: empty receiver", SliceOfPairDatint{}, map[Dat]int{}},
		{"ToMap: nil receiver", nil, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.ToMap()
		assert.Equal(t, cs.want, got, cs.msg)
	}
}
