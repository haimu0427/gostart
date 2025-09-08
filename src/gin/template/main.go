package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name   string
	Age    int
	Gender string
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	//1. 定义模板

	//2. 解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("Failed to parse template:", err)
		return
	}
	//3. 渲染模板
	//3.1 准备数据
	u1 := User{
		Name:   "小王子",
		Age:    20,
		Gender: "男",
	}

	m1 := map[string]any{
		"title": "这是一个标题",
		"count": 100,
		"flag":  true,
	}
	hobbyList := []string{"看书", "旅游", "编程"}
	// 3.2 渲染
	err = t.Execute(w, map[string]any{
		"user":  u1,
		"data":  m1,
		"hobby": hobbyList,
	})
	if err != nil {
		fmt.Println("Failed to render template:", err)
		return
	}

}
func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Failed to start server:", err)
		return
	}

}
