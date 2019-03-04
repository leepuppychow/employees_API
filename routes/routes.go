package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leepuppychow/employees_API/controllers"
)

type Route struct {
	Name        string
	Method      string
	Url         string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	{"PapersIndex", "GET", "/employees", controllers.EmployeesIndex},
	{"PapersIndex", "GET", "/departments", controllers.DepartmentsIndex},
	{"PapersIndex", "GET", "/departmentManagers", controllers.ManagersIndex},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Url).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}
