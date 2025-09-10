package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/redirect", func(c *gin.Context) {
		// c.JSON(http.StatusOK, gin.H{
		// 	"message": "This is a redirect endpoint",
		// })
		c.Redirect(http.StatusMovedPermanently, "https://www.sogo.com/")
		r.Run(":8080")
	})
	r.GET("/a", func(c *gin.Context) {
		c.Request.URL.Path = "/b"
		r.HandleContext(c)
	})
	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "This is endpoint B",
		})
	})
	r.Run(":8080")
}
