package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{

	Route{
		"CreateUserOrganizationById",
		strings.ToUpper("Post"),
		"/api/v1/organization",
		controller.CreateUserOrganizationById,
	},

	Route{
		"DeleteUserOrganizationById",
		strings.ToUpper("Delete"),
		"/api/v1/organization/{organizationId}",
		controller.DeleteUserOrganizationById,
	},

	Route{
		"GetUserOrganizationById",
		strings.ToUpper("Get"),
		"/api/v1/organization/{organizationId}",
		controller.GetUserOrganizationById,
	},

	Route{
		"GetUserOrganizations",
		strings.ToUpper("Get"),
		"/api/v1/organization",
		controller.GetUserOrganizations,
	},

	Route{
		"UpdateUserOrganizationById",
		strings.ToUpper("Put"),
		"/api/v1/organization",
		controller.UpdateUserOrganizationById,
	},
}
