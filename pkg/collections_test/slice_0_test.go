package collections_test

import (
	"testing"

	c "github.com/pvillela/go-light-collections/pkg/collections"
	"github.com/stretchr/testify/assert"
)

////
// Slices used as inputs to functions below. Cloned each time to avoid nasty side-effects
// in tests.

func sAny() c.SliceAny {
	var sFoo SliceFoo = []Foo{{1, "w1"}, {22, "w22"}, {333, "w333"}, {4444, "w4444"}}
	return sFoo.ToSliceAny()
}

func sOtherAny() c.SliceAny {
	var sOtherFoo SliceFoo = []Foo{{0, "w"}, {22, "w22"}, {55555, "w55555"}}
	return sOtherFoo.ToSliceAny()
}

func sEmptyAny() c.SliceAny {
	var emptySliceFoo SliceFoo = []Foo{}
	return emptySliceFoo.ToSliceAny()
}

////
// Tests

func TestLength(t *testing.T) {
	func() {
		got := sAny().Length()
		want := 4
		assert.Equal(t, want, got, "Non-empty slice")
	}()

	func() {
		got := sEmptyAny().Length()
		want := 0
		assert.Equal(t, want, got, "Empty slice")
	}()
}

func TestContains(t *testing.T) {
	func() {
		got := sAny().Contains(Foo{22, "w22"})
		want := true
		assert.Equal(t, want, got, "Present")
	}()

	func() {
		got := sAny().Contains(Foo{22, "xyz"})
		want := false
		assert.Equal(t, want, got, "Absent")
	}()
}

func TestContainsAll(t *testing.T) {
	func() {
		other := append(sAny()[2:2], sAny()[1])
		got := sAny().ContainsAll(other)
		want := true
		assert.Equal(t, want, got, "Subset")
	}()

	func() {
		other := append(sAny()[1:1], Foo{22, "xyz"})
		got := sAny().ContainsAll(other)
		want := false
		assert.Equal(t, want, got, "Not subset")
	}()
}

func TestGet(t *testing.T) {

}

func TestIndexOf(t *testing.T) {

}

func TestIsEmpty(t *testing.T) {

}

func TestLastIndexOf(t *testing.T) {

}

func TestSubSlice(t *testing.T) {

}

func TestAll(t *testing.T) {

}

func TestAny(t *testing.T) {

}

func TestCount(t *testing.T) {

}

func TestDrop(t *testing.T) {

}

func TestDropLast(t *testing.T) {

}

func TestDropLastWhile(t *testing.T) {

}

func TestDropWhile(t *testing.T) {

}

func TestFilter(t *testing.T) {

}

func TestFilterNot(t *testing.T) {

}

func TestFind(t *testing.T) {

}

func TestFirst(t *testing.T) {

}

func TestFlatMap(t *testing.T) {

}

func TestFold(t *testing.T) {

}

func TestForEach(t *testing.T) {

}

func TestGroupBy(t *testing.T) {

}

func TestIndexOfFirst(t *testing.T) {

}

func TestIndexOfLast(t *testing.T) {

}

func TestLast(t *testing.T) {

}

func TestMap(t *testing.T) {

}

func TestMaxWithOrNil(t *testing.T) {

}

func TestMinus(t *testing.T) {

}

func TestMinusElement(t *testing.T) {

}

func TestMinWithOrNil(t *testing.T) {

}

func TestPartition(t *testing.T) {

}

func TestPlus(t *testing.T) {

}

func TestPlusElement(t *testing.T) {

}

func TestReduceOrNil(t *testing.T) {

}

func TestReversed(t *testing.T) {

}

func TestSortedWith(t *testing.T) {

}

func TestTake(t *testing.T) {

}

func TestTakeLast(t *testing.T) {

}

func TestTakeLastWhile(t *testing.T) {

}

func TestTakeWhile(t *testing.T) {

}

func TestToSlice(t *testing.T) {

}

func TestZip(t *testing.T) {

}
