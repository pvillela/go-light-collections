// Code generated -- DO NOT EDIT.

/*
 * Copyright © 2021 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package collections

// ToSet returns a set containing the values in the receiver.
func (s SliceDat) ToSet() map[Dat]bool {
	if s == nil {
		return nil
	}
	set := make(map[Dat]bool, len(s)) // optimize for speed vs space
	for _, x := range s {
		set[x] = true
	}
	return set
}
