package lengthOfLongestSubstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution2(t *testing.T) {
	assert.Equal(t, 3, Solution2("abcabcbb"))
	assert.Equal(t, 1, Solution2("bbbbb"))
	assert.Equal(t, 3, Solution2("pwwkew"))
	assert.Equal(t, 1, Solution2("aa"))
	assert.Equal(t, 0, Solution2(""))
}
