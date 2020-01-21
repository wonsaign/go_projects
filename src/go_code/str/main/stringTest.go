package main

import (
	"fmt"
)

func main() {
	// var i int = 99

	// // 字符类型转换
	// var str = fmt.Sprintf("%d", i)

	// fmt.Printf("str is %T and str value %q \n", str, str)

	// var str2 = strconv.FormatInt(int64(i), 10)

	// fmt.Printf("str2 is %T and str2 value %q \n", str2, str2)

	// var ii, _ = strconv.ParseInt(str2, 10, 8)

	// fmt.Printf("ii is %T and ii value %q", ii, ii)

	var arrt = make([]int32, 2, 10)
	arrt[0] = '你'
	arrt[1] = '好'
	conversionStr := string(arrt) // 这个不是啥方法,别忘记,这个是强转
	fmt.Print(conversionStr)

}
