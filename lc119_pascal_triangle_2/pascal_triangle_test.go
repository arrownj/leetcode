package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPascalTriangle(t *testing.T) {
	want := []int{1}
	got := PascalTriangle(0)
	assert.Equal(t, want, got)
}

func TestPascalTriangle1(t *testing.T) {
	want := []int{1, 1}
	got := PascalTriangle(1)
	assert.Equal(t, want, got)
}

func TestPascalTriangle2(t *testing.T) {
	want := []int{1, 3, 3, 1}
	got := PascalTriangle(3)
	assert.Equal(t, want, got)
}
