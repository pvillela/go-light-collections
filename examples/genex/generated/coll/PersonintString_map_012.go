// Code generated -- DO NOT EDIT.

package coll

func (m MapPersonint) FlatMapString(f func(PairPersonint) SliceString) SliceString {
	if m == nil {
		return nil
	}
	r := make(SliceString, 0, len(m)) // optimizing for speed vs space
	for k, v := range m {
		r = append(r, f(PairPersonint{k, v})...)
	}
	return r
}

func (m MapPersonint) MapString(f func(PairPersonint) String) SliceString {
	if m == nil {
		return nil
	}
	r := make(SliceString, 0, len(m)) // optimizing for speed vs space
	for k, v := range m {
		r = append(r, f(PairPersonint{k, v}))
	}
	return r
}

func (m MapPersonint) MapKeysString(f func(PairPersonint) String) MapStringint {
	if m == nil {
		return nil
	}
	r := make(MapStringint)
	for k, v := range m {
		r[f(PairPersonint{k, v})] = v
	}
	return r
}

func (m MapPersonint) MapValuesString(f func(PairPersonint) String) MapPersonString {
	if m == nil {
		return nil
	}
	r := make(MapPersonString)
	for k, v := range m {
		r[k] = f(PairPersonint{k, v})
	}
	return r
}
