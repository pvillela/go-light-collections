package collections

// Map is the interface for Map operations.
type Map interface {
	Get(k AnyT0) AnyT0
	Put(k AnyT0, v AnyT0)
}

// MakeMap is a factory for Maps.
func MakeMap() Map {
	return MapT01(make(map[AnyT0]AnyT1))
}

func (m MapT01) Get(k AnyT0) AnyT0 {
	return m[k]
}

func (m MapT01) Put(k AnyT0, v AnyT0) {
	m[k] = v
}
