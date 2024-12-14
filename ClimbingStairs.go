package main

import "fmt"

func climbStairs(n int) int {
	var dp = make([]int, n+1)
	ans := recurseClimbStairs(n, dp)
	return ans
}

func recurseClimbStairs(n int, dp []int) int {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 0
	}
	if dp[n] != 0 {
		return dp[n]
	}
	dp[n] = recurseClimbStairs(n-1, dp) + recurseClimbStairs(n-2, dp)
	return dp[n]
}

func main() {
	ans := climbStairs(5)
	fmt.Println(ans)
}
