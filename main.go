package main

import "os"

func main() {
	host := os.Getenv("GO_TEST_HOST")
	port := os.Getenv("GO_TEST_PORT")
}
