// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package service

import (
	endpoint1 "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics/prometheus"
	"github.com/go-kit/kit/tracing/opentracing"
	"github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/kit/transport/http"
	"github.com/lukaszozimek/organization_service/pkg/endpoint"
	http1 "github.com/lukaszozimek/organization_service/pkg/http"
	"github.com/lukaszozimek/organization_service/pkg/service"
	"github.com/oklog/oklog/pkg/group"
	opentracinggo "github.com/opentracing/opentracing-go"
)

func createService(endpoints endpoint.Endpoints) (g *group.Group) {
	g = &group.Group{}
	initHttpHandler(endpoints, g)
	initGRPCHandler(endpoints, g)
	return g
}
func defaultHttpOptions(logger log.Logger, tracer opentracinggo.Tracer) map[string][]http.ServerOption {
	options := map[string][]http.ServerOption{
		"CreateUserOrganizationById": {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "CreateUserOrganizationById", logger))},
		"DeleteUserOrganizationById": {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "DeleteUserOrganizationById", logger))},
		"GetUserOrganizationById":    {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetUserOrganizationById", logger))},
		"GetUserOrganizations":       {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "GetUserOrganizations", logger))},
		"UpdateUserOrganizationById": {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "UpdateUserOrganizationById", logger))},
		"Health":                     {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "Health", logger))},
	}
	return options
}
func defaultGRPCOptions(logger log.Logger, tracer opentracinggo.Tracer) map[string][]grpc.ServerOption {
	options := map[string][]grpc.ServerOption{
		"CreateUserOrganizationById": {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "CreateUserOrganizationById", logger))},
		"DeleteUserOrganizationById": {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "DeleteUserOrganizationById", logger))},
		"GetUserOrganizationById":    {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "GetUserOrganizationById", logger))},
		"GetUserOrganizations":       {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "GetUserOrganizations", logger))},
		"Health":                     {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "Health", logger))},
		"UpdateUserOrganizationById": {grpc.ServerErrorLogger(logger), grpc.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateUserOrganizationById", logger))},
	}
	return options
}
func addDefaultEndpointMiddleware(logger log.Logger, duration *prometheus.Summary, mw map[string][]endpoint1.Middleware) {
	mw["CreateUserOrganizationById"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "CreateUserOrganizationById")), endpoint.InstrumentingMiddleware(duration.With("method", "CreateUserOrganizationById"))}
	mw["DeleteUserOrganizationById"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "DeleteUserOrganizationById")), endpoint.InstrumentingMiddleware(duration.With("method", "DeleteUserOrganizationById"))}
	mw["GetUserOrganizationById"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "GetUserOrganizationById")), endpoint.InstrumentingMiddleware(duration.With("method", "GetUserOrganizationById"))}
	mw["GetUserOrganizations"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "GetUserOrganizations")), endpoint.InstrumentingMiddleware(duration.With("method", "GetUserOrganizations"))}
	mw["UpdateUserOrganizationById"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "UpdateUserOrganizationById")), endpoint.InstrumentingMiddleware(duration.With("method", "UpdateUserOrganizationById"))}
	mw["Health"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "Health")), endpoint.InstrumentingMiddleware(duration.With("method", "Health"))}

}
func addDefaultServiceMiddleware(logger log.Logger, mw []service.Middleware) []service.Middleware {
	return append(mw, service.LoggingMiddleware(logger))
}
func addEndpointMiddlewareToAllMethods(mw map[string][]endpoint1.Middleware, m endpoint1.Middleware) {
	methods := []string{"CreateUserOrganizationById", "DeleteUserOrganizationById", "GetUserOrganizationById", "GetUserOrganizations", "UpdateUserOrganizationById", "Health"}
	for _, v := range methods {
		mw[v] = append(mw[v], m)
	}
}
