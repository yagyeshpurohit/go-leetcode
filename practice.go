package main

import (
	"fmt"
	"math"
	"sort"
)

func isAcronym(words []string, s string) bool {
	if len(words) != len(s) {
		return false
	}
	for i, word := range words {
		if word[0] != s[i] {
			return false
		}
	}
	return true
}

func maximizeTheProfit(n int, offers [][]int) int {

	return 0
}

//https://leetcode.com/problems/house-robber/
func rob(nums []int) int {
	return 0
}

//https://leetcode.com/problems/3sum-closest/
func threeSumClosest(nums []int, target int) int {

	// [6, 1, 2, 5, -2, 4] tar:3
	//sum := 0
	for i := 0; i < len(nums); i++ {

	}

	return 0
}

//https://leetcode.com/problems/merge-two-sorted-lists/
type ListNode struct {
	Val  int
	Next *ListNode
}

//Iterative Approach
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var (
		mergedListHead *ListNode
		curr           *ListNode
	)
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	for list1 != nil || list2 != nil {
		if list1 == nil {
			curr.Next = list2
			list2 = list2.Next
		} else if list2 == nil {
			curr.Next = list1
			list1 = list1.Next
		} else if list1.Val <= list2.Val {
			if mergedListHead == nil {
				mergedListHead = &ListNode{}
				curr = mergedListHead
			} else {
				curr.Next = list1
			}
			list1 = list1.Next
		} else {
			if mergedListHead == nil {
				mergedListHead = &ListNode{}
				curr = mergedListHead
			} else {
				curr.Next = list2
			}
			list2 = list2.Next
		}
	}
	return mergedListHead.Next
}

//https://leetcode.com/problems/longest-consecutive-sequence
func longestConsecutive(nums []int) int {
	return 0
}

//https://leetcode.com/problems/path-sum/
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum -= root.Val
	if root.Left == nil && root.Right == nil {
		return targetSum == 0
	}
	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
}

//https://leetcode.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	var (
		start           int
		end             int
		maxSubstrLength int
	)
	charSet := make(map[uint8]bool)
	start = 0
	end = 0
	maxSubstrLength = 0

	for ; end < len(s); end++ {
		//check if the character encountered is present in the map
		_, ok := charSet[s[end]]
		if ok {
			// char is already present; duplicate val encountered;
			for {
				delete(charSet, s[start])
				if s[start] == s[end] {
					start++
					break
				}
				start++
			}
		}
		charSet[s[end]] = true
		if maxSubstrLength < (end - start + 1) {
			maxSubstrLength = (end - start + 1)
		}
	}
	return maxSubstrLength
}

//https://leetcode.com/problems/longest-nice-subarray/
func longestNiceSubarray(nums []int) int {
	var (
		start                     int
		end                       int
		bitmap                    int
		longestNiceSubarrayLength int
	)
	start = 0
	end = 0
	bitmap = 0
	longestNiceSubarrayLength = 0
	for ; end < len(nums); end++ {
		/*
			step 2: if the new element encountered (nums[end]) has a 1(or multiple 1s) in the same position(s) as in the bitmap vector, it means the longest nice subarray condition is violated.
			for such cases, we need to slide the window from start, and reset the positions of 1s in bitmap to 0s where 1s are present in nums[start]
			XOR does the job here, as (1 XOR 1 = 0)
			We need not worry about (1 XOR 0 = 1) as there will be no case where there will be 1 in nums[start] and 0 in bitmap for the same position (since earlier we must have done bitmap = nums[end] | bitmap)
		*/
		for ; bitmap&nums[end] != 0; start++ {
			bitmap = bitmap ^ nums[start]
		}
		//step 1: for every element of the array, do the OR with bitmap vector(initially all 0s) to get the position of 1s
		bitmap = bitmap | nums[end]
		//step 3: compute the max of curr longest subarr len and the diff between start and end
		if longestNiceSubarrayLength < (end - start + 1) {
			longestNiceSubarrayLength = (end - start + 1)
		}
	}
	return longestNiceSubarrayLength
}

//https://leetcode.com/problems/minimum-size-subarray-sum/
func minSubArrayLen(target int, nums []int) int {
	var (
		begin           int
		end             int
		subArrSum       int
		minSubArrLength int
	)
	begin = 0
	end = 0
	subArrSum = 0
	minSubArrLength = math.MaxInt
	for ; end < len(nums); end++ {

		subArrSum += nums[end]
		for ; subArrSum > target; begin++ {
			if minSubArrLength > end-begin+1 {
				minSubArrLength = end - begin + 1
			}
			subArrSum -= nums[begin]
		}
	}
	if minSubArrLength == math.MaxInt {
		return 0
	} else {
		return minSubArrLength
	}
}

//https://leetcode.com/problems/minimum-window-substring/
//func minWindow(s string, t string) string {
//	exStr := "tmdRkwaPabRPleaxQmnxd"
//	tar := "pqr"
//	//create a char frequency counter for string t
//	charFreqT := make(map[rune]int)
//	return ""
//}

//https://leetcode.com/problems/maximum-subarray/
func maxSubArray(nums []int) int {
	//Divide and Conquer Solution
	return maxSubArrayDivideAndConquer(&nums, 0, len(nums)-1)
}

func maxSubArrayDivideAndConquer(arr *[]int, l int, r int) int {
	// step 3: identify base condition
	if l == r {
		return (*arr)[l]
	}
	//	step 1: Dividing the array and taking max subarray sum of both the halves
	mid := l + (r-l)/2
	maxLeftSubArraySum := maxSubArrayDivideAndConquer(arr, l, mid)
	maxRightSubArraySum := maxSubArrayDivideAndConquer(arr, mid+1, r)
	//	step 2: Computing the maxSubArray sum for the overlapping subarray, i.e. merging the halves
	maxOverlappingSubArraySum := computeOverlappingSubArraySum(arr, l, mid, r)

	//	return the MAX(maxLeftSubArray, maxRightSubArraySum, maxOverlappingSubArraySum)
	if maxLeftSubArraySum >= maxRightSubArraySum && maxLeftSubArraySum >= maxOverlappingSubArraySum {
		return maxLeftSubArraySum
	} else if maxRightSubArraySum >= maxLeftSubArraySum && maxRightSubArraySum >= maxOverlappingSubArraySum {
		return maxRightSubArraySum
	} else {
		return maxOverlappingSubArraySum
	}

}

func computeOverlappingSubArraySum(arr *[]int, l int, m int, r int) int {
	// left <- mid
	maxLeftSum := (*arr)[m]
	for i, currSum := m, 0; i >= l; i-- {
		currSum += (*arr)[i]
		if currSum > maxLeftSum {
			maxLeftSum = currSum
		}
	}
	// mid -> right
	maxRightSum := (*arr)[m+1]
	for j, currSum := m+1, 0; j <= r; j++ {
		currSum += (*arr)[j]
		if currSum > maxRightSum {
			maxRightSum = currSum
		}
	}
	return maxLeftSum + maxRightSum
}

//https://leetcode.com/problems/majority-element/
func majorityElement(nums []int) int {
	//Freq Hashmap
	/*
		freqCount := make(map[int]int)
		for _, num := range nums {
			freqCount[num]++
			if freqCount[num] > len(nums)/2 {
				return num
			}
		}
		return -1
	*/
	//Sorting
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	fmt.Println(nums)
	size := len(nums)
	return nums[size/2]
}

func stocksRecommendation(stockList []string) {
	//creating a map for {stock:recommendFreq}
	stockFreqMap := make(map[string]int)
	for _, stock := range stockList {
		stockFreqMap[stock]++
	}
	//sorting the stocks based on no of recommendations
	type StockRecommendation struct {
		StockName     string
		RecommendFreq int
	}
	var stockRecommendationList []StockRecommendation

	for key, value := range stockFreqMap {
		stockRecommendationList = append(stockRecommendationList, StockRecommendation{StockName: key, RecommendFreq: value})
	}
	sort.Slice(stockRecommendationList, func(i int, j int) bool {
		return stockRecommendationList[i].RecommendFreq > stockRecommendationList[j].RecommendFreq
	})
	fmt.Println("Most recommended Stocks with their recommendation votes are: ")
	for i := range stockRecommendationList {
		fmt.Println(stockRecommendationList[i])
	}

}
