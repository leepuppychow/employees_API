package main

import (
	"log"
	"net/http"
	"os"

	"github.com/leepuppychow/employees_API/database"
	"github.com/leepuppychow/employees_API/routes"
)

func main() {
	database.Connect("employees")
	startServer()
}

func startServer() {
	port := os.Getenv("GO_TEST_PORT")
	router := routes.NewRouter()
	log.Println("Server running at port:", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
