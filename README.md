## Sample RESTful API with employees MySQL DB:

1. Clone database repo from: https://github.com/datacharmer/test_db
2. run `mysql < employees.sql`
3. Set up environment variables (Linux environment):
  * export GO_TEST_USER=*********
  * export GO_TEST_PASS=*********
  * export GO_TEST_HOST=*********
  * export GO_TEST_PORT=*********
4. Start server with `go run main.go`