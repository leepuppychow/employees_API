package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect(database string) {
	db, err := sql.Open("mysql", dbConnection(database))
	if err != nil {
		log.Fatal(err)
	}
	DB = db
}

func dbConnection(database string) string {
	user := os.Getenv("GO_TEST_USER")
	password := os.Getenv("DB_HOST")
	host := os.Getenv("GO_TEST_HOST")

	return user + ":" + password + "@tcp(" + host + ":3306)" + "/" + database
}
