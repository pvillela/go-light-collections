// Code generated -- DO NOT EDIT.

/*
 * Copyright © 2021 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license
 * that can be found in the LICENSE file.
 */

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
			Dat{1, "w1"}: true, Dat{22, "w22"}: true, Dat{333, "w333"}: true,
			Dat{4444, "w4444"}: true}},
		{"ToSet: empty receiver", SliceDat{}, map[Dat]bool{}},
		{"ToSet: nil receiver", nil, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.ToSet()
		assert.Equal(t, cs.want, got, cs.msg)
	}
}
