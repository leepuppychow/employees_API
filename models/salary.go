package models

import (
	"log"

	"github.com/leepuppychow/employees_API/database"
)

type Salary struct {
	EmployeeId string `json:"employee_id"`
	Salary     string `json:"salary"`
	FromDate   string `json:"from_date"`
	ToDate     string `json:"to_date"`
}

func AllSalaries() ([]Salary, error) {
	var salaries []Salary
	var s Salary

	q := "SELECT * FROM salaries"
	rows, err := database.DB.Query(q)
	if err != nil {
		log.Println("Error with all salaries query", err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&s.EmployeeId,
			&s.Salary,
			&s.FromDate,
			&s.ToDate,
		)
		if err != nil {
			log.Println(err)
		}
		salaries = append(salaries, s)
	}
	if err != nil {
		return salaries, err
	} else {
		log.Println("Successful GET /salaries")
		return salaries, nil
	}
}
