package lc1071_greatest_common_divisor_of_strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMuliplay(t *testing.T) {
	assert.Equal(t, "abcabc", multiply("abc", 2))
}

func TestSolution(t *testing.T) {
	assert.Equal(t, "ABC", Solution("ABCABC", "ABC"))
	assert.Equal(t, "AB", Solution("ABABAB", "ABAB"))
	assert.Equal(t, "", Solution("LEET", "CODE"))
}
