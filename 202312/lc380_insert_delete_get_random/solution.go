package main

import "math/rand"

type RandomizedSet struct {
	values map[int]int
	keys   []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		values: make(map[int]int),
		keys:   []int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.values[val]; ok {
		return false
	}
	this.keys = append(this.keys, val)
	this.values[val] = len(this.keys) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if index, ok := this.values[val]; ok {
		delete(this.values, val)
		this.keys[index] = this.keys[len(this.keys)-1]
		if index != len(this.keys)-1 {
			this.values[this.keys[index]] = index
		}
		this.keys = this.keys[:len(this.keys)-1]
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	index := rand.Intn(len(this.keys))
	return this.keys[index]
}
