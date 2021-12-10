// Code generated -- DO NOT EDIT.

/*
 * Copyright © 2021 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package coll

func (m MapStringint) FlatMapint(f func(PairMpStringint) []int) []int {
	if m == nil {
		return nil
	}
	r := make([]int, 0, len(m)) // optimizing for speed vs space
	for k, v := range m {
		r = append(r, f(PairMpStringint{k, v})...)
	}
	return r
}

func (m MapStringint) Mapint(f func(PairMpStringint) int) []int {
	if m == nil {
		return nil
	}
	r := make([]int, 0, len(m)) // optimizing for speed vs space
	for k, v := range m {
		r = append(r, f(PairMpStringint{k, v}))
	}
	return r
}

func (m MapStringint) MapValuesint(f func(PairMpStringint) int) map[String]int {
	if m == nil {
		return nil
	}
	r := make(map[String]int)
	for k, v := range m {
		r[k] = f(PairMpStringint{k, v})
	}
	return r
}
