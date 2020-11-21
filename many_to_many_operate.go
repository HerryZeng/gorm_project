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

	// 增加 方法一
	//tag1 := relation_tables.Tag{
	//	Name: "标签1",
	//	Desc: "标签描述1",
	//}
	//
	//tag2 := relation_tables.Tag{
	//	Name: "标签2",
	//	Desc: "标签描述2",
	//}

	//article2 := relation_tables.Article2{
	//	Title:   "文章标题",
	//	Content: "文章内容",
	//	Desc:    "文章描述",
	//	Tags: []relation_tables.Tag{
	//		tag1, tag2,
	//	},
	//}
	//
	//db.Create(&article2)

	// 增 方法二
	//var tag relation_tables.Tag
	//db.First(&tag, 1)
	//article2 := relation_tables.Article2{
	//	Title:   "文章标题2",
	//	Content: "文章内容2",
	//	Desc:    "文章描述2",
	//	Tags: []relation_tables.Tag{
	//		tag,
	//	},
	//}
	//db.Create(&article2)

	// 查
	var article2_1 relation_tables.Article2
	db.Debug().Preload("Tags").First(&article2_1, 1)
	fmt.Println(article2_1)

	var article2_2 relation_tables.Article2
	db.First(&article2_2, 1)
	db.Model(&article2_2).Association("Tags").Find(&article2_2.Tags)
	fmt.Println(article2_2)

	var article2_3 relation_tables.Article2
	db.First(&article2_3, 1)
	db.Model(&article2_3).Related(&article2_3.Tags, "Tags")
	fmt.Println(article2_3)

	// 改
	var article2_4 relation_tables.Article2
	db.Find(&article2_4, 1)
	db.Model(&article2_4).Where("id=?", 1).Update("title", "Django")
	db.Model(&article2_4.Tags).Where("id=?", 1).Update("name", "Django")

	// 删除
	var article2_5 relation_tables.Article2
	db.Preload("Tags").Find(&article2_5, 1)
	db.Model(&article2_5).Where("id=?", 1).Delete(&article2_5.Tags)
}
