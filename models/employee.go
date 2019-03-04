package models

import (
	"log"

	"github.com/leepuppychow/employees_API/database"
)

type Employee struct {
	EmployeeId int    `json:"employee_id"`
	BirthDate  string `json:"birth_date"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Gender     string `json:"gender"`
	HireDate   string `json:"hire_date"`
	Salary     string `json:"salary"`
	Title      string `json:"title"`
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
			&employee.EmployeeId,
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
		log.Println("Successful GET /employees")
		return employees, nil
	}
}

func AllEmployeesComplete() ([]Employee, error) {
	var employees []Employee
	var employee Employee
	q := `
		SELECT employees.*, title, salary FROM employees
		INNER JOIN titles ON employees.emp_no = titles.emp_no
		INNER JOIN salaries ON employees.emp_no = salaries.emp_no
	`
	rows, err := database.DB.Query(q)
	if err != nil {
		log.Println("Error with all employees query", err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&employee.EmployeeId,
			&employee.BirthDate,
			&employee.FirstName,
			&employee.LastName,
			&employee.Gender,
			&employee.HireDate,
			&employee.Title,
			&employee.Salary,
		)
		if err != nil {
			log.Println(err)
		}
		employees = append(employees, employee)
	}
	if err != nil {
		return employees, err
	} else {
		log.Println("Successful GET /employees-complete")
		return employees, nil
	}
}

func EmployeesOfDepartment(departmentId string) ([]Employee, error) {
	var employees []Employee
	var employee Employee
	q := `
		SELECT DISTINCT employees.* FROM employees
		INNER JOIN dept_emp ON employees.emp_no = dept_emp.emp_no
		WHERE dept_no = ?
	`
	rows, err := database.DB.Query(q, departmentId)
	if err != nil {
		log.Println("Error with employees of department query", err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&employee.EmployeeId,
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
		log.Println("Successful GET /employees-by-department")
		return employees, nil
	}
}
