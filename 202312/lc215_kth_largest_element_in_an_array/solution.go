package main

import "fmt"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topk(nums []int, k int) int {
	minHeap := Constructor()
	for _, num := range nums {
		minHeap.Push(num)
	}
	for i := 0; i < len(nums)-k; i++ {
		minHeap.Pop()
	}
	return minHeap.Top()
}

type MinHeap struct {
	data []int
}

func Constructor() MinHeap {
	return MinHeap{
		data: []int{},
	}
}

func (this *MinHeap) Push(value int) {
	this.data = append(this.data, value)
	this.up(len(this.data) - 1)
	fmt.Printf("Push %d, %v\n", value, this.data)
}

func (this *MinHeap) Pop() int {
	top := this.data[0]
	this.data[0] = this.data[len(this.data)-1]

	fmt.Printf("Pop1 %d, %v\n", top, this.data)
	this.data = this.data[0 : len(this.data)-1]
	fmt.Printf("Pop2 %d, %v\n", top, this.data)
	this.down(0)
	fmt.Printf("Pop3 %d, %v\n", top, this.data)
	return top
}

func (this *MinHeap) Top() int {
	return this.data[0]
}

func (this *MinHeap) Len() int {
	return len(this.data)
}

func (this *MinHeap) up(index int) {
	parent := (index - 1) / 2
	for parent >= 0 && this.data[parent] > this.data[index] {
		this.data[parent], this.data[index] = this.data[index], this.data[parent]
		index = parent
		parent = (index - 1) / 2
	}
}

func (this *MinHeap) down(index int) {
	for {
		left := 2*index + 1
		right := 2*index + 2
		if left >= len(this.data) {
			break
		}
		replacing := left
		if right < len(this.data) && this.data[left] > this.data[right] {
			replacing = right
		}
		if this.data[index] <= this.data[replacing] {
			break
		}
		this.data[index], this.data[replacing] = this.data[replacing], this.data[index]
		index = replacing
	}
}
