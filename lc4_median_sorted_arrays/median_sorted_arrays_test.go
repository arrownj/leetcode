package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMedianSortedArrays(t *testing.T) {
	want := 2.0
	got := FindMedianSortedArrays([]int{1, 3}, []int{2})
	assert.Equal(t, want, got)
}
