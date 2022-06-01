package main

import (
	"fmt"
	"math"
	"math/big"
)

func equal(a, b float64) bool {
	return math.Nextafter(a, b) == b
}

func test_equal() {
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	fmt.Printf("%0.18f + %0.18f = %0.18f\n", a, b, a+b)
	fmt.Printf("%0.18f == %0.18f : %v\n", c, a+b, equal(a+b, c))

	a = 0.0000000004
	b = 0.0000000002
	c = 0.0000000007

	fmt.Printf("%g == %g :%v\n", c, a+b, equal(a+b, c))
}

func big_float() {
	a, _ := new(big.Float).SetString("0.1")
	b, _ := new(big.Float).SetString("0.2")
	c, _ := new(big.Float).SetString("0.3")

	d := new(big.Float).Add(a, b)

	fmt.Println(a, b, c, d)
	fmt.Println(c.Cmp(d))
	// a, b 모두 big.Float일 경우, a.Cmp(b)를 통해 비교 가능, a가 b보다 작으면 -1, 같으면 0, 크면 1
}
