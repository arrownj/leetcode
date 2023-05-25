package lc605_can_place_flowers

func Solution(flowerbed []int, n int) bool {
	adjacentZero := 0
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 {
			adjacentZero++
		} else {
			if adjacentZero > 2 {
				n -= (adjacentZero - 1) / 2
			}
			adjacentZero = 0
		}
	}
	if adjacentZero > 1 {
		n -= adjacentZero / 2

	}
	return n <= 0
}
