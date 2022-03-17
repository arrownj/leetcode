package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestValidParentheses(t *testing.T) {
	want := 0
	got := LongestValidParentheses("")
	assert.Equalf(t, want, got, "want: %v, got: %v", want, got)
}

func TestLongestValidParentheses1(t *testing.T) {
	want := 2
	got := LongestValidParentheses("(()")
	assert.Equalf(t, want, got, "want: %v, got: %v", want, got)
}

func TestLongestValidParentheses2(t *testing.T) {
	want := 4
	got := LongestValidParentheses(")()())")
	assert.Equalf(t, want, got, "want: %v, got: %v", want, got)
}
