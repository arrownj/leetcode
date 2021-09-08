package two_sum

import (
	"reflect"
	"testing"
)

func TestTwoSum1(t *testing.T) {
	got := TwoSum([]int{3, 3}, 6)
	want := []int{0, 1}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestTwoSum2(t *testing.T) {
	got := TwoSum([]int{2, 7, 11, 15}, 9)
	want := []int{0, 1}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestTwoSum3(t *testing.T) {
	got := TwoSum([]int{3, 2, 4}, 6)
	want := []int{1, 2}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestTwoSum4(t *testing.T) {
	got := TwoSum([]int{1, 2}, 4)
	if got != nil {
		t.Errorf("want nil, got %v", got)
	}
}
