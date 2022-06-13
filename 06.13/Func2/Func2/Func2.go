package main

import (
	"fmt"
	"os"
)

func test1(i ...int) { //여러 값 : ...TYPE : 내부에서 슬라이스 타입으로 처리됨
	sum := 0
	for _, v := range i {
		sum += v
	}
	fmt.Printf("Type : %T\nSum : %d\n", i, sum)
}

func test_test1() {
	test1(1, 2, 3, 4, 5, 6, 7, 8, 9)
	test1(10, 20)
}

func test2(agrs ...interface{}) { //...interface{}를 통해 여러 타입 받을 수 있음, 내부에서는 Type 별로 Handle해야함
	for _, v := range agrs {
		switch f := v.(type) {
		case bool:
			val := v.(bool)
			fmt.Printf("%v, %v", val, f)
		case float64:
			val := v.(float64)
			fmt.Printf("%v, %v", val, f)
		case int:
			val := v.(int)
			fmt.Printf("%v, %v", val, f)
		case string:
			val := v.(string)
			fmt.Printf("%v, %v", val, f)
		}
	}
}

func test3() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to Create File")
		return
	}

	defer fmt.Println("Must be Called") //defer는 마지막에 역순으로 진행됨
	defer f.Close()
	defer fmt.Println("File Closed")

	fmt.Println("Write Hello World to File")
	fmt.Fprintln(f, "Hello World")
}

func test4_1(a, b int) int {
	return a + b
}

func test4_2(a, b int) int {
	return a * b
}

func test4_3(op string) func(int, int) int { //return type 에 함수 가능
	if op == "+" {
		return test4_1
	} else if op == "*" {
		return test4_2
	} else {
		return nil
	}
}

func test_test4() {
	var op func(int, int) int //operand 설정 필요
	op = test4_3("*")
	var result = op(3, 4) //결과값 따로 저장
	fmt.Println(result)
	op = test4_3("+")
	result = op(3, 4)
	fmt.Println(result)
}

type opFunc func(int, int) int

func test5(op string) opFunc {
	if op == "+" {
		return func(a, b int) int {
			return a + b
		}
	} else if op == "*" {
		return func(a, b int) int { // return func 후 선언 가능 :: 리터럴 이용 가능
			return a * b
		}
	} else {
		return nil
	}
}

func test_test5() {
	add := test5("+") //선언 후 가져오기 가능
	mul := test5("*")

	var a = add(3, 4)
	var b = mul(3, 4)
	fmt.Println(a, b)

	var c = test5("+")(3, 4) //바로 가져오기 가능
	var d = test5("*")(3, 4)
	fmt.Println(c, d)
}

func test6_1() {
	f := make([]func(), 3)
	fmt.Println("Value Loop")
	for i := 0; i < 3; i++ {
		f[i] = func() {
			fmt.Println(i)
		}
	}
	for i := 0; i < 3; i++ {
		f[i]() //호출 되는 순간의 변수 값을 참조하기 때문에 3 3 3 이 출력됨
	}
}

func test6_2() {
	f := make([]func(), 3)
	fmt.Println("Value Loop")
	for i := 0; i < 3; i++ {
		v := i
		f[i] = func() {
			fmt.Println(v)
		}
	}
	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func test6() {
	test6_1()
	test6_2()
}

type Writer func(string)

func test7_1(writer Writer) {
	writer("Helo World")
}

func test7() {
	f, err := os.Create("text.txt")
	if err != nil {
		fmt.Println("Error!")
		return
	}

	defer f.Close()

	test7_1(func(msg string) {
		fmt.Fprintln(f, msg)
	})
}

func main() {
	// test_test1()
	// test3()
	// test_test4()
	// test_test5()
	// test6()
	test7()
}
