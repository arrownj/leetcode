package lengthOfLongestSubstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution4(t *testing.T) {
	assert.Equal(t, 2, Solution4("ab"))
	assert.Equal(t, 3, Solution4("abcabcbb"))
	assert.Equal(t, 1, Solution4("bbbbb"))
	assert.Equal(t, 3, Solution4("pwwkew"))
}
