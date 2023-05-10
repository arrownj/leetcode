package add_two_numbers

import (
	"math"
)

type Node struct {
	Val  int
	Next *Node
}

func AddTwoNumbers(node1, node2 *Node) *Node {
	carry := 0
	var current, header *Node
	for true {
		val := node1.Val + node2.Val + carry
		carry = val / 10
		val = val % 10
		node := &Node{
			Val:  val,
			Next: nil,
		}
		if header == nil {
			header = node
		}
		if current != nil {
			current.Next = node
		}
		current = node

		node1 = node1.Next
		node2 = node2.Next

		if node1 == nil || node2 == nil {
			break
		}
	}

	for node1 != nil {
		val := node1.Val + carry
		carry = val / 10
		val = val % 10
		node := &Node{
			Val:  val,
			Next: nil,
		}
		if current != nil {
			current.Next = node
		}
		current = node

		node1 = node1.Next
	}

	for node2 != nil {
		val := node2.Val + carry
		carry = val / 10
		val = val % 10
		node := &Node{
			Val:  val,
			Next: nil,
		}
		if current != nil {
			current.Next = node
		}
		current = node

		node2 = node2.Next
	}
	if carry != 0 {
		node := &Node{
			Val:  carry,
			Next: nil,
		}
		if current != nil {
			current.Next = node
		}
	}
	return header
}

func GetValue(node *Node) int {
	value := 0
	length := 0
	curr := node
	for curr != nil {
		value += curr.Val * int(math.Pow10(length))
		length += 1
		curr = curr.Next
	}
	return value
}

func GetNode(value int) *Node {
	if value == 0 {
		return &Node{
			Val:  0,
			Next: nil,
		}
	}

	var current, header *Node
	for value != 0 {
		node := &Node{
			Val:  value % 10,
			Next: nil,
		}
		if header == nil {
			header = node
		}
		if current != nil {
			current.Next = node
		}
		current = node
		value = value / 10
	}
	return header
}
