package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	var first, second *ListNode = dummy, dummy
	for second.Next != nil {
		second = second.Next
		n--
		if n < 0 {
			first = first.Next
		}
	}
	first.Next = first.Next.Next
	return dummy.Next
}
