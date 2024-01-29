package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	curr := head
	var previous *ListNode = nil
	for curr != nil {
		first := curr
		second := curr.Next
		if second == nil {
			break
		}
		if previous != nil {
			previous.Next = second
		}
		previous = first
		curr = second.Next
		first.Next, second.Next = second.Next, first
		if head == first {
			head = second
		}
	}
	return head
}
