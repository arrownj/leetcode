package solution

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func InorderRecursive(root *Node) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}

	result = append(result, InorderRecursive(root.Left)...)
	result = append(result, root.Value)
	result = append(result, InorderRecursive(root.Right)...)
	return result
}

func InorderIterative(root *Node) []int {
	stack := []*Node{}
	result := []int{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, root.Value)
		root = root.Right
	}
	return result
}

func PreorderRecursive(root *Node) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	result = append(result, root.Value)
	result = append(result, PreorderRecursive(root.Left)...)
	result = append(result, PreorderRecursive(root.Right)...)
	return result
}

func PreorderIterative(root *Node) []int {
	stack := []*Node{}
	result := []int{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			result = append(result, root.Value)
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = root.Right
	}
	return result
}

func PostorderRecursive(root *Node) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	result = append(result, PostorderRecursive(root.Left)...)
	result = append(result, PostorderRecursive(root.Right)...)
	result = append(result, root.Value)
	return result
}

func PostorderIterative(root *Node) []int {
	stack := []*Node{}
	result := []int{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			result = append(result, root.Value)
			stack = append(stack, root)
			root = root.Right
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = root.Left
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result

}

func LevelOrder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	queue := []*Node{}
	result := []int{}
	queue = append(queue, root)
	for len(queue) > 0 {
		root = queue[0]
		queue = queue[1:]
		result = append(result, root.Value)
		if root.Left != nil {
			queue = append(queue, root.Left)
		}
		if root.Right != nil {
			queue = append(queue, root.Right)
		}
	}
	return result
}
