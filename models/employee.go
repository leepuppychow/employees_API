package models

import (
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/leepuppychow/employees_API/database"
)

type Employee struct {
	EmployeeID int       `json:"employee_id"`
	BirthDate  time.Time `json:"birth_date"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Gender     string    `json:"gender"`
	HireDate   time.Time `json:"hire_date"`
}

func AllEmployees() ([]Employee, error) {
	var employees []Employee
	var (
		emp_no     int
		birth_date mysql.NullTime
		first_name string
		last_name  string
		gender     string
		hire_date  mysql.NullTime
	)

	q := "SELECT * FROM employees"
	rows, err := database.DB.Query(q)
	if err != nil {
		log.Println("Error with all employees query", err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&emp_no,
			&birth_date,
			&first_name,
			&last_name,
			&gender,
			&hire_date,
		)
		if err != nil {
			log.Println(err)
		}
		emp := Employee{
			EmployeeID: emp_no,
			BirthDate:  birth_date.Time,
			FirstName:  first_name,
			LastName:   last_name,
			Gender:     gender,
			HireDate:   hire_date.Time,
		}
		employees = append(employees, emp)
	}
	if err != nil {
		return employees, err
	} else {
		log.Println("Sucessful GET /employees")
		return employees, nil
	}
}
