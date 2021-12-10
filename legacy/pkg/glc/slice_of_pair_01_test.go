/*
 * Copyright © 2021 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package glc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlice_ToMap(t *testing.T) {
	data := SliceOfPairT0T1{{Dat{1, "w1"}, 10}, {Dat{22, "w22"}, 42}, {Dat{1, "w1"}, 9}}

	cases := []struct {
		msg      string
		receiver SliceOfPairT0T1
		want     map[T0]T1
	}{
		{"ToMap: nonempty receiver", data, map[T0]T1{Dat{1, "w1"}: 9, Dat{22, "w22"}: 42}},
		{"ToMap: empty receiver", SliceOfPairT0T1{}, map[T0]T1{}},
		{"ToMap: nil receiver", nil, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.ToMap()
		assert.Equal(t, cs.want, got, cs.msg)
	}
}
