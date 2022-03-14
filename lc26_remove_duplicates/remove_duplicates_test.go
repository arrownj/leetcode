package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates1(t *testing.T) {
	want := 2
	got := RemoveDuplicates([]int{0, 0, 1})
	assert.Equalf(t, want, got, "want: %v, got: %v", want, got)
}

func TestRemoveDuplicates2(t *testing.T) {
	want := 5
	got := RemoveDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	assert.Equalf(t, want, got, "want: %v, got: %v", want, got)
}

func TestRemoveDuplicates3(t *testing.T) {
	want := 0
	got := RemoveDuplicates([]int{})
	assert.Equalf(t, want, got, "want: %v, got: %v", want, got)
}
