package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanJump(t *testing.T) {
	got := CanJump2([]int{2, 3, 1, 1, 4})
	assert.True(t, got)
}

func TestCanJump1(t *testing.T) {
	got := CanJump2([]int{3, 2, 1, 0, 4})
	assert.False(t, got)
}
