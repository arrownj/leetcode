package twoSum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution2(t *testing.T) {
	assert.Equal(t, []int{0, 1}, Solution2([]int{3, 3}, 6))
	assert.Equal(t, []int{0, 2}, Solution2([]int{2, 3, 4}, 6))
	assert.Equal(t, []int{0, 1}, Solution2([]int{2, 7, 11, 13}, 9))
}
