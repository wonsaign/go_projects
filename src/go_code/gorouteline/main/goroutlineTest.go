package main

import (
	"fmt"
)

var x int = 0

func main() {

	go add()

	add()

	// 结果可能是这样的,为什么呢,因为主线程结束后,不会等待其他gorouteline,所以答应的数据会不一样
	/**
	0
	1
	2
	3
	4
	5
	6
	7
	8
	9
	0
	1
	2
	*/
}

func add() {
	for index := 0; index < 10; index++ {
		fmt.Printf("%v\n", index)
	}
	i := 1
	i++
}
