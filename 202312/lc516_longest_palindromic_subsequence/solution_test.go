package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindromeSubseq(t *testing.T) {
	assert.Equal(t, 4, longestPalindromeSubseq("bbbab"))
	assert.Equal(t, 2, longestPalindromeSubseq("cbbd"))
}
