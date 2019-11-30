package http

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kit/kit/transport/http"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/lukaszozimek/organization_service/pkg/endpoint"
	http1 "net/http"
)

// makeCreateUserOrganizationByIdHandler creates the handler logic
func makeCreateUserOrganizationByIdHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/api/v1/organization").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.CreateUserOrganizationByIdEndpoint, decodeCreateUserOrganizationByIdRequest, encodeCreateUserOrganizationByIdResponse, options...)))
}

// decodeCreateUserOrganizationByIdRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeCreateUserOrganizationByIdRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.CreateUserOrganizationByIdRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeCreateUserOrganizationByIdResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeCreateUserOrganizationByIdResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteUserOrganizationByIdHandler creates the handler logic
func makeDeleteUserOrganizationByIdHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("DELETE").Path("/api/v1/organization/{organizationId}").Handler(handlers.CORS(handlers.AllowedMethods([]string{"DELETE"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.DeleteUserOrganizationByIdEndpoint, decodeDeleteUserOrganizationByIdRequest, encodeDeleteUserOrganizationByIdResponse, options...)))
}

// decodeDeleteUserOrganizationByIdRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteUserOrganizationByIdRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.DeleteUserOrganizationByIdRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeDeleteUserOrganizationByIdResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteUserOrganizationByIdResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetUserOrganizationByIdHandler creates the handler logic
func makeGetUserOrganizationByIdHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/api/v1/organization/{organizationId}").Handler(handlers.CORS(handlers.AllowedMethods([]string{"GET"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.GetUserOrganizationByIdEndpoint, decodeGetUserOrganizationByIdRequest, encodeGetUserOrganizationByIdResponse, options...)))
}

// decodeGetUserOrganizationByIdRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetUserOrganizationByIdRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.GetUserOrganizationByIdRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetUserOrganizationByIdResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetUserOrganizationByIdResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetUserOrganizationsHandler creates the handler logic
func makeGetUserOrganizationsHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/api/v1/organization").Handler(handlers.CORS(handlers.AllowedMethods([]string{"GET"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.GetUserOrganizationsEndpoint, decodeGetUserOrganizationsRequest, encodeGetUserOrganizationsResponse, options...)))
}

// decodeGetUserOrganizationsRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetUserOrganizationsRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.GetUserOrganizationsRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetUserOrganizationsResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetUserOrganizationsResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdateUserOrganizationByIdHandler creates the handler logic
func makeUpdateUserOrganizationByIdHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("PUT").Path("/api/v1/organization").Handler(handlers.CORS(handlers.AllowedMethods([]string{"PUT"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.UpdateUserOrganizationByIdEndpoint, decodeUpdateUserOrganizationByIdRequest, encodeUpdateUserOrganizationByIdResponse, options...)))
}

// decodeUpdateUserOrganizationByIdRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateUserOrganizationByIdRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.UpdateUserOrganizationByIdRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeUpdateUserOrganizationByIdResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateUserOrganizationByIdResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
func ErrorEncoder(_ context.Context, err error, w http1.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
func ErrorDecoder(r *http1.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

// This is used to set the http status, see an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/pkg/addtransport/http.go#L133
func err2code(err error) int {
	return http1.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}
