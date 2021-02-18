package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/main", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})
	r.GET("/a", func(context *gin.Context) {
		//跳转到/b对应的路由处理函数
		context.Request.URL.Path="/b"	//把请求的URL修改
		r.HandleContext(context)	//继续后续的处理
	})
	r.GET("/b", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message" : "b",
		})
	})
	r.Run()
}
