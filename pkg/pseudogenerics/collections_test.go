package pseudogenerics_test

import (
	"fmt"

	pg "github.com/pvillela/GoSimpleCollections/pkg/pseudogenerics"
)

func ExampleMap_Get() {
	m := pg.MakeMap()
	m.Put(1, "a")
	fmt.Println(m.Get(1))
	// Output: a
}
