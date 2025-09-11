package router

import (
	"net/http"
	"proj/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")
	r.GET("/", controller.IndexHandler)
	v1Group := r.Group("v1")
	{
		v1Group.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "hello bubble",
			})
		})
		//代办的添加
		v1Group.POST("/todo", controller.CreatTask)
		//代办的查看(所有)
		v1Group.GET("/todo", controller.GetTodoList)
		//代办的查看(单个)
		v1Group.GET("/todo/:id", func(c *gin.Context) {

		})
		//代办的修改
		v1Group.PUT("/todo/:id", controller.UpdateTodo)
		//代办的删除
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)
	}
	return r
}
