package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})
	r.POST("/upload", func(context *gin.Context) {
		//从请求中读取文件
		f, err := context.FormFile("f1")	//从请求中获取携带的参数一样的
		if err != nil {
			context.JSON(http.StatusOK,gin.H{
				"error" : err.Error(),
			})
		} else {
			//将读取到的文件保存到本地（服务端本地）
			dst := fmt.Sprintf("./%s",f.Filename)
			context.SaveUploadedFile(f,dst)
			context.JSON(http.StatusOK, gin.H{
				"state" : "ok",
			})
		}
	})
	r.Run()
}
