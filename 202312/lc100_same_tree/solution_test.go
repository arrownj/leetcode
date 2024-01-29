package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var node1_2 = &TreeNode{
	Val:   2,
	Left:  nil,
	Right: nil,
}

var node1_3 = &TreeNode{
	Val:   3,
	Left:  nil,
	Right: nil,
}

var node1_1 = &TreeNode{
	Val:   1,
	Left:  node1_2,
	Right: node1_3,
}

var node2_2 = &TreeNode{
	Val:   2,
	Left:  nil,
	Right: nil,
}

var node2_3 = &TreeNode{
	Val:   3,
	Left:  nil,
	Right: nil,
}

var node2_1 = &TreeNode{
	Val:   1,
	Left:  node2_2,
	Right: node2_3,
}

func TestIsSameTreeRecursive(t *testing.T) {
	assert.True(t, isSameTreeRecursive(node1_1, node2_1))
	assert.False(t, isSameTreeRecursive(node1_1, node1_2))
}
