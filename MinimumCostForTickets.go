package main

import (
	"fmt"
	"leetcode/utils"
)

//https://leetcode.com/problems/minimum-cost-for-tickets/

func recurseMinCost(days []int, costs []int, dp *[]int, travelDays map[int]bool, currDay int) int {
	//base case: if the current day is greater than the last travel day, return
	if currDay > days[len(days)-1] {
		return 0
	}
	// check if the current day is a travelling day. If it is not, then we don't need to buy a ticket that day, and simply move to currDay + 1
	//if !travelDays[currDay] {
	//	recurseMinCost(days, costs, dp, travelDays, currDay+1)
	//}
	// If we have already computed for this day, simply return its value
	if (*dp)[currDay] != -1 {
		return (*dp)[currDay]
	}
	if !travelDays[currDay] {
		return recurseMinCost(days, costs, dp, travelDays, currDay+1)
	}
	// recursion: compute all the possible costs for travelling on current day
	oneDayTicketCost := costs[0] + recurseMinCost(days, costs, dp, travelDays, currDay+1)
	oneWeekTicketCost := costs[1] + recurseMinCost(days, costs, dp, travelDays, currDay+7)
	oneMonthTicketCost := costs[2] + recurseMinCost(days, costs, dp, travelDays, currDay+30)
	// look for the minimum cost amongst these and return it
	minCheckList := []int{oneDayTicketCost, oneWeekTicketCost, oneMonthTicketCost}
	(*dp)[currDay] = utils.Min(minCheckList)
	return (*dp)[currDay]
}

func mincostTickets(days []int, costs []int) int {
	// dp[i] denotes the min travel cost if we start from day i
	dp := make([]int, days[len(days)-1]+1)
	for i := range dp {
		dp[i] = -1
	}
	travelDays := make(map[int]bool)
	for _, day := range days {
		travelDays[day] = true
	}

	minTravelCost := recurseMinCost(days, costs, &dp, travelDays, 0)
	return minTravelCost
}

func main() {
	//days := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}
	days := []int{1, 4, 6, 7, 8, 20}
	costs := []int{2, 7, 15}
	fmt.Println(mincostTickets(days, costs))
}
