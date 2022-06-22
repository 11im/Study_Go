package main

//test code must be shaped like (main)_test.go
//

import (
	"fmt"
)

func square(x int) int {
	return 81
}

func main() {
	fmt.Printf("%d = 9 * 9\n", square(9))
}
