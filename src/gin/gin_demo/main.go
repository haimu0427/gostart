package main

import (
	_ "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/thinkerou/favicon"
)

func main() {
	// 创建默认的路由引擎
	r := gin.Default()
	// 处理请求
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "<h1>Hello Gin</h1>",
		})
	})
	r.POST("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "<h1>Hello Gin POST</h1>",
		})
	})
	r.PUT("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "<h1>Hello Gin PUT</h1>",
		})
	})
	r.DELETE("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "<h1>Hello Gin DELETE</h1>",
		})
	})
	// 启动HTTP服务，默认在8080端口启动服务
	r.Run(":8080")
}
