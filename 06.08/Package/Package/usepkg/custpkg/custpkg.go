package custpkg

import "fmt"

const ( //대문자로 시작하는 Const만 공개됨
	PI = 3.14
	pi = 3.14
)

var ScreenSize int = 1080 //변수도 마찬가지
var screeninch int = 27

func PublicFunc() {
	const MyConst = 100 //함수 내에서 선언된 Const는 공개 안됨
	fmt.Println("Public Function", MyConst)
}

func privatefunc() { //함수도
	fmt.Println("Private Function")
}

type Myint int
type mystring string //type도

type MyStruct struct {
	Age  int
	name string //공개되는 구조체 내에서도 소문자로 시작하면 외부 공개 안됨
}

type myprivatestruct struct {
	Age  int //공개되지 않는 구조체의 필드이기 때문에 공개 안됨
	name string
}

func (m MyStruct) PublicMethod() {
	fmt.Println("Public Method")
}

func (m MyStruct) privatemethod() { //Method도 마찬가지, 소문자로 시작하면 공개 안됨
	fmt.Println("Private Method")
}

//Print Custom
func PrintCustom() {
	fmt.Println("This is Custom Package")
}

func init() { //init() 함수는 사용 여부에 관계 없이 해당 패키지의 전역 변수 초기화 이후 수행됨
	fmt.Println("Initiating CustPKG", ScreenSize)
}
