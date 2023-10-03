package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `json:"username"`
	PassWord string `json:"password"`
}
type Post struct {
	gorm.Model
	Title   string
	Content string `gorm:"type:text"`
	Tag     string
}
