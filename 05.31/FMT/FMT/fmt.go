package main

import (
	"bufio"
	"fmt"
	"os"
)

func pprint() {

	x := "ss"
	str := "Hello\tGo\tWorld\n\"GO\"is Golang\n"

	fmt.Printf("%t\n", x)
	fmt.Printf(str)
	fmt.Printf("%s", str)
	fmt.Printf("%q", str)
}

func sScan() {
	var a int
	var b int

	// n은 성공적으로 입력 받은 개수, err은 error
	n, err := fmt.Scan(&a, &b)

	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}
}

func fScan() {
	var a int
	var b int

	n, err := fmt.Scanf("%d %d\n", &a, &b)

	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}
}

func lScan() {
	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)

	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}
}

func CScan() {
	stdin := bufio.NewReader(os.Stdin)

	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)

	if err != nil {
		fmt.Println(err)
		stdin.ReadString('\n')
	} else {
		fmt.Println(n, a, b)
	}

	n, err = fmt.Scanln(&a, &b)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n, a, b)
	}
}

func main() {
	// pprint()
	// sScan()
	// fScan()
	// lScan()
	CScan()
}
