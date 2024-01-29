package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		levelLength := len(queue)
		for i := 0; i < levelLength; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			if i != levelLength-1 {
				queue[i].Next = queue[i+1]
			}
		}
		queue = queue[levelLength:]
	}
	return root
}
