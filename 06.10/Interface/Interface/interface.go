package main

import (
	fedex "Interface/Interface/Parcel/fedex"
	koreanpost "Interface/Interface/Parcel/koreanpost"
	"fmt"
)

/*

type 이름 interface{
	Method1()				//{} 내부에 Method 집합
	Method2(i int) int
}

ex)

type DuckInterface interface{
	Fly()
	Walk (distance float64) float64
}

*/

type Stringer interface {
	String() string
}

type Student struct {
	Name string
	Age  int
}

func (s Student) String() string {
	return fmt.Sprintf("Name : %s Age : %d", s.Name, s.Age)
}

func test1() {
	A := Student{"정환", 27}

	str := A.String()

	fmt.Println(str)
}

// func SendBook(name string, sender *fedex.FedexSender) {
// 	sender.Send(name)
// }

type Sender interface {
	Send(parcel string)
}

func SendBook(parcel string, sender Sender) {
	sender.Send(parcel)
}

func ex1() {
	sender := &fedex.FedexSender{}
	SendBook("어린 왕자", sender)
	SendBook("그리스인 조르바", sender)
}

func ex2() { // Error! Sendbook은 fedex.FedexSender를 인수로 받음!
	sender := &koreanpost.PostSender{}
	SendBook("어린 왕자", sender)
	SendBook("그리스인 조르바", sender)
}

func ex3() {
	ex1()
	ex2()
}

func PrintAge(stringer Stringer) {
	s := stringer.(*Student)
	fmt.Println(s.Age)
}

func test2() {
	s := &Student{Age: 15}
	PrintAge(s)
}

func main() {
	// test1()
	// ex1()
	// ex2()
	// ex3()
	test2()
}
