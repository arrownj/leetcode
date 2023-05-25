package lc238_product_of_array_except_self

func Solution(nums []int) []int {
	product := 1
	zeroCount := 0
	for _, n := range nums {
		if n == 0 {
			zeroCount++
		} else {
			product *= n
		}
	}
	productExceptSelf := make([]int, 0, len(nums))
	for _, n := range nums {
		if zeroCount > 1 {
			productExceptSelf = append(productExceptSelf, 0)
		} else if zeroCount == 1 {
			if n == 0 {
				productExceptSelf = append(productExceptSelf, product)
			} else {
				productExceptSelf = append(productExceptSelf, 0)
			}
		} else {
			productExceptSelf = append(productExceptSelf, product/n)
		}
	}
	return productExceptSelf
}
