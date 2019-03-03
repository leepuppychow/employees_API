package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect(database string) {
	DB, _ := sql.Open("mysql", DBConnection(database))
	err := DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func DBConnection(database string) string {
	user := os.Getenv("GO_TEST_USER")
	password := os.Getenv("GO_TEST_PASS")
	return user + ":" + password + "@/" + database
}
