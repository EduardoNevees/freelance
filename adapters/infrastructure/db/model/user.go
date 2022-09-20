package model

import "github.com/jinzhu/gorm"

type Users struct {
	gorm.Model
	Name     string `gorm:"column:Name;"`
	Role     string `gorm:"column:Roles;"`
	Email    string `gorm:"column:Email;"`
	Password string `gorm:"column:Password;"`
}

func (Users) TableName() string {
	return "Users"
}
