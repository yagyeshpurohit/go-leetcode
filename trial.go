package main

import (
	"fmt"
	"math"
)

const EPSILON = 1e-2

func main() {
	a := 23.6
	b := 23.62
	diff := math.Abs(b - a)
	if diff > EPSILON {
		fmt.Println("diff is large", diff)
	} else {
		fmt.Println("diff is small", diff)
	}

}
