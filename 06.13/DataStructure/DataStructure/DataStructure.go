package main

import (
	Q "DataStructure/DataStructure/Queue"
	Stack "DataStructure/DataStructure/Stack"
	"container/list"
	"fmt"
)

type Element struct { //Linked List
	Value interface{}
	Next  *Element
	Prev  *Element
}

func test1() {
	v := list.New()
	e4 := v.PushBack(4)
	e1 := v.PushBack(1)
	v.InsertBefore(3, e4)
	v.InsertAfter(2, e1)

	for e := v.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}

	fmt.Println()
	for e := v.Back(); e != nil; e = e.Prev() {
		fmt.Print(e.Value, " ")
	}
}

func test2() {
	queue := Q.NewQueue()

	for i := 1; i < 5; i++ {
		queue.Push(i)
	}

	v := queue.Pop()
	for v != nil {
		fmt.Printf("%v ->", v)
		v = queue.Pop()
	}
}

func test3() {
	stack := Stack.NewStack()

	for i := 1; i < 5; i++ {
		stack.Push(i)
	}

	val := stack.Pop()
	for val != nil {
		fmt.Printf("%v ->", val)
		val = stack.Pop()
	}
}
func main() {
	// test1()
	// test2()
	test3()
}
