package controllers

import (
	"net/http"

	"github.com/leepuppychow/employees_API/models"
)

func EmployeesIndex(w http.ResponseWriter, r *http.Request) {
	data, err := models.AllEmployees()
}
