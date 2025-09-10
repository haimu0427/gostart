package main

//querystring
import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/web", func(c *gin.Context) {
		//querystring
		//name := c.DefaultQuery("query", "someone")
		//name, ok := c.GetQuery("query") // same as above
		name := c.Query("query") // same as above
		age := c.Query("age")

		// if !ok {
		// 	//key doesn't exist
		// 	name = "someone1"
		// }

		c.JSON(200, gin.H{
			"name": name,
			"age":  age,
		})
	})

	r.Run(":8080")
}
