package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInorderTraversal(t *testing.T) {
	want := []int{}
	got := InorderTraversal(nil)
	assert.Equalf(t, want, got, "want: %v, got: %v", want, got)
}

func TestInorderTraversal1(t *testing.T) {
	want := []int{1}
	root := TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	got := InorderTraversal(&root)
	assert.Equalf(t, want, got, "want: %v, got: %v", want, got)
}

func TestInorderTraversal2(t *testing.T) {
	want := []int{2, 1, 3}
	left := TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	right := TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	root := TreeNode{
		Val:   1,
		Left:  &left,
		Right: &right,
	}
	got := InorderTraversal(&root)
	assert.Equalf(t, want, got, "want: %v, got: %v", want, got)
}
