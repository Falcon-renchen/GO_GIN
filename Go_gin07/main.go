package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Username string `form:"username" json:"username"`
	Password string	`form:"password" json:"password"`
} 

func main()  {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("/user", func(context *gin.Context) {
		var u UserInfo
		err := context.ShouldBind(&u)
		if err!=nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error" : err.Error(),
			})
		} else {
			fmt.Printf("%v", u)
			context.JSON(http.StatusOK, gin.H{
				"status" : "ok",
			})
		}
	})
	r.POST("/form", func(context *gin.Context) {
		var u UserInfo
		err := context.ShouldBind(&u)
		if err!=nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error" : err.Error(),
			})
		} else {
			fmt.Printf("%v", u)
			context.JSON(http.StatusOK, gin.H{
				"status" : "ok",
			})
		} 
	})
	//r.GET("/user", func(context *gin.Context) {
	//	username := context.Query("username")
	//	password := context.Query("password")
	//	u := UserInfo{
	//		username: username,
	//		password: password,
	//	}
	//	fmt.Printf("%v\n",u)
	//	context.JSON(http.StatusOK, gin.H{
	//		"message" : "ok",
	//	})
	//})
	r.Run()
}
