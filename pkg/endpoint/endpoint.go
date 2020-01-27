package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/lukaszozimek/organization_service/pkg/grpc/pb"
	"github.com/lukaszozimek/organization_service/pkg/model"
	"github.com/lukaszozimek/organization_service/pkg/service"
	"reflect"
)

// CreateUserOrganizationByIdRequest collects the request parameters for the CreateUserOrganizationById method.
type CreateUserOrganizationByIdRequest struct {
	Organization model.Organization `json:"organization"`
}

// CreateUserOrganizationByIdResponse collects the response parameters for the CreateUserOrganizationById method.
type CreateUserOrganizationByIdResponse struct {
	Organization model.Organization `json:"organization"`
	Err          error              `json:"err"`
}

// MakeCreateUserOrganizationByIdEndpoint returns an endpoint that invokes CreateUserOrganizationById on the service.
func MakeCreateUserOrganizationByIdEndpoint(s service.OrganizationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserOrganizationByIdRequest)
		organization, err := s.CreateUserOrganizationById(ctx, req.Organization)
		return CreateUserOrganizationByIdResponse{
			Err:          err,
			Organization: organization,
		}, nil
	}
}

// Failed implements Failer.
func (r CreateUserOrganizationByIdResponse) Failed() error {
	return r.Err
}

// DeleteUserOrganizationByIdRequest collects the request parameters for the DeleteUserOrganizationById method.
type DeleteUserOrganizationByIdRequest struct {
	OrganizationId string `json:"organization_id"`
}

// DeleteUserOrganizationByIdResponse collects the response parameters for the DeleteUserOrganizationById method.
type DeleteUserOrganizationByIdResponse struct {
	Organization model.Organization `json:"organization"`
	Err          error              `json:"err"`
}

// MakeDeleteUserOrganizationByIdEndpoint returns an endpoint that invokes DeleteUserOrganizationById on the service.
func MakeDeleteUserOrganizationByIdEndpoint(s service.OrganizationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteUserOrganizationByIdRequest)
		organization, err := s.DeleteUserOrganizationById(ctx, req.OrganizationId)
		return DeleteUserOrganizationByIdResponse{
			Err:          err,
			Organization: organization,
		}, nil
	}
}

// Failed implements Failer.
func (r DeleteUserOrganizationByIdResponse) Failed() error {
	return r.Err
}

// GetUserOrganizationByIdRequest collects the request parameters for the GetUserOrganizationById method.
type GetUserOrganizationByIdRequest struct {
	OrganizationId string `json:"organization_id"`
}

// GetUserOrganizationByIdResponse collects the response parameters for the GetUserOrganizationById method.
type GetUserOrganizationByIdResponse struct {
	Organization model.Organization `json:"organization"`
	Err          error              `json:"err"`
}

// MakeGetUserOrganizationByIdEndpoint returns an endpoint that invokes GetUserOrganizationById on the service.
func MakeGetUserOrganizationByIdEndpoint(s service.OrganizationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserOrganizationByIdRequest)
		organization, err := s.GetUserOrganizationById(ctx, req.OrganizationId)
		return GetUserOrganizationByIdResponse{
			Err:          err,
			Organization: organization,
		}, nil
	}
}

// Failed implements Failer.
func (r GetUserOrganizationByIdResponse) Failed() error {
	return r.Err
}

// GetUserOrganizationsRequest collects the request parameters for the GetUserOrganizations method.
type GetUserOrganizationsRequest struct{}

// GetUserOrganizationsResponse collects the response parameters for the GetUserOrganizations method.
type GetUserOrganizationsResponse struct {
	Organization []model.Organization `json:"organization"`
	Err          error                `json:"err"`
}

// MakeGetUserOrganizationsEndpoint returns an endpoint that invokes GetUserOrganizations on the service.
func MakeGetUserOrganizationsEndpoint(s service.OrganizationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		organization, err := s.GetUserOrganizations(ctx)
		return GetUserOrganizationsResponse{
			Err:          err,
			Organization: organization,
		}, nil
	}
}

// Failed implements Failer.
func (r GetUserOrganizationsResponse) Failed() error {
	return r.Err
}

// UpdateUserOrganizationByIdRequest collects the request parameters for the UpdateUserOrganizationById method.
type UpdateUserOrganizationByIdRequest struct {
	Organization model.Organization `json:"organization"`
}

// UpdateUserOrganizationByIdResponse collects the response parameters for the UpdateUserOrganizationById method.
type UpdateUserOrganizationByIdResponse struct {
	Organization model.Organization `json:"organization"`
	Err          error              `json:"err"`
}

// MakeUpdateUserOrganizationByIdEndpoint returns an endpoint that invokes UpdateUserOrganizationById on the service.
func MakeUpdateUserOrganizationByIdEndpoint(s service.OrganizationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateUserOrganizationByIdRequest)
		organization, err := s.UpdateUserOrganizationById(ctx, req.Organization)
		return UpdateUserOrganizationByIdResponse{
			Err:          err,
			Organization: organization,
		}, nil
	}
}

// Failed implements Failer.
func (r UpdateUserOrganizationByIdResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

func MakeHealthEndpoint(s service.OrganizationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		rs, err := s.Health(ctx)
		return HealthResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// HealthRequest collects the request parameters for the Health method.
type HealthRequest struct{}

// ResetFinishResponse collects the response parameters for the ResetFinish method.
type HealthResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// CreateUserOrganizationById implements Service. Primarily useful in a client.
func (e Endpoints) CreateUserOrganizationById(ctx context.Context, organization model.Organization) (res model.Organization, err error) {
	request := CreateUserOrganizationByIdRequest{Organization: organization}
	response, err := e.CreateUserOrganizationByIdEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateUserOrganizationByIdResponse).Organization, response.(CreateUserOrganizationByIdResponse).Err
}

// DeleteUserOrganizationById implements Service. Primarily useful in a client.
func (e Endpoints) DeleteUserOrganizationById(ctx context.Context, organizationId string) (organization model.Organization, err error) {
	request := DeleteUserOrganizationByIdRequest{OrganizationId: organizationId}
	response, err := e.DeleteUserOrganizationByIdEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteUserOrganizationByIdResponse).Organization, response.(DeleteUserOrganizationByIdResponse).Err
}

// GetUserOrganizationById implements Service. Primarily useful in a client.
func (e Endpoints) GetUserOrganizationById(ctx context.Context, organizationId string) (res model.Organization, err error) {
	request := GetUserOrganizationByIdRequest{OrganizationId: organizationId}
	response, err := e.GetUserOrganizationByIdEndpoint(ctx, request)
	if err != nil {
		return
	}
	if reflect.TypeOf(response).Elem() == reflect.TypeOf(new(pb.CreateUserOrganizationByIdReply)).Elem() {
		return model.Organization{}, nil
	}
	return response.(GetUserOrganizationByIdResponse).Organization, response.(GetUserOrganizationByIdResponse).Err

}

// GetUserOrganizations implements Service. Primarily useful in a client.
func (e Endpoints) GetUserOrganizations(ctx context.Context) (organization []model.Organization, err error) {
	request := GetUserOrganizationsRequest{}
	response, err := e.GetUserOrganizationsEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetUserOrganizationsResponse).Organization, response.(GetUserOrganizationsResponse).Err
}

// UpdateUserOrganizationById implements Service. Primarily useful in a client.
func (e Endpoints) UpdateUserOrganizationById(ctx context.Context, organization model.Organization) (res model.Organization, err error) {
	request := UpdateUserOrganizationByIdRequest{Organization: organization}
	response, err := e.UpdateUserOrganizationByIdEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdateUserOrganizationByIdResponse).Organization, response.(UpdateUserOrganizationByIdResponse).Err
}

// Health implements Service. Primarily useful in a client.
func (e Endpoints) Health(ctx context.Context) (rs string, err error) {
	request := HealthRequest{}
	response, err := e.HealthEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(HealthResponse).Rs, response.(HealthResponse).Err
}
