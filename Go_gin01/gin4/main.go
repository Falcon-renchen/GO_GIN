package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/posts/index", func(context *gin.Context) {
		// http请求
		context.HTML(http.StatusOK, "posts/index.tmpl", gin.H{	//模版渲染
			"title" : "wyp.com",
		})
	})

	r.GET("/users/index", func(context *gin.Context) {
		// http请求
		context.HTML(http.StatusOK, "users/index.tmpl", gin.H{	//模版渲染
			"title" : "wrh.com",
		})
	})
	r.Run(":8080")
}
