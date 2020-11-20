package modules

import (
	"github.com/jinzhu/gorm"
	"time"
)

type GormMode struct {
	gorm.Model
	Name          string
	DBACreateTime time.Time
}

func (GormMode) TableName() string {
	return "test_gorm_model"
}
