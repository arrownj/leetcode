package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func Solution1(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	var result *ListNode
	var current *ListNode
	for l1 != nil && l2 != nil {
		val := l1.Val + l2.Val + carry
		if val >= 10 {
			carry = int(val / 10)
			val = val % 10
		} else {
			carry = 0
		}
		node := &ListNode{
			Val:  val,
			Next: nil,
		}
		if result == nil {
			result = node
		}
		if current != nil {
			current.Next = node
		}
		current = node
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		val := l1.Val + carry
		if val >= 10 {
			carry = int(val / 10)
			val = val % 10
		} else {
			carry = 0
		}
		node := &ListNode{
			Val:  val,
			Next: nil,
		}
		current.Next = node
		current = node
		l1 = l1.Next
	}

	for l2 != nil {
		val := l2.Val + carry
		if val >= 10 {
			carry = int(val / 10)
			val = val % 10
		} else {
			carry = 0
		}
		node := &ListNode{
			Val:  val,
			Next: nil,
		}
		current.Next = node
		current = node
		l2 = l2.Next
	}

	if carry > 0 {
		node := &ListNode{
			Val:  carry,
			Next: nil,
		}
		current.Next = node
	}
	return result
}
