package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	r := gin.Default()

	r.GET("/web", func(context *gin.Context) {

		//name := context.Query("query")
		//name := context.DefaultQuery("query","somebody")//取不到就用指定的默认值
		name , ok := context.GetQuery("query")
		if !ok {
			//取不到
			name = "somebody"
		}

		context.JSON(http.StatusOK, gin.H{
			"name" : name,
		})
	})
	r.Run()
}
