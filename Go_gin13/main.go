package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	ID int64
	Name string `gorm:"default:'小王子'"`
	Age int64
}

func main()  {

	db, err := gorm.Open("mysql","root:root@(localhost:3306)/db1?charset=utf8&parseTime=True&loc=Local")
	if err!=nil {
		panic(err)
	}
	defer db.Close()

	//把模型与数据库中的表对应起来
	db.AutoMigrate(&User{})
/*	//创建
	u := &User{Age:21}
	fmt.Println(db.NewRecord(&u))	//判断主键是否为空
	db.Debug().Create(&u)
	db.Create(&u)	//创建数据表
	fmt.Println(db.NewRecord(&u))	//判断主键是否为空*/

	//查询
	/*var user User
	db.First(&user)
	fmt.Println("%v\n",user)
	*/

/*	var user []User
	db.Find(&user)
	db.Debug().Find(&user)
	fmt.Println(&user)*/

/*	var user User
	db.First(&user,3)
	db.Debug().Find(&user)
	fmt.Println(&user)*/

/*
	var user User
	db.Where(&User{Name:"wrh",Age:11}).First(&user)
	db.Debug().Find(&user)
	fmt.Println(&user)*/

	var user User
	db.Where(map[string]interface{}{"Name":"wrh","Age":11}).Find(&user)
	db.Debug().Find(&user)
	fmt.Println(&user)
}
