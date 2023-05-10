package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumTotal(t *testing.T) {
	want := -10
	got := MinimumTotal([][]int{[]int{-10}})
	assert.Equal(t, want, got)
}

func TestMinimumTotal1(t *testing.T) {
	want := 11
	got := MinimumTotal([][]int{
		[]int{2},
		[]int{3, 4},
		[]int{6, 5, 7},
		[]int{4, 1, 8, 3},
	})
	assert.Equal(t, want, got)
}
