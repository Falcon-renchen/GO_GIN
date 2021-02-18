package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
func main()  {
	r := gin.Default()
	r.Any("/user", func(context *gin.Context) {
		switch context.Request.Method {
		case "GET":
			context.JSON(http.StatusOK, gin.H{"method":"GET"})
		case http.MethodPost:
			context.JSON(http.StatusOK, gin.H{"method":"POST"})
		}
	})
	r.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound,gin.H{
			"method" : "wyp",
		})
	})

	videoGroup := r.Group("/video")
	{
		videoGroup.GET("/index", func(context *gin.Context) {
			context.JSON(http.StatusOK,gin.H{"method":"/video/index"})
		})
		videoGroup.POST("/main", func(context *gin.Context) {
			context.JSON(http.StatusOK,gin.H{"method":"/video/main"})
		})
	}
	r.Run()
}
