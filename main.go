package main

import (
	"log"
	"net/http"
	"os"

	"github.com/leepuppychow/employees_API/database"
	"github.com/leepuppychow/employees_API/routes"
)

func main() {
	host := os.Getenv("GO_TEST_HOST")
	port := os.Getenv("GO_TEST_PORT")
	address := host + ":" + port

	router := routes.NewRouter()
	database.Connect("employees")
	log.Println("Server running at:", address)

	log.Fatal(http.ListenAndServe(address, router))
}
