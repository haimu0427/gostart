package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO",
			"tag":  "<br>",
		}
		var msg struct {
			Name    string `json:"name"`
			Message string `json:"message"`
			Age     int    `json:"age"`
		}
		msg.Name = "Lena"
		msg.Message = "hello"
		msg.Age = 18
		c.JSON(http.StatusOK, msg)
		c.JSON(http.StatusOK, data)

	})
	r.Run(":8080")
}
