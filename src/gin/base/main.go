package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, h *http.Request) {
	t, err := template.ParseFiles("./base.tmpl", "./index.tmpl")
	if err != nil {
		fmt.Println("err is:", err)
		return
	}
	msg := "Welcome to the Block Index Page"
	t.Execute(w, msg)
}
func main() {
	http.HandleFunc("/index", index)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("err is:", err)
		return
	}
	// Start the server

}
