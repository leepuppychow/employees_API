## Sample RESTful API with employees MySQL DB:

1. Clone database repo from: https://github.com/datacharmer/test_db
2. Run\
  `mysql < employees.sql`
3. Clone this repository
4. Have Docker installed
5. From root directory run `docker build -t employees .`
6. Start the Docker container with `docker run --publish 8080:8080 --name employees --rm employees`
7. Stop the Docker container with `docker stop employees` (from another Shell)
