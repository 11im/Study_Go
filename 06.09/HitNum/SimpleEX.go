package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func gen_rand() int {
	rand.Seed(time.Now().UnixNano()) //현재 시간을 시드로

	rand := rand.Intn(100) // 0~99의 무작위 정수

	return rand
}

var stdin = bufio.NewReader(os.Stdin)

func insert_num() (int, error) {
	var num int
	_, err := fmt.Scanln(&num)

	if err != nil {
		stdin.ReadString('\n')
	}
	return num, err
}

func main() {

	var target_num int = gen_rand()
	var count int = 0

	for {
		fmt.Printf("Insert Num : ")

		num, err := insert_num()

		if err != nil { //에러 검사
			fmt.Println("Only Number PLZ")
		}

		if num == target_num { //맞추면
			fmt.Println("Hit! Total Count :", count)
			break
		} else if num > target_num {
			count++
			fmt.Println("Your Number is Bigger! You've Tried : ", count)
		} else if num < target_num {
			count++
			fmt.Println("Your Number is Smaller! You've Tried : ", count)
		}
	}
}
