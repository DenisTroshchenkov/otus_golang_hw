package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type cacheItem struct {
	mapKey Key
	value  interface{}
}

type lruCache struct {
	mu       sync.Mutex
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	mapElem, exists := c.items[key]

	if exists {
		c.queue.MoveToFront(mapElem)
		mapElem.Value = cacheItem{key, value}
	} else {
		if c.queue.Len() == c.capacity {
			lastQueueElem := c.queue.Back()
			lastCacheElem, ok := lastQueueElem.Value.(cacheItem)
			if !ok {
				return false
			}
			delete(c.items, lastCacheElem.mapKey)
			c.queue.Remove(lastQueueElem)
		}
		c.queue.PushFront(cacheItem{key, value})
		c.items[key] = c.queue.Front()
	}

	return exists
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	cacheElem, exists := c.items[key]
	if exists {
		var value interface{}
		if item, ok := cacheElem.Value.(cacheItem); ok {
			value = item.value
		} else {
			return nil, false
		}
		c.queue.MoveToFront(cacheElem)
		return value, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}
