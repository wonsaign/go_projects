package main

import (
	"fmt"
	"time"
)

func main() {

	var now = time.Now()

	fmt.Print(now.YearDay())
}
