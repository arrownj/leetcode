package main

func lemonadeChange(bills []int) bool {
	changes := map[int]int{
		5:  0,
		10: 0,
		20: 0,
	}
	for _, bill := range bills {
		if count, ok := changes[bill]; ok {
			changes[bill] = count + 1
		}
		if bill == 20 {
			tenCount, _ := changes[10]
			fiveCount, _ := changes[5]
			if tenCount >= 1 && fiveCount >= 1 {
				changes[10] = tenCount - 1
				changes[5] = fiveCount - 1
			} else if fiveCount >= 3 {
				changes[5] = fiveCount - 3
			} else {
				return false
			}
		} else if bill == 10 {
			fiveCount, _ := changes[5]
			if fiveCount >= 1 {
				changes[5] = fiveCount - 1
			} else {
				return false
			}
		}
	}
	return true
}
