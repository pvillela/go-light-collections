/*
 * Copyright © 2021 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package glc

// SliceInt is a wrapper type to enable extension methods.
type SliceInt []int

// ToSliceAny converts to
func (s SliceInt) ToSliceAny() SliceT0 {
	r := make(SliceT0, len(s))
	for i, x := range s {
		r[i] = x
	}
	return r
}

// ToSliceInt is a conversion function.
func ToSliceInt(s SliceT0) SliceInt {
	r := make(SliceInt, len(s))
	for i, x := range s {
		r[i] = x.(int)
	}
	return r
}
