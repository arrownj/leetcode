package main

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func PrintListNode(l *ListNode) {
	for curr := l; curr != nil; curr = curr.Next {
		if curr.Next != nil {
			fmt.Printf("%d -> ", curr.Val)
		} else {
			fmt.Printf("%d", curr.Val)
			fmt.Println()
		}
	}
}

func MakeListNode(val int) *ListNode {
	if val == 0 {
		result := &ListNode{
			Val:  0,
			Next: nil,
		}
		PrintListNode(result)
		return result
	}

	var result *ListNode
	var current *ListNode
	for val > 0 {
		v := val % 10
		node := &ListNode{
			Val:  v,
			Next: nil,
		}
		if result == nil {
			result = node
		}
		if current != nil {
			current.Next = node
		}
		current = node
		val = int(val / 10)
	}

	PrintListNode(result)

	return result
}

func ListNodeValue(l *ListNode) int {
	PrintListNode(l)
	idx := 0
	val := 0
	for l != nil {
		val += l.Val * int((math.Pow(10, float64(idx))))
		l = l.Next
		idx++
	}
	return val
}

func TestSolution1(t *testing.T) {
	assert.Equal(t, 807, ListNodeValue(Solution1(MakeListNode(342), MakeListNode(465))))
	assert.Equal(t, 1, ListNodeValue(Solution1(MakeListNode(0), MakeListNode(1))))
	assert.Equal(t, 101, ListNodeValue(Solution1(MakeListNode(2), MakeListNode(99))))
	assert.Equal(t, 1234567891, ListNodeValue(Solution1(MakeListNode(1234567890), MakeListNode(1))))
}
