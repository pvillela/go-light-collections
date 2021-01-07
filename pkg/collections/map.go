package collections

// MapAnyAny is a type wrapper, implements Map interface.
type MapAnyAny map[AnyT0]AnyT0

// Map is the interface for Map operations.
type Map interface {
	Get(k AnyT0) AnyT0
	Put(k AnyT0, v AnyT0)
}

// MakeMap is a factory for Maps.
func MakeMap() Map {
	return MapAnyAny(make(map[AnyT0]AnyT0))
}

func (m MapAnyAny) Get(k AnyT0) AnyT0 {
	return m[k]
}

func (m MapAnyAny) Put(k AnyT0, v AnyT0) {
	m[k] = v
}
