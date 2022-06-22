package main

import (
	"fmt"
)

func test1() {
	var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 26.2} // 선언 -> var (이름) [배열 길이] (타입) = [길이](타입){ 내용물 }

	for i := 0; i < 5; i++ {
		fmt.Print(t[i], ", ")
	}
}

func test2() {
	var t2 [5]int = [5]int{1: 10, 3: 30} // 대입 시 {인덱스 : 값, 인덱스 : 값}으로 원하는 인덱스에 넣을 수 있음
	t3 := [...]int{1, 2, 3}              // ...을 이용해 배열 요소 개수 생략 가능
	for i := 0; i < 5; i++ {
		fmt.Print(t2[i], ", ")
	}
	for i := 0; i < 3; i++ {
		fmt.Print(t3[i], ", ")
	}
}

func test3() {
	var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 26.2}

	for i, v := range t { //range는 인덱스와 배열 값을 return
		fmt.Println(i, v)
	}
}

func test4() {
	a := [...]int{1, 2, 3, 4, 5}
	b := [...]int{10, 20, 30, 40, 50}

	for i := 0; i < 5; i++ {
		fmt.Println(a[i])
	}

	for i := 0; i < 5; i++ {
		fmt.Println(b[i])
	}

	b = a // 우변 메모리 내용을 좌변 메모리에 복사, 배열, 타입이 모두 같아야함

	for i := 0; i < 5; i++ {
		fmt.Println(b[i])
	}
}

func test5() {
	a := [2][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}

	for _, arr := range a {
		for _, v := range arr {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}

}

func main() {
	// test1()
	// test2()
	// test3()
	// test4()
	test5()
}
