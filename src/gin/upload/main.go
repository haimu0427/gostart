package main

import (
	_ "log"
	_ "net/http"

	_ "github.com/gin-gonic/gin"
)

func main() {
	// r := gin.Default()
	// r.LoadHTMLFiles("index.html")
	// r.GET("/upload", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.html", nil)
	// })
	// r.POST("/upload", func(c *gin.Context) {
	// 	form, _ := c.MultipartForm()
	// 	files := form.File["files"]
	// 	dst := "./" + files.Filename
	// 	for _, file := range files {

	// 		log.Println(file.Filename)
	// 	}
	// 	// 这里的 "files" 是指前端 form 表单中 input 标签的 name 属性
	// 	// 例如：<input type="file" name="files" multiple>
	// 	// 上传文件到指定的路径
	// 	err := c.SaveUploadedFile(files, dst)
	// 	if err != nil {
	// 		c.String(http.StatusInternalServerError, "Failed to upload file")
	// 		return
	// 	}
	// 	c.JSON(http.StatusOK, gin.H{"message": "upload success"})
	// })
	// r.Run(":8080")
}
