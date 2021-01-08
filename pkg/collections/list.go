package collections

// List (obviously) defines the methods to be implemented by List implementations.
type List interface {
	Length() int
	Contains(elem AnyT0) bool
	ContainsAll(elems SliceT0) bool
	Get(index int) AnyT0
	IndexOf(elem AnyT0) int
	IsEmpty() bool
	LastIndexOf(elem AnyT0) int
	SubSlice(fromIndex int, toIndex int) SliceT0
	All(pred func(AnyT0) bool) bool
	Any(pred func(AnyT0) bool) bool
	Count(pred func(AnyT0) bool) int
	Drop(n int) SliceT0
	DropLast(n int) SliceT0
	DropLastWhile(pred func(AnyT0) bool) SliceT0
	DropWhile(pred func(AnyT0) bool) SliceT0
	Filter(pred func(AnyT0) bool) SliceT0
	FilterNot(pred func(AnyT0) bool) SliceT0
	Find(elem AnyT0) AnyT0
	// // FindLast(elem Any) Any
	First() AnyT0
	FlatMap(func(AnyT0) []AnyT1) []AnyT1
	// // Flatten -- defined in interface ListOfList
	Fold(z AnyT1, op func(AnyT1, AnyT0) AnyT1) AnyT1
	ForEach(f func(AnyT0))
	GroupBy(keySelector func(AnyT0) AnyT1) map[AnyT0][]AnyT1
	IndexOfFirst(pred func(AnyT0) bool) int
	IndexOfLast(pred func(AnyT0) bool) int
	// IsNotEmpty() bool
	Last() AnyT0
	Map(f func(AnyT0) AnyT1) []AnyT1
	MaxWithOrNil(comparator func(AnyT0, AnyT0) int) AnyT0
	Minus(other SliceT0) SliceT0
	MinusElement(elem AnyT0) SliceT0
	MinWithOrNil(comparator func(AnyT0, AnyT0) int) AnyT0
	Partition(pred func(AnyT0) bool) (SliceT0, SliceT0)
	Plus(other SliceT0) SliceT0
	PlusElement(elem AnyT0) SliceT0
	// Reduce(op func(AnyT, AnyT) AnyT) AnyT
	ReduceOrNil(op func(AnyT0, AnyT0) AnyT0) AnyT0
	Reversed() SliceT0
	// // RunningFold(z Any, op func(AnyT0, AnyT1) AnyT1) SliceAny
	// // RunningReduce(op func(AnyT0, AnyT1) AnyT1) SliceAny
	SortedWith(comparator func(AnyT0, AnyT0) int) SliceT0
	Take(n int) SliceT0
	TakeLast(n int) SliceT0
	TakeLastWhile(pred func(AnyT0) bool) SliceT0
	TakeWhile(pred func(AnyT0) bool) SliceT0
	ToSlice() SliceT0
	// // ToSet() Set
	Zip(other SliceT1) SliceTPair01
}

type ListOfList interface {
	Flatten() SliceT0
}

type ListOfPair interface {
	ToMap() map[AnyT0]AnyT1
}

func validateListInterface(s SliceT0) {
	f := func(itf List) {}
	f(s)
}

func validateListOfListInterface(s Slice2T0) {
	f := func(itf ListOfList) {}
	f(s)
}

func validateListOfPairInterface(s SliceTPair01) {
	f := func(itf ListOfPair) {}
	f(s)
}