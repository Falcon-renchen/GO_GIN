package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserInfo struct {
	ID uint
	Name string
	Gender string
	Hobby string
}

func main()  {
	//连接数据库
	db, err := gorm.Open("mysql","root:root@/db1?charset=utf8&parseTime=True&loc=Local")
	if err!=nil {
		panic(err)
	}
	defer db.Close()

	//创建表
	db.AutoMigrate(&UserInfo{})
/*	//添加数据行
	u1 := UserInfo{1,"武宇鹏","男","咏春"}
	db.Create(u1)*/
	//查询
	var u UserInfo
	db.First(&u)
	fmt.Println("%v\n",u)
/*	//更新
	db.Model(&u).Update("hobby","双色球")*/
	//删除
	db.Delete(&u)
}
