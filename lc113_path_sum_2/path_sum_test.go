package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathSum(t *testing.T) {
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
	want := [][]int{}
	got := PathSum(&root, 0)
	assert.Equal(t, want, got)
}

func TestPathSum1(t *testing.T) {
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
	want := [][]int{}
	got := PathSum(&root, 5)
	assert.Equal(t, want, got)
}

func TestPathSum2(t *testing.T) {
	left4_1 := TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}
	right4_2 := TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	left3_1 := TreeNode{
		Val:   11,
		Left:  &left4_1,
		Right: &right4_2,
	}

	left3_2 := TreeNode{
		Val:   13,
		Left:  nil,
		Right: nil,
	}

	left4_3 := TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	right4_4 := TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	right3_3 := TreeNode{
		Val:   4,
		Left:  &left4_3,
		Right: &right4_4,
	}

	left2_1 := TreeNode{
		Val:   4,
		Left:  &left3_1,
		Right: nil,
	}

	right2_2 := TreeNode{
		Val:   8,
		Left:  &left3_2,
		Right: &right3_3,
	}

	root := TreeNode{
		Val:   5,
		Left:  &left2_1,
		Right: &right2_2,
	}

	want := [][]int{[]int{5, 4, 11, 2}, []int{5, 8, 4, 5}}
	got := PathSum(&root, 22)
	assert.Equal(t, want, got)
}
