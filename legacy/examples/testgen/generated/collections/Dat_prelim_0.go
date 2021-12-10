// Code generated -- DO NOT EDIT.

/*
 * Copyright © 2021 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package collections

////
// Preliminaries

// Dat is an example data structure.
type Dat struct {
	V1 int
	V2 string
}

// Type conversion functions.

func toDat(a Any) Dat { return Any(a).(Dat) }

func toInt(a Any) int { return Any(a).(int) }

func toString(a Any) string { return Any(a).(string) }
