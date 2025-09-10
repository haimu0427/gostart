package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name     string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

func main() {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		//bind query string
		var user User
		err := c.ShouldBindJSON(&user) // Use ShouldBindJSON for JSON data
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {
			fmt.Println(user, "bind by ShouldBindJSON")
		}

		// username := c.Query("username")
		// password := c.Query("password")
		// user1 := User{
		// 	Name:     username,
		// 	Password: password,
		// }
		// fmt.Println(user1)
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})
	r.Run(":8080")
}
