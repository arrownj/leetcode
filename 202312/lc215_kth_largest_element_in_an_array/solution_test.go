package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopk(t *testing.T) {
	//assert.Equal(t, 4, topk([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
	assert.Equal(t, 2, topk([]int{-1, 2, 0}, 1))
}
