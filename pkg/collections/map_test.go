package collections_test

import (
	"fmt"

	c "github.com/pvillela/GoSimpleCollections/pkg/collections"
)

func ExampleMap_Get() {
	m := c.MakeMap()
	m.Put(1, "a")
	fmt.Println(m.Get(1))
	// Output: a
}