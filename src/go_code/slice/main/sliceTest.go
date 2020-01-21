package main

import "fmt"

func main() {

	var slice []string = make([]string, 5, 10)
	fmt.Printf("%p \n", slice)
	fmt.Printf("%p", &slice)

	var sliceP []*string = make([]*string, 5, 10)
	fmt.Printf("%p", sliceP)
	fmt.Printf("%v \n", len(slice))
	fmt.Print(cap(slice))

	for index, s := range slice {
		// ASCII表中对应的A
		s = string(65 + index)
		slice[index] = s
		fmt.Print(&index)
		fmt.Print(s)
	}
	var pSlice *[]string = &slice

	fmt.Printf("%p", pSlice)
}
