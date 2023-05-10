package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanToInt(t *testing.T) {
	want := 1
	got := RomanToInt("I")
	assert.Equalf(t, want, got, "want: %v, got: %v", want, got)
}

func TestRomanToInt1(t *testing.T) {
	want := 3
	got := RomanToInt("III")
	assert.Equalf(t, want, got, "want: %v, got: %v", want, got)
}

func TestRomanToInt2(t *testing.T) {
	want := 58
	got := RomanToInt("LVIII")
	assert.Equalf(t, want, got, "want: %v, got: %v", want, got)
}

func TestRomanToInt3(t *testing.T) {
	want := 1994
	got := RomanToInt("MCMXCIV")
	assert.Equalf(t, want, got, "want: %v, got: %v", want, got)
}
