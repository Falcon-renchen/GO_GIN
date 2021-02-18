package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name string
	Gender string
	Age int64
}

func sayHello(w http.ResponseWriter, r *http.Request)  {
	//解析模版
	t, err := template.ParseFiles("./hello.tmpl")
	if err!=nil {
		fmt.Println("templates is err %v",err)
	}
	//渲染模版
	u := &User{
		Name:   "小王子",
		Gender: "男",
		Age:    10,
	}
	m := map[string]interface{}{
		"Name" : "大王子",
		"Gender" : "女",
		"Age" : 20,
	}
	hobbyList := []string{
		"篮球",
		"足球",
		"双色球",
	}

	err = t.Execute(w,map[string]interface{}{
		"u" : u,
		"m" : m,
		"hobby" : hobbyList,
	})
	if err!=nil {
		fmt.Println("render templates is err %v",err)
		return
	}
}


func main()  {
	http.HandleFunc("/hello",sayHello)
	err := http.ListenAndServe(":9090",nil)
	if err!=nil {
		fmt.Println("HTTP server start failed err %v",err)
		return
	}
}
