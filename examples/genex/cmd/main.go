package main

import (
	"fmt"

	"github.com/pvillela/go-light-collections/examples/genex/generated/coll"
	"github.com/pvillela/go-light-collections/examples/genex/pkga"
)

// Simple example of usage of generated collections
func main() {
	sPerson := coll.SlicePerson{{Name: "Peter", Age: 70}, {Name: "Paul", Age: 60},
		{Name: "Mary", Age: 65}}
	f := func(p pkga.Person) int { return len(p.Name) + p.Age }
	res := sPerson.Mapint(f)
	fmt.Printf("%T %v\n", res, res)
	// Output: coll.Sliceint [75 64 69]
}
