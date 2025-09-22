package main

import (
	"fmt"
	"sync"
)

var (
	JiechengNum map[int]int = make(map[int]int, 20)
	Lock        sync.Mutex
	//synchorinzed 同步
	//Mutexsync.Mutex //互斥锁
)

func main() {
	for i := 1; i <= 200; i++ {
		//并发错误,共同读取同一个内存空间fatal error: concurrent map writes
		go jiecheng(i)
	}
	for i, v := range JiechengNum {
		fmt.Println("Factorial of", i, "is", v)
	}
}
func jiecheng(n int) {
	Lock.Lock()
	if n == 1 {
		JiechengNum[n] = 1
	} else {
		JiechengNum[n] = JiechengNum[n-1] * n
	}
	Lock.Unlock()
}
