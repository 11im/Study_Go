package main

import (
	"bufio"
	"fmt"
	"os"
)

func test1() {
	for i := 0; i < 10; i++ {
		fmt.Print(i, ", ")
	}

	for x := 0; x > 1; {
		//code
	}
	x := 1
	for x > 1 {
		//code
	}

	for x > 1 {
		//code
	}

	for true {
		//code
	}

	for {
		//code
	}
}

func test2() {
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("입력!")
		var num int
		_, err := fmt.Scanln((&num))
		if err != nil {
			fmt.Println("숫자!")

			//Clean Buffer
			stdin.ReadString('\n')
			continue
		}
		fmt.Printf("입력 : %d\n", num)
		if num%2 == 0 {
			break
		}
	}
	fmt.Println("끝!")
}

func flag_for() {
	a := 1
	b := 1

	found := false
	for ; a <= 9; a++ {
		for b = 1; b <= 9; b++ {
			if a*b == 45 {
				found = true // ❶ 찾았음을 표시하고 break
				break
			}
		}
		if found { // ❷ 바깥쪽 for문에서 찾았는지 검사해서 break
			break
		}
	}
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
}

func label_for() {
	a := 1
	b := 1
OuterFor:
	for ; a <= 9; a++ {
		for b = 1; b <= 9; b++ {
			if a*b == 45 {
				break OuterFor
			}
		}
	}
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
}

func find_45(a int) (int, bool) {
	for b := 1; b <= 9; b++ {
		if a*b == 45 {
			return b, true
		}
	}
	return 0, false
}

func func_for() {
	a := 1
	b := 0

	for ; a <= 9; a++ {
		var found bool
		if b, found = find_45(a); found {
			break
		}
	}
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
}

func main() {
	//test1()
	//test2()

	flag_for()
	label_for()
	func_for()
}
