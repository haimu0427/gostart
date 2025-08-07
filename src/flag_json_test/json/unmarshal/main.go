package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Birthday string  `json:"birthday"`
	Sal      float64 `json:"sal"`
	Skill    string  `json:"skill"`
}

func unmarshalStruct() {
	//str一般是通过网络传输过程中读取到的
	str := "{\"name\":\"牛魔王\",\"age\":500,\"birthday\":\"2011-11-11\",\"sal\":8000,\"skill\":\"牛魔拳\"}"
	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Println("json unmarshal err = ", err)
		return
	}
	fmt.Println("unmarshal result = ", monster)
}
func unmarshalMap() {
	str := "{\"address\":\"火焰山火云洞\",\"age\":18,\"name\":\"红孩儿\"}"
	var m map[string]any
	err := json.Unmarshal([]byte(str), &m) //反序列化底层会make一个map
	//make操作封装到了json.Unmarshal函数中
	if err != nil {
		fmt.Println("json unmarshal err = ", err)
		return
	}
	fmt.Println("unmarshal map = ", m)

}
func unmarshalSlice() {
	str := "[{\"name\":\"牛魔王\",\"age\":500,\"birthday\":\"2011-11-11\"," +
		"\"sal\":8000,\"skill\":\"牛魔拳\"},{\"name\":\"红孩儿\",\"age\":18,\"birthday\":\"2012-12-12\",\"sal\":5000,\"skill\":\"火云邪神\"}]"
	var slice []Monster
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Println("json unmarshal err = ", err)
		return
	}
	fmt.Println("unmarshal slice = ", slice)
}
func main() {
	unmarshalStruct()
	unmarshalMap()
	unmarshalSlice()
}
