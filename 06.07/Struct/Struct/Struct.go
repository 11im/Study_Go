package main

import (
	"fmt"
)

type Student struct { //대문자로 선언하면 package 외부로 노출됨
	Name  string
	Class int
	Age   int
	No    int
	Score float64
}

func test1() {
	var a Student
	a.Name = "임정환"
	a.Class = 2
	a.Age = 27
	a.No = 222
	a.Score = 3.7

	fmt.Println("이름 : ", a.Name)
	fmt.Println("반 : ", a.Class)
	fmt.Println("나이 : ", a.Age)
	fmt.Println("번호 : ", a.No)
	fmt.Println("성적 : ", a.Score)
}

func test2() {
	var a Student

	fmt.Println("이름 : ", a.Name)
	fmt.Println("반 : ", a.Class)
	fmt.Println("나이 : ", a.Age)
	fmt.Println("번호 : ", a.No)
	fmt.Println("성적 : ", a.Score)

	a = Student{"임정환", 2, 27, 222, 3.7}

	fmt.Println("이름 : ", a.Name)
	fmt.Println("반 : ", a.Class)
	fmt.Println("나이 : ", a.Age)
	fmt.Println("번호 : ", a.No)
	fmt.Println("성적 : ", a.Score)
}

func test3() {
	var a = Student{Name: "임정환", Score: 3.7}

	fmt.Println("이름 : ", a.Name)
	fmt.Println("반 : ", a.Class)
	fmt.Println("나이 : ", a.Age)
	fmt.Println("번호 : ", a.No)
	fmt.Println("성적 : ", a.Score)
}

type University struct {
	Studentinfo   Student
	Major         string
	StudentNumber int
}

type SeoulUniv struct {
	UnivInfo University
	Address  string
}

func test4() {
	stud := Student{"임정환", 27, 27, 27, 2.7}
	univ := University{stud, "DS", 10}
	su := SeoulUniv{univ, "공릉동"}

	fmt.Printf("%s, %d, %d, %d, %f, %s, %d, %s \n",
		su.UnivInfo.Studentinfo.Name,
		su.UnivInfo.Studentinfo.No,
		su.UnivInfo.Studentinfo.Age,
		su.UnivInfo.Studentinfo.Class,
		su.UnivInfo.Studentinfo.Score,
		su.UnivInfo.Major,
		su.UnivInfo.StudentNumber,
		su.Address,
	)
}

func test5() {
	stud1 := Student{"임정환", 27, 27, 27, 2.7}

	fmt.Printf("%s, %d, %d, %d, %f\n",
		stud1.Name,
		stud1.Class,
		stud1.No,
		stud1.Age,
		stud1.Score,
	)

	stud2 := stud1

	fmt.Printf("%s, %d, %d, %d, %f\n",
		stud2.Name,
		stud2.Class,
		stud2.No,
		stud2.Age,
		stud2.Score,
	)
}
func main() {
	// test1()
	// test2()
	// test3()
	// test4()
	test5()
}
