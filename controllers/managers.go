package controllers

import (
	"net/http"

	"github.com/leepuppychow/employees_API/helpers"
	"github.com/leepuppychow/employees_API/models"
)

func ManagersIndex(w http.ResponseWriter, r *http.Request) {
	data, err := models.AllManagers()
	helpers.WriteResponse(data, err, 400, w)
}
