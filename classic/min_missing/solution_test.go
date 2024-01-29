package min_missing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, 3, Solution([]int{1, 2, 4, 5, 6}))
	assert.Equal(t, 4, Solution([]int{1, 2, 0, 3, -1}))
	assert.Equal(t, 1, Solution([]int{4, 3, 5}))
	assert.Equal(t, 2, Solution([]int{1, 3, 5, 1000010109, 19085390, -1, -2}))
}
