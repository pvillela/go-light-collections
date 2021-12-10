/*
 * Copyright © 2021 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package glc

// SliceString is a wrapper type to enable extension methods.
type SliceString []string

// ToSliceAny is defined to implement ConvertibleToSliceAny.
func (s SliceString) ToSliceAny() SliceT0 {
	r := make(SliceT0, len(s))
	for i, x := range s {
		r[i] = x
	}
	return r
}

// ToSliceString is a conversion function.
func ToSliceString(s SliceT0) SliceString {
	r := make(SliceString, len(s))
	for i, x := range s {
		r[i] = x.(string)
	}
	return r
}
