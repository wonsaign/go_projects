package main

import "fmt"

// Person struct has public  name age sex
type Person struct {
	Name string // var Name string 会报错
	Age  byte
	Sex  bool
}

func main() {
	var per Person // 不会是nil,这点就说明了是值变量,其他语言,比如java,就会是null
	per.Name = "old wang"
	per.Age = 18
	per.Sex = true

	var p2 Person = Person{
		Name: "papap",
		Age:  19,
		Sex:  true, //这最后的逗号,是因为 换行的缘故啊 坑爹啊
	}
	var p3 Person = Person{
		Name: "papap",
		Age:  19,
		Sex:  true} // 看见没啊

	a := 1
	b := 2
	a,b = b,a


	fmt.Print(p2)
	fmt.Print(p3)

	fmt.Print(per)
}
