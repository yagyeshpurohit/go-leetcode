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

func Make2DMatrix(rows int, cols int) [][]int {
	matrix := make([][]int, rows)
	for i, _ := range matrix {
		matrix[i] = make([]int, cols)
	}
	return matrix
}

func Make2DMatrixWithDefaultVal[T any](rows int, cols int, defaultVal T) [][]T {
	matrix := make([][]T, rows)
	for i, _ := range matrix {
		matrix[i] = make([]T, cols)
		for j, _ := range matrix[i] {
			matrix[i][j] = defaultVal
		}
	}
	return matrix
}
