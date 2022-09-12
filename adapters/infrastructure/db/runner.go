package db

import (
	"errors"
	"fmt"
	"log"

	"github.com/EduardoNevesRamos/frelancer.git/common/env"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Start() {
	err := tryOpenConnection()
	if err != nil {
		log.Fatal(err)
	}
}

func tryOpenConnection() error {

	dsn := getUrlConnection()

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return errors.New("there was an error connecting to the database")
	}

	db = database
	return nil
}

func getUrlConnection() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		env.DataBase.USER,
		env.DataBase.PASSWORD,
		env.DataBase.HOST,
		env.DataBase.SQL_PORT,
		env.DataBase.DB_NAME,
	)
}

func Get() *gorm.DB {
	return db
}