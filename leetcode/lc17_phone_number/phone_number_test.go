package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLetterCombinations(t *testing.T) {
	want := []string{}
	got := LetterCombinations("")
	assert.Equal(t, want, got)
}

func TestLetterCombinations1(t *testing.T) {
	want := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
	got := LetterCombinations("23")
	assert.Equal(t, want, got)
}
