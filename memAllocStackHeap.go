package main

//func main() {
//	//a := readA()
//	//fmt.Println(a)
//	//b := make([]byte, 16)
//	//readB(b)
//	//n := answer()
//	//fmt.Println(n)
//	s := make([]int, 5)
//	fmt.Printf("%d, %d, %v, %p \n", len(s), cap(s), s, &s)
//	s = append(s, []int{1, 2, 3, 4, 5, 6}...)
//	fmt.Printf("%d, %d, %v, %p \n", len(s), cap(s), s, &s)
//}

func readA() []byte {
	a := make([]byte, 32)
	return a
}

func readB(b []byte) {
	b = append(b, 29)
}

//func answer() *int {
//	x := 7
//	return &x
//
//}
