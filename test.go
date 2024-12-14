package main

import "fmt"

/*
// sorted array - repetitions are allowed
find the elements whose freq > n/4; n is the no of elements

ex: [1,1,1,3,3,3,3,6,6,9,9,9,9,9,13,16]
1 -> [0,2]
3 -> [3,6]
6 -> [7,8]
9 -> [9-13]
13 -> [14,14]
16 -> [15,15]
*/

func upperBound(arr []int, start int, end int) (lastIndex int) {
	//firstIndex := start
	fmt.Println("upperBound for element: ", arr[start])
	for start < end {
		mid := start + (end-start)/2
		if arr[mid] > arr[start] {
			end = mid
		} else {
			start = mid + 1
		}
		fmt.Println(start, end)
	}
	lastIndex = start
	return lastIndex
}

func frequentElement(nums []int) []int {
	var ans []int
	//start := 0
	end := len(nums) - 1
	//for {
	//
	//}
	var lastIndex int
	for start := 0; start < len(nums); start = lastIndex + 1 {
		firstIndex := start
		lastIndex = upperBound(nums, start, end)
		fmt.Println(nums[start], ": ", firstIndex, lastIndex)
		if (lastIndex - firstIndex + 1) >= len(nums)/4 {
			ans = append(ans, nums[start])
		}
	}

	return ans
}
