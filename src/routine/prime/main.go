package main

import (
	"fmt"
	"time"
)

func main() {

	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000)
	exitChan := make(chan bool, 4)
	//记录时间
	start := time.Now().Unix()
	go putNum(intChan)
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		end := time.Now().Unix()
		fmt.Printf("耗时：%d秒\n", end-start)
		close(primeChan)

	}()
	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Println(res)
	}
	close(exitChan)
	fmt.Println("所有的素数计算完毕")

}
func putNum(intChan chan int) {
	for i := 2; i < 8000; i++ {
		intChan <- i
	}
	close(intChan)
}
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	for {
		num, ok := <-intChan
		if !ok {
			break
		}
		if isPrime(num) {
			primeChan <- num
		}
	}
	exitChan <- true
}

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	if num == 2 {
		return true
	}
	if num%2 == 0 {
		return false
	}
	for i := 3; i*i <= num; i += 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}
