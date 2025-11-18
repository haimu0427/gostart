package main

// 现在我们要编写一个程序，计算一个数的阶乘
import "fmt"

func main() {
	var n int
	fmt.Println("请输入一个整数:")
	fmt.Scanln(&n)
	fmt.Println("这个数的阶乘是:", factorial(n))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
