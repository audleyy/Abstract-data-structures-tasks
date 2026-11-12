package main

type CacheEntry struct {
	key       int32
	value     int32
	frequency int32
	lastUsed  int32
}

type LFUCache struct {
	entries  []CacheEntry
	capacity int
	size     int
	time     int32
}

func newLFUCache(capacity int) *LFUCache {
	if capacity < 0 {
		capacity = 0
	}

	return &LFUCache{
		entries:  make([]CacheEntry, capacity),
		capacity: capacity,
	}
}

func (cache *LFUCache) findKey(key int32) int {
	for i := 0; i < cache.size; i++ {
		if cache.entries[i].key == key {
			return i
		}
	}
	return -1
}

func (cache *LFUCache) findLeastUsed() int {
	leastUsed := 0

	for i := 1; i < cache.size; i++ {
		if cache.entries[i].frequency < cache.entries[leastUsed].frequency ||
			(cache.entries[i].frequency == cache.entries[leastUsed].frequency &&
				cache.entries[i].lastUsed < cache.entries[leastUsed].lastUsed) {
			leastUsed = i
		}
	}

	return leastUsed
}

func (cache *LFUCache) recordAccess(index int) {
	cache.entries[index].frequency++
	cache.time++
	cache.entries[index].lastUsed = cache.time
}

func (cache *LFUCache) set(key int32, value int32) {
	if cache.capacity == 0 {
		return
	}

	index := cache.findKey(key)
	if index != -1 {
		cache.entries[index].value = value
		cache.recordAccess(index)
		return
	}

	if cache.size < cache.capacity {
		index = cache.size
		cache.size++
	} else {
		index = cache.findLeastUsed()
	}

	cache.time++
	cache.entries[index] = CacheEntry{key, value, 1, cache.time}
}

func (cache *LFUCache) get(key int32) int32 {
	index := cache.findKey(key)
	if index == -1 {
		return -1
	}

	cache.recordAccess(index)
	return cache.entries[index].value
}
