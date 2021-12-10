// Code generated -- DO NOT EDIT.

/*
 * Copyright © 2021 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package collections

func (m Mapintstring) FlatMapint(f func(PairMpintstring) []int) []int {
	if m == nil {
		return nil
	}
	r := make([]int, 0, len(m)) // optimizing for speed vs space
	for k, v := range m {
		r = append(r, f(PairMpintstring{k, v})...)
	}
	return r
}

func (m Mapintstring) Mapint(f func(PairMpintstring) int) []int {
	if m == nil {
		return nil
	}
	r := make([]int, 0, len(m)) // optimizing for speed vs space
	for k, v := range m {
		r = append(r, f(PairMpintstring{k, v}))
	}
	return r
}

func (m Mapintstring) MapValuesint(f func(PairMpintstring) int) map[int]int {
	if m == nil {
		return nil
	}
	r := make(map[int]int)
	for k, v := range m {
		r[k] = f(PairMpintstring{k, v})
	}
	return r
}
