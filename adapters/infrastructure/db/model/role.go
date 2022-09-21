package model

import (
	"time"
)

type Role struct {
	ID        uint       `gorm:"column:Id;index"`
	Name      string     `gorm:"name"`
	CreatedAt time.Time  `gorm:"column:CreatedAt;"`
	UpdatedAt time.Time  `gorm:"column:UpdatedAt;"`
	DeletedAt *time.Time `sql:"index"`
}

func (Role) TableName() string {
	return "Role"
}
