package main

import "fmt"

func main() {

	var x, y = functionTest()

	// how to write=> 使用'_'占位
	var a, _ = functionTest()

	var f = paramTest3

	fmt.Println(f(10000))

	fmt.Print(x, y)
	fmt.Print(a)

	// 匿名函数调用方式一
	var alf = func(i int) int {
		return i
	}
	fmt.Println(alf(10))
	// 匿名函数调方式二
	var alf2 = func(i int) int {
		return i
	}(10)
	fmt.Println(alf2)

	// 无返回值到函数
	noReturn(2)

}

func functionTest() (x int, y int) {
	i := 10
	var ii = 20
	return i, ii
}

func paramTest(s1 string, s2 string) string {
	return s1 + s2
}

func paramTest2(s1 string, s2 string) (string, string) {
	return s1 + s2, s1 + s2
}

func paramTest3(i int) int {
	return i
}

func noReturn(i int) {
	fmt.Print(i)
}
