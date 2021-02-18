package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//注意URL的匹配不要冲突
func main() {
	r := gin.Default()
	r.GET("/user/:name/:age", func(context *gin.Context) {
		name := context.Param("name")
		age := context.Param("age")
		context.JSON(http.StatusOK, gin.H{
			"name" : name,
			"age" : age,
		})
	})
	r.GET("/blog/:month/:year", func(context *gin.Context) {
		month := context.Param("month")
		year := context.Param("year")
		context.JSON(http.StatusOK, gin.H{
			"month" : month,
			"year" : year,
		})
	})
	r.Run()
}
