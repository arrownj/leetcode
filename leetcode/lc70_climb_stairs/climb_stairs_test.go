package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClimbStairs(t *testing.T) {
	want := 1
	got := ClimbStairs(1)
	assert.Equalf(t, want, got, "want: %v, got: %v", want, got)
}

func TestClimbStairs1(t *testing.T) {
	want := 2
	got := ClimbStairs(2)
	assert.Equalf(t, want, got, "want: %v, got: %v", want, got)
}

func TestClimbStairs2(t *testing.T) {
	want := 3
	got := ClimbStairs(3)
	assert.Equalf(t, want, got, "want: %v, got: %v", want, got)
}
