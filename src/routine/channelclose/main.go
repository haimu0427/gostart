package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	intChan := make(chan int, 50)
	go writeDate(intChan, &wg)
	go readDate(intChan, &wg)
	wg.Wait() // 等待所有 goroutine 完成
}
func writeDate(intChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 50; i++ {
		intChan <- i
		fmt.Println("写入数据:", i)
	}
	close(intChan)
}

func readDate(intChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range intChan {
		fmt.Println("读取数据:", v)
	}
}
