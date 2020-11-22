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

	// FirstOrInit

	var user relation_tables.User
	//db.FirstOrInit(&user, relation_tables.User{Name: "Herry"})
	db.Debug().FirstOrInit(&user, relation_tables.User{Name: "Hank"})
	fmt.Println(user)

	var ages []int
	var users []relation_tables.User
	db.Debug().Where("age >?", 18).Find(&users).Pluck("age", &ages)
	fmt.Println(ages)
}
