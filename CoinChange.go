package main

import (
	"fmt"
)

//https://leetcode.com/problems/coin-change/

func coinChange(coins []int, amount int) int {
	/* RECURSION + DP (TOP DOWN)
	if amount == 0 {
		return 0
	}
	// dp[i] denotes the min no of coins required to make the sum amount = i
	dp := make([]int, amount+1)
	minNoOfCoins := recurseCoinChange(coins, amount, &dp)
	if minNoOfCoins < 0 {
		return -1
	}
	return minNoOfCoins

	*/
	dp := make([]int, amount+1)
	return 0
}

//func recurseCoinChange(coins []int, amount int, dp *[]int) int {
//	var noOfCoins []int
//	// base case1: if amount is negative, means we can't get the amount from the given recursive summation of coins
//	if amount < 0 {
//		return math.MinInt
//	}
//	// base case2: we have got the amount from the given recursive summation of coins
//	if amount == 0 {
//		return 0
//	}
//	if (*dp)[amount] != 0 {
//		return (*dp)[amount]
//	}
//	// recursion
//	for _, coin := range coins {
//		noOfCoins = append(noOfCoins, 1+recurseCoinChange(coins, amount-coin, dp))
//	}
//	(*dp)[amount] = utils.MinPositive(noOfCoins)
//	return (*dp)[amount]
//}

func main() {
	coins := []int{2, 5}
	amount := 0
	fmt.Println(coinChange(coins, amount))
}
