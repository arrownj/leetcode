package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSameTree(t *testing.T) {
	tree1 := TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	tree2 := TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	got := SameTree(&tree1, &tree2)
	assert.Truef(t, got, "want: true, got: %v", got)
}
