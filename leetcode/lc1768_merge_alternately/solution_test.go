package lc1768_merge_alternately

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, "apbqcr", Solution("abc", "pqr"))
	assert.Equal(t, "apbqrs", Solution("ab", "pqrs"))
	assert.Equal(t, "apbqcd", Solution("abcd", "pq"))

}
