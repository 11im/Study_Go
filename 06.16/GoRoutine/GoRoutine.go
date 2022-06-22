package main

import (
	"fmt"
	"sync"
	"time"
)

func test1_1() {
	korean := []rune{'가', '나', '다'}
	for _, v := range korean {
		fmt.Printf("%c ", v)
		time.Sleep(500 * time.Millisecond)

	}
}

func test1_2() {
	for i := 1; i < 4; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(400 * time.Millisecond)
	}
}

func test1() {
	go test1_1()
	go test1_2()

	time.Sleep(3 * time.Second)
}

var wg sync.WaitGroup

func test2_1(a, b int) {
	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Printf("sum %d to %d : %d\n", a, b, sum)
	wg.Done()
}
func test2() {
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go test2_1(1, 10000000)
	}
	wg.Wait()
}

var mutex sync.Mutex

type Account struct {
	Balance float64
}

func test3_1(account *Account) {
	mutex.Lock()
	defer mutex.Unlock()

	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance should not be below 0 : %f", account.Balance))
	}
	account.Balance += 1000
	time.Sleep(time.Millisecond)
	account.Balance -= 1000
}

func test3() {
	var wg1 sync.WaitGroup
	account := &Account{0}
	wg1.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for {
				test3_1(account)
			}
			wg1.Done() // 에러! 닿을 수 없는 코드 -> wg.Done 작동이 안됨
		}()
	}
	wg.Wait()
}

type Job interface {
	Do()
}

type SquareJob struct {
	index int
}

func (j *SquareJob) Do() {
	fmt.Printf("Work %d Start!\n", j.index)
	time.Sleep(1 * time.Second)
	fmt.Printf("%d Work Finished! - %d\n", j.index, j.index*j.index)
}

func test4() {
	var jobList [10]Job

	for i := 0; i < 10; i++ {
		jobList[i] = &SquareJob{i}
	}

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		job := jobList[i]
		go func() {
			job.Do()
			wg.Done()
		}()
	}
	wg.Wait()
}

func main() {
	// test1()
	// test2()
	// test3()
	test4()
}
