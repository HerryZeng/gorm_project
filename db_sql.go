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

	var users []relation_tables.User

	db.Raw("select * from users where name =?", "dlzeng").Find(&users)
	fmt.Print(users)

	//db.Exec("insert into users(name,age,address) values(?,?,?)", "Abc", 22, "西乡")
	db.Exec("update users set address=? where name =?", "固戍", "Abc")
	db.Exec("delete from users where id =?", 5)
}
