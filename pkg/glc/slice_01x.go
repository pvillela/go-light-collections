/*
 * Copyright © 2021 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package glc

// GroupByT1 returns a map whose keys are outputs of the keySelector function applied to
// the items in the receiver and whose values are slices containing the items in the
// receiver that correspond to each key obtained with the keySelector function.
func (s SliceT0) GroupByT1(keySelector func(T0) T1) map[T1]SliceT0 {
	if s == nil {
		return nil
	}
	m := make(map[T1]SliceT0, len(s)/2) // optimizing for speed vs space
	for _, x := range s {
		k := keySelector(x)
		lst, ok := m[k]
		if !ok {
			lst = make(SliceT0, 0, 1)
		}
		lst = append(lst, x)
		m[k] = lst
	}
	return m
}
