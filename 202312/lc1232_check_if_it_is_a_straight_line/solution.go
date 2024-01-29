package main

func checkStraightLine(coordinates [][]int) bool {
	if len(coordinates) <= 2 {
		return true
	}
	x0, y0 := coordinates[0][0], coordinates[0][1]
	x1, y1 := coordinates[1][0], coordinates[1][1]
	for i := 2; i < len(coordinates); i++ {
		xi, yi := coordinates[i][0], coordinates[i][1]
		if xi-x0 != 0 && x1-x0 != 0 {
			if float64(yi-y0)/float64(xi-x0) != float64(y1-y0)/float64(x1-x0) {
				return false
			}
		} else if xi-x0 == 0 && x1-x0 == 0 {

		} else {
			return false
		}

	}
	return true
}
