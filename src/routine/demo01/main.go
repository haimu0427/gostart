package main

import (
	"fmt"
)

func main() {
	//通过函数传参的时候设置管道的使用权限
	fmt.Println("Hello, World!")
	//双向通道
	var chan1 chan int
	chan1 = make(chan int, 3)
	chan1 <- 1000
	//只写
	var chan2 chan<- int
	chan2 = make(chan int, 3)
	chan2 <- 20
	//只读
	var chan3 <-chan int
	chan3 = chan1
	fmt.Println(<-chan3)
}
