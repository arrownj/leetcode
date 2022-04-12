package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountNumber(t *testing.T) {
	got := CountNumber2(0)
	assert.Equal(t, 1, got)
}

func TestCountNumber1(t *testing.T) {
	got := CountNumber2(2)
	assert.Equal(t, 91, got)
}
