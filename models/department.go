package models

import (
	"log"

	"github.com/leepuppychow/employees_API/database"
)

type Department struct {
	DepartmentId   string `json:"department_id"`
	DepartmentName string `json:"department_name"`
}

func AllDepartments() ([]Department, error) {
	var depts []Department
	var d Department

	q := "SELECT * FROM departments"
	rows, err := database.DB.Query(q)
	if err != nil {
		log.Println("Error with all departments query", err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&d.DepartmentId,
			&d.DepartmentName,
		)
		if err != nil {
			log.Println(err)
		}
		depts = append(depts, d)
	}
	if err != nil {
		return depts, err
	} else {
		log.Println("Successful GET /departments")
		return depts, nil
	}
}

func DepartmentOfEmployee(employeeId string) ([]Department, error) {
	var depts []Department
	var d Department

	q := `
		SELECT departments.* FROM departments
		INNER JOIN dept_emp ON departments.dept_no = dept_emp.dept_no
		WHERE emp_no=?
	`
	rows, err := database.DB.Query(q, employeeId)
	if err != nil {
		log.Println("Error with department by employee query", err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&d.DepartmentId,
			&d.DepartmentName,
		)
		if err != nil {
			log.Println(err)
		}
		depts = append(depts, d)
	}
	if err != nil {
		return depts, err
	} else {
		log.Println("Successful GET /departments-by-employee/:id")
		return depts, nil
	}
}
