package main

//中间件
import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	fmt.Println("indexHandler in...")
	c.String(http.StatusOK, "Hello, World!")
}
func authMiddleWare(c *gin.Context) {
	fmt.Println("authMIddleWare in...")
	//需要登陆哦
}
func authMiddleWareToo(doCheck bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if doCheck {
			fmt.Println("authMMiddleWareToo in...")
		}
	}
}
func loginHandler(c *gin.Context) {
	fmt.Println("loginHandler in...")
	c.JSON(http.StatusOK, gin.H{
		"message": "login success",
	})
}
func m1(c *gin.Context) {
	//业务逻辑
	fmt.Println("m1 in...")
	//计时开始
	start := time.Now()
	c.Next() //调用下一个中间件或处理器
	cost := time.Since(start)
	fmt.Println("cost:", cost)
	//go funcXX()
	//在goroutine中使用c会报错，因为c不是并发安全的
	//只能使用c的副本
	//cCopy := c.Copy() 只读对象
	//go func() {
	//	cCopy.Abort()
	//}()
	c.Abort()
	fmt.Println("m1 out...")
}
func main() {
	r := gin.Default()                                 // gin.Default()默认使用了Logger和Recovery中间件，其中：
	r.Use(m1, authMiddleWare, authMiddleWareToo(true)) //注册全局中间件
	r.GET("/index", indexHandler)                      //先走m1再走indexHandler
	r.GET("/login", loginHandler)
	r.Run(":8080")
}
