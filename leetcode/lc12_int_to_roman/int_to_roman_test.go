package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntToRoman(t *testing.T) {
	want := "III"
	got := IntToRoman(3)
	assert.Equalf(t, want, got, "want: %v, got: %v", want, got)
}

func TestIntToRoman1(t *testing.T) {
	want := "LVIII"
	got := IntToRoman(58)
	assert.Equalf(t, want, got, "want: %v, got: %v", want, got)
}

func TestIntToRoman2(t *testing.T) {
	want := "MCMXCIV"
	got := IntToRoman(1994)
	assert.Equalf(t, want, got, "want: %v, got: %v", want, got)
}
