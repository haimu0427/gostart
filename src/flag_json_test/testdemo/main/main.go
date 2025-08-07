package main

import (
	"fmt"
)

func addUpper(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}

	return res
}

func main() {
	fmt.Println(addUpper(10))
}
