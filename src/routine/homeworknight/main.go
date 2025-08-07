package main

import (
	"fmt"
	"sync"
)

type resMap struct {
	sync.Mutex
	m map[int]int
}

func main() {
	//同步工具使用
	var wg sync.WaitGroup
	//声明数据

	var numChan = make(chan int, 2000)
	var resChan = make(chan int, 2000)
	// 使用并发安全的 map 实现
	var resNum = resMap{m: make(map[int]int)}
	//操作主体
	master := 1
	worker := 8
	wg.Add(master) // 8个popChan + 1个pushChan
	go pushChan(numChan, &wg)
	wg.Add(worker)
	for i := 0; i < worker; i++ {
		go popChan(numChan, resChan, &resNum, &wg)
	}

	//遍历输出
	wg.Wait()
	// 关闭结果通道
	close(resChan)
	resNum.Lock()
	for i, v := range resNum.m {
		fmt.Println("sum of ", i, " is ", v)
	}
	resNum.Unlock()
}
func pushChan(numChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 2000; i++ {
		numChan <- i
	}
}
func popChan(numChan chan int, resChan chan int, resNum *resMap, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range numChan {
		resNum.Lock()
		resChan <- calSum(v, resNum.m)
		resNum.Unlock()
	}
}
func calSum(n int, resNum map[int]int) int {

	if n == 1 {
		resNum[1] = 1
		return 1
	} else {
		resNum[n] = resNum[n-1] + n
	}
	return resNum[n]
}
