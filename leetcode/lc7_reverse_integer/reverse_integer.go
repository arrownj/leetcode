package reverse_integer

import (
	"math"
)

func ReverseInteger(x int) int {
	nums := []int{}
	for n := x; n != 0; n /= 10 {
		nums = append(nums, n%10)
	}
	value := 0
	for i, num := range nums {
		value += num * int(math.Pow10(len(nums)-i-1))
	}
	if value > int(math.Pow(2, 31))-1 || value < int(math.Pow(2, 31))*-1 {
		value = 0
	}
	return value
}
