package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func test1_1(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()
	rd := bufio.NewReader(file)
	line, _ := rd.ReadString('\n')
	return line, nil
}

func test1_2(filename string, line string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = fmt.Fprintln(file, line)
	return err
}

func test_test1() {
	filename := "data.txt"

	line, err := test1_1(filename)
	if err != nil {
		err = test1_2(filename, "This is WriteFile")
		if err != nil {
			fmt.Println("Failed to Create File ", err)
			return
		}
		line, err = test1_1(filename)
		if err != nil {
			fmt.Println("Failed to Read File ", err)
			return
		}
	}
	fmt.Println("Contents :", line)
}

func test2(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("Negative! f:%g", f)

	}
	return math.Sqrt(f), nil
}

func test_test2() {
	sqrt, err := test2(-1)
	if err != nil {
		fmt.Printf("Err : %v\n", err)
		return
	}
	fmt.Println(sqrt)
}

type PWError struct {
	Len   int
	RQLen int
}

func (err PWError) Error() string {
	return "Too Short!"
}

func test3(id, pw string) error {
	if len(pw) < 8 {
		return PWError{len(pw), 8}
	}
	return nil
}

func test_test3() {
	err := test3("myId", "myPw")
	if err != nil {
		if errInfo, ok := err.(PWError); ok {
			fmt.Printf("%v Len:%d RQLen:%d\n", errInfo, errInfo.Len, errInfo.RQLen)
		}
	} else {
		fmt.Println("Registered")
	}

}

func test4_1(str string) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(str)) //스캐너 생성
	scanner.Split(bufio.ScanWords)                      //한 단어씩 끊어 읽기

	pos := 0
	a, n, err := test4_2(scanner)
	if err != nil {
		return 0, fmt.Errorf("Failed to readNextInt(), pos:%d, err:%w\n", pos, err) //에러 감싸기
	}

	pos = n + 1
	b, n, err := test4_2(scanner)
	if err != nil {
		return 0, fmt.Errorf("Failed to readNextInt(), pos:%d, err:%w\n", pos, err) //에러 감싸기
	}
	return a * b, nil
}

func test4_2(scanner *bufio.Scanner) (int, int, error) {
	if !scanner.Scan() { //단어 읽기
		return 0, 0, fmt.Errorf("Failed to Scan")
	}

	word := scanner.Text()

	number, err := strconv.Atoi(word) //str -> int
	if err != nil {
		return 0, 0, fmt.Errorf("Failed to Convert word:%s,err:%w", word, err)
	}

	return number, len(word), nil
}

func test4_3(eq string) {
	rst, err := test4_1(eq)
	if err == nil {
		fmt.Println(rst)
	} else {
		fmt.Println(err)
		var numError *strconv.NumError
		if errors.As(err, &numError) {
			fmt.Println("NumberError:", numError)
		}
	}
}

func test_test4() {
	test4_3("123 3")
	test4_3("123 abc")
}

func test5(a, b int) {
	if b == 0 {
		panic("b cannot be 0")
	}
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
}

func test_test5() {
	test5(9, 3)
	test5(9, 0)
}

func test6_1() {
	fmt.Println("func 1 start")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic Revoer -", r) // Recovered
		}
	}()

	test6_2() // 2->3 순서로 호출
	fmt.Println("func 1 end")
}

func test6_2() {
	fmt.Printf("9 / 3 = %d\n", test6_3(9, 3))
	fmt.Printf("9 / 0 = %d\n", test6_3(9, 0)) // panic
}

func test6_3(a, b int) int {
	if b == 0 {
		panic("b cannot be 0") //panic!
	} else {
		return a / b
	}
}

func test_test6() {
	test6_1()
	fmt.Println("Program Keeps Running")
}

func main() {
	// test_test1()
	// test_test2()
	// test_test3()
	// test_test4()
	// test_test5()
	test_test6()
}
