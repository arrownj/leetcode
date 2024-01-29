package main

import "math"

type MinStack struct {
	list    []int
	minList []int
}

func Constructor() MinStack {
	return MinStack{
		list:    []int{},
		minList: []int{math.MaxInt64},
	}
}

func (this *MinStack) Push(val int) {
	this.list = append(this.list, val)
	min := this.minList[len(this.minList)-1]
	if val <= min {
		this.minList = append(this.minList, val)
	}
}

func (this *MinStack) Pop() {
	top := this.list[len(this.list)-1]
	this.list = this.list[:len(this.list)-1]
	min := this.minList[len(this.minList)-1]
	if top == min {
		this.minList = this.minList[:len(this.minList)-1]
	}
}

func (this *MinStack) Top() int {
	return this.list[len(this.list)-1]
}

func (this *MinStack) GetMin() int {
	return this.minList[len(this.minList)-1]
}
