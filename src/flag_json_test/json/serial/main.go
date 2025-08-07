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

func main() {
	// testStruct()
	// testMap()
	// testSlice()
	testFloat64()
}
func testStruct() {
	monster := Monster{
		Name:     "牛魔王",
		Age:      500,
		Birthday: "2011-11-11",
		Sal:      8000.0,
		Skill:    "牛魔拳",
	}
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println("json marshal err=", err)
		return
	}
	fmt.Printf("json marshal data=%v\n", string(data))
}
func testMap() {
	var m map[string]any
	m = make(map[string]any)
	m["name"] = "红孩儿"
	m["age"] = 18
	m["adress"] = "火焰山火云洞"
	data, err := json.Marshal(&m)
	if err != nil {
		fmt.Println("json marshal err =", err)
		return
	}
	fmt.Println("json marshal data=", string(data))
}
func testSlice() {

	var slice []map[string]any
	slice = make([]map[string]any, 3)
	slice[0] = make(map[string]any)
	slice[0]["name"] = "孙悟空"
	slice[0]["age"] = 18
	slice[0]["adress"] = "花果山"
	var m1 map[string]any
	m1 = make(map[string]any)
	m1["name"] = "猪八戒"
	m1["age"] = 28
	m1["adress"] = "高老庄"
	slice = append(slice, m1)
	slice[2] = make(map[string]any)
	slice[2]["name"] = "沙和尚"
	slice[2]["age"] = 38
	slice[2]["adress"] = "流沙河"
	data, err := json.Marshal(&slice)
	if err != nil {
		fmt.Println("json marshal err =", err)
		return
	}
	fmt.Println("json marshal data=", string(data))
}
func testFloat64() {
	var num1 float64 = 2345.67
	data, err := json.Marshal(&num1)
	if err != nil {
		fmt.Println("json marshal err =", err)
		return
	}
	fmt.Println("json marshal data=", string(data))
}
