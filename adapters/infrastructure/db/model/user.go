package model

import "github.com/jinzhu/gorm"

type Users struct {
	gorm.Model
	ID       uint64 `gorm:"column:Id; index"`
	Name     string `gorm:"column:Name;"`
	Email    string `gorm:"column:Email;"`
	Password string `gorm:"column:Password;"`
	Role     string `gorm:"column:role;"`
}

func (Users) TableName() string {
	return "Users"
}
