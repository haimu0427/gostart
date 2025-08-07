package util

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

func (m Monster) Store() {
	//序列化过程
	data, err := json.Marshal(&m)
	if err != nil {
		fmt.Println("json marshal err=", err)
		return
	}
	fmt.Println("json marshal succ:", string(data))
	//将序列化后的数据写入文件
	file, err := os.OpenFile("in.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file err= ", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(string(data))
	if err != nil {
		fmt.Println("write file err = ", err)
		return
	}
	err = writer.Flush()
	if err != nil {
		fmt.Println("write file err = ", err)
		return
	}
}
func (m Monster) Restore() {
	//读取文件
	content, err := os.ReadFile("in.txt")
	if err != nil {
		fmt.Println("read file err = ", err)
		return
	}
	fmt.Println("file content:", string(content))
	//反序列化过程
	err1 := json.Unmarshal([]byte(content), &m)
	if err1 != nil {
		fmt.Println("json unmarshal err = ", err1)
		return
	}
	fmt.Println("json unmarshal succ:", m)

}
