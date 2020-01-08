// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package http

import (
	"github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/lukaszozimek/organization_service/pkg/endpoint"

	http1 "net/http"
)

// NewHTTPHandler returns a handler that makes a set of endpoints available on
// predefined paths.
func NewHTTPHandler(endpoints endpoint.Endpoints, options map[string][]http.ServerOption) http1.Handler {
	m := mux.NewRouter()
	makeCreateUserOrganizationByIdHandler(m, endpoints, options["CreateUserOrganizationById"])
	makeDeleteUserOrganizationByIdHandler(m, endpoints, options["DeleteUserOrganizationById"])
	makeGetUserOrganizationByIdHandler(m, endpoints, options["GetUserOrganizationById"])
	makeGetUserOrganizationsHandler(m, endpoints, options["GetUserOrganizations"])
	makeUpdateUserOrganizationByIdHandler(m, endpoints, options["UpdateUserOrganizationById"])
	makeHealthHandler(m, endpoints, options["Health"])
	return m
}
