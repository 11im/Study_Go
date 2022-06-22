package main

import (
	"fmt"
)

func test1() {
	var a int
	var p *int
	p = &a
	*p = 20

	fmt.Printf("%v, %d\n", p, a)
}

func test2() {
	var a int = 10
	var p *int
	var q *int
	p = &a
	q = p
	*p = 20

	fmt.Printf("%v\n", p == q)
}

func test3() {
	var a int
	var p *int
	p = &a
	*p = 20
	var q *int
	if q == nil {
		fmt.Printf("%v\n", q)
	} else {
		fmt.Printf("%v, %d\n", p, a)
	}
}

type Data struct {
	id   int
	data [200]int
}

func test_pointer(arg Data) {
	arg.id = 999
	arg.data[100] = 999
}

func test_pointer2(arg *Data) {
	arg.id = 999
	arg.data[100] = 999
}

func test4() {
	var data Data

	test_pointer(data)

	fmt.Printf("%v, %v\n", data.id, data.data[100])

	test_pointer2(&data)

	fmt.Printf("%v, %v\n", data.id, data.data[100])

}

func test5() {
	var data Data
	var p *Data = &data

	fmt.Printf("%v\n", p)

	var q *Data = &Data{}

	fmt.Printf("%v\n", q)

}

func test6() {
	var p1 *Data = &Data{}
	var p2 *Data = p1
	var p3 *Data = p1
	//이렇게 선언하면 인스턴스(메모리에 있는 데이터 실체)가 하나만 생성됨

	fmt.Println(p1, p2, p3)

	var d1 Data
	var d2 Data = d1
	var d3 Data = d1
	//이렇게 선언하면 인스턴스가 3개 생성됨, d2,d3는 d1을 복사해서 새 인스턴스 생성
	fmt.Println(d1, d2, d3)

	q1 := &Data{id: 3}
	var q2 = new(Data) // 원하는 값으로 초기화는 불가능

	fmt.Println(q1, q2)

}

func main() {
	// test1()
	// test2()
	// test3()
	// test4()
	// test5()
	// test6()
}
