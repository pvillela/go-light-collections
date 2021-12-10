/*
 * Copyright © 2021 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package main

import (
	"fmt"

	"github.com/pvillela/go-light-collections/pkg/glc"
)

// Map used as input to functions below.
func mBase() glc.Map[int, string] {
	return glc.Map[int, string]{1: "w1", 22: "w22", 333: "w333", 4444: "w4444"}
}

func main() {

	func() {
		fmt.Println("\n*** map example")
		m := glc.Map[int, string]{1: "w1", 22: "w22", 333: "w333", 4444: "w4444"}
		fmt.Println(m)
		fmt.Println(mBase())
	}()
}
