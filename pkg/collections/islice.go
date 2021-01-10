package collections

// ISliceT0 defines the methods to be implemented by the concrete type SliceT0 that only
// depend on type T0.
type ISliceT0 interface {
	Length() int
	Contains(elem T0) bool
	ContainsAll(elems SliceT0) bool
	Get(index int) T0
	IndexOf(elem T0) int
	IsEmpty() bool
	LastIndexOf(elem T0) int
	SubSlice(fromIndex int, toIndex int) SliceT0
	All(pred func(T0) bool) bool
	Any(pred func(T0) bool) bool
	Count(pred func(T0) bool) int
	Drop(n int) SliceT0
	DropLast(n int) SliceT0
	DropLastWhile(pred func(T0) bool) SliceT0
	DropWhile(pred func(T0) bool) SliceT0
	Filter(pred func(T0) bool) SliceT0
	FilterNot(pred func(T0) bool) SliceT0
	Find(elem T0) T0
	First() T0
	ForEach(f func(T0))
	IndexOfFirst(pred func(T0) bool) int
	IndexOfLast(pred func(T0) bool) int
	IsNotEmpty() bool
	Last() T0
	MaxWith(comparator func(T0, T0) int) T0
	Minus(other SliceT0) SliceT0
	MinusElement(elem T0) SliceT0
	MinWith(comparator func(T0, T0) int) T0
	Partition(pred func(T0) bool) (SliceT0, SliceT0)
	Plus(other SliceT0) SliceT0
	PlusElement(elem T0) SliceT0
	Reduce(op func(T0, T0) T0) T0
	Reversed() SliceT0
	SortedWith(comparator func(T0, T0) int) SliceT0
	Take(n int) SliceT0
	TakeLast(n int) SliceT0
	TakeLastWhile(pred func(T0) bool) SliceT0
	TakeWhile(pred func(T0) bool) SliceT0
	ToSlice() SliceT0
	// ToSet() Set
}

type ISliceT0T1 interface {
	FlatMapAnyT1(func(T0) []T1) []T1
	FoldAnyT1(z T1, op func(T1, T0) T1) T1
	GroupByAnyT1(keySelector func(T0) T1) map[T0][]T1
	MapAnyT1(f func(T0) T1) []T1
	ZipAnyT1(other SliceT1) []PairT0T1
}

type ISlice2T0 interface {
	Flatten() SliceT0
}

type ISliceOfPairT0T1 interface {
	ToMap() map[T0]T1
}

func validateListInterface(s SliceT0) {
	f := func(itf ISliceT0) {}
	f(s)
	g := func(itf ISliceT0T1) {}
	g(s)
}

func validateListOfListInterface(s Slice2T0) {
	f := func(itf ISlice2T0) {}
	f(s)
}

func validateListOfPairInterface(s SliceOfPairT0T1) {
	f := func(itf ISliceOfPairT0T1) {}
	f(s)
}
