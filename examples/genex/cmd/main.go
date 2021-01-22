package main

import (
	"fmt"

	"github.com/pvillela/go-light-collections/examples/genex/generated/coll"
	"github.com/pvillela/go-light-collections/examples/genex/pkga"
)

// Simple example of usage of generated collections
func main() {
	persons := coll.SlicePerson{{Name: "John", Age: 80}, {Name: "Paul", Age: 78},
		{Name: "George", Age: 77}, {Name: "Ringo", Age: 80}}

	personsWithEvenNames := persons.Filter(func(p pkga.Person) bool { return len(p.Name)%2 == 0 })

	namesOfPWEN := personsWithEvenNames.MapString(func(p pkga.Person) string { return p.Name })

	agesOfPWEN := personsWithEvenNames.Mapint(func(p pkga.Person) int { return p.Age })

	namesToAgesOfPWEN := namesOfPWEN.Zipint(agesOfPWEN).ToMap()

	fmt.Println(namesToAgesOfPWEN["Paul"])
	// Output: 78

	fmt.Println(namesToAgesOfPWEN["Ringo"])
	// Output: 0

	fmt.Println(namesToAgesOfPWEN)
	// Output: map[George:77 John:80 Paul:78]

	beatlesMap := namesToAgesOfPWEN.Add("Ringo", 80)

	fmt.Println(beatlesMap)
	// Output: map[George:77 John:80 Paul:78 Ringo:80]

	beatlesMapEvenAges := beatlesMap.FilterValues(func(age int) bool { return age%2 == 0 })

	fmt.Println(beatlesMapEvenAges)
	// Output: map[John:80 Paul:78 Ringo:80]

	mungedMap := beatlesMapEvenAges.MapValuesint(func(p coll.PairStringint) int { return len(p.X1) + p.X2 })

	fmt.Println(mungedMap)
	// Output: map[John:84 Paul:82 Ringo:85]
}
