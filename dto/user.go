package dto

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint64         `json:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	Role      string         `json:"Role"`
	CreatedAt time.Time      `json:"created"`
	UpdatedAt time.Time      `json:"updated"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted"`
}

// type Role struct {
// 	ID        uint64         `json:"id"`
// 	Name      string         `json:"name"`
// 	CreatedAt time.Time      `json:"created"`
// 	UpdatedAt time.Time      `json:"updated"`
// 	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted"`
// }

// func (u User) WithUserRole() User {
// 	u.Role = Role{
// 		ID:   1,
// 		Name: "commom",
// 	}
// 	return u
// }
