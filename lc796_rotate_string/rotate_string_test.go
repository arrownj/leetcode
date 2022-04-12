package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateString(t *testing.T) {
	assert.True(t, RotateString("", ""))
}

func TestRotateString1(t *testing.T) {
	assert.True(t, RotateString("abcde", "cdeab"))
}

func TestRotateString2(t *testing.T) {
	assert.False(t, RotateString("abcde", "abced"))
}
