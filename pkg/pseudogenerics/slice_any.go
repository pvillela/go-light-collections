package pseudogenerics

import (
	"reflect"
	"sort"
)

// SliceAny is a type wrapper, implements SliceAny interface.
type SliceAny []AnyT0

// ConvertibleToSliceAny is used as a type for slice parameters in pseudo-polymorphic functions.
type ConvertibleToSliceAny interface {
	ToSliceAny() SliceAny
}

/////////////////////
// Methods

func (s SliceAny) Length() int {
	return len(s)
}

func (s SliceAny) Contains(elem AnyT0) bool {
	return s.IndexOf(elem) >= 0
}

func (s SliceAny) ContainsAll(elems SliceAny) bool {
	for _, e := range elems {
		if !s.Contains(e) {
			return false
		}
	}
	return true
}

func (s SliceAny) Get(index int) AnyT0 {
	return s[index]
}

func (s SliceAny) IndexOf(elem AnyT0) int {
	pred := func(a AnyT0) bool { return reflect.DeepEqual(elem, a) }
	return s.IndexOfFirst(pred)
}

func (s SliceAny) IsEmpty() bool {
	return s == nil || len(s) == 0
}

func (s SliceAny) LastIndexOf(elem AnyT0) int {
	pred := func(a AnyT0) bool { return reflect.DeepEqual(elem, a) }
	return s.IndexOfLast(pred)
}

func (s SliceAny) SubSlice(fromIndex int, toIndex int) SliceAny {
	return s[fromIndex:toIndex]
}

func (s SliceAny) All(pred func(AnyT0) bool) bool {
	for _, x := range s {
		if !pred(x) {
			return false
		}
	}
	return true
}

func (s SliceAny) Any(pred func(AnyT0) bool) bool {
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

func (s SliceAny) Count(pred func(AnyT0) bool) int {
	return s.Fold(0, countOp(pred)).(int)
}

func (s SliceAny) Drop(n int) SliceAny {
	if n > len(s) {
		return s[:0]
	}
	return s[n:]
}

func (s SliceAny) DropLast(n int) SliceAny {
	if n > len(s) {
		return s[:0]
	}
	return s[:len(s)-n]
}

func (s SliceAny) DropLastWhile(pred FuncAnyBool) SliceAny {
	last := len(s)
	for i := len(s) - 1; i >= 0 && pred(s[i]); i-- {
		last = i
	}
	return s[:last]
}

func (s SliceAny) DropWhile(pred FuncAnyBool) SliceAny {
	first := 0
	for i := 0; i < len(s) && pred(s[i]); i++ {
		first = i + 1
	}
	return s[first:]
}

// Filter returns a new slice containing only the elements in the given slice for which
// the application of the function pred returns true.
func (s SliceAny) Filter(pred func(AnyT0) bool) SliceAny {
	output := make(SliceAny, 0)
	for i, a := range s {
		if pred(s[i]) {
			output = append(output, a)
		}
	}
	return output
}

func (s SliceAny) FilterNot(pred func(AnyT0) bool) SliceAny {
	return s.Filter(func(a AnyT0) bool { return pred(a) })
}

func (s SliceAny) Find(elem AnyT0) AnyT0 {
	if s.IndexOf(elem) != -1 {
		return elem
	}
	return nil
}

func (s SliceAny) First() AnyT0 {
	return s[0]
}

// Fold returns the accumulated value obtained by applying the operation op to z,
// and the first element of the given slice, then applying op to the result of the first
// operation and the second element of the given slice, and so on.
func (s SliceAny) Fold(z AnyT0, op func(AnyT0, AnyT0) AnyT0) AnyT0 {
	result := z
	for _, a := range s {
		result = op(result, a)
	}
	return result
}

func (s SliceAny) ForEach(f func(AnyT0)) {
	for _, x := range s {
		f(x)
	}
}

func (s SliceAny) GroupBy(keySelector func(AnyT0) AnyT0) map[AnyT0]SliceAny {
	m := make(map[AnyT0]SliceAny)
	for _, x := range s {
		k := keySelector(x)
		lst, ok := m[k]
		if !ok {
			lst = make(SliceAny, 1)
			m[k] = lst
		}
		lst = append(lst, x)
	}
	return m
}

func (s SliceAny) IndexOfFirst(pred func(AnyT0) bool) int {
	for i, x := range s {
		if pred(x) {
			return i
		}
	}
	return -1
}

func (s SliceAny) IndexOfLast(pred func(AnyT0) bool) int {
	for i := len(s) - 1; i >= 0; i-- {
		if pred(s[i]) {
			return i
		}
	}
	return -1
}

func (s SliceAny) Last() AnyT0 {
	return s[len(s)]
}

// Map returns a new slice resulting from the application of a given function to
// each element of a given slice.
func (s SliceAny) Map(f func(AnyT0) AnyT0) SliceAny {
	output := make(SliceAny, len(s))
	for i, a := range s {
		output[i] = f(a)
	}
	return output
}

func (s SliceAny) MaxWithOrNil(comparator func(AnyT0, AnyT0) int) AnyT0 {
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

func (s SliceAny) Minus(other SliceAny) SliceAny {
	r := s
	for _, x := range other {
		r = r.MinusElement(x)
	}
	return r
}

func (s SliceAny) MinusElement(elem AnyT0) SliceAny {
	return s.Filter(func(a AnyT0) bool { return reflect.DeepEqual(a, elem) })
}

func (s SliceAny) MinWithOrNil(comparator func(AnyT0, AnyT0) int) AnyT0 {
	reverseComp := func(a1 AnyT0, a2 AnyT0) int { return -comparator(a1, a2) }
	return s.MaxWithOrNil(reverseComp)
}

func (s SliceAny) Partition(pred func(AnyT0) bool) (SliceAny, SliceAny) {
	output1 := make(SliceAny, 0)
	output2 := make(SliceAny, 0)
	for i, a := range s {
		if pred(s[i]) {
			output1 = append(output1, a)
		} else {
			output2 = append(output2, a)
		}
	}
	return output1, output2
}

func (s SliceAny) Plus(other SliceAny) SliceAny {
	return append(s, other...)
}

func (s SliceAny) PlusElement(elem AnyT0) SliceAny {
	return append(s, elem)
}

func (s SliceAny) ReduceOrNil(op func(AnyT0, AnyT0) AnyT0) AnyT0 {
	if len(s) == 0 {
		return nil
	}
	return s.Fold(s[0], op)
}

func (s SliceAny) Reversed() SliceAny {
	r := make(SliceAny, len(s))
	size := len(s)
	for i, x := range s {
		r[size-i-1] = x
	}
	return r
}

type sortable struct {
	comparator func(AnyT0, AnyT0) int
	slice      SliceAny
}

func (x sortable) Len() int               { return len(x.slice) }
func (x sortable) Less(i int, j int) bool { return x.comparator(x.slice[i], x.slice[j]) < 0 }
func (x sortable) Swap(i int, j int)      { x.slice[i], x.slice[j] = x.slice[j], x.slice[i] }

func (s SliceAny) SortedWith(comparator func(AnyT0, AnyT0) int) SliceAny {
	r := make(SliceAny, len(s))
	srt := sortable{comparator: comparator, slice: r}
	sort.Sort(srt)
	return r
}

func (s SliceAny) Take(n int) SliceAny {
	if n > len(s) {
		return s
	}
	return s[:n]
}

func (s SliceAny) TakeLast(n int) SliceAny {
	if n > len(s) {
		return s
	}
	return s[len(s)-n:]
}
