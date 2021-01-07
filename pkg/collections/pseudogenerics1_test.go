package collections_test

import (
	"strconv"
	"testing"

	pg "github.com/pvillela/GoSimpleCollections/pkg/collections"
	"github.com/stretchr/testify/assert"
)

func TestSlice_Map(t *testing.T) {
	// Create a specific map function from the pseudo-generic one.
	sliceIntMapToInt := func(s []int, f func(int) int) []int {
		sa := pg.SliceInt(s).ToSliceAny()
		fa := func(a pg.AnyT) pg.AnyT { return f(a.(int)) }
		ra := sa.Map(fa)
		return pg.ToSliceInt(ra)
	}

	// Use the specific map function
	xin := []int{1, 22, 333}
	f := func(i int) int { return i * 10 }
	xout := sliceIntMapToInt(xin, f)

	assert.Equal(t, xout, []int{10, 220, 3330})
}

func TestSlice_Filter(t *testing.T) {
	// Create a specific filter function from the pseudo-generic cone.
	sliceStringFilter := func(s []string, pred func(string) bool) []string {
		sa := pg.SliceString(s).ToSliceAny()
		preda := func(a pg.AnyT) bool { return pred(a.(string)) }
		ra := sa.Filter(preda)
		return pg.ToSliceString(ra)
	}

	// Use the specific filter function.
	xin := []string{"1", "22", "333"}
	pred := func(str string) bool {
		i, _ := strconv.Atoi(str)
		return i%2 != 0
	}
	xout := sliceStringFilter(xin, pred)

	assert.Equal(t, xout, []string{"1", "333"})
}

func TestSlice_FoldLeft(t *testing.T) {
	// Create a specific fold function from the pseudo-generic one.
	sliceStringFoldInt := func(s []string, z int, op func(int, string) int) int {
		sa := pg.SliceString(s).ToSliceAny()
		opa := func(za pg.AnyT, a pg.AnyT) pg.AnyT { return op(za.(int), a.(string)) }
		ra := sa.Fold(z, opa)
		return ra.(int)
	}

	// Use the specific fold function.
	xin := []string{"1", "22", "333"}
	op := func(z int, str string) int { return z + len(str) }
	xout := sliceStringFoldInt(xin, 0, op)

	assert.Equal(t, xout, 6)
}
