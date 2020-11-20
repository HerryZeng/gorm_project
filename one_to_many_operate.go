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

	// 增 方法一
	//var user2 = relation_tables.User2{
	//	Name:    "dlzeng",
	//	Age:     18,
	//	Address: "深圳",
	//	Article: []relation_tables.Article{
	//		{
	//			Title:   "Beego详解1",
	//			Content: "Beego内容1,xxx",
	//			Desc:    "beego详解描述1",
	//		},
	//		{
	//			Title:   "Beego详解2",
	//			Content: "Beego内容2,xxx",
	//			Desc:    "beego详解描述2",
	//		},
	//	},
	//}
	// 增 方法二
	//article3 := relation_tables.Article{
	//	Title:   "Beego详解3",
	//	Content: "Beego内容3,xxx",
	//	Desc:    "beego详解描述3",
	//}
	//article4 := relation_tables.Article{
	//	Title:   "Beego详解4",
	//	Content: "Beego内容4,xxx",
	//	Desc:    "beego详解描述4",
	//}
	//article5 := relation_tables.Article{
	//	Title:   "Beego详解5",
	//	Content: "Beego内容5,xxx",
	//	Desc:    "beego详解描述5",
	//}
	//var user2 = relation_tables.User2{
	//	Name:    "dlzeng",
	//	Age:     18,
	//	Address: "深圳",
	//	Article: []relation_tables.Article{
	//		article3, article5,
	//	},
	//}
	//db.Create(&user2)

	// 查
	var user2 relation_tables.User2
	db.Debug().Preload("Article").Find(&user2, 1)
	fmt.Println(user2)

	var user2_1 relation_tables.User2
	db.First(&user2_1, 2)
	db.Model(&user2_1).Related(&user2_1.Article, "Article")
	fmt.Println(user2_1)

	var user2_2 relation_tables.User2
	db.First(&user2_2, 3)
	db.Model(&user2_2).Association("Article").Find(&user2_2.Article)
	fmt.Println(user2_2)

	// 改
	var user2_4 relation_tables.User2
	//db.First(&user2_4,3)
	db.Preload("Article").Find(&user2_4, 2)
	db.Model(&user2_4.Article).Where("id=?", 1).
		Update("title", "Gin详解1")
	db.Model(&user2_4.Article).Where("id=?", 2).
		Update("title", "GoWeb详解1")
	// 删除
	var user3_5 relation_tables.User2
	db.First(&user3_5, 4)
	db.Debug().Model(&user3_5).Related(&user3_5.Article, "Article").Where("id=?", 7).Delete(&user3_5.Article)
}
