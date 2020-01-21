package main

import "fmt"

func main() {

	itWillBeWrong()

}

func itWillBeWrong() {

	//defer catchPanic()

	defer func() {
		var rec = recover()

		if rec != nil {
			fmt.Print(rec)
		}
	}()

	var i int8 = 1
	var zero int8 = 0
	if zero == 0 {
		panic("zero cant devided")
	}

	var x = i / zero

	fmt.Print(x)
}

func catchPanic() {
	var rec = recover()

	if rec != nil {
		fmt.Print(rec)
	}
}
