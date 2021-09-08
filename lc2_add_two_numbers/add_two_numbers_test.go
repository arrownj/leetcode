package add_two_numbers

import "testing"

func TestAddTwoNumbers1(t *testing.T) {
	a := GetNode(243)
	b := GetNode(564)

	sum := AddTwoNumbers(a, b)
	if GetValue(sum) != 807 {
		t.Errorf("want %d, got %d", 807, GetValue(sum))
	}
}

func TestAddTwoNumbers2(t *testing.T) {
	a := GetNode(9999999)
	b := GetNode(9999)
	if GetValue(a) != 9999999 {
		t.Errorf("want %d, got %d", 9999999, GetValue(a))
	}
	if GetValue(b) != 9999 {
		t.Errorf("want %d, got %d", 9999, GetValue(b))
	}

	sum := AddTwoNumbers(a, b)
	if GetValue(sum) != 10009998 {
		t.Errorf("want %d, got %d", 10009998, GetValue(sum))
	}
}
