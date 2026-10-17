package main

import (
	"container/list"
	"sync"
)

type LRU[K comparable, V any] interface {
	Get(key K) (value V, ok bool)
	Set(key K, value V)
}

type entry[K comparable, V any] struct {
	key   K
	value V
}

type LRUCache[K comparable, V any] struct {
	capacity int
	mu       sync.Mutex
	ll       list.List
	items    map[K]*list.Element
}

func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	return &LRUCache[K, V]{
		capacity: capacity,
		items:    make(map[K]*list.Element),
	}
}

func (c *LRUCache[K, V]) Get(key K) (value V, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	element, exists := c.items[key]
	if !exists {
		var zero V
		return zero, false
	}

	c.ll.MoveToFront(element)

	return element.Value.(*entry[K, V]).value, true
}

func (c *LRUCache[K, V]) Set(key K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.capacity <= 0 {
		return
	}

	element, exists := c.items[key]
	if exists {
		element.Value.(*entry[K, V]).value = value
		return
	}

	if len(c.items) >= c.capacity {
		c.evictOldest()
	}

	c.items[key] = c.ll.PushFront(&entry[K, V]{key: key, value: value})
}

// evictOldest вызывается с захваченным мьютексом
func (c *LRUCache[K, V]) evictOldest() {
	oldest := c.ll.Back()
	item := oldest.Value.(*entry[K, V])
	delete(c.items, item.key)
	c.ll.Remove(oldest)
}
