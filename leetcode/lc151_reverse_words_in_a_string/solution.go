package lc151_reverse_words_in_a_string

import (
	"github.com/golang-collections/collections/stack"
)

func Solution(s string) string {
	stk1 := stack.New()
	stk2 := stack.New()
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c != ' ' {
			stk1.Push(c)
		} else {
			if stk1.Len() > 0 && stk1.Peek() != ' ' {
				stk1.Push(' ')
			}
		}
	}
	output := ""
	if stk1.Peek() == ' ' {
		stk1.Pop()
	}
	for stk1.Len() > 0 {
		c := stk1.Pop()
		if c != ' ' {
			stk2.Push(c)
		} else {
			for stk2.Len() > 0 {
				popChar := stk2.Pop().(byte)
				output += string(popChar)
			}
			output += " "
		}
	}
	for stk2.Len() > 0 {
		output += string(stk2.Pop().(byte))
	}
	return output
}
