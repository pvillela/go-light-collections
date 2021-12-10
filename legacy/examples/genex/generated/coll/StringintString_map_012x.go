// Code generated -- DO NOT EDIT.

/*
 * Copyright © 2021 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package coll

func (m MapStringint) MapKeysString(f func(PairMpStringint) String) map[String]int {
	if m == nil {
		return nil
	}
	r := make(map[String]int)
	for k, v := range m {
		r[f(PairMpStringint{k, v})] = v
	}
	return r
}
