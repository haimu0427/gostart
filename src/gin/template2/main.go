package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	//定义函数 kua
	//要么只有一个返回值，要么两个返回值且第二个返回值是error
	kua := func(name string) (string, error) {
		return name + "酷毙了", nil
	}

	// 1. 定义模板
	//需要在解析模板之前注册函数
	t := template.New("f.tmpl")
	t.Funcs(template.FuncMap{
		"kua": kua,
	})
	// 2. 解析模板
	_, err := t.ParseFiles("./f.tmpl", "./ul.tmpl") //这里a引用了b,所以要把b放在后面,同时也解析进来
	if err != nil {
		fmt.Println("Failed to parse template:", err)
		return
	}
	// 3. 渲染模板
	// 3.1 准备数据
	name := "小王子"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Println("Failed to render template:", err)
	}

}
func main() {
	http.HandleFunc("/", f1)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Failed to start server:", err)
		return
	}

}
