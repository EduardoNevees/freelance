package model

import "time"

type Users struct {
	ID        uint       `gorm:"column:Id;index"`
	Name      string     `gorm:"column:Name;"`
	Role      Role       `gorm:"foreignkey:Id;references:Id"`
	Email     string     `gorm:"column:Email;"`
	Password  string     `gorm:"column:Password;"`
	CreatedAt time.Time  `gorm:"column:CreatedAt;"`
	UpdatedAt time.Time  `gorm:"column:UpdatedAt;"`
	DeletedAt *time.Time `sql:"index"`
}

func (Users) TableName() string {
	return "Users"
}
