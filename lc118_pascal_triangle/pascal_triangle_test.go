package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPascalTriangle(t *testing.T) {
	want := [][]int{[]int{1}}
	got := PascalTriangle(1)
	assert.Equal(t, want, got)
}

func TestPascalTriangle1(t *testing.T) {
	want := [][]int{
		[]int{1},
		[]int{1, 1},
		[]int{1, 2, 1},
		[]int{1, 3, 3, 1},
		[]int{1, 4, 6, 4, 1},
	}
	got := PascalTriangle(5)
	assert.Equal(t, want, got)
}
