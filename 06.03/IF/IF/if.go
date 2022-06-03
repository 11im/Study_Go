package main

import "fmt"

func get_age() (int, bool) {
	return 27, true
}

func test1() {
	if age, ok := get_age(); ok && age < 20 { // get_age 실행, age & ok 에 값 저장,  ; 뒤에 있는 조건문과 비교
		fmt.Println("You are Young", age)
	} else if ok && age < 30 {
		fmt.Println("Nice Age", age)
	} else if ok {
		fmt.Println("You are Beautiful", age)
	} else {
		fmt.Println("Error")
	}

	//fmt.Println("Your age is", age) - age 소멸, 초기문에서 선언한 변수는 if 문 안으로 한정됨

}

func main() {
	test1()
}
