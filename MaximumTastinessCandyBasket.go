package main

import "sort"

func maximumTastiness(price []int, k int) int {
	//sort
	/*
		sort.SliceStable(price, func(i, j int) bool {
			return price[i] < price[j]
		})
	*/
	sort.Ints(price)
	n := len(price)
	maxPossibleTastiness := price[n-1] - price[0] // we only need 2 candies, the cheapest and the most expensive
	minPossibleTastiness := 0                     // we can have 2 or more candies of same price
	maximumTastinessAns := 0                      // ans to be returned
	//we start with maxBasketSize = 2 and aim to reach upto k
	maxBasketSize := 2
	for begin, end := minPossibleTastiness, maxPossibleTastiness; begin <= end; {
		mid := begin + (end-begin)/2
		maxBasketSize = maximumSubsetHavingGreaterDiff(price, mid)
		if maxBasketSize < k {
			end = mid - 1
		} else {
			begin = mid + 1
			maximumTastinessAns = mid
		}
	}
	return maximumTastinessAns
}

func maximumSubsetHavingGreaterDiff(nums []int, diff int) int {
	//Assuming nums[] will always be sorted
	noOfSubsetElems := 1
	lastSubsetElem := nums[0]
	for i, num := range nums {
		if i == 0 {
			continue
		}
		if num-lastSubsetElem >= diff {
			lastSubsetElem = num
			noOfSubsetElems++
		}
	}
	return noOfSubsetElems
}
