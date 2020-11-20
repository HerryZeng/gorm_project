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

	// 增加
	//userProfile := relation_tables.UserProfile{
	//	Pic:   "1.jpg",
	//	CPic:  "2.jpg",
	//	Phone: "1234567",
	//	User: relation_tables.User{
	//		Name:    "dlzeng",
	//		Age:     23,
	//		Address: "xxx",
	//	},
	//}
	//db.Create(&userProfile)

	// 第一种查询 Association
	fmt.Println("第一种查询")
	var user_profile relation_tables.UserProfile
	db.Debug().First(&user_profile, 1)
	fmt.Println(user_profile)
	db.Debug().Model(&user_profile).Association("User").Find(&user_profile.User)
	fmt.Println(user_profile)

	// 第二种查询 Preload
	fmt.Println("第二种查询")
	var user_profile2 relation_tables.UserProfile
	db.Debug().Preload("User").Find(&user_profile2, 1)
	fmt.Println(user_profile2)

	// 第二种查询 Related
	fmt.Println("第三种查询")
	var user_profile3 relation_tables.UserProfile
	db.First(&user_profile3, 1)

	//var user relation_tables.User
	db.Debug().Model(&user_profile3).Related(&user_profile3.User, "User")
	fmt.Println(user_profile3)
	//fmt.Println(user)

	// 改
	var user_profile4 relation_tables.UserProfile
	db.Preload("User").First(&user_profile4, 2)
	fmt.Println(user_profile4)
	//db.Model(&user_profile4.User).Update("name", "李四")
	db.Model(&user_profile4.User).Update(relation_tables.User{Name: "Herry", Age: 44, Address: "深圳"})
	fmt.Println(user_profile4)

	var user_profile5 relation_tables.UserProfile
	db.Preload("User").First(&user_profile5, 1)
	fmt.Println(user_profile5)
	db.Debug().Delete(&user_profile5.User)
}
