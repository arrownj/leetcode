package main

import "strconv"

func calPoints(operations []string) int {
	scores := []int{}
	for _, operation := range operations {
		if "D" == operation {
			scores = append(scores, scores[len(scores)-1]*2)
		} else if "C" == operation {
			scores = scores[:len(scores)-1]
		} else if "+" == operation {
			scores = append(scores, scores[len(scores)-2]+scores[len(scores)-1])
		} else {
			score, _ := strconv.Atoi(operation)
			scores = append(scores, score)
		}
	}

	sum := 0
	for i := 0; i < len(scores); i++ {
		sum += scores[i]
	}
	return sum
}
