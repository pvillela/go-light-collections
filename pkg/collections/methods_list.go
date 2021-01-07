package collections

// methods is just an inventory of the methods of SliceAny.
type methods interface {
	Length() int
	Contains(elem AnyT0) bool
	ContainsAll(elems SliceAny) bool
	Get(index int) AnyT0
	IndexOf(elem AnyT0) int
	IsEmpty() bool
	LastIndexOf(elem AnyT0) int
	SubSlice(fromIndex int, toIndex int) SliceAny
	All(pred func(AnyT0) bool) bool
	Any(pred func(AnyT0) bool) bool
	Count(pred func(AnyT0) bool) int
	Drop(n int) SliceAny
	DropLast(n int) SliceAny
	DropLastWhile(pred func(AnyT0) bool) SliceAny
	DropWhile(pred func(AnyT0) bool) SliceAny
	Filter(pred func(AnyT0) bool) SliceAny
	FilterNot(pred func(AnyT0) bool) SliceAny
	Find(elem AnyT0) AnyT0
	// // FindLast(elem Any) Any
	First() AnyT0
	// // FlatMap -- no clean implementation without generics
	// // Flatten -- no clean implementation without generics
	Fold(z AnyT0, op func(AnyT0, AnyT0) AnyT0) AnyT0
	ForEach(f func(AnyT0))
	GroupBy(keySelector func(AnyT0) AnyT0) map[AnyT0]SliceAny
	IndexOfFirst(pred func(AnyT0) bool) int
	IndexOfLast(pred func(AnyT0) bool) int
	// IsNotEmpty() bool
	Last() AnyT0
	Map(f func(AnyT0) AnyT0) SliceAny
	MaxWithOrNil(comparator func(AnyT0, AnyT0) int) AnyT0
	Minus(other SliceAny) SliceAny
	MinusElement(elem AnyT0) SliceAny
	MinWithOrNil(comparator func(AnyT0, AnyT0) int) AnyT0
	Partition(pred func(AnyT0) bool) (SliceAny, SliceAny)
	Plus(other SliceAny) SliceAny
	PlusElement(elem AnyT0) SliceAny
	// Reduce(op func(AnyT, AnyT) AnyT) AnyT
	ReduceOrNil(op func(AnyT0, AnyT0) AnyT0) AnyT0
	Reversed() SliceAny
	// // RunningFold(z Any, op FuncAnyAnyAny) SliceAny
	// // RunningReduce(op FuncAnyAnyAny) SliceAny
	SortedWith(comparator func(AnyT0, AnyT0) int) SliceAny
	Take(n int) SliceAny
	TakeLast(n int) SliceAny
	// TakeLastWhile(pred func(AnyT) bool) SliceAny
	// TakeWhile(pred func(AnyT) bool) SliceAny
	// ToSlice() SliceAny
	// // ToMap() Map
	// // ToSet() Set
	// Zip(other SliceAny) SliceAny
}

func validateInterface(s SliceAny) {
	f := func(itf methods) {}
	f(s)
}
