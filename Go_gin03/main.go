package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	r := gin.Default()

	r.GET("/json", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"name" : "wyp",
			"sex" : "boy",
			"age" : 21,
		})
	})
	type Person struct {
		Name string	`json:"name"`
		Sex string
		Age int64
	}
	
	r.GET("/person", func(context *gin.Context) {
		data := &Person{
			Name: "WYP",
			Sex:  "男",
			Age:  20,
		}

		context.JSON(http.StatusOK, data)
	})
	r.Run()
}
