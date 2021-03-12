/*
 * Copyright © 2021 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package glc

// PairT1T0 defines a pseudo-generic pair.
type PairT1T0 struct {
	X1 T1
	X2 T0
}

// SliceOfPairT1T0 is a type wrapper.
type SliceOfPairT1T0 []PairT1T0

// MapT1T0 is a type wrapper, implements Map interface.
type MapT1T0 map[T1]T0

// MapT1SliceT0 is a type wrapper.
type MapT1SliceT0 map[T1]SliceT0

// MapT1SetT0 is a type wrapper.
type MapT1SetT0 map[T1]SetT0

// SetOfPairT1T0 is a type wrapper, implements Set interface.
type SetOfPairT1T0 map[PairT1T0]bool
