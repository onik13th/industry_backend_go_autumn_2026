package main

type Cache[K comparable, V any] struct {
	capacity int
	items    map[K]V
}

func NewCache[K comparable, V any](capacity int) *Cache[K, V] {
	return &Cache[K, V]{
		capacity: capacity,
		items:    make(map[K]V),
	}
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
	val, found := c.items[key]

	return val, found
}

func (c *Cache[K, V]) Set(key K, value V) bool {
	if c.capacity <= 0 {
		return false
	}

	if _, exists := c.items[key]; !exists && len(c.items) >= c.capacity {
		return false
	}

	c.items[key] = value

	return true
}
