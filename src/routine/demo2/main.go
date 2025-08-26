package main

import (
	"fmt"
)

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from panic:", err)
		}
	}()
	var amap map[int]string
	amap[1] = "test" // This will cause a panic because amap is nil
}

func main() {
	go test()
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	strChan := make(chan string, 10)
	for i := 0; i < 10; i++ {
		strChan <- fmt.Sprintf("str%v", i)
	}
	for {
		select {
		case v := <-intChan:
			fmt.Println("intChan:", v)
		case v := <-strChan:
			fmt.Println("strChan:", v)
		}
	}
}
