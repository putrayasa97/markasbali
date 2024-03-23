package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func OpenConn() (*sql.DB, error) {
	// user_db:password_db@tcp(host_db:port_db)/nama_db
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	// fmt.Println(connString)

	db, err := sql.Open("mysql", connString)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return db, nil
}
