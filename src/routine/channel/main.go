package main

import (
	"fmt"
)

// 多取和多放都会导致死锁阻塞
func main() {
	//写入数据
	num := 10
	var intChan chan int = make(chan int, 4)
	fmt.Println(intChan)
	intChan <- 10
	intChan <- 20
	intChan <- 30
	intChan <- num
	fmt.Println(len(intChan), cap(intChan))
	//读取数据
	num2 := <-intChan
	<-intChan
	num2 = <-intChan
	num2 = <-intChan
	fmt.Println(num2)
	fmt.Println(len(intChan), cap(intChan))
}
