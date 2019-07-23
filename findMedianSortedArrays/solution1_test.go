package findMedianSortedArrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution1(t *testing.T) {
	assert.Equal(t, float64(2), Solution1([]int{1, 3}, []int{2}))
	assert.Equal(t, 2.5, Solution1([]int{1, 2}, []int{3, 4}))
}
