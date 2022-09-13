package db

import (
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/EduardoNevesRamos/frelancer.git/common/env"
	"gorm.io/driver/postgres"
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

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return errors.New("there was an error connecting to the database")
	}

	db = database
	return nil
}

func getUrlConnection() string {

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
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
