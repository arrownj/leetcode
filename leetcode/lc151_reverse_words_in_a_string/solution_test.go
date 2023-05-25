package lc151_reverse_words_in_a_string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, "blue is sky the", Solution("the sky is blue"))
	assert.Equal(t, "world hello", Solution("  hello world  "))
	assert.Equal(t, "example good a", Solution("a good   example"))
}
