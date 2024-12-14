package main

import "fmt"

func main() {
	/*
		arr := []interface{}{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, 5, 1}

		root := createTree(arr)
		fmt.Println("InOrder traversal.......")
		inOrder(root)
		fmt.Println()
		fmt.Println("LevelOrder traversal.....")
		levelOrderSeq := levelOrder(root)
		fmt.Println(levelOrderSeq)
		fmt.Println()
	*/

	/*
		allETFs := []string{"SCHG", "SCHD", "VTI", "VOO", "QQQM", "VUG", "SCHG", "XLK", "VGT", "VONG", "FTEC", "MGK", "SCHG", "SCHD", "SCHX", "VUG", "VONG", "VGT", "VIOV", "VGT", "XLK", "IYW", "PSI", "SMH", "VOO", "VIG", "QQQ", "XLV", "SCHA", "AGG", "SPY", "QQQ", "SCHD", "SPMO"}
		sNp500ETFs := []string{"SPY", "VOO", "SPY", "IVV", "VOO", "SPY", "IVV", "SWPPX"}
		nasdaqETS := []string{"NASDX", "QQQ", "QQQ", "QQQM", "ONEQ", "QQQE", "QQQJ", "QTEC"}
		dividendETFs := []string{"SCHD", "VYM", "SDY", "HDV", "SCHD", "HDV", "SPYD", "VYMI", "SPHD", "PEY", "SCHD"}
		chineseETFs := []string{"KWEB", "MCH", "FXI", "CNXT", "MCHI", "FLCH", "CHIQ", "ASHR"}
		realEstateETFs := []string{"USRT", "XLRE", "BBRE", "VNQ", "NURE", "RSPR", "REZ", "JPRE", "INDS"}
		stocksRecommendation(allETFs)
	*/
	/*
		rand.Seed(time.Now().UnixNano())
		sendMsg("hello")
		go sendMsg("test1")
		go sendMsg("test2")
		go sendMsg("test3")
		sendMsg("main")
	*/
	/*
		go spinner(100 * time.Millisecond)
		const n = 45
		fibN := fib(n) // slow
		fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
	*/
	/*
		allETFs := []string{"SCHG", "SCHD", "VTI", "VOO", "QQQM", "VUG", "SCHG", "XLK", "VGT", "VONG", "FTEC", "MGK", "SCHG", "SCHD", "SCHX", "VUG", "VONG", "VGT", "VIOV", "VGT", "XLK", "IYW", "PSI", "SMH", "VOO", "VIG", "QQQ", "XLV", "SCHA", "AGG", "SPY", "QQQ", "SCHD", "SPMO"}
		stocksRecommendation(allETFs)

	*/

	//nums := []int{1, 1, 1, 3, 3, 3, 3, 6, 6, 9, 9, 9, 9, 9, 13, 16}
	//result := frequentElement(nums)
	//fmt.Println(result)

	maxHeightArr := []int{2, 3, 4, 3}
	ans := findMaxHeightSum(maxHeightArr)
	fmt.Println(ans)
}

//func sendMsg(msg string) {
//	pause()
//	fmt.Println(msg)
//}
//
//func pause() {
//	pauseTime := rand.Intn(100)
//	fmt.Println("pauseTime: ", pauseTime)
//	time.Sleep(time.Duration(pauseTime) * time.Millisecond)
//}
//
//func spinner(delay time.Duration) {
//	for {
//		for _, r := range `-\|/` {
//			fmt.Printf("\r%c", r)
//			time.Sleep(delay)
//		}
//	}
//}
//
//func fib(x int) int {
//	if x < 2 {
//		return x
//	}
//	return fib(x-1) + fib(x-2)
//}
