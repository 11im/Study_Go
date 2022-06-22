package main

import "fmt"

func test1() {

	a := 3

	switch a {
	case 1:
		fmt.Println("a == 1")
	case 2:
		fmt.Println("a == 2")
	case 3:
		fmt.Println("a == 3")
	case 4:
		fmt.Println("a == 4")
	default:
		fmt.Println("a > 4")
	}
}

func test2() {

	day := "Friday"

	switch day {
	case "Monday", "Tuesday":
		fmt.Println("수업... 가야지?")
	case "Wednesday", "Thursday", "Friday":
		fmt.Println("실습... 가야지?")
	default:
		fmt.Println("안해")
	}
}

func test3() {
	temp := 18

	switch true { //true 생략 가능
	case temp < 10, temp > 30:
		fmt.Println("집 밖은 위험해")
	case temp >= 10 && temp < 20:
		fmt.Println("겉옷 필요")
	case temp >= 15 && temp < 25: // 이미 Switch 끝남.
		fmt.Println("연구실 가야지")
	default:
		fmt.Println("따듯")
	}
}

func my_age() int {
	return 27
}
func test4() {
	switch age := my_age(); age {
	case 10:
		fmt.Println("Teeage")
	case 33:
		fmt.Println("Pair 3")
	default:
		fmt.Println("My age is", age)
	}
}

func test5() {
	switch age := my_age(); {
	case age < 10:
		fmt.Println("Child")
	case age < 20:
		fmt.Println("Teenager")
	case age < 30:
		fmt.Println("20s")
	default:
		fmt.Println("My age is", age)
	}
}

type ColorType int //별칭 정의, Const 열거
const (
	Red ColorType = iota
	Blue
	Green
	Yellow
)

func color_2_string(color ColorType) string {
	switch color {
	case Red:
		return "Red"
	case Blue:
		return "Blue"
	case Green:
		return "Green"
	case Yellow:
		return "Yellow"
	default:
		return "Undefined"
	}
}

func get_my_F() ColorType {
	return Blue
}

func test_color() {
	fmt.Println("My Favorite Color is", color_2_string(get_my_F()))
}

func test_ft() {
	temp := 18

	switch true { //true 생략 가능
	case temp < 10, temp > 30:
		fmt.Println("집 밖은 위험해")
	case temp >= 10 && temp < 20:
		fmt.Println("겉옷 필요")
		fallthrough // 다음 case도 진행
	case temp >= 15 && temp < 25: // 이미 Switch 끝남.
		fmt.Println("연구실 가야지")
	default:
		fmt.Println("따듯")
	}
}

func main() {
	test1()
	test2()
	test3()
	test4()
	test5()
	test_color()
	test_ft()
}
