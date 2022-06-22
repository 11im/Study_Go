package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func test1_1(wg *sync.WaitGroup, ch chan int) {
	n := <-ch
	time.Sleep(time.Second)
	fmt.Printf("%d\n", n*n)
	wg.Done()
}

func test1() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go test1_1(&wg, ch)
	}
	ch <- 9
	ch <- 3
	ch <- 4
	wg.Wait()
}

func test2_1(wg *sync.WaitGroup, ch chan int) {
	for n := range ch {
		fmt.Printf("%d\n", n*n)
		time.Sleep(time.Second)
	}
	wg.Done()
}

func test2() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go test2_1(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	close(ch)
	wg.Wait()
}

func test3_1(wg *sync.WaitGroup, ch chan int, quit chan bool) {
	for {
		select {
		case n := <-ch:
			fmt.Printf("%d\n", n*n)
			time.Sleep(time.Second)
		case <-quit:
			wg.Done()
			return
		}
	}
}

func test3() {
	var wg sync.WaitGroup
	ch := make(chan int)
	quit := make(chan bool)

	wg.Add(1)
	go test3_1(&wg, ch, quit)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}

	quit <- true
	wg.Wait()
}

func test4_1(wg *sync.WaitGroup, ch chan int) {
	tick := time.Tick(time.Second)
	terminate := time.After(10 * time.Second)

	for {
		select {
		case <-tick:
			fmt.Println("Tick")
		case <-terminate:
			fmt.Println("Terminated")
			wg.Done()
			return
		case n := <-ch:
			fmt.Printf("%d\n", n*n)
			time.Sleep(time.Second)
		}
	}
}

func test4() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go test4_1(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}

	wg.Wait()
}

type Car struct {
	Body  string
	Tire  string
	Color string
}

var wg sync.WaitGroup
var startTime = time.Now()

func test5_1(tireCh chan *Car) {
	tick := time.Tick(time.Second)
	after := time.After(10 * time.Second)

	for {
		select {
		case <-tick:
			car := &Car{}
			car.Body = "Sports Car"
			tireCh <- car
		case <-after:
			close(tireCh)
			wg.Done()
			return
		}
	}
}

func test5_2(tireCh, paintCh chan *Car) {
	for car := range tireCh {
		time.Sleep(time.Second)
		car.Tire = "Winter tire"
		paintCh <- car
	}
	wg.Done()
	close(paintCh)
}

func test5_3(paintch chan *Car) {
	for car := range paintch {
		time.Sleep(time.Second)
		car.Color = "Red"
		duration := time.Now().Sub(startTime)
		fmt.Printf("%.2f Complete Car : %s, %s, %s\n", duration.Seconds(), car.Body, car.Color, car.Tire)
	}
	wg.Done()
}

func test5() {
	tireCh := make(chan *Car)
	paintCh := make(chan *Car)

	fmt.Printf("Start\n")

	wg.Add(3)
	go test5_1(tireCh)
	go test5_2(tireCh, paintCh)
	go test5_3(paintCh)

	wg.Wait()
	fmt.Println("Close")
}

func test6() {
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	go test6_1(ctx)
	time.Sleep(5 * time.Second)
	cancel()

	wg.Wait()
}

func test6_1(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-ctx.Done():
			wg.Done()
			return
		case <-tick:
			fmt.Println("Tick")
		}
	}

}

func test7() {
	wg.Add(1)

	ctx := context.WithValue(context.Background(), "number", 9) //
	go test7_1(ctx)

	wg.Wait()
}

func test7_1(ctx context.Context) {
	if v := ctx.Value("number"); v != nil {
		n := v.(int)
		fmt.Printf("%d", n*n)
	}
	wg.Done()
}

func main() {
	// test1()
	// test2()
	// test3()
	// test4()
	// test5()
	// test6()
	test7()
}
