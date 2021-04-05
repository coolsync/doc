package models

import "github.com/jinzhu/gorm"

type GormModel struct {
	gorm.Model
	Name string
}

func (GormModel) TableName() string {
	return "test_gorm_model"
}

