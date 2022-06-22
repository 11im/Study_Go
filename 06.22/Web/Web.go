package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

//기본 http handler
func test1() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!")
	})

	http.ListenAndServe(":3000", nil)
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	name := values.Get("name")
	if name == "" {
		name = "World"
	}
	id, _ := strconv.Atoi(values.Get("id"))
	fmt.Fprintf(w, "Hello %s! id : %d", name, id)
}

//http://localhost:3000/bar?id=5&name=aaa 이런 식으로 주소 뒤에 인수를 주는 경우
func test2() {
	http.HandleFunc("/bar", barHandler)
	http.ListenAndServe(":3000", nil)
}

// 기본 주소면 Hello World, /bar이면 Hello Bar
func test3() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!")
	})

	mux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Bar!")
	})
	http.ListenAndServe(":3000", mux)
}

//파일을 보이게 함. StripPrefix는 Handle에 전해진 인수를 빼고 http 전달함. localhost:3000/static/gopher.jpg 입력 시 이미지 확인 가능
func test4() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":3000", nil)
}

func MakeWebHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!")
	})

	mux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Bar!")
	})
	return mux
}

//Test용 코드
func test5() {
	http.ListenAndServe(":3000", MakeWebHandler())
}

type Student struct {
	Name  string
	Age   int
	Score int
}

func MakeWebHandler2() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/student", StudentHandler)
	return mux
}

func StudentHandler(w http.ResponseWriter, r *http.Request) {
	var student = Student{"aaa", 16, 87}
	data, _ := json.Marshal(student)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
}

//Json
func test6() {
	http.ListenAndServe(":3000", MakeWebHandler2())
}

func main() {
	// test1()
	// test2()
	// test3()
	// test4()
	// test5()
	// test6()
}
