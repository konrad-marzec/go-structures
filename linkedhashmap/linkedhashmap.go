package linkedhashmap

import (
	"github.com/konrad-marzec/go-lru-cache/linkedlist"
)

type Map[T comparable] struct {
	capacity int
	list     *linkedlist.List[string]
	table    map[string]T
}

func New[T comparable](capacity int) *Map[T] {
	return &Map[T]{
		table:    make(map[string]T),
		capacity: capacity,
		list:     linkedlist.New[string](),
	}
}

func (m *Map[T]) Get(key string) (T, error) {
	if v, ok := m.table[key]; ok {
		m.list.Remove(key)
		m.list.Insert(key)

		return v, nil
	}

	return *new(T), nil
}

func (m *Map[T]) Set(key string, value T) {
	if _, ok := m.table[key]; ok {
		m.list.Remove(key)
	} else {

		if m.list.Size() >= m.capacity {
			k := m.list.Shift()
			delete(m.table, k.Value())
		}
	}

	m.list.Insert(key)
	m.table[key] = value
}

func (m *Map[T]) List() *linkedlist.List[string] {
	return m.list
}
