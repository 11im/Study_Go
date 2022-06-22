package main

import (
	"fmt"
)

func sample() {
	const Pi1 float64 = 3.141592
	var Pi2 float64 = 3.141592

	//Pi1 = 10
	Pi2 = 10

	fmt.Println("Pi", Pi1)
	fmt.Println("Pi", Pi2)
}

func Iota() {
	const (
		Red   int = iota
		Blue  int = iota
		Green int = iota
	)
	fmt.Println(Red, Blue, Green)

	const (
		C1 uint = iota + 1
		C2
		C3
	)
	fmt.Println(C1, C2, C3)

	const (
		F1 uint = 1 << iota
		F2
		F3
	)
	fmt.Println(F1, F2, F3)

}

func type_test() {
	const PI = 3.14
	const PI2 float64 = 3.14

	var a int = PI * 100       // Const 선언 시 타입을 지정하지 않음 -> 복사될 때 타입 지정됨
	var b int = int(PI2 * 100) // Const 선언 시 타입을 지정하였기 때문에 형변환 필요

	fmt.Println(a, b)
}

func main() {
	sample()
	Iota()
	type_test()
}
