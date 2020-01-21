package main

import "fmt"

func main() {
	var a int = 1
	var b *int
	b = &a
	*b = 2
	fmt.Println(b, *b, &b, a, &a)

	// var x *int = int * (0xc000014078)
	var x *int = nil
	fmt.Println(x)
}

func test() {

	var a int = 1

	fmt.Printf("原参数的指针: %x\n", &a)

	assignTest(&a)
	//assignTest2(a)

	fmt.Printf("原参数修改后的指针: %x\n", &a)
	fmt.Printf("原参数的值域有没有被修改: %x\n", a)
}

func assignTest(a *int) {
	fmt.Printf("此时进入函数时候指针的值是: %x\n", a)
	var i int = 10
	// 注意 下面两周操作 都是对形参的修改，都是值传递操作，与java的传惨一致

	// 地址换了，函数，都是赋值操作，类似java的参数，形参是实参的地址值的赋值,这个操作只是将形参的【地址】变了，并没有改变实参
	//a = &i
	// 还是上面的操作，这个操作将形参【地址】所对应的【值】改了，由于形参和实参都使用同一个地址，所以实参的值也变了
	*a = i

	fmt.Printf("指针赋值结束后指针的值是: %x\n", a)
	/* 使用指针访问值 */
	fmt.Printf("*a 变量的值: %d\n", a)
}

func assignTest2(a int) {
	fmt.Printf("此时进入函数时候指针的值是: %x\n", a)
	var i int = 10
	// 注意 下面两周操作 都是对形参的修改，都是值传递操作，与java的传惨一致

	// 地址换了，函数，都是赋值操作，类似java的参数，形参是实参的地址值的赋值,这个操作只是将形参的【地址】变了，并没有改变实参
	//a = &i
	// 还是上面的操作，这个操作将形参【地址】所对应的【值】改了，由于形参和实参都使用同一个地址，所以实参的值也变了
	a = i

	fmt.Printf("指针赋值结束后指针的值是: %x\n", a)
	/* 使用指针访问值 */
	fmt.Printf("a变量的地址值: %d\n", &a)
}
