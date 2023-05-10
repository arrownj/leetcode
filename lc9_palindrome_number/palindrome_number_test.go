package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	assert.True(t, IsPalindrome(0))
}

func TestIsPalindrome1(t *testing.T) {
	assert.True(t, IsPalindrome(121))
}

func TestIsPalindrome2(t *testing.T) {
	assert.False(t, IsPalindrome(-121))
}

func TestIsPalindrome3(t *testing.T) {
	assert.False(t, IsPalindrome(10))
}
