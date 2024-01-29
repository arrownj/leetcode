package main

func isHappy(num int) bool {
	nums := map[int]bool{
		num: true,
	}
	for num != 1 {
		sum := 0
		for num != 0 {
			sum += (num % 10) * (num % 10)
			num /= 10
		}
		num = sum
		if _, ok := nums[num]; ok {
			return false
		}
	}
	return true

}
