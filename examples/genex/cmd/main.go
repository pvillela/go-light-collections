package main

import (
	"fmt"

	"github.com/pvillela/go-light-collections/examples/genex/generated/coll"
	"github.com/pvillela/go-light-collections/examples/genex/pkga"
)

// Example of usage of generated collections
func main() {

	// Examples with various collections involving the Person type.

	var persons coll.SlicePerson = []coll.Person{{Name: "John", Age: 80}, {Name: "Paul", Age: 78},
		{Name: "George", Age: 77}, {Name: "Ringo", Age: 80}}
	fmt.Printf("*** persons: %#v\n", persons)

	personsWithEvenNames := persons.Filter(func(p pkga.Person) bool { return len(p.Name)%2 == 0 })
	fmt.Printf("*** personsWithEvenNames: %#v\n", personsWithEvenNames)

	var namesOfPWEN coll.SliceString = personsWithEvenNames.MapString(func(p pkga.Person) string { return p.Name })
	fmt.Printf("*** namesOfPWEN: %#v\n", namesOfPWEN)

	var agesOfPWEN coll.Sliceint = personsWithEvenNames.Mapint(func(p pkga.Person) int { return p.Age })
	fmt.Printf("*** agesOfPWEN: %#v\n", agesOfPWEN)

	var nameAgePairsOfPWEN coll.SliceOfPairStringint = namesOfPWEN.Zipint(agesOfPWEN)
	var namesToAgesOfPWEN coll.MapStringint = nameAgePairsOfPWEN.ToMap()
	fmt.Printf("*** namesToAgesOfPWEN: %#v\n", namesToAgesOfPWEN)
	fmt.Println(namesToAgesOfPWEN["Paul"])
	fmt.Println(namesToAgesOfPWEN["Ringo"])

	beatlesMap := namesToAgesOfPWEN.Add("Ringo", 80)
	fmt.Printf("*** beatlesMap: %#v\n", beatlesMap)

	beatlesMapEvenAges := beatlesMap.FilterValues(func(age int) bool { return age%2 == 0 })
	fmt.Printf("*** beatlesMapEvenAges: %#v\n", beatlesMapEvenAges)

	mungedMap := beatlesMapEvenAges.MapValuesint(func(p coll.PairSlStringint) int { return len(p.X1) + p.X2 })
	fmt.Printf("*** mungedMap: %#v\n", mungedMap)

	// Example with nested collection type.
}
