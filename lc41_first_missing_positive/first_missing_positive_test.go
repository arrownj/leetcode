package main

import "testing"

func TestFirstMissingPositive1(t *testing.T) {
	want := 3
	nums := []int{1, 2, 0}
	got := FirstMissingPositive(nums)
	if want != got {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestFirstMissingPositive2(t *testing.T) {
	want := 2
	nums := []int{3, 4, -1, 1}
	got := FirstMissingPositive(nums)
	if want != got {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestFirstMissingPositive3(t *testing.T) {
	want := 1
	nums := []int{7, 8, 9, 11, 12}
	got := FirstMissingPositive(nums)
	if want != got {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestAnotherFirstMissingPositive1(t *testing.T) {
	want := 3
	nums := []int{1, 2, 0}
	got := AnotherFirstMissingPositive(nums)
	if want != got {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestAnotherFirstMissingPositive2(t *testing.T) {
	want := 2
	nums := []int{3, 4, -1, 1}
	got := AnotherFirstMissingPositive(nums)
	if want != got {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestAnotherFirstMissingPositive3(t *testing.T) {
	want := 1
	nums := []int{7, 8, 9, 11, 12}
	got := AnotherFirstMissingPositive(nums)
	if want != got {
		t.Errorf("want: %d, got: %d", want, got)
	}
}
