package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./login.html","./index.html")
	r.GET("/login", func(context *gin.Context) {
		context.HTML(http.StatusOK, "login.html", nil)
	})
	r.POST("/login", func(context *gin.Context) {
/*		//获取form表单中的数据
		username := context.PostForm("username")
		password := context.PostForm("password")*/
/*		username := context.DefaultPostForm("username","somebody")
		password := context.DefaultPostForm("password","***")*/
		username, ok := context.GetPostForm("username")
		if !ok {
			username = "sb"
		}
		password, ok := context.GetPostForm("password")
		if !ok {
			password = "***"
		}

		context.HTML(http.StatusOK, "index.html", gin.H{
			"Name" : username,
			"Password" : password,
		})
	})
	r.Run(":8080")
}
