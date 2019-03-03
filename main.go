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
	log.Println("Server running on port:", port)
	database.Connect("employees")

	log.Fatal(http.ListenAndServe(address, router))
	log.Println("Server running at:", address)
}
