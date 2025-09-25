//https://leetcode.com/problems/3sum-closest/

package main

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	//Sort
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	closeness := math.MaxFloat64 //how close your sum is to the target
	closestSum := 0              //global closest sum
	//Fix one element for each iteration
	for i, _ := range nums {
		for start, end := i+1, len(nums)-1; start < end; {
			tempSum := nums[i] + nums[start] + nums[end]                  //a+b+c
			tempCloseness := math.Abs(float64(target) - float64(tempSum)) // t - (a+b+c)
			//check if t - (a+b+c) is closer than the already existing closeness
			if tempCloseness < closeness {
				closeness = tempCloseness
				closestSum = tempSum //(a+b+c) is the new closest sum for now
			}
			if tempSum > target {
				end--
			} else {
				start++
			}
		}

	}
	return closestSum
}

//func main() {
//	arr := []int{5, 10, 15, 23, 25}
//	fmt.Println(threeSumClosest(arr, 54))
//}
