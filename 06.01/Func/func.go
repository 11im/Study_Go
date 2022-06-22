package main

import "fmt"

func Sample(x int, y int) int {
	return x + y
	// Sample - 함수 명
	// x, y - 매개 변수, 매개 변수 뒤에 타입
	// int - 반환 타입, 맨 뒤에
}

func test_sample() {
	fmt.Println(Sample(1, 2))
}

func Devide(x, y int) (result int, success bool) {
	if y == 0 {
		result = 0
		success = false
		return
	} else {
		result = x / y
		success = true
		return
	}
	//매개 변수 x, y를 한 번에 지정
	// result, success처럼 반환값을 각각 지정할 수 있음
}

func test_devide() {
	c, suc := Devide(4, 2)
	fmt.Println(c, suc)
	d, succ := Devide(5, 0)
	fmt.Println(d, succ)
}

func recur(a int) {
	if a == 0 {
		return
	}

	fmt.Println(a)
	recur(a - 1)
	fmt.Println("After", a)
}

func test_recur() {
	recur(3)
}

func main() {
	test_sample()
	test_devide()
	test_recur()
}
