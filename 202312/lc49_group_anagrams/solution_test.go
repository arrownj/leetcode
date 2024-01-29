package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupAnagrams(t *testing.T) {
	assert.Equal(t, [][]string{[]string{"a"}}, groupAnagrams([]string{"a"}))
	assert.Equal(t, [][]string{[]string{"bat"}, []string{"nat", "tan"}, []string{"ate", "eat", "tea"}}, groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
