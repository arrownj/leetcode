package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var node3_1 = &TreeNode{
	Val:   3,
	Left:  nil,
	Right: nil,
}

var node4_1 = &TreeNode{
	Val:   4,
	Left:  nil,
	Right: nil,
}

var node2_1 = &TreeNode{
	Val:   2,
	Left:  node3_1,
	Right: node4_1,
}

var node3_2 = &TreeNode{
	Val:   3,
	Left:  nil,
	Right: nil,
}

var node4_2 = &TreeNode{
	Val:   4,
	Left:  nil,
	Right: nil,
}

var node2_2 = &TreeNode{
	Val:   2,
	Left:  node4_2,
	Right: node3_2,
}

var node1 = &TreeNode{
	Val:   1,
	Left:  node2_1,
	Right: node2_2,
}

func TestIsSymmetric(t *testing.T) {
	assert.True(t, IsSymmetric(nil))
	assert.False(t, IsSymmetric(node2_1))
	assert.True(t, IsSymmetric(node1))
}
