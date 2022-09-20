package model

import "github.com/jinzhu/gorm"

type Users struct {
	gorm.Model
	Name     string `gorm:"column:Name;"`
	Email    string `gorm:"column:Email;"`
	Password string `gorm:"column:Password;"`
	Role     Role   `gorm:"reference:id;"`
}

func (Users) TableName() string {
	return "Users"
}
