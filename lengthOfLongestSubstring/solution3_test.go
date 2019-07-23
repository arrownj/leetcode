package lengthOfLongestSubstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution3(t *testing.T) {
	assert.Equal(t, 3, Solution3("abcabcbb"))
	assert.Equal(t, 1, Solution3("bbbbb"))
	assert.Equal(t, 3, Solution3("pwwkew"))
}
