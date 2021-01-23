// Code generated -- DO NOT EDIT.

package collections

func (m Mapintstring) Entries() map[PairMpintstring]bool {
	if m == nil {
		return nil
	}
	entries := make(map[PairMpintstring]bool, len(m))
	for k, v := range m {
		entries[PairMpintstring{k, v}] = true
	}
	return entries
}

func (m Mapintstring) Values() map[string]bool {
	if m == nil {
		return nil
	}
	values := make(map[string]bool, len(m))
	for _, v := range m {
		values[v] = true
	}
	return values
}
