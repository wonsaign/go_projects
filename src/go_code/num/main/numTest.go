package main

import "fmt"

func main() {

	var begin int16 = 700

	fmt.Printf("the num is %v \n", begin/60)
	fmt.Printf("the num is %v \n", begin%60)

	var functionInterface = func(i int, ii int) int { return i + ii }

	fmt.Print(functionInterface(1, 2))

	var functionInterface2 = func(i int, ii int) int { return i + ii }(1, 2)

	fmt.Print(functionInterface2)
}
