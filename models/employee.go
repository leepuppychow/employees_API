package models

import (
	"fmt"
	"log"

	"github.com/leepuppychow/employees_API/database"
)

type Employee struct {
	EmployeeID int    `json:"employee_id"`
	BirthDate  string `json:"birth_date"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Gender     string `json:"gender"`
	HireDate   string `json:"hire_date"`
}

func AllEmployees() ([]Employee, error) {
	var employees []Employee
	var employee Employee

	q := "SELECT * FROM employees"
	rows, err := database.DB.Query(q)
	if err != nil {
		log.Println("Error with all employees query", err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&employee.EmployeeID,
			&employee.BirthDate,
			&employee.FirstName,
			&employee.LastName,
			&employee.Gender,
			&employee.HireDate,
		)
		if err != nil {
			log.Println(err)
		}
		employees = append(employees, employee)
	}
	if err != nil {
		return employees, err
	} else {
		log.Println("Sucessful GET /employees")
		return employees, nil
	}
}
