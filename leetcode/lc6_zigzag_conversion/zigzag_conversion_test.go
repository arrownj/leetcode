package zigzag_conversion

import "testing"

func TestConvertZigzag1(t *testing.T) {
	got := ConvertZigzag("THECHOICEISYOURS", 4)
	want := "TIOHOCYUEHESRCIS"
	if got != want {
		t.Errorf("want: %s, got: %s", want, got)
	}
}

func TestConvertZigzag2(t *testing.T) {
	got := ConvertZigzag("PAYPALISHIRING", 3)
	want := "PAHNAPLSIIGYIR"
	if got != want {
		t.Errorf("want: %s, got: %s", want, got)
	}
}

func TestConvertZigzag3(t *testing.T) {
	got := ConvertZigzag("PAYPALISHIRING", 4)
	want := "PINALSIGYAHRPI"
	if got != want {
		t.Errorf("want: %s, got: %s", want, got)
	}
}

func TestConvertZigzag4(t *testing.T) {
	got := ConvertZigzag("A", 1)
	want := "A"
	if got != want {
		t.Errorf("want: %s, got: %s", want, got)
	}
}
