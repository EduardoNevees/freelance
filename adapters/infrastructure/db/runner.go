package db

import (
	"fmt"
	"log"
	"time"

	"github.com/EduardoNevesRamos/frelancer.git/common/env"
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

var db *gorm.DB

func Start() {
	tryOpenConnection()
}

func tryOpenConnection() {

	dsn := getUrlConnection()

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error: ", err)
	}

	db = database

	config, _ := db.DB()
	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	RunMigrations(db)
}

func getUrlConnection() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)%s?parseTime=true&loc=Local",
		env.DataBase.USER,
		env.DataBase.PASSWORD,
		env.DataBase.HOST,
		env.DataBase.PORT,
		env.DataBase.DB_NAME,
	)
}

func Get() *gorm.DB {
	return db
}
