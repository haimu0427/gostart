package main

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

// 什么是静态文件, .css,.js,.jpg,.png等文件
// 如何加载静态文件呢
// 什么是模板文件, .tmpl,.html等文件
func main() {
	// Create a gin router with default middleware:
	r := gin.Default()

	// Set a custom template function
	r.SetFuncMap(template.FuncMap{
		"safe": func(s string) template.HTML {
			return template.HTML(s)
		},
	})

	// load static files
	r.Static("/statics", "./statics") //静态文件的解析

	// Load HTML templates from the "templates" directory
	// r.LoadHTMLFiles("./templates/posts/index.tmpl") //模板的解析
	r.LoadHTMLGlob("./templates/**/*") //模板的解析, 支持使用通配符,**是匹配任意层级的目录,*是匹配任意文件名，包含.tmpl和.html文件

	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(200, "index.tmpl", gin.H{ //模板的渲染
			"title": "<a href='https://liwenzhou.com'>liwenzhou.com</a>",
		})
	})
	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(200, "index.tmpl", gin.H{ //模板的渲染
			"title": "users index website",
		})
	})
	r.GET("/home", func(c *gin.Context) {
		c.HTML(200, "home.html", nil)
	})

	r.Run(":8080") // 启动HTTP服务，默认在8080端口启动服务go
}
