// Code generated -- DO NOT EDIT.

/*
 * Copyright © 2021 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package collections

func (m Mapintstring) MapKeysint(f func(PairMpintstring) int) map[int]string {
	if m == nil {
		return nil
	}
	r := make(map[int]string)
	for k, v := range m {
		r[f(PairMpintstring{k, v})] = v
	}
	return r
}
