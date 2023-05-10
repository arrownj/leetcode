package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrap(t *testing.T) {
	want := 0
	got := Trap([]int{1})
	assert.Equal(t, want, got)
}
