package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Connect(database string) *sql.DB {
	db, err := sql.Open("mysql", DBConnection(database))
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func DBConnection(database string) string {
	user := os.Getenv("GO_TEST_USER")
	password := os.Getenv("GO_TEST_PASS")
	return user + ":" + password + "@/" + database
}
