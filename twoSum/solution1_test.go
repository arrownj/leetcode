package twoSum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution1(t *testing.T) {
	assert.Equal(t, []int{0, 1}, Solution1([]int{2, 7, 11, 13}, 9))
}
