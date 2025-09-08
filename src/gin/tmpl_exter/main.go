package main

//永远都不要相信你的用户输入的数据
//防止XSS攻击
//1.模板引擎会自动帮你转义html标签
//2.如果你想让某些内容不被转义，可以使用template.HTML类型
import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, h *http.Request) {
	// 1.创建模板，并设置定界符
	t := template.New("index.tmpl").Delims("{[", "]}")
	// 2.解析模板文件
	t, err := t.ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Println("err is:", err)
		return
	}
	// 3.渲染模板
	// 传入数据
	msg := "Welcome to the External Template Index Page"
	err = t.Execute(w, msg)
	if err != nil {
		fmt.Println("err is:", err)
		return
	}
}
func xss(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("xss.tmpl").Funcs(template.FuncMap{
		"safe": func(s string) template.HTML {
			return template.HTML(s)
		},
	}).ParseFiles("./xss.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	jsStr := `<script>alert('嘿嘿嘿')</script>`
	// 在模板文件 xss.tmpl 中使用 {{ safe . }} 来调用 safe 函数
	err = tmpl.Execute(w, jsStr)
	if err != nil {
		fmt.Println(err)
	}
}
func xxs(w http.ResponseWriter, h *http.Request) {
	t, err := template.ParseFiles("./xxs.tmpl")
	if err != nil {
		fmt.Println("err is:", err)
		return
	}
	// 使用template.HTML类型来绕过HTML转义，允许XSS攻击（仅用于学习目的）
	msg := template.HTML("<script>alert('XSS攻击成功!');</script>")
	err = t.Execute(w, msg)
	if err != nil {
		fmt.Println("err is:", err)
		return
	}
}
func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/xss", xss)
	http.HandleFunc("/xxs", xxs)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("err is:", err)
		return
	}
}
