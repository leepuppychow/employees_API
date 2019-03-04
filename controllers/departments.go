package controllers

import (
	"net/http"

	"github.com/leepuppychow/employees_API/helpers"
	"github.com/leepuppychow/employees_API/models"
)

func DepartmentsIndex(w http.ResponseWriter, r *http.Request) {
	data, err := models.AllDepartments()
	helpers.WriteResponse(data, err, 400, w)
}
