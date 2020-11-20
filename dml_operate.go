package main

import (
	"fmt"
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
	//db.Create(&modules.User{Id: 1, Name: "张三", Age: 18, Address: "xxx", Pic: "/static/upload/pic11.jpg", Phone: "11111111111"})

	// 查
	var user modules.User
	db.First(&user, 1)
	fmt.Println(user)
	db.First(&user, "name=?", "张三").Update("phone", "1234567890")
	fmt.Println(user)

	// 改
	db.Model(&user).Update("age", 20).Update("address", "yyy-zs")
	//db.Model(&user).Update("address", "xxx-zs")

	// 删除
	db.First(&user, 1).Delete(&user)
	//db.Delete(&user)
}
