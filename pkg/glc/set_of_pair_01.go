/*
 * Copyright © 2021 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package glc

// SetOfPairT0T1 is a type wrapper, implements ISetOfPairT0T1 interface.
type SetOfPairT0T1 map[struct {
	X1 T0
	X2 T1
}]bool

// ToMap returns a map whose keys are the first components in the elements of the receiver and
// whose values are the corresonding second components in the elements of the receiver.
// If multiple elements in the receiver have the same first component, the corresponding
// value in the resulting map will be picked from one of them.
func (s SetOfPairT0T1) ToMap() MapT0T1 {
	if s == nil {
		return nil
	}
	m := make(map[T0]T1, len(s))
	for p := range s {
		m[p.X1] = p.X2
	}
	return m
}
