package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var node2 = &TreeNode{
	Val:   2,
	Left:  nil,
	Right: nil,
}
var node3 = &TreeNode{
	Val:   3,
	Left:  nil,
	Right: nil,
}
var node1 = &TreeNode{
	Val:   1,
	Left:  node2,
	Right: node3,
}

func TestSumNumbers(t *testing.T) {
	assert.Equal(t, 25, sumNumbers(node1))
}
