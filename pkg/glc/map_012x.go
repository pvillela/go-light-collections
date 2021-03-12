/*
 * Copyright © 2021 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package glc

func (m MapT0T1) MapKeysT2(f func(PairMpT0T1) T2) map[T2]T1 {
	if m == nil {
		return nil
	}
	r := make(map[T2]T1)
	for k, v := range m {
		r[f(PairMpT0T1{k, v})] = v
	}
	return r
}
