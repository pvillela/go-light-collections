package collections

// IMap is the interface for IMap operations.
type IMap interface {
	Get(k T0) T0
	Put(k T0, v T0)
}

// MakeMap is a factory for Maps.
func MakeMap() IMap {
	return MapT0T1(make(map[T0]T1))
}

func (m MapT0T1) Get(k T0) T0 {
	return m[k]
}

func (m MapT0T1) Put(k T0, v T0) {
	m[k] = v
}
