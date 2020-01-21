package main

import (
	"fmt"
	"go_code/inter"
	"go_code/inter/impl"
)

func main() {

	var computer impl.UsbImpl = impl.UsbImpl{}

	computer.Connect()

	var ii inter.T = 65

	fmt.Print(ii)

	// 这时候就需要断言了
	var isStr bool
	var ss string
	ss, isStr = ii.(string)

	if isStr {
		fmt.Print(ss)
	} else {
		fmt.Print("类型转换错误")
	}

}
