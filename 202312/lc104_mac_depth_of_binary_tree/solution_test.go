package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var node1 = &TreeNode{
	Val:   1,
	Left:  nil,
	Right: nil,
}

var node2 = &TreeNode{
	Val:   2,
	Left:  nil,
	Right: nil,
}

var node3 = &TreeNode{
	Val:   3,
	Left:  node1,
	Right: node2,
}

var node4 = &TreeNode{
	Val:   4,
	Left:  nil,
	Right: nil,
}

var node5 = &TreeNode{
	Val:   5,
	Left:  node3,
	Right: node4,
}

func TestMaxDepthRecursive1(t *testing.T) {
	assert.Equal(t, 0, MaxDepthRecursive(nil))
}

func TestMaxDepthRecursive2(t *testing.T) {
	assert.Equal(t, 2, MaxDepthRecursive(node3))
}

func TestMaxDepthRecursive3(t *testing.T) {
	assert.Equal(t, 3, MaxDepthRecursive(node5))
}

//BFS
func TestMaxDepthBFS1(t *testing.T) {
	assert.Equal(t, 0, MaxDepthBFS(nil))
}

func TestMaxDepthBFS2(t *testing.T) {
	assert.Equal(t, 2, MaxDepthBFS(node3))
}

func TestMaxDepthBFS3(t *testing.T) {
	assert.Equal(t, 3, MaxDepthBFS(node5))
}
