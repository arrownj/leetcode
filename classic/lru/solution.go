package lru

import (
	"errors"
)

type LRUCache struct {
	Capacity int
	Keys     []string
	Values   map[string]string
}

func NewLRUCache(capacity int) *LRUCache {
	keys := make([]string, 0, capacity)
	values := make(map[string]string)
	return &LRUCache{
		Capacity: capacity,
		Values:   values,
	}
}

func (l *LRUCache) Get(key string) (string, error) {
	if value, ok := l.Values[key]; ok {
		l.resort(key)
		return value, nil
	}
	return "", errors.New("not in cache")
}

func (l *LRUCache) Put(key, value string) {
	if len(l.Keys) == l.Capaticy {
		toBeDeleteKey := l.Keys[len(l.Keys)-1]
		delete(l.Values, toBeDeleteKey)
	}
	l.resort(key)
	l.Values[key] = value
}

func (l *LRUCache) resort(key string) {
}
