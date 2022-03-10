package generate_parentheses

import (
	"reflect"
	"testing"
	"github.com/deckarep/golang-set"
)

func TestGenerateParentheses(t *testing.T) {
	got := GenerateParentheses(1)
	want := []string{"()"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v, got: %v", want, got)
	}
}

func TestGenerateParentheses1(t *testing.T) {
	got := mapset.NewSetFromSlice(GenerateParentheses(2))
	want := mapset.NewSetFromSlice([]string{"()()", "(())"})
	if !want.Equal(got) {
		t.Errorf("want: %v, got: %v", want, got)
	}
}

func TestGenerateParentheses2(t *testing.T) {
	got := GenerateParentheses(3)
	want := []string{"((()))", "(()())", "(())()", "()(())", "()()()"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v, got: %v", want, got)
	}
}
