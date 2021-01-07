package collections

import (
	"reflect"
	"sort"
)

// SliceAny0 is a type wrapper, implements List interface.
type SliceAny0 []AnyT0

// SliceAny1 is a type wrapper.
// Used to clarify method signatures and facilitate replacement for code generation.
type SliceAny1 []AnyT1

// Slice2Any1 is a type wrapper.
// Used to clarify method signatures and facilitate replacement for code generation.
type Slice2Any1 []SliceAny1

// ConvertibleToSliceAny is used as a type for slice parameters in pseudo-polymorphic functions.
type ConvertibleToSliceAny interface {
	ToSliceAny() SliceAny0
}

/////////////////////
// Methods

func (s SliceAny0) Length() int {
	return len(s)
}

func (s SliceAny0) Contains(elem AnyT0) bool {
	return s.IndexOf(elem) >= 0
}

func (s SliceAny0) ContainsAll(elems SliceAny0) bool {
	for _, e := range elems {
		if !s.Contains(e) {
			return false
		}
	}
	return true
}

func (s SliceAny0) Get(index int) AnyT0 {
	return s[index]
}

func (s SliceAny0) IndexOf(elem AnyT0) int {
	pred := func(a AnyT0) bool { return reflect.DeepEqual(elem, a) }
	return s.IndexOfFirst(pred)
}

func (s SliceAny0) IsEmpty() bool {
	return s == nil || len(s) == 0
}

func (s SliceAny0) LastIndexOf(elem AnyT0) int {
	pred := func(a AnyT0) bool { return reflect.DeepEqual(elem, a) }
	return s.IndexOfLast(pred)
}

func (s SliceAny0) SubSlice(fromIndex int, toIndex int) SliceAny0 {
	return s[fromIndex:toIndex]
}

func (s SliceAny0) All(pred func(AnyT0) bool) bool {
	for _, x := range s {
		if !pred(x) {
			return false
		}
	}
	return true
}

func (s SliceAny0) Any(pred func(AnyT0) bool) bool {
	for _, x := range s {
		if pred(x) {
			return true
		}
	}
	return false
}

func countOp(pred func(AnyT0) bool) func(AnyT0, AnyT0) AnyT0 {
	return func(n AnyT0, a AnyT0) AnyT0 {
		if pred(a) {
			return n.(int) + 1
		}
		return n.(int)
	}
}

func (s SliceAny0) Count(pred func(AnyT0) bool) int {
	return s.Fold(0, countOp(pred)).(int)
}

func (s SliceAny0) Drop(n int) SliceAny0 {
	if n > len(s) {
		return s[:0]
	}
	return s[n:]
}

func (s SliceAny0) DropLast(n int) SliceAny0 {
	if n > len(s) {
		return s[:0]
	}
	return s[:len(s)-n]
}

func (s SliceAny0) DropLastWhile(pred FuncAnyBool) SliceAny0 {
	last := len(s)
	for i := len(s) - 1; i >= 0 && pred(s[i]); i-- {
		last = i
	}
	return s[:last]
}

func (s SliceAny0) DropWhile(pred FuncAnyBool) SliceAny0 {
	first := 0
	for i := 0; i < len(s) && pred(s[i]); i++ {
		first = i + 1
	}
	return s[first:]
}

// Filter returns a new slice containing only the elements in the given slice for which
// the application of the function pred returns true.
func (s SliceAny0) Filter(pred func(AnyT0) bool) SliceAny0 {
	output := make(SliceAny0, 0)
	for i, a := range s {
		if pred(s[i]) {
			output = append(output, a)
		}
	}
	return output
}

func (s SliceAny0) FilterNot(pred func(AnyT0) bool) SliceAny0 {
	return s.Filter(func(a AnyT0) bool { return pred(a) })
}

func (s SliceAny0) Find(elem AnyT0) AnyT0 {
	if s.IndexOf(elem) != -1 {
		return elem
	}
	return nil
}

func (s SliceAny0) First() AnyT0 {
	return s[0]
}

// Fold returns the accumulated value obtained by applying the operation op to z,
// and the first element of the given slice, then applying op to the result of the first
// operation and the second element of the given slice, and so on.
func (s SliceAny0) Fold(z AnyT0, op func(AnyT0, AnyT0) AnyT0) AnyT0 {
	result := z
	for _, a := range s {
		result = op(result, a)
	}
	return result
}

func (s SliceAny0) ForEach(f func(AnyT0)) {
	for _, x := range s {
		f(x)
	}
}

func (s SliceAny0) GroupBy(keySelector func(AnyT0) AnyT0) map[AnyT0]SliceAny0 {
	m := make(map[AnyT0]SliceAny0)
	for _, x := range s {
		k := keySelector(x)
		lst, ok := m[k]
		if !ok {
			lst = make(SliceAny0, 1)
			m[k] = lst
		}
		lst = append(lst, x)
	}
	return m
}

func (s SliceAny0) IndexOfFirst(pred func(AnyT0) bool) int {
	for i, x := range s {
		if pred(x) {
			return i
		}
	}
	return -1
}

func (s SliceAny0) IndexOfLast(pred func(AnyT0) bool) int {
	for i := len(s) - 1; i >= 0; i-- {
		if pred(s[i]) {
			return i
		}
	}
	return -1
}

func (s SliceAny0) Last() AnyT0 {
	return s[len(s)]
}

// Map returns a new slice resulting from the application of a given function to
// each element of a given slice.
func (s SliceAny0) Map(f func(AnyT0) AnyT0) SliceAny0 {
	output := make(SliceAny0, len(s))
	for i, a := range s {
		output[i] = f(a)
	}
	return output
}

func (s SliceAny0) MaxWithOrNil(comparator func(AnyT0, AnyT0) int) AnyT0 {
	if len(s) == 0 {
		return nil
	} else {
		max := s[0]
		for i := 1; i < len(s); i++ {
			if comparator(max, s[i]) < 0 {
				max = s[i]
			}
		}
		return max
	}
}

func (s SliceAny0) Minus(other SliceAny0) SliceAny0 {
	r := s
	for _, x := range other {
		r = r.MinusElement(x)
	}
	return r
}

func (s SliceAny0) MinusElement(elem AnyT0) SliceAny0 {
	return s.Filter(func(a AnyT0) bool { return reflect.DeepEqual(a, elem) })
}

func (s SliceAny0) MinWithOrNil(comparator func(AnyT0, AnyT0) int) AnyT0 {
	reverseComp := func(a1 AnyT0, a2 AnyT0) int { return -comparator(a1, a2) }
	return s.MaxWithOrNil(reverseComp)
}

func (s SliceAny0) Partition(pred func(AnyT0) bool) (SliceAny0, SliceAny0) {
	output1 := make(SliceAny0, 0)
	output2 := make(SliceAny0, 0)
	for i, a := range s {
		if pred(s[i]) {
			output1 = append(output1, a)
		} else {
			output2 = append(output2, a)
		}
	}
	return output1, output2
}

func (s SliceAny0) Plus(other SliceAny0) SliceAny0 {
	return append(s, other...)
}

func (s SliceAny0) PlusElement(elem AnyT0) SliceAny0 {
	return append(s, elem)
}

func (s SliceAny0) ReduceOrNil(op func(AnyT0, AnyT0) AnyT0) AnyT0 {
	if len(s) == 0 {
		return nil
	}
	return s.Fold(s[0], op)
}

func (s SliceAny0) Reversed() SliceAny0 {
	r := make(SliceAny0, len(s))
	size := len(s)
	for i, x := range s {
		r[size-i-1] = x
	}
	return r
}

type sortable struct {
	comparator func(AnyT0, AnyT0) int
	slice      SliceAny0
}

func (x sortable) Len() int               { return len(x.slice) }
func (x sortable) Less(i int, j int) bool { return x.comparator(x.slice[i], x.slice[j]) < 0 }
func (x sortable) Swap(i int, j int)      { x.slice[i], x.slice[j] = x.slice[j], x.slice[i] }

func (s SliceAny0) SortedWith(comparator func(AnyT0, AnyT0) int) SliceAny0 {
	r := make(SliceAny0, len(s))
	srt := sortable{comparator: comparator, slice: r}
	sort.Sort(srt)
	return r
}

func (s SliceAny0) Take(n int) SliceAny0 {
	if n > len(s) {
		return s
	}
	return s[:n]
}

func (s SliceAny0) TakeLast(n int) SliceAny0 {
	if n > len(s) {
		return s
	}
	return s[len(s)-n:]
}
