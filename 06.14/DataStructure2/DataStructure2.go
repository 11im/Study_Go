package main

import (
	"container/ring" //ring은 container/ring에 있음
	"fmt"
)

func test1() {
	r := ring.New(5)
	n := r.Len()

	for i := 0; i < n; i++ {
		r.Value = 'A' + i
		r = r.Next()
	}

	for j := 0; j < n; j++ {
		fmt.Printf("%c ", r.Value)
		r = r.Next()
	}

	fmt.Println()

	for j := 0; j < n; j++ {
		fmt.Printf("%c ", r.Value)
		r = r.Prev()
	}
}

func test2() {
	m := make(map[string]string) //map은 기본 기능으로 지원함
	m["김하나"] = "서울시 중구"
	m["김둘"] = "서울시 노원구"
	m["김셋"] = "서울시 중랑구"

	fmt.Printf("%s\n", m["김하나"])

	m["김하나"] = "서울시 노원구"

	fmt.Printf("%s\n", m["김하나"])

	for k, v := range m {
		fmt.Println(k, v)
	}

}

type Product struct {
	Name  string
	Price int
}

func test3() {
	m := make(map[int]Product)
	m[1] = Product{"Pen", 1000}
	m[2] = Product{"Note", 2000}
	m[3] = Product{"Eraser", 500}
	m[4] = Product{"Sharp Pencil", 10000}
	m[5] = Product{"Pencil", 300}

	for k, v := range m {
		fmt.Println(k, v)
	}
}

func test4() {
	m := make(map[int]int)
	m[1] = 1
	m[2] = 2
	m[3] = 3
	m[4] = 4
	m[5] = 5

	delete(m, 3)
	delete(m, 4)
	fmt.Println(m[3])
	fmt.Println(m[1])

	for i := 1; i <= 5; i++ {
		v, ok := m[i]
		fmt.Println(v, ok)
	}

}

func main() {
	// test1()
	// test2()
	// test3()
	test4()
}
