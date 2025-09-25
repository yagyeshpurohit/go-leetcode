package main

import (
	"fmt"
	"sort"
	"strings"
)

const planetMassThreshold = 10.0
const milkyWayThreshold = 21.3

type Planet struct {
	mass float64
}

type Star struct {
	mass         float64
	isInMilkyWay bool
}

func test(mass, distance float64) {
	if mass <= planetMassThreshold {
		p := Planet{mass: mass}
		fmt.Println("Is a planet", p)
	} else {
		s := Star{
			mass:         mass,
			isInMilkyWay: false,
		}
		if distance <= milkyWayThreshold {
			s.isInMilkyWay = true
			fmt.Println("Is a Star in the milky way", s)
		} else {
			fmt.Println("Is a Star, but not in milky way", s)
		}
	}
}

type Candidate struct {
	name string
	rank int
}

func sortStructs(candidateList []Candidate) {
	sort.Slice(candidateList, func(i, j int) bool {
		return candidateList[i].rank < candidateList[j].rank
	})
}

func twoSum(nums []int, target int) []int {
	numIndexMap := make(map[int]int)
	for i, num := range nums {
		j, ok := numIndexMap[target-num]
		fmt.Printf("%d ", j)
		if ok {
			return []int{j, i}
		}
		numIndexMap[num] = i
	}
	return []int{}
}

func getSchemeName(columnName string) string {
	// format -> "Amount ( SchemeName ) "
	start := strings.Index(columnName, "(")
	end := strings.LastIndex(columnName, ")")
	fmt.Println(start, end)
	return columnName[start+2 : end-1]
}

//func main() {
//	//test(10.01, 13)
//	//fmt.Println(twoSum([]int{4, 3, 7, 4, 3, 5}, 6))
//	//sl := []int{2, 3, 5}
//	//fmt.Printf("len:%d cap:%d", len(sl), cap(sl))
//	fmt.Println(getSchemeName("Amount ( Ad hoc Payment ADDITION-OPERATOR_ATTENDANCE_INCENTIVE CREDIT )"))
//}
