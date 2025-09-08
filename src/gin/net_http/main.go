package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

var a int = 0

func sayHello(w http.ResponseWriter, r *http.Request) {
	file, _ := os.ReadFile("./hello.txt")
	_, _ = fmt.Fprintln(w, string(file))
	fmt.Println("hello golang", a)
	a++
}

func main() {
	resp, err := http.Get("https://www.liwenzhou.com")
	if err != nil {
		fmt.Println("http.Get err:", err)
		return
	}
	defer resp.Body.Close()
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("io.ReadAll err:", err)
		return
	}
	http.HandleFunc("/hello", sayHello)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("http server start err:", err)
		return
	}
	//fmt.Println(string(body))
}
