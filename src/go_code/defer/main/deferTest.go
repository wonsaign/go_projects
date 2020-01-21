package main

import "fmt"

func main() {
	//sum(1, 2)
	var i, ii *int // 还未分配地址 默认为nil
	i = new(int)   // 给予地址
	ii = new(int)  // 给予地址
	*i = 1
	*ii = 2
	sumPoint(i, ii)

	fmt.Printf("i %v\n", *i)
	fmt.Printf("ii %v\n", *ii)
}

func sum(i int, ii int) {

	defer fmt.Print(i)

	defer fmt.Print(ii)

	var sum = i + ii

	fmt.Print(sum)

}

func sumPoint(i *int, ii *int) {

	defer setTen(i)

	defer setNine(ii)

	var sum = *i + *ii

	fmt.Printf("sum %v\n", sum)

}

func setTen(i *int) {
	*i = 10
}
func setNine(i *int) {
	*i = 9
}
