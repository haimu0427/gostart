package controller

import (
	"net/http"
	"proj/dao"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Hello Bubble",
	})
}
func CreatTask(c *gin.Context) {
	// 前端页面填写代办事项, 后端处理
	// 1. 从请求中把数据拿出来
	var todo dao.Todo
	c.BindJSON(&todo) // 2. 存入数据库
	// 3. 返回响应
	err := dao.DB.Create(&todo).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "创建成功",
			"msg":     todo,
		})
	}

}

func GetTodoList(c *gin.Context) {
	var todoList []dao.Todo
	err := dao.DB.Find(&todoList).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, todoList)
}
func UpdateTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "无效的ID",
		})
		return
	}
	var todo dao.Todo
	err := dao.DB.Where("id = ?", id).First(&todo).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.BindJSON(&todo)
	if err = dao.DB.Save(&todo).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
}
func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	err := dao.DB.Delete(&dao.Todo{}, id).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
}
