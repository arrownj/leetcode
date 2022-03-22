package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidBinarySearchTree(t *testing.T) {
	want := true
	got := IsValidBinarySearchTree(nil)
	assert.Equalf(t, want, got, "want: %v, got: %v", want, got)
}

func TestIsValidBinarySearchTree1(t *testing.T) {
	want := true
	root := TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	got := IsValidBinarySearchTree(&root)
	assert.Equalf(t, want, got, "want: %v, got: %v", want, got)
}

func TestIsValidBinarySearchTree2(t *testing.T) {
	want := true
	left := TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	right := TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	root := TreeNode{
		Val:   2,
		Left:  &left,
		Right: &right,
	}
	got := IsValidBinarySearchTree(&root)
	assert.Equalf(t, want, got, "want: %v, got: %v", want, got)
}
