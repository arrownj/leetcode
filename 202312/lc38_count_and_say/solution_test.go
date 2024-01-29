package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountAndSay(t *testing.T) {
	assert.Equal(t, "1", countAndSay(1))
	assert.Equal(t, "11", countAndSay(2))
}
