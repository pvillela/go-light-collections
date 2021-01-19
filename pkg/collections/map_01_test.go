package collections

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

////
// Preliminaries

// Map used as input to functions below.
func mBase() MapT0T1 {
	return MapT0T1{1: "w1", 22: "w22", 333: "w333", 4444: "w4444"}
}

// func pairsBase() SetOfPairT0T1 {
// 	return SetOfPairT0T1{
// 		PairT0T1{1, "w1"}: true, PairT0T1{22, "w22"}: true, PairT0T1{333, "w333"}: true,
// 		PairT0T1{4444, "w4444"}: true,
// 	}
// }

// func intX1(a Any) int { return Any(Any(a).(PairT0T1).X1).(int) }

// func stringX2(a Any) string { return Any(Any(a).(PairT0T1).X1).(string) }

func toInt(a Any) int { return Any(a).(int) }

func toString(a Any) string { return Any(a).(string) }

////
// TestMap_s

func TestMap_Copy(t *testing.T) {
	cases := []struct {
		msg      string
		receiver MapT0T1
	}{
		{"Copy: non-empty map", mBase()},
		{"Copy: empty map", MapT0T1{}},
		{"Copy: nil map", nil},
	}

	for _, cs := range cases {
		got := cs.receiver.Copy()
		assert.Equal(t, cs.receiver, got, cs.msg)
		assert.True(t, &cs.receiver != &got, cs.msg)
	}
}

func TestMap_Entries(t *testing.T) {
	cases := []struct {
		msg      string
		receiver MapT0T1
		want     SetOfPairT0T1
	}{
		{"Entries: non-empty map", mBase(), SetOfPairT0T1{
			PairT0T1{1, "w1"}: true, PairT0T1{22, "w22"}: true, PairT0T1{333, "w333"}: true,
			PairT0T1{4444, "w4444"}: true}},
		{"Entries: empty map", MapT0T1{}, SetOfPairT0T1{}},
		{"Entries: nil map", nil, SetOfPairT0T1{}},
	}

	for _, cs := range cases {
		got := cs.receiver.Entries()
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_Keys(t *testing.T) {
	cases := []struct {
		msg      string
		receiver MapT0T1
		want     SetT0
	}{
		{"Keys: non-empty map", mBase(), SetT0{1: true, 22: true, 333: true, 4444: true}},
		{"Keys: empty map", MapT0T1{}, SetT0{}},
		{"Keys: nil map", nil, SetT0{}},
	}

	for _, cs := range cases {
		got := cs.receiver.Keys()
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_LengthSize(t *testing.T) {
	cases := []struct {
		msg      string
		receiver MapT0T1
		want     int
	}{
		{"Length and Size: non-empty map", mBase(), 4},
		{"Length and Size: empty map", MapT0T1{}, 0},
		{"Length and Size: nil map", nil, 0},
	}

	for _, cs := range cases {
		got1 := cs.receiver.Length()
		assert.Equal(t, cs.want, got1, cs.msg)
		got2 := cs.receiver.Size()
		assert.Equal(t, cs.want, got2, cs.msg)
	}
}

func TestMap_Values(t *testing.T) {
	cases := []struct {
		msg      string
		receiver MapT0T1
		want     SetT1
	}{
		{"Values: non-empty map", mBase(), SetT1{"w1": true, "w22": true, "w333": true,
			"w4444": true}},
		{"Values: empty map", MapT0T1{}, SetT1{}},
		{"Values: nil map", nil, SetT1{}},
	}

	for _, cs := range cases {
		got := cs.receiver.Values()
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_ContainsKey(t *testing.T) {
	cases := []struct {
		msg      string
		receiver MapT0T1
		arg      int
		want     bool
	}{
		{"CotainsKey: present", mBase(), 22, true},
		{"ContainsKey: absent", mBase(), 0, false},
		{"ContainsKey: empty map", MapT0T1{}, 22, false},
		{"ContainsKey: nil map", nil, 22, false},
	}

	for _, cs := range cases {
		got := cs.receiver.ContainsKey(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_ContainsValue(t *testing.T) {
	cases := []struct {
		msg      string
		receiver MapT0T1
		arg      string
		want     bool
	}{
		{"ContainsValue: present", mBase(), "w22", true},
		{"ContainsValue: absent", mBase(), "w0", false},
		{"ContainsValue: empty map", MapT0T1{}, "w22", false},
		{"ContainsValue: nil map", nil, "w22", false},
	}

	for _, cs := range cases {
		got := cs.receiver.ContainsValue(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_Count(t *testing.T) {
	pred1 := func(a PairT0T1) bool { return toInt(a.X1) > 0 }
	pred2 := func(a PairT0T1) bool { return toInt(a.X1)%2 == 0 }
	pred3 := func(a PairT0T1) bool { return toInt(a.X1) < 0 }

	cases := []struct {
		msg      string
		receiver MapT0T1
		arg      func(PairT0T1) bool
		want     int
	}{
		{"Count: pred matches all", mBase(), pred1, len(mBase())},
		{"Count: pred matches some", mBase(), pred2, 2},
		{"Count: pred matches none", mBase(), pred3, 0},
		{"Count: empty receiver", MapT0T1{}, pred2, 0},
		{"Count: nil receiver", nil, pred2, 0},
	}

	for _, cs := range cases {
		got := cs.receiver.Count(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_Get(t *testing.T) {
	cases := []struct {
		msg      string
		receiver MapT0T1
		arg      int
		want     T1
		wok      bool
	}{
		{"Get: key exists", mBase(), 22, mBase()[22], true},
		{"Get: key doesn't exist", mBase(), 0, mBase()[0], false},
		{"Get: empty map", MapT0T1{}, 22, "", false},
		{"Get: nil map", nil, 2, "", false},
	}

	for _, cs := range cases {
		got, ok := cs.receiver.Get(cs.arg)
		assert.Equal(t, cs.wok, ok, cs.msg)
		if ok {
			assert.Equal(t, cs.want, got, cs.msg)
		}
	}
}

func TestMap_IsEmpty(t *testing.T) {
	cases := []struct {
		msg      string
		receiver MapT0T1
		want     bool
	}{
		{"IsEmpty: non-empty", mBase(), false},
		{"IsEmpty: empty", MapT0T1{}, true},
		{"IsEmpty: nil", nil, true},
	}

	for _, cs := range cases {
		got := cs.receiver.IsEmpty()
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_All(t *testing.T) {
	pred1 := func(a PairT0T1) bool { return toInt(a.X1) > 0 }
	pred2 := func(a PairT0T1) bool { return toInt(a.X1)%2 == 0 }
	pred3 := func(a PairT0T1) bool { return toInt(a.X1) < 0 }

	cases := []struct {
		msg      string
		receiver MapT0T1
		arg      func(PairT0T1) bool
		want     bool
	}{
		{"All: pred matches all", mBase(), pred1, true},
		{"All: pred matches some", mBase(), pred2, false},
		{"All: pred matches none", mBase(), pred3, false},
		{"All: empty receiver", MapT0T1{}, pred2, true},
		{"All: nil receiver", nil, pred2, true},
	}

	for _, cs := range cases {
		got := cs.receiver.All(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_Any(t *testing.T) {
	pred1 := func(a PairT0T1) bool { return toInt(a.X1) > 0 }
	pred2 := func(a PairT0T1) bool { return toInt(a.X1)%2 == 0 }
	pred3 := func(a PairT0T1) bool { return toInt(a.X1) < 0 }

	cases := []struct {
		msg      string
		receiver MapT0T1
		arg      func(PairT0T1) bool
		want     bool
	}{
		{"Any: pred matches all", mBase(), pred1, true},
		{"Any: pred matches some", mBase(), pred2, true},
		{"Any: pred matches none", mBase(), pred3, false},
		{"Any: empty receiver", MapT0T1{}, pred2, false},
		{"Any: nil receiver", nil, pred2, false},
	}

	for _, cs := range cases {
		got := cs.receiver.Any(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_ToSlice(t *testing.T) {
	pairsBase := SetOfPairT0T1{
		PairT0T1{1, "w1"}: true, PairT0T1{22, "w22"}: true, PairT0T1{333, "w333"}: true,
		PairT0T1{4444, "w4444"}: true,
	}

	toSet := func(sp SliceOfPairT0T1) SetOfPairT0T1 {
		if sp == nil {
			return nil
		}
		set := make(SetOfPairT0T1, len(sp))
		for _, x := range sp {
			set[x] = true
		}
		return set
	}

	cases := []struct {
		msg      string
		receiver MapT0T1
		want     SetOfPairT0T1
	}{
		{"ToSlice: non-empty", mBase(), pairsBase},
		{"ToSlice: empty", MapT0T1{}, SetOfPairT0T1{}},
		{"ToSlice: nil", nil, nil},
	}

	for _, cs := range cases {
		got := toSet(cs.receiver.ToSlice())
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_Filter(t *testing.T) {
	pred1 := func(a PairT0T1) bool { return toInt(a.X1) > 0 }
	pred2 := func(a PairT0T1) bool { return toInt(a.X1)%2 == 0 }
	pred3 := func(a PairT0T1) bool { return toInt(a.X1) < 0 }

	cases := []struct {
		msg      string
		receiver MapT0T1
		arg      func(PairT0T1) bool
		want     MapT0T1
	}{
		{"Filter: pred matches all", mBase(), pred1, mBase()},
		{"Filter: pred matches some", mBase(), pred2, MapT0T1{22: "w22", 4444: "w4444"}},
		{"Filter: pred matches none", mBase(), pred3, MapT0T1{}},
		{"Filter: empty receiver", MapT0T1{}, pred2, MapT0T1{}},
		{"Filter: nil receiver", nil, pred2, MapT0T1{}},
	}

	for _, cs := range cases {
		got := cs.receiver.Filter(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_FilterKeys(t *testing.T) {
	pred1 := func(a T0) bool { return toInt(a) > 0 }
	pred2 := func(a T0) bool { return toInt(a)%2 == 0 }
	pred3 := func(a T0) bool { return toInt(a) < 0 }

	cases := []struct {
		msg      string
		receiver MapT0T1
		arg      func(T0) bool
		want     MapT0T1
	}{
		{"FilterKeys: pred matches all", mBase(), pred1, mBase()},
		{"FilterKeys: pred matches some", mBase(), pred2, MapT0T1{22: "w22", 4444: "w4444"}},
		{"FilterKeys: pred matches none", mBase(), pred3, MapT0T1{}},
		{"FilterKeys: empty receiver", MapT0T1{}, pred2, MapT0T1{}},
		{"FilterKeys: nil receiver", nil, pred2, MapT0T1{}},
	}

	for _, cs := range cases {
		got := cs.receiver.FilterKeys(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_FilterNot(t *testing.T) {
	pred1 := func(a PairT0T1) bool { return toInt(a.X1) > 0 }
	pred2 := func(a PairT0T1) bool { return toInt(a.X1)%2 == 1 }
	pred3 := func(a PairT0T1) bool { return toInt(a.X1) < 0 }

	cases := []struct {
		msg      string
		receiver MapT0T1
		arg      func(PairT0T1) bool
		want     MapT0T1
	}{
		{"FilterNot: pred matches all", mBase(), pred1, MapT0T1{}},
		{"FilterNot: pred matches some", mBase(), pred2, MapT0T1{22: "w22", 4444: "w4444"}},
		{"FilterNot: pred matches none", mBase(), pred3, mBase()},
		{"FilterNot: empty receiver", MapT0T1{}, pred2, MapT0T1{}},
		{"FilterNot: nil receiver", nil, pred2, MapT0T1{}},
	}

	for _, cs := range cases {
		got := cs.receiver.FilterNot(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_FilterValues(t *testing.T) {
	pred1 := func(a T1) bool { return true }
	pred2 := func(a T1) bool { return len(toString(a))%2 == 1 }
	pred3 := func(a T1) bool { return false }

	cases := []struct {
		msg      string
		receiver MapT0T1
		arg      func(T0) bool
		want     MapT0T1
	}{
		{"FilterValues: pred matches all", mBase(), pred1, mBase()},
		{"FilterValues: pred matches some", mBase(), pred2, MapT0T1{22: "w22", 4444: "w4444"}},
		{"FilterValues: pred matches none", mBase(), pred3, MapT0T1{}},
		{"FilterValues: empty receiver", MapT0T1{}, pred2, MapT0T1{}},
		{"FilterValues: nil receiver", nil, pred2, MapT0T1{}},
	}

	for _, cs := range cases {
		got := cs.receiver.FilterValues(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_ForEach(t *testing.T) {
	cases := []struct {
		msg      string
		receiver MapT0T1
		want     SliceT0
	}{
		{"ForEach: non-empty receiver", mBase(), SliceT0{1, 22, 333, 4444}},
		{"ForEach: empty receiver", MapT0T1{}, SliceT0{}},
		{"ForEach: nil receiver", nil, SliceT0{}},
	}

	for _, cs := range cases {
		got := SliceT0{}
		f := func(a PairT0T1) {
			got = append(got, a.X1)
		}

		cs.receiver.ForEach(f)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_GetOrElse(t *testing.T) {
	f := func(a T0) T1 { return fmt.Sprintf("%v", a) }

	cases := []struct {
		msg      string
		receiver MapT0T1
		arg      T0
		want     T1
	}{
		{"GetOrElse: key prsent", mBase(), 22, "w22"},
		{"GetOrElse: key absent", mBase(), 9, "9"},
		{"GetOrElse: empty", MapT0T1{}, 22, "22"},
		{"GetOrElse: nil", nil, 22, "22"},
	}

	for _, cs := range cases {
		got := cs.receiver.GetOrElse(cs.arg, f)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_IsNotEmpty(t *testing.T) {
	cases := []struct {
		msg      string
		receiver MapT0T1
		want     bool
	}{
		{"IsNotEmpty: non-empty", mBase(), true},
		{"IsNotEmpty: empty", MapT0T1{}, false},
		{"IsNotEmpty: nil", nil, false},
	}

	for _, cs := range cases {
		got := cs.receiver.IsNotEmpty()
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_MaxWith(t *testing.T) {
	comp := func(a1 PairT0T1, a2 PairT0T1) int { return toInt(a1.X1) - toInt(a2.X1) }

	cases := []struct {
		msg      string
		receiver MapT0T1
		arg      func(PairT0T1, PairT0T1) int
		want     PairT0T1
		werr     error
	}{
		{"MaxWith: non-empty receiver", mBase(), comp, PairT0T1{4444, "w4444"}, nil},
		{"MaxWith: empty receiver", MapT0T1{}, comp, PairT0T1{}, errors.New("empty or nil map")},
		{"MaxWith: nil receiver", nil, comp, PairT0T1{}, errors.New("empty or nil map")},
	}

	for _, cs := range cases {
		got, err := cs.receiver.MaxWith(cs.arg)
		assert.Equal(t, cs.werr, err, cs.msg)
		if err == nil {
			assert.Equal(t, cs.want, got, cs.msg)
		}
	}
}

func TestMap_MinusKey(t *testing.T) {
	cases := []struct {
		msg      string
		receiver MapT0T1
		arg      T0
		want     MapT0T1
	}{
		{"MinusKey: present", mBase(), 22, MapT0T1{1: "w1", 333: "w333", 4444: "w4444"}},
		{"MinusKey: absent", mBase(), 9, mBase()},
		{"MinusKey: empty map", MapT0T1{}, 22, MapT0T1{}},
		{"MinusKey: nil map", nil, 22, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.MinusKey(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_MinusKeys(t *testing.T) {
	cases := []struct {
		msg      string
		receiver MapT0T1
		arg      SliceT0
		want     MapT0T1
	}{
		{"MinusKeys: subset", mBase(), SliceT0{22, 333}, MapT0T1{1: "w1", 4444: "w4444"}},
		{"MinusKeys: intersects", mBase(), SliceT0{0, 22, 9, 333},
			MapT0T1{1: "w1", 4444: "w4444"}},
		{"MinusKeys: disjoint", mBase(), SliceT0{0, 9, 42}, mBase()},
		{"MinusKeys: empty slice", MapT0T1{}, SliceT0{22, 333}, MapT0T1{}},
		{"MinusKeys: nil slice", nil, SliceT0{22, 333}, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.MinusKeys(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_MinWith(t *testing.T) {
	comp := func(a1 PairT0T1, a2 PairT0T1) int { return -(toInt(a1.X1) - toInt(a2.X1)) }

	cases := []struct {
		msg      string
		receiver MapT0T1
		arg      func(PairT0T1, PairT0T1) int
		want     PairT0T1
		werr     error
	}{
		{"MinWith: non-empty receiver", mBase(), comp, PairT0T1{4444, "w4444"}, nil},
		{"MinWith: empty receiver", MapT0T1{}, comp, PairT0T1{}, errors.New("empty or nil map")},
		{"MinWith: nil receiver", nil, comp, PairT0T1{}, errors.New("empty or nil map")},
	}

	for _, cs := range cases {
		got, err := cs.receiver.MinWith(cs.arg)
		assert.Equal(t, cs.werr, err, cs.msg)
		if err == nil {
			assert.Equal(t, cs.want, got, cs.msg)
		}
	}
}

func TestMap_PlusEntry(t *testing.T) {
	cases := []struct {
		msg      string
		receiver MapT0T1
		arg      PairT0T1
		want     MapT0T1
	}{
		{"PlusEntry: non-empty", MapT0T1{1: "w1", 22: "w22", 4444: "w4444"}, PairT0T1{333, "w333"},
			mBase()},
		{"PlusEntry: empty", MapT0T1{}, PairT0T1{333, "w333"}, MapT0T1{333: "w333"}},
		{"PlusEntry: nil", nil, PairT0T1{333, "w333"}, MapT0T1{333: "w333"}},
	}

	for _, cs := range cases {
		got := cs.receiver.PlusEntry(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_PlusMap(t *testing.T) {
	cases := []struct {
		msg      string
		receiver MapT0T1
		arg      MapT0T1
		want     MapT0T1
	}{
		{"PlusMap: non-empty + non-empty", mBase(), MapT0T1{9: "x9", 333: "x3"},
			MapT0T1{1: "w1", 9: "x9", 22: "w22", 333: "x3", 4444: "w4444"}},
		{"PlusMap: non-empty + empty", mBase(), MapT0T1{}, mBase()},
		{"PlusMap: non-empty + nil", mBase(), nil, mBase()},
		{"PlusMap: empty + non-empty", MapT0T1{}, mBase(), mBase()},
		{"PlusMap: nil + non-empty", nil, mBase(), mBase()},
		{"PlusMap: empty + empty", MapT0T1{}, MapT0T1{}, MapT0T1{}},
		{"PlusMap: empty + nil", MapT0T1{}, nil, MapT0T1{}},
		{"PlusMap: nil + empty", nil, MapT0T1{}, MapT0T1{}},
		{"PlusMap: nil + nil", nil, nil, nil},
	}

	for _, cs := range cases {
		got := cs.receiver.PlusMap(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_PlusSlice(t *testing.T) {
	cases := []struct {
		msg      string
		receiver MapT0T1
		arg      SliceOfPairT0T1
		want     MapT0T1
	}{
		{"PlusSlice: non-empty + non-empty", mBase(), SliceOfPairT0T1{{9, "x9"}, {333, "x3"}},
			MapT0T1{1: "w1", 9: "x9", 22: "w22", 333: "x3", 4444: "w4444"}},
		{"PlusSlice: non-empty + empty", mBase(), SliceOfPairT0T1{}, mBase()},
		{"PlusSlice: empty + non-empty", MapT0T1{}, SliceOfPairT0T1{{9, "x9"}, {333, "x3"}},
			MapT0T1{9: "x9", 333: "x3"}},
		{"PlusSlice: nil + non-empty", nil, SliceOfPairT0T1{{9, "x9"}, {333, "x3"}},
			MapT0T1{9: "x9", 333: "x3"}},
		{"PlusSlice: empty + empty", MapT0T1{}, SliceOfPairT0T1{}, MapT0T1{}},
		{"PlusSlice: empty + nil", MapT0T1{}, nil, MapT0T1{}},
		{"PlusSlice: nil + empty", nil, SliceOfPairT0T1{}, MapT0T1{}},
		{"PlusSlice: nil + nil", nil, nil, MapT0T1{}},
	}

	for _, cs := range cases {
		got := cs.receiver.PlusSlice(cs.arg)
		assert.Equal(t, cs.want, got, cs.msg)
	}
}

func TestMap_Add(t *testing.T) {
	cases := []struct {
		msg      string
		receiver MapT0T1
		arg1     T0
		arg2     T1
		want     MapT0T1
	}{
		{"Add: key present", mBase(), 333, "x3",
			MapT0T1{1: "w1", 22: "w22", 333: "x3", 4444: "w4444"}},
		{"Add: key absent", mBase(), 9, "x9",
			MapT0T1{1: "w1", 9: "x9", 22: "w22", 333: "x3", 4444: "w4444"}},
		{"Add: empty", MapT0T1{}, 333, "w333", MapT0T1{333: "w333"}},
		{"Add: nil", nil, 333, "w333", MapT0T1{333: "w333"}},
	}

	for _, cs := range cases {
		if cs.receiver != nil {
			cs.receiver.Add(cs.arg1, cs.arg2)
			got := cs.receiver
			assert.Equal(t, cs.want, got, cs.msg)
		} else {
			var fPanic assert.PanicTestFunc
			assert.Panics()
		}
	}
}
