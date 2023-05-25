package lc45_jump_game_2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, 2, Solution([]int{2, 3, 1, 1, 4}))
}
