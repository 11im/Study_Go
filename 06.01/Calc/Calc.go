package main

import (
	"fmt"
	"math"
	"math/big"
)

func equal(a, b float64) bool {
	return math.Nextafter(a, b) == b
	//Nextafter로 마지막 비트 비교, 작은 차이는 무시
}

func test_equal() {
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	fmt.Printf("%0.18f + %0.18f = %0.18f\n", a, b, a+b)
	fmt.Printf("%0.18f == %0.18f : %v\n", c, a+b, equal(a+b, c))

	a = 0.0000000004
	b = 0.0000000002
	c = 0.0000000007

	fmt.Printf("%g == %g :%v\n", c, a+b, equal(a+b, c))
}

func big_float() {
	a, _ := new(big.Float).SetString("0.1")
	b, _ := new(big.Float).SetString("0.2")
	c, _ := new(big.Float).SetString("0.3")

	d := new(big.Float).Add(a, b)

	fmt.Println(a, b, c, d)
	fmt.Println(c.Cmp(d))
	// a, b 모두 big.Float일 경우, a.Cmp(b)를 통해 비교 가능, a가 b보다 작으면 -1, 같으면 0, 크면 1
}

func Logic() {
	fmt.Println(5 < 8 && 2 >= 3)
	fmt.Println(5 < 8 && 2 <= 3)
	fmt.Println(5 < 8 || 2 >= 3)
	fmt.Println(!(5 < 8 || 2 >= 3))
}

func Sub() {

	var a int
	a = 10
	fmt.Println(a)
	var b int
	var c int
	b = 10
	c = b
	fmt.Println(b, c)
	//c = b = 10 -> b = 10은 return 값이 없기 때문에 c에 대입 불가.

	var d, e int
	d, e = 3, 4
	fmt.Println(d, e)

	var x int = 10
	var y int = 20
	fmt.Println(x, y)

	x, y = y, x
	fmt.Println(x, y)

	var z int = 10
	z = z + 2
	fmt.Println(z)
	z += 2
	fmt.Println(z)
	z++
	fmt.Println(z)
	z--
	fmt.Println(z)
}

func main() {
	test_equal()
	big_float()
	Logic()
	Sub()
}
