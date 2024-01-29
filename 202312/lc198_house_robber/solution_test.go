package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRob1(t *testing.T) {
	assert.Equal(t, 0, Rob([]int{}))
	assert.Equal(t, 1, Rob([]int{1}))
	assert.Equal(t, 4, Rob([]int{1, 2, 3, 1}))
	assert.Equal(t, 12, Rob([]int{2, 7, 9, 3, 1}))
}
