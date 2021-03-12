/*
 * Copyright © 2021 Paulo Villela. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package main

import (
	"fmt"

	c "github.com/pvillela/go-light-collections/examples/genex/generated/coll"
	"github.com/pvillela/go-light-collections/examples/genex/pkga"
)

// Example of usage of generated collections
func main() {

	// Examples with various collections involving the Person type.

	beatles := []c.Person{{Name: "John", Year: 1940}, {Name: "Paul", Year: 1942},
		{Name: "George", Year: 1943}, {Name: "Ringo", Year: 1940}}
	fmt.Printf("*** beatles: %#v\n", beatles)

	beatlesWithEvenNames :=
		c.SlicePerson(beatles).Filter(func(p pkga.Person) bool { return len(p.Name)%2 == 0 })
	fmt.Printf("*** beatlesWithEvenNames: %#v\n", beatlesWithEvenNames)

	namesOfBWEN := beatlesWithEvenNames.MapString(func(p pkga.Person) string { return p.Name })
	fmt.Printf("*** namesOfBWEN: %#v\n", namesOfBWEN)

	agesOfBWEN := beatlesWithEvenNames.Mapint(func(p pkga.Person) int { return p.Year })
	fmt.Printf("*** agesOfBWEN: %#v\n", agesOfBWEN)

	nameYearPairsOfBWEN := c.SliceString(namesOfBWEN).Zipint(agesOfBWEN)
	fmt.Printf("*** nameYearPairsOfBWEN: %#v\n", nameYearPairsOfBWEN)

	namesToYearsOfBWEN := c.SliceOfPairStringint(nameYearPairsOfBWEN).ToMap()
	fmt.Printf("*** namesToYearsOfBWEN: %#v\n", namesToYearsOfBWEN)
	fmt.Println(namesToYearsOfBWEN["Paul"])
	fmt.Println(namesToYearsOfBWEN["Ringo"])

	beatlesMap := c.MapStringint(namesToYearsOfBWEN).Add("Ringo", 1940)
	fmt.Printf("*** beatlesMap: %#v\n", beatlesMap)

	beatlesMapEvenYears := beatlesMap.FilterValues(func(age int) bool { return age%2 == 0 })
	fmt.Printf("*** beatlesMapEvenYears: %#v\n", beatlesMapEvenYears)

	mungedMap :=
		beatlesMapEvenYears.MapValuesint(func(p c.PairMpStringint) int { return len(p.X1) + p.X2 })
	fmt.Printf("*** mungedMap: %#v\n", mungedMap)

	beatlesFullNameMap := beatlesMap.MapKeysString(func(p c.PairMpStringint) string {
		switch p.X1 {
		case "John":
			return "John Lennon"
		case "Paul":
			return "Paul McCartney"
		case "George":
			return "George Harrison"
		case "Ringo":
			return "Ringo Starr"
		default:
			return "Mick Jagger"
		}
	})
	fmt.Printf("*** beatlesFullNameMap: %#v\n", beatlesFullNameMap)

	comboMap := c.MapStringint(beatlesFullNameMap).PlusMap(mungedMap)
	fmt.Printf("*** comboMap: %#v\n", comboMap)

	// Example with nested collection type.

	sliceOfMaps := []c.MapStringint{beatlesMap, mungedMap}
	fmt.Printf("*** sliceOfMaps: %#v\n", sliceOfMaps)

	sliceOfPerson :=
		c.SliceMapStringint(sliceOfMaps).FlatMapPerson(func(m c.MapStringint) []c.Person {
			var pairs c.SlicePairMpStringint = m.ToSlice()
			return pairs.MapPerson(func(p c.PairMpStringint) c.Person {
				return c.Person{Name: p.X1, Year: p.X2}
			})
		})
	fmt.Printf("*** sliceOfPerson: %#v\n", sliceOfPerson)
}
