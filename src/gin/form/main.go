package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./login.html", "./in.html")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	r.POST("/login", func(c *gin.Context) {
		//username := c.PostForm("username")
		//password := c.PostForm("password")
		// 设定默认值
		// username := c.DefaultPostForm("username", "someone")
		// password := c.DefaultPostForm("password", "******")
		username, ok := c.GetPostForm("username")
		if !ok {
			username = "someone"
		}
		password, ok := c.GetPostForm("password")
		if !ok {
			password = "******"
		}
		c.HTML(http.StatusOK, "in.html", gin.H{
			"username": username,
			"password": password,
		})

	})
	r.Run(":8080")
}
