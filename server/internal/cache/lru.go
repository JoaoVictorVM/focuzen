// Package cache provides an in-memory LRU used to shield the YouTube API quota.
package cache

import (
	"container/list"
	"sync"
	"time"
)

type entry[K comparable, V any] struct {
	key       K
	value     V
	expiresAt time.Time
}

// LRU is a thread-safe, fixed-capacity cache with optional per-entry TTL.
// A non-positive ttl disables expiry.
type LRU[K comparable, V any] struct {
	mu       sync.Mutex
	capacity int
	ttl      time.Duration
	ll       *list.List
	items    map[K]*list.Element
	now      func() time.Time
}

// NewLRU builds an LRU with the given capacity (clamped to a minimum of 1) and ttl.
func NewLRU[K comparable, V any](capacity int, ttl time.Duration) *LRU[K, V] {
	if capacity < 1 {
		capacity = 1
	}
	return &LRU[K, V]{
		capacity: capacity,
		ttl:      ttl,
		ll:       list.New(),
		items:    make(map[K]*list.Element, capacity),
		now:      time.Now,
	}
}

// Get returns the cached value for key. Expired entries are evicted and reported
// as a miss.
func (c *LRU[K, V]) Get(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	el, ok := c.items[key]
	if !ok {
		var zero V
		return zero, false
	}

	ent := el.Value.(*entry[K, V])
	if c.expired(ent) {
		c.remove(el)
		var zero V
		return zero, false
	}

	c.ll.MoveToFront(el)
	return ent.value, true
}

// Add inserts or refreshes a value, evicting the least-recently-used entry when
// the capacity is exceeded.
func (c *LRU[K, V]) Add(key K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if el, ok := c.items[key]; ok {
		ent := el.Value.(*entry[K, V])
		ent.value = value
		ent.expiresAt = c.expiry()
		c.ll.MoveToFront(el)
		return
	}

	el := c.ll.PushFront(&entry[K, V]{key: key, value: value, expiresAt: c.expiry()})
	c.items[key] = el

	if c.ll.Len() > c.capacity {
		c.remove(c.ll.Back())
	}
}

func (c *LRU[K, V]) expiry() time.Time {
	if c.ttl <= 0 {
		return time.Time{}
	}
	return c.now().Add(c.ttl)
}

func (c *LRU[K, V]) expired(ent *entry[K, V]) bool {
	return !ent.expiresAt.IsZero() && c.now().After(ent.expiresAt)
}

func (c *LRU[K, V]) remove(el *list.Element) {
	c.ll.Remove(el)
	delete(c.items, el.Value.(*entry[K, V]).key)
}
