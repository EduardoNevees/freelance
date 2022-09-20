package db

import (
	"github.com/EduardoNevesRamos/frelancer.git/adapters/infrastructure/db/model"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(model.Users{})
	db.AutoMigrate(model.Role{})
}
