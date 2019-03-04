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
	{"EmployeesIndex", "GET", "/employees", controllers.EmployeesIndex},
	{"EmployeesComplete", "GET", "/employees-complete", controllers.EmployeesComplete},
	{"DepartmentsIndex", "GET", "/departments", controllers.DepartmentsIndex},
	{"DepartmentOfEmployee", "GET", "/department-by-employee/{employeeId}", controllers.DepartmentOfEmployee},
	{"ManagersIndex", "GET", "/departmentManagers", controllers.ManagersIndex},
	{"SalariesIndex", "GET", "/salaries", controllers.SalariesIndex},
	{"SalariesIndex", "GET", "/salaries", controllers.SalariesIndex},
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
