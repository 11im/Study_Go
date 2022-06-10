package main

import (
	"fmt"
	"time"
)

/*

func (리시버) 메서드명() 반환타입{
	...
}

ex)
func (r Rabbit) info() int{
	return r.width * r.height
}

*/

type account struct {
	balance int
}

func withdrawFunc(a *account, amount int) { //일반 함수
	a.balance -= amount
}

func (a *account) withdrawMethod(amount int) {
	a.balance -= amount
}

func test1() {
	a1 := account{1000}
	a2 := account{1000}

	withdrawFunc(&a1, 100)
	a2.withdrawMethod(100)

	fmt.Println(a1.balance, a2.balance)

}

type myInt int

func (i myInt) add(j int) int {
	return int(i) + j
}

func test2() {
	var a myInt = 10
	fmt.Println(a.add(50))

	var b int = 10
	fmt.Println(myInt(b).add(50))
}

func (a account) withdrawValue(amount int) {
	a.balance -= amount
}

func (a account) withdrawReturnValue(amount int) account {
	a.balance -= amount
	return a
}

func test3() {
	var a1 *account = &account{100}
	a1.withdrawMethod(30)
	fmt.Println(a1.balance)
	a1.withdrawValue(30)
	var a2 account = a1.withdrawReturnValue(30)
	fmt.Println(a1.balance)
	fmt.Println(a2.balance)
}

type Courier struct {
	Name string
}

type Product struct {
	Name  string
	Price int
	ID    int
}

type Parcel struct {
	Pdt           *Product
	ShippedTime   time.Time
	DeliveredTime time.Time
}

func (c *Courier) SendProduct(pdt *Product) *Parcel {
	var parcel = &Parcel{}
	parcel.Pdt = pdt
	parcel.ShippedTime = time.Now()
	return parcel
}

func (parcel *Parcel) Delivered() *Product {
	parcel.DeliveredTime = time.Now()
	return parcel.Pdt
}

func main() {
	// test1()
	// test2()
	test3()
}
