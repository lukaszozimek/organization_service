package service

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/lukaszozimek/organization_service/pkg/model"
)

// Middleware describes a service middleware.
type Middleware func(OrganizationService) OrganizationService

type loggingMiddleware struct {
	logger log.Logger
	next   OrganizationService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a OrganizationService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next OrganizationService) OrganizationService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) CreateUserOrganizationById(ctx context.Context, organization model.Organization) (res model.Organization, err error) {
	defer func() {
		l.logger.Log("method", "CreateUserOrganizationById", "organization", organization, "res", res, "err", err)
	}()
	return l.next.CreateUserOrganizationById(ctx, organization)
}
func (l loggingMiddleware) DeleteUserOrganizationById(ctx context.Context, organizationId string) (res model.Organization, err error) {
	defer func() {
		l.logger.Log("method", "DeleteUserOrganizationById", "organizationId", organizationId, "res", res, "err", err)
	}()
	return l.next.DeleteUserOrganizationById(ctx, organizationId)
}
func (l loggingMiddleware) GetUserOrganizationById(ctx context.Context, organizationId string) (res model.Organization, err error) {
	defer func() {
		l.logger.Log("method", "GetUserOrganizationById", "organizationId", organizationId, "res", res, "err", err)
	}()
	return l.next.GetUserOrganizationById(ctx, organizationId)
}
func (l loggingMiddleware) GetUserOrganizations(ctx context.Context) (res []model.Organization, err error) {
	defer func() {
		l.logger.Log("method", "GetUserOrganizations", "res", res, "err", err)
	}()
	return l.next.GetUserOrganizations(ctx)
}
func (l loggingMiddleware) UpdateUserOrganizationById(ctx context.Context, organization model.Organization) (res model.Organization, err error) {
	defer func() {
		l.logger.Log("method", "UpdateUserOrganizationById", "organization", organization, "res", res, "err", err)
	}()
	return l.next.UpdateUserOrganizationById(ctx, organization)
}
func (l loggingMiddleware) Health(ctx context.Context) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "Health", "rs", rs, "err", err)
	}()
	return l.next.Health(ctx)
}
