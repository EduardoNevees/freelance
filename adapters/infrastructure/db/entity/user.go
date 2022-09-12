package entity

import "github.com/jinzhu/gorm"

type Users struct {
	gorm.Model
	ID        uint64 `gorm:"column:Id; index"`
	Name      string `gorm:"column:Name;"`
	Email     string `gorm:"column:Email;"`
	Password  string `gorm:"column:Password;"`
	ModelName string `gorm:"column:ModelName;"`
}

func (Users) TableName() string {
	return "Users"
}
