package main

type LRUCache struct {
	capacity int
	cache    map[int]int
	keys     []int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]int),
		keys:     []int{},
	}
}

func (this *LRUCache) Get(key int) int {
	if value, ok := this.cache[key]; ok {
		this.rearrangeKeys(key)
		return value
	}
	return -1
}

func (this *LRUCache) Put(key, value int) {
	this.cache[key] = value
	if value, ok := this.cache[key]; ok {
		this.rearrangeKeys(key)
		this.cache[key] = value
	} else {
		this.cache[key] = value
		if len(this.cache) == this.capacity {
			removeKey := this.keys[0]
			delete(this.cache, removeKey)
			this.keys = this.keys[1:]
		}
		this.keys = append(this.keys, key)
	}
}

func (this *LRUCache) rearrangeKeys(key int) {
	idx := 0
	for idx < len(this.keys) {
		if this.keys[idx] == key {
			break
		}
		idx++
	}
	for i := idx; i < len(this.keys)-1; i++ {
		this.keys[i] = this.keys[i+1]
	}
	this.keys[len(this.keys)-1] = key
}
