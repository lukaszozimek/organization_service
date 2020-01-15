package grpc

import (
	"context"
	"errors"
	"github.com/go-kit/kit/transport/grpc"
	"github.com/lukaszozimek/organization_service/pkg/endpoint"
	"github.com/lukaszozimek/organization_service/pkg/grpc/pb"
	context1 "golang.org/x/net/context"
)

// makeCreateUserOrganizationByIdHandler creates the handler logic
func makeCreateUserOrganizationByIdHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.CreateUserOrganizationByIdEndpoint, decodeCreateUserOrganizationByIdRequest, encodeCreateUserOrganizationByIdResponse, options...)
}

// decodeCreateUserOrganizationByIdResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain CreateUserOrganizationById request.
// TODO implement the decoder
func decodeCreateUserOrganizationByIdRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Organization' Decoder is not impelemented")
}

// encodeCreateUserOrganizationByIdResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeCreateUserOrganizationByIdResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Organization' Encoder is not impelemented")
}
func (g *grpcServer) CreateUserOrganizationById(ctx context1.Context, req *pb.CreateUserOrganizationByIdRequest) (*pb.CreateUserOrganizationByIdReply, error) {
	_, rep, err := g.createUserOrganizationById.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateUserOrganizationByIdReply), nil
}

// makeDeleteUserOrganizationByIdHandler creates the handler logic
func makeDeleteUserOrganizationByIdHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.DeleteUserOrganizationByIdEndpoint, decodeDeleteUserOrganizationByIdRequest, encodeDeleteUserOrganizationByIdResponse, options...)
}

// decodeDeleteUserOrganizationByIdResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain DeleteUserOrganizationById request.
// TODO implement the decoder
func decodeDeleteUserOrganizationByIdRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Organization' Decoder is not impelemented")
}

// encodeDeleteUserOrganizationByIdResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeDeleteUserOrganizationByIdResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Organization' Encoder is not impelemented")
}
func (g *grpcServer) DeleteUserOrganizationById(ctx context1.Context, req *pb.DeleteUserOrganizationByIdRequest) (*pb.DeleteUserOrganizationByIdReply, error) {
	_, rep, err := g.deleteUserOrganizationById.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteUserOrganizationByIdReply), nil
}

// makeGetUserOrganizationByIdHandler creates the handler logic
func makeGetUserOrganizationByIdHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetUserOrganizationByIdEndpoint, decodeGetUserOrganizationByIdRequest, encodeGetUserOrganizationByIdResponse, options...)
}

// decodeGetUserOrganizationByIdResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetUserOrganizationById request.
// TODO implement the decoder
func decodeGetUserOrganizationByIdRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Organization' Decoder is not impelemented")
}

// encodeGetUserOrganizationByIdResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetUserOrganizationByIdResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Organization' Encoder is not impelemented")
}
func (g *grpcServer) GetUserOrganizationById(ctx context1.Context, req *pb.GetUserOrganizationByIdRequest) (*pb.GetUserOrganizationByIdReply, error) {
	_, rep, err := g.getUserOrganizationById.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetUserOrganizationByIdReply), nil
}

// makeGetUserOrganizationsHandler creates the handler logic
func makeGetUserOrganizationsHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetUserOrganizationsEndpoint, decodeGetUserOrganizationsRequest, encodeGetUserOrganizationsResponse, options...)
}

// decodeGetUserOrganizationsResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetUserOrganizations request.
// TODO implement the decoder
func decodeGetUserOrganizationsRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Organization' Decoder is not impelemented")
}

// encodeGetUserOrganizationsResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetUserOrganizationsResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Organization' Encoder is not impelemented")
}
func (g *grpcServer) GetUserOrganizations(ctx context1.Context, req *pb.GetUserOrganizationsRequest) (*pb.GetUserOrganizationsReply, error) {
	_, rep, err := g.getUserOrganizations.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetUserOrganizationsReply), nil
}

// makeUpdateUserOrganizationByIdHandler creates the handler logic
func makeUpdateUserOrganizationByIdHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.UpdateUserOrganizationByIdEndpoint, decodeUpdateUserOrganizationByIdRequest, encodeUpdateUserOrganizationByIdResponse, options...)
}

// decodeUpdateUserOrganizationByIdResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain UpdateUserOrganizationById request.
// TODO implement the decoder
func decodeUpdateUserOrganizationByIdRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Organization' Decoder is not impelemented")
}

// encodeUpdateUserOrganizationByIdResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeUpdateUserOrganizationByIdResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Organization' Encoder is not impelemented")
}
func (g *grpcServer) UpdateUserOrganizationById(ctx context1.Context, req *pb.UpdateUserOrganizationByIdRequest) (*pb.UpdateUserOrganizationByIdReply, error) {
	_, rep, err := g.updateUserOrganizationById.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UpdateUserOrganizationByIdReply), nil
}

// makeHealthHandler creates the handler logic
func makeHealthHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.HealthEndpoint, decodeHealthRequest, encodeHealthResponse, options...)
}

// decodeHealthResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain Health request.
// TODO implement the decoder
func decodeHealthRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Organization' Decoder is not impelemented")
}

// encodeHealthResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeHealthResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Organization' Encoder is not impelemented")
}
func (g *grpcServer) Health(ctx context1.Context, req *pb.HealthRequest) (*pb.HealthReply, error) {
	_, rep, err := g.health.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.HealthReply), nil
}
