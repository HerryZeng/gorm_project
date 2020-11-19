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

	// 创建表
	//db.CreateTable(&User{})
	// 指定表名
	//db.Table("user").CreateTable(&User{})
	// 删除表
	//db.DropTable("user")
	//db.DropTable(&User{})
	// 检查表是否存在
	//b := db.HasTable("user")
	//fmt.Println(b)
	//b2 := db.HasTable(&User{})
	//fmt.Println(b2)
	// 自动迁移
	db.AutoMigrate(&modules.User{})

}
