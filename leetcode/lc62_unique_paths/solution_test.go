package lc62_unique_paths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, 1, Solution(1, 1))
	assert.Equal(t, 28, Solution(3, 7))
	assert.Equal(t, 3, Solution(3, 2))
	assert.Equal(t, 28, Solution(7, 3))
}
