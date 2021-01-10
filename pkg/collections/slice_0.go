package collections

import (
	"reflect"
	"sort"
)

/////////////////////
// Methods

func (s SliceT0) Length() int {
	return len(s)
}

func (s SliceT0) Contains(elem AnyT0) bool {
	return s.IndexOf(elem) >= 0
}

func (s SliceT0) ContainsAll(elems SliceT0) bool {
	for i := range elems {
		e := elems[i]
		if !s.Contains(e) {
			return false
		}
	}
	return true
}

func (s SliceT0) Get(index int) AnyT0 {
	return s[index]
}

func (s SliceT0) IndexOf(elem AnyT0) int {
	pred := func(a AnyT0) bool { return reflect.DeepEqual(elem, a) }
	return s.IndexOfFirst(pred)
}

func (s SliceT0) IsEmpty() bool {
	return s == nil || len(s) == 0
}

func (s SliceT0) LastIndexOf(elem AnyT0) int {
	pred := func(a AnyT0) bool { return reflect.DeepEqual(elem, a) }
	return s.IndexOfLast(pred)
}

func (s SliceT0) SubSlice(fromIndex int, toIndex int) SliceT0 {
	return s[fromIndex:toIndex]
}

func (s SliceT0) All(pred func(AnyT0) bool) bool {
	for _, x := range s {
		if !pred(x) {
			return false
		}
	}
	return true
}

func (s SliceT0) Any(pred func(AnyT0) bool) bool {
	for _, x := range s {
		if pred(x) {
			return true
		}
	}
	return false
}

func countOp(pred func(AnyT0) bool) func(AnyT1, AnyT0) AnyT1 {
	return func(n AnyT1, a AnyT0) AnyT1 {
		if pred(a) {
			return n.(int) + 1
		}
		return n.(int)
	}
}

func (s SliceT0) Count(pred func(AnyT0) bool) int {
	return s.Fold(0, countOp(pred)).(int)
}

func (s SliceT0) Drop(n int) SliceT0 {
	if n > len(s) {
		return s[:0]
	}
	return s[n:]
}

func (s SliceT0) DropLast(n int) SliceT0 {
	if n > len(s) {
		return s[:0]
	}
	return s[:len(s)-n]
}

func (s SliceT0) DropLastWhile(pred func(AnyT0) bool) SliceT0 {
	last := len(s)
	for i := len(s) - 1; i >= 0 && pred(s[i]); i-- {
		last = i
	}
	return s[:last]
}

func (s SliceT0) DropWhile(pred func(AnyT0) bool) SliceT0 {
	first := 0
	for i := 0; i < len(s) && pred(s[i]); i++ {
		first = i + 1
	}
	return s[first:]
}

// Filter returns a new slice containing only the elements in the given slice for which
// the application of the function pred returns true.
func (s SliceT0) Filter(pred func(AnyT0) bool) SliceT0 {
	output := make(SliceT0, 0, len(s)/2) // optimizing for speed vs space
	for i, a := range s {
		if pred(s[i]) {
			output = append(output, a)
		}
	}
	return output
}

func (s SliceT0) FilterNot(pred func(AnyT0) bool) SliceT0 {
	return s.Filter(func(a AnyT0) bool { return !pred(a) })
}

func (s SliceT0) Find(elem AnyT0) AnyT0 {
	if s.IndexOf(elem) != -1 {
		return elem
	}
	return nil
}

func (s SliceT0) First() AnyT0 {
	return s[0]
}

func (s SliceT0) ForEach(f func(AnyT0)) {
	for _, x := range s {
		f(x)
	}
}

func (s SliceT0) IndexOfFirst(pred func(AnyT0) bool) int {
	for i, x := range s {
		if pred(x) {
			return i
		}
	}
	return -1
}

func (s SliceT0) IndexOfLast(pred func(AnyT0) bool) int {
	for i := len(s) - 1; i >= 0; i-- {
		if pred(s[i]) {
			return i
		}
	}
	return -1
}

func (s SliceT0) IsNotEmpty() bool {
	return !s.IsEmpty()
}

func (s SliceT0) Last() AnyT0 {
	return s[len(s)-1]
}

// MaxWith returns the first element in the slice with maximum value, using a comparator function.
// Panics if the slice is empty.
func (s SliceT0) MaxWith(comparator func(AnyT0, AnyT0) int) AnyT0 {
	max := s[0]
	for i := 1; i < len(s); i++ {
		if comparator(max, s[i]) < 0 {
			max = s[i]
		}
	}
	return max
}

func (s SliceT0) minusAllElement(elem AnyT0) SliceT0 {
	return s.FilterNot(func(a AnyT0) bool { return reflect.DeepEqual(a, elem) })
}

func (s SliceT0) Minus(other SliceT0) SliceT0 {
	r := s
	for _, x := range other {
		r = r.minusAllElement(x)
	}
	return r
}

// MinusElement -- if the element passed as an argument is present in the receiver, this
// function returns a slice with the contents of the receiver minus the first occurrence of
// that element.  Otherwise, it returns the original slice.
func (s SliceT0) MinusElement(elem AnyT0) SliceT0 {
	index := s.IndexOfFirst(func(a AnyT0) bool { return reflect.DeepEqual(a, elem) })
	if index == -1 {
		return s
	}
	return append(s[:index], s[index+1:]...)
}

// MinWith returns the first element in the slice with minimum value, using a comparator function.
// Panics if the slice is empty.
func (s SliceT0) MinWith(comparator func(AnyT0, AnyT0) int) AnyT0 {
	reverseComp := func(a1 AnyT0, a2 AnyT0) int { return -comparator(a1, a2) }
	return s.MaxWith(reverseComp)
}

func (s SliceT0) Partition(pred func(AnyT0) bool) (SliceT0, SliceT0) {
	output1 := make(SliceT0, 0, len(s)/2) // optimizing for speed vs space
	output2 := make(SliceT0, 0, len(s)/2) // optimizing for speed vs space
	for i, a := range s {
		if pred(s[i]) {
			output1 = append(output1, a)
		} else {
			output2 = append(output2, a)
		}
	}
	return output1, output2
}

func (s SliceT0) Plus(other SliceT0) SliceT0 {
	return append(s, other...)
}

func (s SliceT0) PlusElement(elem AnyT0) SliceT0 {
	return append(s, elem)
}

// fold0 is used to support Reduce
func (s SliceT0) fold0(z AnyT0, op0 func(AnyT0, AnyT0) AnyT0) AnyT0 {
	op := func(a1 AnyT1, a0 AnyT0) AnyT1 { return op0(a1.(AnyT0), a0) }
	return s.Fold(z.(AnyT1), op)
}

// Reduce returns the accumulated value obtained by applying the operation op to the first
// two elements of the given slice, then applying op to the result of the first
// operation and the third element of the given slice, and so on.
// If the slice has length 1, returns the only element in the slice.
// It is a special case of Fold where the z value is the first element of the receiver and
// the fold is executed on the original slice minus the first element.
// Panics if the slice is empty.
func (s SliceT0) Reduce(op func(AnyT0, AnyT0) AnyT0) AnyT0 {
	return s.Drop(1).fold0(s[0], op)
}

func (s SliceT0) Reversed() SliceT0 {
	r := make(SliceT0, len(s))
	size := len(s)
	for i, x := range s {
		r[size-i-1] = x
	}
	return r
}

type sortable struct {
	comparator func(AnyT0, AnyT0) int
	slice      SliceT0
}

func (x sortable) Len() int               { return len(x.slice) }
func (x sortable) Less(i int, j int) bool { return x.comparator(x.slice[i], x.slice[j]) < 0 }
func (x sortable) Swap(i int, j int)      { x.slice[i], x.slice[j] = x.slice[j], x.slice[i] }

func (s SliceT0) SortedWith(comparator func(AnyT0, AnyT0) int) SliceT0 {
	r := make(SliceT0, len(s))
	copy(r, s)
	srt := sortable{comparator: comparator, slice: r}
	sort.Sort(srt)
	return r
}

func (s SliceT0) Take(n int) SliceT0 {
	if n > len(s) {
		return s
	}
	return s[:n]
}

func (s SliceT0) TakeLast(n int) SliceT0 {
	if n > len(s) {
		return s
	}
	return s[len(s)-n:]
}

func (s SliceT0) TakeLastWhile(pred func(AnyT0) bool) SliceT0 {
	first := len(s)
	for i := len(s) - 1; i >= 0 && pred(s[i]); i-- {
		first = i
	}
	return s[first:]
}

func (s SliceT0) TakeWhile(pred func(AnyT0) bool) SliceT0 {
	last := 0
	for i := 0; i < len(s) && pred(s[i]); i++ {
		last = i + 1
	}
	return s[:last]
}

func (s SliceT0) ToSlice() SliceT0 {
	return s
}
