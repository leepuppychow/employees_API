package controllers

import (
	"net/http"

	"github.com/leepuppychow/employees_API/helpers"
	"github.com/leepuppychow/employees_API/models"
)

func EmployeesIndex(w http.ResponseWriter, r *http.Request) {
	data, err := models.AllEmployees()
	helpers.WriteResponse(data, err, 400, w)
}
