package config

import (
	"fmt"
	"log"
	"os"
	"sekolahbeta/mini-project-3/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlDB struct {
	DB *gorm.DB
}

var Mysql MysqlDB

func OpenDB(env string) {
	DB_NAME := os.Getenv("DB_NAME")
	if env == "default" {
		DB_NAME = os.Getenv("DB_NAME")
	}

	if env == "testing" {
		DB_NAME = os.Getenv("DB_NAME_TESTING")
	}

	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		DB_NAME)

	conn := mysql.Open(connString)
	mysqlConn, err := gorm.Open(conn, &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	Mysql = MysqlDB{
		DB: mysqlConn,
	}

	err = autoMigrate(mysqlConn)
	if err != nil {
		log.Fatal(err)
	}
}

func autoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&model.Book{},
	)

	if err != nil {
		return err
	}

	return nil
}
