package collections

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Create a specific map function from the pseudo-generic one.
func (s SliceInt) MapInt(f func(int) int) []int {
	sa := SliceInt(s).ToSliceAny()
	fa := func(a Any) Any { return f(a.(int)) }
	ra := sa.Map(fa)
	return ToSliceInt(ra)
}

// Use the specific map function
func TestSliceInt_MapInt(t *testing.T) {
	var xin SliceInt = []int{1, 22, 333}
	f := func(i int) int { return i * 10 }
	xout := xin.MapInt(f)

	assert.Equal(t, xout, []int{10, 220, 3330})
}

// Create a specific filter function from the pseudo-generic cone.
func (s SliceString) Filter(pred func(string) bool) SliceString {
	sa := SliceString(s).ToSliceAny()
	preda := func(a Any) bool { return pred(a.(string)) }
	ra := sa.Filter(preda)
	return ToSliceString(ra).Und()
}

// Use the specific filter function
func TestSliceString_Filter(t *testing.T) {
	// Use the specific filter function.
	var xin SliceString = []string{"1", "22", "333"}
	pred := func(str string) bool {
		i, _ := strconv.Atoi(str)
		return i%2 != 0
	}
	var xout []string = xin.Filter(pred)

	assert.Equal(t, xout, []string{"1", "333"})
}

// Create a specific fold function from the pseudo-generic one.
func (s SliceString) FoldInt(z int, op func(int, string) int) int {
	sa := SliceString(s).ToSliceAny()
	opa := func(za Any, a Any) Any { return op(za.(int), a.(string)) }
	ra := sa.Fold(z, opa)
	return ra.(int)
}

// Use the spdcific Fold fuunction.
func TestSliceString_FoldInt(t *testing.T) {
	// Use the specific fold function.
	var xin SliceString = []string{"1", "22", "333"}
	op := func(z int, str string) int { return z + len(str) }
	xout := xin.FoldInt(0, op)

	assert.Equal(t, xout, 6)
}
