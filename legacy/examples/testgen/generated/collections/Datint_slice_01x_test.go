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

func TestSlice_GroupByint(t *testing.T) {
	f := func(a Dat) int { return toDat(a).V1 % 2 }

	cases := []struct {
		msg      string
		receiver SliceDat
		arg      func(Dat) int
		want     map[int]SliceDat
	}{
		{"GroupByint: nonempty receiver", sDat(), f, map[int]SliceDat{
			0: {Dat{22, "w22"}, Dat{4444, "w4444"}, Dat{22, "w22"}},
			1: {Dat{1, "w1"}, Dat{333, "w333"}},
		}},
		{"GroupByint: empty receiver", SliceDat{}, f, map[int]SliceDat{}},
		{"GroupByint: nil receiver", nil, f, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.GroupByint(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}
