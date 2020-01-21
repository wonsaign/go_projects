package main

import "fmt"

func main() {
	fmt.Println("hello",
		"world")

	// deferTest(1, 2)
	pointTest()
	// 函数大写代表了 共有
	// fmt.panic("")
}

// function 请求参数，返回参数  另外支持多个参数返回
func cal(a int, b int) (x int, y string) {
	var i = 10

	var ii string = "10"

	fmt.Println(ii)

	return i + a + b, "hello world"
}

// 延迟调用 很牛逼啊
func deferTest(a int, b int) (z int) {
	//不是很明白，不过很牛逼的样子
	defer func() {
		fmt.Println("延迟操作")
	}()

	fmt.Printf("变量的地址: %x\n", &a)

	z = a + b
	fmt.Println(z)

	return z
}

func pointTest() {
	var i int = 10
	var ip *int

	ip = &i

	//fmt.Printf("指针赋值操作: %x\n", &ip)
	fmt.Printf("指针赋值操作: %x\n", ip)
	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)
}
