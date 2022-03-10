package divide_two_integers

import "testing"

func TestDivideTwoIntegers1(t *testing.T) {
	got := Divide(8, 2)
	want := 4
	if got != want {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestDivideTwoIntegers2(t *testing.T) {
	got := Divide(0, 3)
	want := 0
	if got != want {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestDivideTwoIntegers3(t *testing.T) {
	got := Divide(8, 3)
	want := 2
	if got != want {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestDivideTwoIntegers4(t *testing.T) {
	got := Divide(9, -2)
	want := -4
	if got != want {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestDivideTwoIntegers5(t *testing.T) {
	got := Divide(-9, 2)
	want := -4
	if got != want {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestDivideTwoIntegers6(t *testing.T) {
	got := Divide(-9, -2)
	want := 4
	if got != want {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestDivideTwoIntegers7(t *testing.T) {
	got := Divide(-2147483648, -1)
	want := 2147483647
	if got != want {
		t.Errorf("want: %d, got: %d", want, got)
	}
}
