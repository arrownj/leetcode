package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiply(t *testing.T) {
	assert.Equal(t, "6", multiply("2", "3"))
	assert.Equal(t, "144", multiply("12", "12"))
	assert.Equal(t, "56088", multiply("123", "456"))
}
