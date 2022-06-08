package main

import (
	"fmt"
	Template "html/template"
	_ "math/rand"
	"text/template"

	custpkg "module_test/Package/usepkg/custpkg"

	"github.com/guptarohit/asciigraph"
	"github.com/tuckersGo/musthaveGo/ch16/expkg"
)

func test1() {
	template.New("foo").Parse(`{{define "T"}}HELLO`)
	Template.New("foo").Parse(`{{define "T"}}HELLO`)
}

func test2() {
	custpkg.PrintCustom()
	expkg.PrintSample()
	data := []float64{3, 4, 5, 6, 9, 7, 5, 8, 5, 10, 2, 7, 2, 5, 6}
	graph := asciigraph.Plot(data)
	fmt.Println(graph)
}

func test3() {
	fmt.Println(custpkg.PI)
	custpkg.PublicFunc()

	var myint custpkg.Myint = 10
	fmt.Println(myint)

	var mystruct = custpkg.MyStruct{Age: 10}
	fmt.Println(mystruct.Age)

}

func main() {
	// test2()
	test3()
}
