package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leepuppychow/employees_API/helpers"
	"github.com/leepuppychow/employees_API/models"
)

func EmployeesIndex(w http.ResponseWriter, r *http.Request) {
	data, err := models.AllEmployees()
	helpers.WriteResponse(data, err, 400, w)
}

func EmployeesComplete(w http.ResponseWriter, r *http.Request) {
	data, err := models.AllEmployeesComplete()
	helpers.WriteResponse(data, err, 400, w)
}

func EmployeesOfDepartment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	departmentId := vars["departmentId"]
	data, err := models.EmployeesOfDepartment(departmentId)
	helpers.WriteResponse(data, err, 400, w)
}
