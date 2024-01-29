package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFiindClosestElements(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 4}, findClosestElements([]int{1, 2, 3, 4, 5}, 4, 3))
	assert.Equal(t, []int{1, 2, 3, 4}, findClosestElements([]int{1, 2, 3, 4, 5}, 4, -1))
	assert.Equal(t, []int{7, 8, 9, 10}, findClosestElements([]int{6, 7, 8, 9, 10}, 4, 11))
}
