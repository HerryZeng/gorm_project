package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm_project/modules/relation_tables"
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
	// 统一增加前缀, 只针对使用模型与表名默认规则
	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	return "sys_" + defaultTableName
	//}
	db.AutoMigrate(&relation_tables.User{}, &relation_tables.UserProfile{},
		relation_tables.User2{}, relation_tables.Article{}, relation_tables.Article2{}, relation_tables.Tag{})
	//db.AutoMigrate(&modules.User{}, &modules.GormMode{}, &modules.Article{})
	//db.AutoMigrate(&modules.User{}, &modules.UserInfo{}, &modules.DBUserInfo{}, &modules.DBXXXUserInfo{})

}
