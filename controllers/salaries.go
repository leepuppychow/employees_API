package controllers

import (
	"net/http"

	"github.com/leepuppychow/employees_API/helpers"
	"github.com/leepuppychow/employees_API/models"
)

func SalariesIndex(w http.ResponseWriter, r *http.Request) {
	data, err := models.AllSalaries()
	helpers.WriteResponse(data, err, 400, w)
}
