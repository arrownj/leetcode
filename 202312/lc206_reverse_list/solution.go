package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type Stack struct {
	Queue []*ListNode
}

func NewStack() *Stack {
	return &Stack{
		Queue: []*ListNode{},
	}
}

func (s *Stack) Pop() *ListNode {
	if len(s.Queue) == 0 {
		return nil
	}
	ret := s.Queue[len(s.Queue)-1]
	s.Queue = s.Queue[:len(s.Queue)-1]
	return ret
}

func (s *Stack) Push(node *ListNode) {
	s.Queue = append(s.Queue, node)
}

func (s *Stack) Empty() bool {
	return len(s.Queue) == 0
}

func ReverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	stack := NewStack()
	for curr := head; curr != nil; curr = curr.Next {
		stack.Push(curr)
	}
	retHead := stack.Pop()
	curr := retHead
	for !stack.Empty() {
		curr.Next = stack.Pop()
		curr = curr.Next
	}
	curr.Next = nil
	return retHead
}
