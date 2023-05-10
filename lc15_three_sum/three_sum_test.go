package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSum(t *testing.T) {
	want := [][]int{}
	assert.Equal(t, want, ThreeSum([]int{}))
}
