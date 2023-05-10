package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAtoi(t *testing.T) {
	assert.Equal(t, 0, Atoi(""))
}

func TestAtoi1(t *testing.T) {
	assert.Equal(t, 42, Atoi("42"))
}

func TestAtoi2(t *testing.T) {
	assert.Equal(t, -42, Atoi("   -42"))
}

func TestAtoi3(t *testing.T) {
	assert.Equal(t, 4193, Atoi("4193 with words"))
}
