package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func sayHello(c *gin.Context)  {
	c.JSON(200,gin.H{
		"message" : "hello world!",
	})
}

func main()  {
	//创建默认路由引擎
	r := gin.Default()
	//使用get请求
	r.GET("/hello",sayHello)
	r.POST("/book", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message" : "POST",
		})
	})
	r.PUT("/book", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message" : "PUT",
		})
	})
	r.DELETE("/book", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message" : "DELETE",
		})
	})

	r.Run()		//如果要转换端口   r.Run(":9090")
}