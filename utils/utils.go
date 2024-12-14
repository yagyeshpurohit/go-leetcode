package utils

import "math"

func Min(arr []int) int {
	minVal := math.MaxInt
	for _, val := range arr {
		if val < minVal {
			minVal = val
		}
	}
	return minVal
}

func Max(arr []int) int {
	maxVal := math.MinInt
	for _, val := range arr {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}

func MinPositive(arr []int) int {
	var positiveArr []int
	for _, num := range arr {
		if num > 0 {
			positiveArr = append(positiveArr, num)
		}
	}
	if len(positiveArr) == 0 {
		return Min(arr)
	} else {
		return Min(positiveArr)
	}
}
