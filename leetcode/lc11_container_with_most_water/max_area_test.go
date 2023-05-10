package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxArea(t *testing.T) {
	assert.Equal(t, 1, MaxArea([]int{1, 1}))
}

func TestMaxArea1(t *testing.T) {
	assert.Equal(t, 49, MaxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
