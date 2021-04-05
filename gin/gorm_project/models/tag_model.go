package models

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	//Title string `gorm:"not null;unique_index"`
	Title string `gorm:"not null;unique;size:64"`
	Content string `gorm:"column:a_content"`
	Desc string `gorm:"type:int(11)"`
	Test string `gorm:"-"`
}
