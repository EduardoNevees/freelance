package db

import (
	"github.com/EduardoNevesRamos/frelancer.git/model"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(model.User{})
}
