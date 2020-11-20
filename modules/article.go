package modules

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Title   string `gorm:"not null;unique_index"`
	Content string `gorm:"column:a_content"`
	Desc    string `gorm:"size:128"`
	UserId  string `gorm:"type:int(18)"`
	Test    string `gorm:"-"`
}
