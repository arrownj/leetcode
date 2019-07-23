package lengthOfLongestSubstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution1(t *testing.T) {
	assert.Equal(t, 3, Solution1("abcabcbb"))
	assert.Equal(t, 1, Solution1("bbbbb"))
	assert.Equal(t, 3, Solution1("pwwkew"))
}
