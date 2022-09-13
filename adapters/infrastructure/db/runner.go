package db

import (
	"errors"
	"log"

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

	dsn := "postgres://dev:dev@localhost:5432/teste_dev?sslmode=disable"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return errors.New("there was an error connecting to the database")
	}

	db = database
	return nil
}

// func getUrlConnection() string {
// 	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
// 		env.DataBase.HOST,
// 		env.DataBase.PORT,
// 		env.DataBase.USER,
// 		env.DataBase.DB_NAME,
// 		env.DataBase.PASSWORD,
// 	)
// }

func Get() *gorm.DB {
	return db
}
