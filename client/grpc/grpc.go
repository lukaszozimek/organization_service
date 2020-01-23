package grpc

import (
	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	endpoint1 "github.com/lukaszozimek/organization_service/pkg/endpoint"
	grpcHanlder "github.com/lukaszozimek/organization_service/pkg/grpc"
	"github.com/lukaszozimek/organization_service/pkg/grpc/pb"
	"github.com/lukaszozimek/organization_service/pkg/service"
	"google.golang.org/grpc"
)

// New returns an AddService backed by an HTTP server living at the remote
// instance. We expect instance to come from a service discovery system, so
// likely of the form "host:port".
func New(conn *grpc.ClientConn) (service.OrganizationService, error) {

	var createUserOrganizationByIdEndpoint endpoint.Endpoint
	{
		createUserOrganizationByIdEndpoint = grpctransport.NewClient(
			conn, "Organization", "CreateUserOrganizationById",
			grpcHanlder.EncodeCreateUserOrganizationByIdRequest,
			grpcHanlder.DecodeCreateUserOrganizationByIdResponse,
			pb.CreateUserOrganizationByIdReply{},
		).Endpoint()
	}

	var deleteUserOrganizationByIdEndpoint endpoint.Endpoint
	{
		deleteUserOrganizationByIdEndpoint = grpctransport.NewClient(
			conn, "Organization", "DeleteUserOrganizationById",
			grpcHanlder.EncodeDeleteUserOrganizationByIdRequest,
			grpcHanlder.DecodeDeleteUserOrganizationByIdResponse,
			pb.DeleteUserOrganizationByIdReply{},
		).Endpoint()
	}

	var getUserOrganizationByIdEndpoint endpoint.Endpoint
	{
		getUserOrganizationByIdEndpoint = grpctransport.NewClient(
			conn, "Organization", "GetUserOrganizationById",
			grpcHanlder.EncodeGetUserOrganizationByIdRequest,
			grpcHanlder.DecodeGetUserOrganizationByIdResponse,
			pb.GetUserOrganizationByIdReply{},
		).Endpoint()
	}

	var getUserOrganizationsEndpoint endpoint.Endpoint
	{
		getUserOrganizationsEndpoint = grpctransport.NewClient(
			conn, "Organization", "GetUserOrganizations",
			grpcHanlder.EncodeGetUserOrganizationsRequest,
			grpcHanlder.DecodeGetUserOrganizationsResponse,
			pb.GetUserOrganizationsReply{},
		).Endpoint()
	}

	var updateUserOrganizationByIdEndpoint endpoint.Endpoint
	{
		updateUserOrganizationByIdEndpoint = grpctransport.NewClient(
			conn, "Organization", "UpdateUserOrganizationById",
			grpcHanlder.EncodeUpdateUserOrganizationByIdRequest,
			grpcHanlder.DecodeUpdateUserOrganizationByIdResponse,
			pb.UpdateUserOrganizationByIdReply{},
		).Endpoint()
	}

	return endpoint1.Endpoints{
		CreateUserOrganizationByIdEndpoint: createUserOrganizationByIdEndpoint,
		DeleteUserOrganizationByIdEndpoint: deleteUserOrganizationByIdEndpoint,
		GetUserOrganizationByIdEndpoint:    getUserOrganizationByIdEndpoint,
		GetUserOrganizationsEndpoint:       getUserOrganizationsEndpoint,
		UpdateUserOrganizationByIdEndpoint: updateUserOrganizationByIdEndpoint,
	}, nil
}
