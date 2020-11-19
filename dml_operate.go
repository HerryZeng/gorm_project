package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm_project/modules"
)

func main() {
	connStr := "root:Kaka@2019@/gorm_project?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", connStr)
	if err != nil {
		panic(err)
	}

	defer db.Close()
	// 增
	db.Create(&modules.User{Name: "张三", Age: 18, Address: "xxx", Pic: "/static/upload/pic11.jpg", Phone: "11111111111"})

}
