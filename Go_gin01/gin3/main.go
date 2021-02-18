package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request)  {
	//定义模版
	//解析模版
	t := template.New("hello.tmpl")
	_, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("get a error")
		return
	}
	name := "小王子"
	//渲染模版
	t.Execute(w,name)
}

func main()  {
	http.HandleFunc("/f1",f1)
	err := http.ListenAndServe(":9000",nil)
	if err!=nil {
		fmt.Println("HTTP server start failed err %v",err)
		return
	}
}
