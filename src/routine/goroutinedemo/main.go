package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	num := runtime.NumCPU()
	fmt.Println("Number of CPUs:", num)
	runtime.GOMAXPROCS(num - 2)
	go test()
	for i := range 10 {
		fmt.Println("Hello, World!", i)
		time.Sleep(100 * time.Millisecond)
	}
}
func test() {
	for i := range 10 {
		fmt.Println("hello golang", i, " ", i)
		time.Sleep(100 * time.Millisecond)
	}
}
