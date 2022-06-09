package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	Balnace = 1000
	Earn    = 500
	Fail    = 100
	Victory = 5000
	Lose    = 0
)

var stdin = bufio.NewReader(os.Stdin)

func gen_rand() int {
	rand.Seed(time.Now().UnixNano())

	rand := rand.Intn(5) + 1

	return rand
}

func get_num() (int, error) {
	var num int
	_, err := fmt.Scanln(&num)

	if err != nil {
		stdin.ReadString('\n')
	}

	return num, err
}

func main() {
	money := Balnace
	fmt.Println("Welcome to Slot Machine! You Start with ", money)

	for {

		fmt.Printf("Insert Numer : ")
		target_num := gen_rand()
		input_num, err := get_num()

		if err != nil || input_num > 5 || input_num < 1 {
			fmt.Println("Only Number 1~5 Plz")
			continue
		}

		if target_num == input_num {
			money += Earn
			fmt.Println("Congraturation! Now You Have ", money)
		} else {
			money -= Fail
			fmt.Println("I'm Sorry. Now You Have ", money)
		}

		if money >= Victory || money <= Lose {
			break
		}
	}

}
