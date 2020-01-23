package grpc

import (
	"bytes"
	"context"
	"encoding/json"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	grpcHanlder "github.com/lukaszozimek/organization_service/pkg/grpc"
	"google.golang.org/grpc"
	"io/ioutil"
	"net/url"
)

// New returns an AddService backed by an HTTP server living at the remote
// instance. We expect instance to come from a service discovery system, so
// likely of the form "host:port".
func New(conn *grpc.ClientConn) (service.OrganizationService, error) {

	var createUserOrganizationByIdEndpoint endpoint.Endpoint
	{
		createUserOrganizationByIdEndpoint = grpctransport.NewClient(
			conn, "Lorem", "Lorem",
			grpcHanlder.EncodeCreateUserOrganizationByIdRequest,
			grpcHanlder.DecodeCreateUserOrganizationByIdResponse,
			pb.LoremResponse{},
		).Endpoint()
	}

	var deleteUserOrganizationByIdEndpoint endpoint.Endpoint
	{
		deleteUserOrganizationByIdEndpoint = grpctransport.NewClient(
			conn, "Lorem", "Lorem",
			grpcHanlder.EncodeCreateUserOrganizationByIdResponse,
			grpcHanlder.DecodeCreateUserOrganizationByIdResponse,
			pb.LoremResponse{},
		).Endpoint()
	}

	var getUserOrganizationByIdEndpoint endpoint.Endpoint
	{
		getUserOrganizationByIdEndpoint = grpctransport.NewClient(
			conn, "Lorem", "Lorem",
			lorem_grpc.EncodeGRPCLoremRequest,
			lorem_grpc.DecodeGRPCLoremResponse,
			pb.LoremResponse{},
		).Endpoint()
	}

	var getUserOrganizationsEndpoint endpoint.Endpoint
	{
		getUserOrganizationsEndpoint = grpctransport.NewClient(
			conn, "Lorem", "Lorem",
			lorem_grpc.EncodeGRPCLoremRequest,
			lorem_grpc.DecodeGRPCLoremResponse,
			pb.LoremResponse{},
		).Endpoint()
	}

	var updateUserOrganizationByIdEndpoint endpoint.Endpoint
	{
		updateUserOrganizationByIdEndpoint = grpctransport.NewClient(
			conn, "Lorem", "Lorem",
			lorem_grpc.EncodeGRPCLoremRequest,
			lorem_grpc.DecodeGRPCLoremResponse,
			pb.LoremResponse{},
		).Endpoint()
	}

	var healthEndpoint endpoint.Endpoint
	{
		healthEndpoint = grpctransport.NewClient(
			conn, "Lorem", "Lorem",
			lorem_grpc.EncodeGRPCLoremRequest,
			lorem_grpc.DecodeGRPCLoremResponse,
			pb.LoremResponse{},
		).Endpoint()
	}

	return endpoint1.Endpoints{
		CreateUserOrganizationByIdEndpoint: createUserOrganizationByIdEndpoint,
		DeleteUserOrganizationByIdEndpoint: deleteUserOrganizationByIdEndpoint,
		GetUserOrganizationByIdEndpoint:    getUserOrganizationByIdEndpoint,
		GetUserOrganizationsEndpoint:       getUserOrganizationsEndpoint,
		HealthEndpoint:                     healthEndpoint,
		UpdateUserOrganizationByIdEndpoint: updateUserOrganizationByIdEndpoint,
	}, nil
}

// EncodeHTTPGenericRequest is a transport/http.EncodeRequestFunc that
// SON-encodes any request to the request body. Primarily useful in a client.
func encodeHTTPGenericRequest(_ context.Context, r *http1.Request, request interface{}) error {
	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

// decodeCreateUserOrganizationByIdResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeCreateUserOrganizationByIdResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.CreateUserOrganizationByIdResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeDeleteUserOrganizationByIdResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeDeleteUserOrganizationByIdResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.DeleteUserOrganizationByIdResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeGetUserOrganizationByIdResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeGetUserOrganizationByIdResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.GetUserOrganizationByIdResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeGetUserOrganizationsResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeGetUserOrganizationsResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.GetUserOrganizationsResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeUpdateUserOrganizationByIdResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeUpdateUserOrganizationByIdResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.UpdateUserOrganizationByIdResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeHealthResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeHealthResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.HealthResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}
func copyURL(base *url.URL, path string) (next *url.URL) {
	n := *base
	n.Path = path
	next = &n
	return
}
