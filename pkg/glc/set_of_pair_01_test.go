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

func TestSet_ToMap(t *testing.T) {
	data := SetOfPairT0T1{{22, "42"}: true, {1, "9"}: true}

	cases := []struct {
		msg      string
		receiver SetOfPairT0T1
		want     MapT0T1
	}{
		{"ToMap: nonempty receiver", data, MapT0T1{1: "9", 22: "42"}},
		{"ToMap: empty receiver", SetOfPairT0T1{}, MapT0T1{}},
		{"ToMap: nil receiver", nil, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.ToMap()
		assert.Equal(t, cs.want, got, cs.msg)
	}
}
