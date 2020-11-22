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

	// First
	var user relation_tables.User
	db.First(&user, 1) // 默认使用id查询
	fmt.Println(user)
	db.First(&user, "name=?", "Herry")
	fmt.Println(user)
	ret := db.Where("name=?", "Herry").First(&user)
	fmt.Println(user)
	fmt.Println(ret.RowsAffected)

	// FirstOrCreate
	var user5 relation_tables.User
	user4 := relation_tables.User{
		Name:    "dlzeng",
		Age:     22,
		Address: "宝安",
		//PId:     1,
	}
	ret2 := db.FirstOrCreate(&user5, user4)
	fmt.Println(ret2.Error)
	fmt.Println(ret2.RowsAffected)
	fmt.Println(user5)

	var user6 relation_tables.User
	db.Debug().Last(&user6)
	fmt.Println(user6)

	var user7 relation_tables.User
	db.Debug().Take(&user7, 2)
	fmt.Println(user7)

	// Not
	fmt.Println("Not")
	var user9 []relation_tables.User
	//db.Where("name = ?", "dlzeng").Find(&user9)
	db.Debug().Not("name = ?", "Herry").Order("p_id desc").Find(&user9)
	fmt.Println(user9)

	type UserJoinUserProfile struct {
		Id      int
		Name    string
		Age     int
		Address string
		PId     int
		Pic     string
		CPic    string
		Phone   string
	}

	var userJoinUserProfile1 []UserJoinUserProfile
	var user10 []relation_tables.User
	db.Debug().Select("users.*,user_profiles.*").Joins("left join user_profiles on users.id = user_profiles.id").
		Find(&user10).Scan(&userJoinUserProfile1)
	fmt.Println(userJoinUserProfile1)
	//fmt.Println(user10)
}
