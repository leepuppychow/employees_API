package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leepuppychow/employees_API/helpers"
	"github.com/leepuppychow/employees_API/models"
)

func DepartmentsIndex(w http.ResponseWriter, r *http.Request) {
	data, err := models.AllDepartments()
	helpers.WriteResponse(data, err, 400, w)
}

func DepartmentOfEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	employeeId := vars["employeeId"]
	data, err := models.DepartmentOfEmployee(employeeId)
	helpers.WriteResponse(data, err, 400, w)
}
