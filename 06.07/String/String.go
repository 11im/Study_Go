package main

import (
	"fmt"
)

func test1() {
	str1 := "Hello\t'World'\n"
	str2 := `Hello\t'World'\n`
	fmt.Print(str1)
	fmt.Print(str2)

	str3 := "Hello\nWorld\n"
	str4 := `Hello
	World
	`
	str5 := "Hello\\nWorld\n"

	fmt.Print(str3)
	fmt.Print(str4)
	fmt.Print(str5)
}

func test2() {
	var sample rune = '한'

	fmt.Printf("%T\n", sample)
	fmt.Printf("%c\n", sample)
	fmt.Println(sample)
}

func test3() {
	var ko string = "가나다라마"
	var en string = "abcde"

	fmt.Printf("%d\n", len(ko))
	fmt.Printf("%d\n", len(en))
}

func test4() {
	str1 := "Hello World"
	runes1 := []rune{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100}

	fmt.Println(str1)
	fmt.Println(string(runes1))

	str2 := "Hello 월드"
	runes2 := []rune(str2)

	fmt.Println(str2)
	fmt.Println(len(str2))
	fmt.Println(string(runes2))
	fmt.Println(len(runes2))

}

func test5() {
	str1 := "Hello 월드!"
	for i := 0; i < len(str1); i++ {
		fmt.Printf("타입:%T, 값:%d, 문자값:%c\n", str1[i], str1[i], str1[i])
	}
}

func test6() {
	str1 := "Hello 월드!"
	arr := []rune(str1)

	for i := 0; i < len(arr); i++ {
		fmt.Printf("타입:%T, 값:%d, 문자값:%c\n", arr[i], arr[i], arr[i])
	}
}

func test7() {
	str1 := "Hello 월드!"
	for _, v := range str1 {
		fmt.Printf("타입:%T, 값:%d, 문자값:%c\n", v, v, v)
	}
}

func test8() {
	str1 := "BBB"
	str2 := "aaaaAAA"
	str3 := "BBBD"
	str4 := "ZZZ"

	fmt.Println(str1 > str2)
	fmt.Println(str1 < str3)
	fmt.Println(str1 < str4)
}

func main() {
	// test1()
	// test2()
	// test3()
	// test4()
	// test5()
	// test6()
	// test7()
	test8()
}
