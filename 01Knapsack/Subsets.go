package _1Knapsack

import "leetcode/utils"

//https://leetcode.com/problems/partition-equal-subset-sum/

/*
	RECURSION + MEMOIZATION
*/
func canPartition(nums []int) bool {
	n := len(nums)
	arraySum := 0
	for _, num := range nums {
		arraySum += num
	}
	if arraySum&1 == 1 {
		//	odd array sum
		return false
	}
	subsetSum := arraySum / 2
	dp := utils.Make2DMatrixWithDefaultVal(n+1, subsetSum+1, -1)

	return canPartitionRecurse(nums, dp, subsetSum, n)

}

func canPartitionRecurse(wt []int, dp [][]int, w int, n int) bool {
	var flag bool
	if w == 0 {
		return true
	}
	if n == 0 {
		return w == 0
	}
	if dp[n][w] != -1 {
		return dp[n][w] != 0
	}
	if wt[n-1] <= w {
		flag = canPartitionRecurse(wt, dp, w-wt[n-1], n-1) || canPartitionRecurse(wt, dp, w, n-1)
	} else {
		flag = canPartitionRecurse(wt, dp, w, n-1)
	}
	if flag {
		dp[n][w] = 1
	} else {
		dp[n][w] = 0
	}
	return flag
}

/*
----------------------------------------------------------------------------------
----------------------------------------------------------------------------------
*/

//https://leetcode.com/problems/partition-to-k-equal-sum-subsets/description/

/*
	BACKTRACKING
*/
func canPartitionKSubsets(nums []int, k int) bool {
	arraySum := 0
	for _, num := range nums {
		arraySum += num
	}
	if arraySum%k != 0 {
		//cannot divide array sum into k equal subset sum
		return false
	}
	subsetSum := arraySum / k
	n := len(nums)
	visited := make([]bool, n)
	return canPartitionBacktrack(nums, visited, 0, k, 0, subsetSum)

}

func canPartitionBacktrack(nums []int, visited []bool, index int, k int, currSum int, target int) bool {
	if k == 1 {
		return true
	}
	if currSum > target {
		return false
	}
	if currSum == target {
		return canPartitionBacktrack(nums, visited, 0, k-1, 0, target)
	}
	for i := index; i < len(nums); i++ {
		if visited[i] == false {
			visited[i] = true
			if canPartitionBacktrack(nums, visited, i+1, k, currSum+nums[i], target) {
				return true
			}
			visited[i] = false
		}
	}
	return false
}
