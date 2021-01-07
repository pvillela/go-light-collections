package pseudogenerics_test

import (
	"strconv"
	"testing"

	pg "github.com/pvillela/GoSimpleCollections/pkg/pseudogenerics"
	"github.com/stretchr/testify/assert"
)

func TestMapSlice(t *testing.T) {
	// Create a specific map function from the pseudo-generic one.
	mapSliceIntToInt := func(s []int, f func(int) int) []int {
		sa := pg.SliceInt(s)
		fa := func(a pg.AnyT0) pg.AnyT0 { return f(a.(int)) }
		ra := pg.MapSlice(sa, fa)
		return pg.ToSliceInt(ra)
	}

	// Use the specific map function
	xin := []int{1, 22, 333}
	f := func(i int) int { return i * 10 }
	xout := mapSliceIntToInt(xin, f)

	assert.Equal(t, xout, []int{10, 220, 3330})
}

func TestFilterSlice(t *testing.T) {
	// Create a specific filter function from the pseudo-generic cone.
	filterSliceInt := func(s []string, pred func(string) bool) []string {
		sa := pg.SliceString(s)
		preda := func(a pg.AnyT0) bool { return pred(a.(string)) }
		ra := pg.FilterSlice(sa, preda)
		return pg.ToSliceString(ra)
	}

	// Use the specific filter function.
	xin := []string{"1", "22", "333"}
	pred := func(str string) bool {
		i, _ := strconv.Atoi(str)
		return i%2 != 0
	}
	xout := filterSliceInt(xin, pred)

	assert.Equal(t, xout, []string{"1", "333"})
}

func TestFoldLeftSlice(t *testing.T) {
	// Create a specific fold function from the pseudo-generic one.
	foldLeftIntString := func(s []string, z int, op func(int, string) int) int {
		sa := pg.SliceString(s)
		opa := func(za pg.AnyT0, a pg.AnyT0) pg.AnyT0 { return op(za.(int), a.(string)) }
		ra := pg.FoldLeftSlice(sa, z, opa)
		return ra.(int)
	}

	// Use the specific fold function.
	xin := []string{"1", "22", "333"}
	op := func(z int, str string) int { return z + len(str) }
	xout := foldLeftIntString(xin, 0, op)

	assert.Equal(t, xout, 6)
}
