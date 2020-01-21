package main

import "fmt"

func main() {

	var s string = "hello world"

	for index := 0; index < len(s); index++ {
		fmt.Println(s[index])
	}

	for _, val := range s {
		fmt.Println(val)
	}

	for {
		fmt.Println(1)
	}

}

func test(a int, b int) {

}

// func test(a int) {

// }
