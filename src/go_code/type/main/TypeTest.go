package main

import (
	"fmt"
	"go_code/type/model"
)

// GoFunc type指定一个函数的别名,不能指定是实现,这个类似接口中的方法啊
type GoFunc func(a int, b int) int

func main() {
	// golang导入的是包名,而不是文件名,引用的时候要带上包名.
	var stu model.Student
	stu.SetName("wangs")
	fmt.Print(stu)
	a := 10
	b := 20

	// 匿名函数可以直接调用函数中的变量
	// 我这个变量声明也只是一个壳,只是告诉你,goFunc是一个自定义的(type)的类型
	// 类似java8 中的function函数  ,比如 (e1,e2)->return e1.compare(e2)
	// goFunc的赋值
	goFucn := func(i int, ii int) int {
		return a + b
	}

	calculateFun(10, goFucn)

}

// 使用goFunc声明
func calculateFun(i int, gofunc GoFunc) {
	// 切记,要传参,要不然就是函数变量
	var goFN GoFunc = gofunc
	// 传参后才是返回值
	res := goFN(1, 2)
	if i > res {
		fmt.Print("")
	}
}
