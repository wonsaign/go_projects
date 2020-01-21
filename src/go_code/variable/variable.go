package variable

import "fmt"

// 全局变量声明
//var x, y, z = 10, 20, 30
var x = 100
var y = 1000
var z = 10000

//  按照视频上来操作 是报错的
var (
	x1 = 1
	y1 = 2
	z1 = 3
)

func main() {
	fmt.Print(x1, y1, z1, x, y, z)
}

// 单变量声明
func singleDeclare() {

	// 变量推导
	var i = 10
	var ii int // 默认值是0

	var iii int = 20
	iiii := 30 // 等价于 var iiii = 30

	fmt.Println(i, ii, iii, iiii)
}

// 多变量声明
func multiDeclare() (ix int, iix int) {

	// wrong way
	//var i int, ii int ,iii int , iiii int

	// 变量推导
	var i, ii, iii, iiii int
	//	 多变量 不能指定声明每个变量类型， 但是可以使用类型推导
	var a1, a2, a3 = "wangsai", 18, 1

	// var 与 := 等价关系
	b1, b2, b3 := "hello", "ketty", 18

	fmt.Println(i, ii, iii, iiii)
	fmt.Println(a1, a2, a3)
	fmt.Println(b1, b2, b3)
	return i, ii
}
