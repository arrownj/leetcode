package lc443_string_compression

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, 6, Solution([]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}))
	assert.Equal(t, 1, Solution([]byte{'a'}))
	assert.Equal(t, 4, Solution([]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}))
}
