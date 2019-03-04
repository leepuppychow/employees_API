## Sample RESTful API with employees MySQL DB:

1. Clone database repo from: https://github.com/datacharmer/test_db
2. Run\
  `mysql < employees.sql`
3. Set up environment variables (Example for Linux environment):\
  `export GO_TEST_USER=root`\
  `export GO_TEST_PASS=password`\
  `export GO_TEST_HOST=localhost`\
  `export GO_TEST_PORT=8080`\
4. Start server with:\
  `go run main.go`
