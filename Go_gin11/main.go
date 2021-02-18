package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandler(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"method" : "index",
	})
}
func m1(c *gin.Context)  {
	fmt.Println("m1 in...")
}

func main()  {
	r := gin.Default()
	r.GET("/index",m1,IndexHandler)
	r.Run()
}
