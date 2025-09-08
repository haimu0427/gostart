package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, h *http.Request) {
	t, err := template.ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Println("err is:", err)
		return
	}
	msg := "Welcome to the Block Index Page"
	t.Execute(w, msg)

}
func home(w http.ResponseWriter, h *http.Request) {
	t, err := template.ParseFiles("./home.tmpl")
	if err != nil {
		fmt.Println("err is:", err)
		return
	}
	msg := "Welcome to the Block Home Page"
	t.Execute(w, msg)
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("err is:", err)
		return
	}

}
