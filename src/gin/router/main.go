package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//访问什么,
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	// 写写路由组
	videoGroup := r.Group("/video")
	{
		videoGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "This is the video index",
			})

		})
		videoGroup.POST("/upload", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Video uploaded successfully",
			})
		})
		videoGroup.PUT("/update", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Video updated successfully",
			})
		})
	}
	r.Any("/any", func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodGet:
			c.JSON(http.StatusOK, gin.H{
				"message": "This is a GET request",
			})
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{
				"message": "This is an POST request",
			})
		case http.MethodPut:
			c.JSON(http.StatusOK, gin.H{
				"message": "This is a PUT request",
			})
		}
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "404 not found",
		})
	})

	r.Run(":8080")
}
