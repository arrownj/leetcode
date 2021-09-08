package longest_palindrome

import "testing"

func TestGetLongestPalindrome1(t *testing.T) {
	got := GetLongestPalindrome("babad")
	want := "bab"
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestGetLongestPalindrome2(t *testing.T) {
	got := GetLongestPalindrome("cbbd")
	want := "bb"
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestGetLongestPalindrome3(t *testing.T) {
	got := GetLongestPalindrome("a")
	want := "a"
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestGetLongestPalindrome4(t *testing.T) {
	got := GetLongestPalindrome("abc")
	want := "a"
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}

}
