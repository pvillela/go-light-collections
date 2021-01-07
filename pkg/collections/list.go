package collections

// List (obviously) defines the methods to be implemented by List implementations.
type List interface {
	Length() int
	Contains(elem AnyT0) bool
	ContainsAll(elems SliceAny0) bool
	Get(index int) AnyT0
	IndexOf(elem AnyT0) int
	IsEmpty() bool
	LastIndexOf(elem AnyT0) int
	SubSlice(fromIndex int, toIndex int) SliceAny0
	All(pred func(AnyT0) bool) bool
	Any(pred func(AnyT0) bool) bool
	Count(pred func(AnyT0) bool) int
	Drop(n int) SliceAny0
	DropLast(n int) SliceAny0
	DropLastWhile(pred func(AnyT0) bool) SliceAny0
	DropWhile(pred func(AnyT0) bool) SliceAny0
	Filter(pred func(AnyT0) bool) SliceAny0
	FilterNot(pred func(AnyT0) bool) SliceAny0
	Find(elem AnyT0) AnyT0
	// // FindLast(elem Any) Any
	First() AnyT0
	// // FlatMap -- no clean implementation without generics
	// // Flatten -- defined in interface ListOfList
	Fold(z AnyT0, op func(AnyT0, AnyT0) AnyT0) AnyT0
	ForEach(f func(AnyT0))
	GroupBy(keySelector func(AnyT0) AnyT0) map[AnyT0]SliceAny0
	IndexOfFirst(pred func(AnyT0) bool) int
	IndexOfLast(pred func(AnyT0) bool) int
	// IsNotEmpty() bool
	Last() AnyT0
	Map(f func(AnyT0) AnyT0) SliceAny0
	MaxWithOrNil(comparator func(AnyT0, AnyT0) int) AnyT0
	Minus(other SliceAny0) SliceAny0
	MinusElement(elem AnyT0) SliceAny0
	MinWithOrNil(comparator func(AnyT0, AnyT0) int) AnyT0
	Partition(pred func(AnyT0) bool) (SliceAny0, SliceAny0)
	Plus(other SliceAny0) SliceAny0
	PlusElement(elem AnyT0) SliceAny0
	// Reduce(op func(AnyT, AnyT) AnyT) AnyT
	ReduceOrNil(op func(AnyT0, AnyT0) AnyT0) AnyT0
	Reversed() SliceAny0
	// // RunningFold(z Any, op FuncAnyAnyAny) SliceAny
	// // RunningReduce(op FuncAnyAnyAny) SliceAny
	SortedWith(comparator func(AnyT0, AnyT0) int) SliceAny0
	Take(n int) SliceAny0
	TakeLast(n int) SliceAny0
	// TakeLastWhile(pred func(AnyT) bool) SliceAny
	// TakeWhile(pred func(AnyT) bool) SliceAny
	// ToSlice() SliceAny
	// // ToMap() Map
	// // ToSet() Set
	// Zip(other SliceAny) SliceAny
}

type ListOfList interface {
	Flatten() SliceAny0
}

func validateInterface(s SliceAny0) {
	f := func(itf List) {}
	f(s)
}
