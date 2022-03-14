package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateParentheses(t *testing.T) {
	want := []string{"()"}
	got := GenerateParentheses(1)
	assert.ElementsMatchf(t, want, got, "want: %v, got: %v", want, got)
}

func TestGenerateParentheses1(t *testing.T) {
	want := []string{"()()", "(())"}
	got := GenerateParentheses(2)
	assert.ElementsMatchf(t, want, got, "want: %v, got: %v", want, got)
}

func TestGenerateParentheses2(t *testing.T) {
	want := []string{"((()))", "(()())", "(())()", "()(())", "()()()"}
	got := GenerateParentheses(3)
	assert.ElementsMatchf(t, want, got, "want: %v, got: %v", want, got)
}
