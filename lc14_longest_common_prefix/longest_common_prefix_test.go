package longest_common_prefix

import "testing"

func TestGetLongestCommonPrefix1(t *testing.T) {
	got := GetLongestCommonPrefix([]string{"I", "love", "U"})
	want := ""
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestGetLongestCommonPrefix2(t *testing.T) {
	got := GetLongestCommonPrefix([]string{})
	want := ""
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestGetLongestCommonPrefix3(t *testing.T) {
	got := GetLongestCommonPrefix([]string{"", "I", "love", "U"})
	want := ""
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestGetLongestCommonPrefix4(t *testing.T) {
	got := GetLongestCommonPrefix([]string{"loveI", "love", "loveU"})
	want := "love"
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestGetLongestCommonPrefix5(t *testing.T) {
	got := GetLongestCommonPrefix([]string{"flower", "flow", "flight"})
	want := "fl"
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}
