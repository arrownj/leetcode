package lc1431_kids_with_candies

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, []bool{true}, Solution([]int{1}, 1))
	assert.Equal(t, []bool{true, true, true, false, true}, Solution([]int{2, 3, 5, 1, 3}, 3))
	assert.Equal(t, []bool{true, false, false, false, false}, Solution([]int{4, 2, 1, 1, 2}, 1))
	assert.Equal(t, []bool{true, false, true}, Solution([]int{12, 1, 12}, 10))
}
