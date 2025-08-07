package main

import (
	"fmt"
)

type cat struct {
	name  string
	age   int
	color string
}

func main() {
	// 创建一个有缓冲的通道
	// 并存储不同类型的数据
	var allChan chan any = make(chan any, 10)
	cat1 := cat{
		name:  "Tom",
		age:   3,
		color: "white",
	}
	cat2 := cat{
		name:  "Jerry",
		age:   2,
		color: "black",
	}
	allChan <- cat1
	allChan <- cat2
	allChan <- "Hello, World!"
	allChan <- 100

	close(allChan) // 关闭通道

	cat11 := <-allChan
	fmt.Println(cat11.(cat).name)
	newcat := <-allChan
	a := newcat.(cat)
	fmt.Println(a.name)

}
