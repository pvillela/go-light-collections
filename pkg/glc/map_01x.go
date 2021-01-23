package glc

func (m MapT0T1) Entries() map[PairMpT0T1]bool {
	if m == nil {
		return nil
	}
	entries := make(map[PairMpT0T1]bool, len(m))
	for k, v := range m {
		entries[PairMpT0T1{k, v}] = true
	}
	return entries
}

func (m MapT0T1) Values() map[T1]bool {
	if m == nil {
		return nil
	}
	values := make(map[T1]bool, len(m))
	for _, v := range m {
		values[v] = true
	}
	return values
}
