package main

import "fmt"

func findMaxHeightSum(maximumHeight []int) int {
	//maximumHeight = [2,3,4,3]
	heightFreqMap := make(map[int]bool)
	sum := 0
	for _, height := range maximumHeight {
		//height encountered for the first time
		if heightFreqMap[height] != true {
			sum += height
			heightFreqMap[height] = true
			continue
		} else {
			var h int
			for h = height; heightFreqMap[h] == true; h-- {
				//	todo: condition change

			}
			sum += h
			heightFreqMap[h] = true
		}

	}
	if heightFreqMap[0] == true {
		return -1
	}
	return sum
}

func main() {
	maxHeightArr := []int{4, 4, 4}
	ans := findMaxHeightSum(maxHeightArr)
	fmt.Println(ans)
}
