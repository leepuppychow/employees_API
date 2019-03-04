package models

import (
	"log"

	"github.com/leepuppychow/employees_API/database"
)

type Manager struct {
	DepartmentId string `json:"department_id"`
	EmployeeId   string `json:"employee_id"`
	FromDate     string `json:"from_date"`
	ToDate       string `json:"to_date"`
}

func AllManagers() ([]Manager, error) {
	var managers []Manager
	var m Manager

	q := "SELECT * FROM dept_manager"
	rows, err := database.DB.Query(q)
	if err != nil {
		log.Println("Error with all managers query", err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&m.EmployeeId,
			&m.DepartmentId,
			&m.FromDate,
			&m.ToDate,
		)
		if err != nil {
			log.Println(err)
		}
		managers = append(managers, m)
	}
	if err != nil {
		return managers, err
	} else {
		log.Println("Successful GET /departmentManagers")
		return managers, nil
	}
}
