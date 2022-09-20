package model

import "github.com/jinzhu/gorm"

type Role struct {
	gorm.Model
	ID   uint64 `gorm:"column:Id; index"`
	Name string `gorm:"column:name;"`
}

func (Role) TableName() string {
	return "Role"
}
