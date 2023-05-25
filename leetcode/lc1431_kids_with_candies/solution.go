package lc1431_kids_with_candies

func Solution(candies []int, extraCandies int) []bool {
	if len(candies) == 0 {
		return []bool{}
	}
	max := candies[0]
	for _, num := range candies {
		if num > max {
			max = num
		}
	}
	ret := make([]bool, len(candies), len(candies))
	for i, num := range candies {
		if num+extraCandies >= max {
			ret[i] = true
		} else {
			ret[i] = false
		}
	}
	return ret
}
