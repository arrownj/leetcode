package lc238_product_of_array_except_self

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, []int{24, 12, 8, 6}, Solution([]int{1, 2, 3, 4}))
	assert.Equal(t, []int{0, 0, 9, 0, 0}, Solution([]int{-1, 1, 0, -3, 3}))
}
