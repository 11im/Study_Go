package main

import (
	"fmt"
	"sort"
)

func test1() {
	var arr = [...]int{1, 2, 3}
	var slice = []int{1, 2, 3}

	fmt.Println(arr, "\n", slice)

	var slice2 = make([]int, 3)

	fmt.Println(slice2)

	for _, v := range slice {
		v += 2
		v *= 2
		fmt.Println(v)
	}
}

func test2() {
	var slice1 = []int{1, 2, 3}
	var slice2 = []int{}
	var slice3 = []int{}

	for _, v := range slice1 {
		slice2 = append(slice2, v*2)
	}

	slice3 = append(slice1, slice2...)

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
}

func test3() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[1:4]

	fmt.Println(array)
	fmt.Println(slice)

	array[1] = 100

	fmt.Println(array)
	fmt.Println(slice, len(slice), cap(slice))

	slice = append(slice, 100)

	fmt.Println(array)
	fmt.Println(slice, len(slice), cap(slice))
}

func test4() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[1:2]
	slice2 := array[1:2:3]

	fmt.Println(slice, len(slice), cap(slice))
	fmt.Println(slice2, len(slice2), cap(slice2))
}

func test5() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, len(slice1))
	slice3 := []int{}
	slice4 := make([]int, len(slice1))

	for i, v := range slice1 {
		slice2[i] = v
	}

	slice3 = append(slice3, slice1...)

	cnt := copy(slice4, slice1) //copy하기 위해서는 Slice에 Cap이 있어야함

	slice1[1] = 100

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(cnt, slice4)
}

func test6() {
	slice := []int{1, 2, 3, 4, 5, 6}

	slice1 := make([]int, len(slice))
	copy(slice1, slice)
	slice2 := make([]int, len(slice))
	copy(slice2, slice)

	for i := 3; i < len(slice1); i++ {
		slice1[i-1] = slice1[i]
	}
	slice1 = slice1[:len(slice1)-1]
	fmt.Println(slice1)

	slice2 = append(slice2[:2], slice2[3:]...)
	fmt.Println(slice2)
}

func test7() {
	slice := []int{1, 2, 3, 4, 5, 6}
	idx := 2

	slice1 := make([]int, len(slice))
	copy(slice1, slice)
	slice2 := make([]int, len(slice))
	copy(slice2, slice)

	slice1 = append(slice1, 0)
	for i := len(slice1) - 2; i >= idx; i-- {
		slice1[i+1] = slice1[i]
	}

	slice1[idx] = 100
	fmt.Println(slice1)

	slice2 = append(slice2[:idx], append([]int{100}, slice2[idx:]...)...)
	fmt.Println(slice2)
}

type Student struct {
	Name string
	Age  int
}

type Students []Student

func (s Students) Len() int           { return len(s) }
func (s Students) Less(i, j int) bool { return s[i].Age < s[j].Age }
func (s Students) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func test8() {
	s := []int{5, 9, 7, 5, 21, 1, 11, 3, 5, 4, 8, 63, 13}
	sort.Ints(s)
	fmt.Println(s)

	s2 := []Student{
		{"정환", 27}, {"재범", 28}, {"민선", 25}}

	sort.Sort(Students(s2))

	fmt.Println(s2)

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
