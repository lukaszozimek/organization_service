package grpc

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	endpoint1 "github.com/lukaszozimek/organization_service/pkg/endpoint"
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
			conn, "pb.Organization", "CreateUserOrganizationById",
			EncodeCreateUserOrganizationByIdRequest,
			DecodeCreateUserOrganizationByIdResponse,
			pb.CreateUserOrganizationByIdReply{},
		).Endpoint()
	}

	var deleteUserOrganizationByIdEndpoint endpoint.Endpoint
	{
		deleteUserOrganizationByIdEndpoint = grpctransport.NewClient(
			conn, "pb.Organization", "DeleteUserOrganizationById",
			EncodeDeleteUserOrganizationByIdRequest,
			DecodeDeleteUserOrganizationByIdResponse,
			pb.DeleteUserOrganizationByIdReply{},
		).Endpoint()
	}

	var getUserOrganizationByIdEndpoint endpoint.Endpoint
	{
		getUserOrganizationByIdEndpoint = grpctransport.NewClient(
			conn, "pb.Organization", "GetUserOrganizationById",
			EncodeGetUserOrganizationByIdRequest,
			DecodeGetUserOrganizationByIdResponse,
			*pb.GetUserOrganizationByIdReply{},
		).Endpoint()
	}

	var getUserOrganizationsEndpoint endpoint.Endpoint
	{
		getUserOrganizationsEndpoint = grpctransport.NewClient(
			conn, "pb.Organization", "GetUserOrganizations",
			EncodeGetUserOrganizationsRequest,
			DecodeGetUserOrganizationsResponse,
			pb.GetUserOrganizationsReply{},
		).Endpoint()
	}

	var updateUserOrganizationByIdEndpoint endpoint.Endpoint
	{
		updateUserOrganizationByIdEndpoint = grpctransport.NewClient(
			conn, "pb.Organization", "UpdateUserOrganizationById",
			EncodeUpdateUserOrganizationByIdRequest,
			DecodeUpdateUserOrganizationByIdResponse,
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

// decodeGetUserOrganizationByIdResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetUserOrganizationById request.
// TODO implement the decoder
func EncodeGetUserOrganizationByIdRequest(_ context.Context, r interface{}) (interface{}, error) {
	//resp := r.(GetUserOrganizationByIdRequest)
	return &pb.GetUserOrganizationByIdRequest{}, nil
}

// encodeGetUserOrganizationByIdResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func DecodeGetUserOrganizationByIdResponse(_ context.Context, r interface{}) (interface{}, error) {
	//resp := r.(*pb.GetUserOrganizationByIdReply)
	return &pb.GetUserOrganizationByIdReply{}, nil
}

// decodeDeleteUserOrganizationByIdResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain DeleteUserOrganizationById request.
// TODO implement the decoder
func EncodeDeleteUserOrganizationByIdRequest(_ context.Context, r interface{}) (interface{}, error) {
	//resp := r.(DeleteUserOrganizationByIdRequest)
	return &pb.DeleteUserOrganizationByIdRequest{}, nil
}

func DecodeDeleteUserOrganizationByIdResponse(_ context.Context, r interface{}) (interface{}, error) {
	//resp := r.(DeleteUserOrganizationByIdReply)
	return &pb.DeleteUserOrganizationByIdReply{}, nil
}
func EncodeCreateUserOrganizationByIdRequest(_ context.Context, r interface{}) (interface{}, error) {
	//resp := r.(CreateUserOrganizationByIdRequest)
	return &pb.CreateUserOrganizationByIdRequest{}, nil
}
func DecodeCreateUserOrganizationByIdResponse(_ context.Context, r interface{}) (interface{}, error) {
	//resp := r.(CreateUserOrganizationByIdReply)
	return &pb.CreateUserOrganizationByIdReply{}, nil
}

// decodeUpdateUserOrganizationByIdResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain UpdateUserOrganizationById request.
// TODO implement the decoder
func EncodeUpdateUserOrganizationByIdRequest(_ context.Context, r interface{}) (interface{}, error) {
	//resp := r.(UpdateUserOrganizationByIdRequest)
	return &pb.UpdateUserOrganizationByIdRequest{}, nil
}

// encodeUpdateUserOrganizationByIdResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func DecodeUpdateUserOrganizationByIdResponse(_ context.Context, r interface{}) (interface{}, error) {
	//resp := r.(UpdateUserOrganizationByIdReply)
	return &pb.UpdateUserOrganizationByIdReply{}, nil
} // decodeGetUserOrganizationsResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetUserOrganizations request.
// TODO implement the decoder
func EncodeGetUserOrganizationsRequest(_ context.Context, r interface{}) (interface{}, error) {
	//resp := r.(GetUserOrganizationsRequest)
	return &endpoint1.GetUserOrganizationsRequest{}, nil
}

// encodeGetUserOrganizationsResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func DecodeGetUserOrganizationsResponse(_ context.Context, r interface{}) (interface{}, error) {
	//resp := r.(GetUserOrganizationsReply)
	return &pb.GetUserOrganizationsReply{}, nil
}
