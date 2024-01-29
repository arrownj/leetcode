package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumSquares(t *testing.T) {
	assert.Equal(t, 1, numSquares(1))
	assert.Equal(t, 2, numSquares(2))
	assert.Equal(t, 3, numSquares(3))
	assert.Equal(t, 1, numSquares(4))
	assert.Equal(t, 3, numSquares(12))
	assert.Equal(t, 2, numSquares(13))
}
