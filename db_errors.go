package main

import (
	"fmt"
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

	// Errors

	var users []relation_tables.User
	result := db.Where("name == ?", "Herry2").Find(&users)
	fmt.Println(result.RowsAffected)
	fmt.Println(users)
	//fmt.Println(result.Error)
	fmt.Println(result.GetErrors())

	if result.Error != nil {
		// 发生错误
	}
	//fmt.Println(result.GetErrors())

	ct := db.Begin()
	ret3 := ct.Commit()
	if ret3.Error != nil {
		ct.Rollback()
	}
}
