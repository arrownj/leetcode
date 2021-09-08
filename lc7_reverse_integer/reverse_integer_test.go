package reverse_integer

import "testing"

func TestReverseInteger1(t *testing.T) {
	got := ReverseInteger(0)
	want := 0
	if got != want {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestReverseInteger2(t *testing.T) {
	got := ReverseInteger(123)
	want := 321
	if got != want {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestReverseInteger3(t *testing.T) {
	got := ReverseInteger(-123)
	want := -321
	if got != want {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestReverseInteger4(t *testing.T) {
	got := ReverseInteger(120)
	want := 21
	if got != want {
		t.Errorf("want: %d, got: %d", want, got)
	}
}
