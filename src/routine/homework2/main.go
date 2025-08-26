package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"sync"
	"time"
)

func main() {

	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	file, err := os.OpenFile("data.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	//
	var wg sync.WaitGroup
	var producter int = 1
	var consumer int = 1

	wg.Add(producter + consumer)
	go writeDataTofile(file, &wg)
	go sortData(file, &wg)

	wg.Wait()

}
func writeDataTofile(file *os.File, wg *sync.WaitGroup) {
	defer wg.Done()
	writer := bufio.NewWriter(file)
	for i := 0; i < 1000; i++ {
		data := rand.Int()
		_, err := fmt.Fprintln(writer, data)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
	writer.Flush()
}
func sortData(file *os.File, wg *sync.WaitGroup) {
	defer wg.Done()
	var data []int = make([]int, 0, 1000)
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println("End of file reached")
				break
			}
			fmt.Println("Error reading from file:", err)
			return
		}
		str := string(line)
		str = str[:len(str)-1] // 去掉换行符
		if str == "" {
			continue
		}
		i, _ := strconv.Atoi(str)
		data = append(data, i)
	}
	sort.Ints(data)
	writer := bufio.NewWriter(file)
	for _, v := range data {
		_, err := fmt.Fprintln(writer, v)
		if err != nil {
			fmt.Println("Error writing sorted data to file:", err)
			return
		}
	}
	writer.Flush()

}
