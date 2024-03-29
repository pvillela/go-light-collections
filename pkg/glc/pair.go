/*
 *  Copyright © 2021 Paulo Villela. All rights reserved.
 *  Use of this source code is governed by the Apache 2.0 license
 *  that can be found in the LICENSE file.
 */

package glc

// Pair is a generic pair type.
type Pair[T1 any, T2 any] struct {
	X1 T1
	X2 T2
}
