package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasPathSum(t *testing.T) {
	root := TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}
	got := HasPathSum(&root, 0)
	assert.Truef(t, got, "want: true, got: %v", got)
}

func TestHasPathSum1(t *testing.T) {
	left := TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	root := TreeNode{
		Val:   1,
		Left:  &left,
		Right: nil,
	}
	got := HasPathSum(&root, 0)
	assert.Falsef(t, got, "want: false, got: %v", got)
}

func TestHasPathSum2(t *testing.T) {
	left := TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	right := TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	root := TreeNode{
		Val:   2,
		Left:  &left,
		Right: &right,
	}
	got := HasPathSum(&root, 1)
	assert.Falsef(t, got, "want: false, got: %v", got)
}
