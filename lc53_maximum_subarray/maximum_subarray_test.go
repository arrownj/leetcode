package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumSubarray(t *testing.T) {
	want := 1
	got := MaximumSubarray([]int{1})
	assert.Equal(t, want, got)
}

func TestMaximumSubarray1(t *testing.T) {
	want := 6
	got := MaximumSubarray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	assert.Equal(t, want, got)
}

func TestMaximumSubarray2(t *testing.T) {
	want := 23
	got := MaximumSubarray([]int{5, 4, -1, 7, 8})
	assert.Equal(t, want, got)
}
