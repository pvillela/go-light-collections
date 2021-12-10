// Code generated -- DO NOT EDIT.

/*
 * Copyright © 2021 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package collections

// GroupByint returns a map whose keys are outputs of the keySelector function applied to
// the items in the receiver and whose values are slices containing the items in the
// receiver that correspond to each key obtained with the keySelector function.
func (s SliceDat) GroupByint(keySelector func(Dat) int) map[int]SliceDat {
	if s == nil {
		return nil
	}
	m := make(map[int]SliceDat, len(s)/2) // optimizing for speed vs space
	for _, x := range s {
		k := keySelector(x)
		lst, ok := m[k]
		if !ok {
			lst = make(SliceDat, 0, 1)
		}
		lst = append(lst, x)
		m[k] = lst
	}
	return m
}
