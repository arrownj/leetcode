package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSymmetricTree(t *testing.T) {
	assert.True(t, IsSymmetricTree(nil))
}

func TestIsSymmetricTree1(t *testing.T) {
	left_2_1 := TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	left_2_2 := TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}

	right_2_3 := TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	right_2_4 := TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}

	left_1_1 := TreeNode{
		Val:   2,
		Left:  &left_2_1,
		Right: &left_2_2,
	}
	right_1_2 := TreeNode{
		Val:   2,
		Left:  &right_2_3,
		Right: &right_2_4,
	}

	root := TreeNode{
		Val:   1,
		Left:  &left_1_1,
		Right: &right_1_2,
	}
	assert.True(t, IsSymmetricTree(&root))
}

func TestIsSymmetricTree2(t *testing.T) {
	left_2_2 := TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}

	right_2_4 := TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}

	left_1_1 := TreeNode{
		Val:   2,
		Left:  nil,
		Right: &left_2_2,
	}
	right_1_2 := TreeNode{
		Val:   2,
		Left:  nil,
		Right: &right_2_4,
	}

	root := TreeNode{
		Val:   1,
		Left:  &left_1_1,
		Right: &right_1_2,
	}
	assert.False(t, IsSymmetricTree(&root))
}
