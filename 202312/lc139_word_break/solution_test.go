package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordBreak(t *testing.T) {
	assert.True(t, wordBreak("leetcode", []string{"leet", "code"}))
	assert.True(t, wordBreak("applepenapple", []string{"apple", "pen"}))
	assert.False(t, wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}
