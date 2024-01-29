package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = reverseList(l1)
	l2 = reverseList(l2)
	curr1, curr2 := l1, l2
	carry := 0
	var previous *ListNode = nil
	var head *ListNode = nil
	for curr1 != nil || curr2 != nil {
		value := carry
		if curr1 != nil {
			value += curr1.Val
			curr1 = curr1.Next
		}
		if curr2 != nil {
			value += curr2.Val
			curr2 = curr2.Next
		}
		carry = value / 10
		value = value % 10
		node := &ListNode{
			Val:  value,
			Next: nil,
		}
		if previous != nil {
			previous.Next = node
		}
		previous = node
		if head == nil {
			head = node
		}
	}
	if carry > 0 {
		node := &ListNode{
			Val:  carry,
			Next: nil,
		}
		if previous != nil {
			previous.Next = node
		}
	}
	return reverseList(head)
}

func reverseList(l *ListNode) *ListNode {
	var previous *ListNode = nil
	curr := l
	for curr != nil {
		next := curr.Next
		curr.Next = previous
		previous = curr
		curr = next
	}
	return previous
}
