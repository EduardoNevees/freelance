package model

import (
	"github.com/jinzhu/gorm"
)

type Role struct {
	gorm.Model
	Name string `gorm:"column:name;"`
}

func (Role) TableName() string {
	return "Role"
}
