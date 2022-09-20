package db

import (
	"github.com/EduardoNevesRamos/frelancer.git/dto"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(dto.User{})
}
