package main

func merge(intervals [][]int) [][]int {
	result := [][]int{}
	for _, interval := range intervals {
		if len(result) == 0 {
			result = append(result, interval)
		} else {
			for _, existInterval := range result {
				if hasOverlap(interval, existInterval) {

				}

			}
		}

	}

}
