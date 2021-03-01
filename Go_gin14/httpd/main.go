package main

import (
	"Restful/platform/newfeeds"
	"fmt"
)
func main() {
	//r := gin.Default()
	//r.GET("/ping", handler.PingGet)
	//r.GET("/ping2",handler.PingGet2())
	//r.Run()

	feed := newfeeds.New()

	fmt.Println(feed)

	feed.Add(newfeeds.Item{"Hello","how ya",})

	fmt.Println(feed)
}
